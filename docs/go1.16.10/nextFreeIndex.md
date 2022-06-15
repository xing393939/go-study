```go
/*
初始状态，s.freeindex是0，allocCache全是1，allocBits全是0
第一次分配，freeindex为1，allocCache为01111...
第二次分配，freeindex是2，allocCache为00111...
span被清扫后...
第一次分配，freeindex是2，allocCache是001010...
第二次分配，freeindex是4，allocCache是000010...

当s.allocCache全是0时，通过s.refillAllocCache从allocBits取空闲状态，
首先将freeindex取整到64倍数，并填充allocCache的位数。
 */
func (s *mspan) nextFreeIndex() uintptr {
	sfreeindex := s.freeindex
	snelems := s.nelems
	if sfreeindex == snelems {
		return sfreeindex
	}
	if sfreeindex > snelems {
		throw("s.freeindex > s.nelems")
	}
	// aCache是uint64，二进制1表示空闲，0表示已用
	aCache := s.allocCache

	bitIndex := sys.Ctz64(aCache)
	for bitIndex == 64 {
		// Move index to start of next cached bits.
		sfreeindex = (sfreeindex + 64) &^ (64 - 1)
		if sfreeindex >= snelems {
			s.freeindex = snelems
			return snelems
		}
		whichByte := sfreeindex / 8
		// Refill s.allocCache with the next 64 alloc bits.
		s.refillAllocCache(whichByte)
		aCache = s.allocCache
		bitIndex = sys.Ctz64(aCache)
	}
	result := sfreeindex + uintptr(bitIndex)
	if result >= snelems {
		s.freeindex = snelems
		return snelems
	}

	s.allocCache >>= uint(bitIndex + 1)
	sfreeindex = result + 1

	if sfreeindex%64 == 0 && sfreeindex != snelems {
		whichByte := sfreeindex / 8
		s.refillAllocCache(whichByte)
	}
	s.freeindex = sfreeindex
	return result
}

func (s *mspan) refillAllocCache(whichByte uintptr) {
	bytes := (*[8]uint8)(unsafe.Pointer(s.allocBits.bytep(whichByte)))
	aCache := uint64(0)
	aCache |= uint64(bytes[0])
	aCache |= uint64(bytes[1]) << (1 * 8)
	aCache |= uint64(bytes[2]) << (2 * 8)
	aCache |= uint64(bytes[3]) << (3 * 8)
	aCache |= uint64(bytes[4]) << (4 * 8)
	aCache |= uint64(bytes[5]) << (5 * 8)
	aCache |= uint64(bytes[6]) << (6 * 8)
	aCache |= uint64(bytes[7]) << (7 * 8)
	s.allocCache = ^aCache
}
```