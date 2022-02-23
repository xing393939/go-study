package test

type cache struct {
	bucketCount uint64
	close       chan struct{}
}

type Opt func(opt *cache)

func NewCache2(opts ...Opt) {
	c := &cache{
		close: make(chan struct{}),
	}
	for _, each := range opts {
		each(c)
	}
}

func SetShardCount(count uint64) Opt {
	return func(opt *cache) {
		opt.bucketCount = count
	}
}

func main() {
	NewCache2(SetShardCount(256))
}
