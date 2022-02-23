package test

type HashFunc func()

type options struct {
	hashFunc    HashFunc
	bucketCount uint64
}

type Option interface {
	apply(*options)
}

type Bucket struct {
	count uint64
}

func (b Bucket) apply(opts *options) {
	opts.bucketCount = b.count
}

func WithBucketCount(count uint64) Option {
	return Bucket{
		count: count,
	}
}

type Hash struct {
	hashFunc HashFunc
}

func (h Hash) apply(opts *options) {
	opts.hashFunc = h.hashFunc
}

func WithHashFunc(hashFunc HashFunc) Option {
	return Hash{hashFunc: hashFunc}
}

func NewCache(opts ...Option) *cache {
	o := &options{
		hashFunc:    func() {},
		bucketCount: 0,
	}
	for _, each := range opts {
		each.apply(o)
	}
	return &cache{}
}

func main() {
	NewCache(WithBucketCount(128), WithHashFunc(func() {}))
}
