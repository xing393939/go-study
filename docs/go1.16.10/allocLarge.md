```go
func (c *mcache) allocLarge(size uintptr, needzero bool, noscan bool) *mspan {
	if size+_PageSize < size {
		throw("out of memory")
	}
	npages := size >> _PageShift
	if size&_PageMask != 0 {
		npages++
	}

	// 此次分配需要扣除gc信用
	deductSweepCredit(npages*_PageSize, npages)

	spc := makeSpanClass(0, noscan)
	s := mheap_.alloc(npages, spc, needzero)
	if s == nil {
		throw("out of memory")
	}
	stats := memstats.heapStats.acquire()
	atomic.Xadduintptr(&stats.largeAlloc, npages*pageSize)
	atomic.Xadduintptr(&stats.largeAllocCount, 1)
	memstats.heapStats.release()

	// Update heap_live and revise pacing if needed.
	atomic.Xadd64(&memstats.heap_live, int64(npages*pageSize))
	if trace.enabled {
		// Trace that a heap alloc occurred because heap_live changed.
		traceHeapAlloc()
	}
	if gcBlackenEnabled != 0 {
		gcController.revise()
	}

	// Put the large span in the mcentral swept list so that it's
	// visible to the background sweeper.
	mheap_.central[spc].mcentral.fullSwept(mheap_.sweepgen).push(s)
	s.limit = s.base() + size
	heapBitsForAddr(s.base()).initSpan(s)
	return s
}
```