package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	Three := m.BuckHashSys + m.GCSys + m.OtherSys
	// printKB(m.BuckHashSys, m.GCSys, m.OtherSys) // 证明这三个不是8K对齐的
	// sys总内存 = Three + m.HeapSys + m.StackSys + m.MCacheSys + m.MSpanSys
	printKB(m.Sys, Three+m.HeapSys+m.StackSys+m.MCacheSys+m.MSpanSys)
	// 堆内内存 = m.HeapInuse + m.HeapIdle
	printKB(m.HeapSys, m.HeapInuse+m.HeapIdle)
	// 栈内存 = StackSys = StackInuse
	printKB(m.StackSys, m.StackInuse)

	mapNone := getMapNone()
	printKB(m.Sys + mapNone)
	for {
	}
}

func printKB(args ...uint64) {
	newArgs := make([]interface{}, len(args))
	for k, v := range args {
		newArgs[k] = float64(v) / 1024
	}
	fmt.Println(newArgs)
}

func getMapNone() uint64 {
	contents, err := ioutil.ReadFile("/proc/self/maps")
	if err != nil {
		return 0
	}
	mapNone := uint64(0)
	lines := strings.Split(string(contents), "\n")
	for k := range lines {
		row := strings.Split(lines[k], " ")
		if k < 4 {
			mapNone += getLineSub(row[0])
		}
		if len(row) > 2 && row[1] == "---p" {
			mapNone += getLineSub(row[0])
		}
	}
	return mapNone
}

func getLineSub(two string) uint64 {
	lines := strings.Split(two, "-")
	aLen := len(lines[0])
	if aLen > 8 {
		lines[0] = lines[0][4:aLen]
		lines[1] = lines[1][4:aLen]
	}
	a, _ := strconv.ParseUint(lines[0], 16, 32)
	b, _ := strconv.ParseUint(lines[1], 16, 32)
	c := b - a
	return c
}
