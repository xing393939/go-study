package backtrace

import "strconv"

func zzzPrintP(u int64) {
	print("0x", strconv.FormatInt(u, 16), " ")
}

func zzzPrintLn() {
	println()
}

func ZzzPrintTrace()
