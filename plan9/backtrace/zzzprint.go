package backtrace

func zzzPrintP(n int64) {
	h := [16]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	i16 := int64(16)
	arr := [16]byte{}
	for i := 0; n > 0; n /= i16 {
		lsb := h[n%i16]
		i++
		arr[i] = lsb
	}
	result := make([]byte, 16, 16)
	for i := 0; i < 16; i++ {
		result[i] = arr[16-i-1]
	}
	print("0x", string(result), " ")
}

func zzzPrintLn() {
	println()
}

func ZzzPrintTrace()
