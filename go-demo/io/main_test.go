package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"testing"
	"unsafe"
)

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, _ := os.Open("test.txt")
		reader := bufio.NewReaderSize(file, 1024)
		all := bytes.Buffer{}
		for {
			tmp := make([]byte, 1024)
			_, err := reader.Read(tmp)
			if err == io.EOF {
				break
			}
			all.Write(tmp)
		}
		s := all.String()
		_ = s
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reader, _ := os.Open("test.txt")
		all := bytes.Buffer{}
		for {
			tmp := make([]byte, 1024)
			_, err := reader.Read(tmp)
			if err == io.EOF {
				break
			}
			all.Write(tmp)
		}
		by := all.Bytes()
		s := *(*string)(unsafe.Pointer(&by))
		_ = s
	}
}

func BenchmarkC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reader, _ := os.Open("test.txt")
		all := bytes.Buffer{}
		for {
			tmp := make([]byte, 1024)
			_, err := reader.Read(tmp)
			if err == io.EOF {
				break
			}
			all.Write(tmp)
		}
		s := all.String()
		_ = s
	}
}
