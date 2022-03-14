package backtrace

import "unsafe"

func zzzPrintP(u unsafe.Pointer) {
	print(u, " ")
}

func zzzPrintLn() {
	println()
}

func ZzzPrintTrace()
