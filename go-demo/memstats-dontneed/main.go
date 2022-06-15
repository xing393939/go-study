package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	useDontNeed := flag.Bool("dontneed", false, "use MADV_DONTNEED instead of MADV_FREE")
	flag.Parse()

	// anonymous mapping
	m, _ := syscall.Mmap(-1, 0, 10<<20, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_PRIVATE|syscall.MAP_ANON)
	printStats("After anon mmap:", m)

	// page fault by accessing it
	for i := 0; i < len(m); i += pageSize {
		m[i] = 42
	}
	printStats("After anon fault:", m)

	// use different strategy
	if *useDontNeed {
		_ = syscall.Madvise(m, syscall.MADV_DONTNEED)
		printStats("After MADV_DONTNEED:", m)
	} else {
		_ = syscall.Madvise(m, 0x5)
		printStats("After MADV_FREE:", m)
	}
	stat, _ := getSmaps()
	fmt.Printf("VSS: %d MiB, RSS: %d MiB, PSS: %d MiB, USS: %d MiB\n",
		stat.VSS/(1<<20), stat.RSS/(1<<20), stat.PSS/(1<<20), stat.USS/(1<<20))
	runtime.KeepAlive(m)
}

func printStats(ident string, m []byte) {
	fmt.Print(ident, " ", rss(), " MiB RSS\n")
}

var pageSize = syscall.Getpagesize()

func rss() int {
	data, err := ioutil.ReadFile("/proc/self/stat")
	if err != nil {
		log.Fatal(err)
	}
	fs := strings.Fields(string(data))
	rss, err := strconv.ParseInt(fs[23], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(uintptr(rss) * uintptr(pageSize) / (1 << 20)) // MiB
}

type mmapStat struct {
	Size           uint64
	RSS            uint64
	PSS            uint64
	PrivateClean   uint64
	PrivateDirty   uint64
	PrivateHugetlb uint64
}

func getMmaps() (*[]mmapStat, error) {
	var ret []mmapStat
	contents, err := ioutil.ReadFile("/proc/self/smaps")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(contents), "\n")
	// function of parsing a block
	getBlock := func(block []string) (mmapStat, error) {
		m := mmapStat{}
		for _, line := range block {
			if strings.Contains(line, "VmFlags") ||
				strings.Contains(line, "Name") {
				continue
			}
			field := strings.Split(line, ":")
			if len(field) < 2 {
				continue
			}
			v := strings.Trim(field[1], " kB") // remove last "kB"
			t, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return m, err
			}
			switch field[0] {
			case "Size":
				m.Size = t
			case "Rss":
				m.RSS = t
			case "Pss":
				m.PSS = t
			case "Private_Clean":
				m.PrivateClean = t
			case "Private_Dirty":
				m.PrivateDirty = t
			case "Private_Hugetlb":
				m.PrivateHugetlb = t
			}
		}
		return m, nil
	}
	blocks := make([]string, 16)
	for _, line := range lines {
		if strings.HasSuffix(strings.Split(line, " ")[0], ":") == false {
			if len(blocks) > 0 {
				g, err := getBlock(blocks)
				if err != nil {
					return &ret, err
				}
				ret = append(ret, g)
			}
			blocks = make([]string, 16)
		} else {
			blocks = append(blocks, line)
		}
	}
	return &ret, nil
}

type smapsStat struct {
	VSS uint64 // bytes
	RSS uint64 // bytes
	PSS uint64 // bytes
	USS uint64 // bytes
}

func getSmaps() (*smapsStat, error) {
	mmaps, err := getMmaps()
	if err != nil {
		panic(err)
	}
	smaps := &smapsStat{}
	for _, mmap := range *mmaps {
		smaps.VSS += mmap.Size * 1014
		smaps.RSS += mmap.RSS * 1024
		smaps.PSS += mmap.PSS * 1024
		smaps.USS += mmap.PrivateDirty*1024 + mmap.PrivateClean*1024 + mmap.PrivateHugetlb*1024
	}
	return smaps, nil
}
