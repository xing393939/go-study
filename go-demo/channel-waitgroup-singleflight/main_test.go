package channel_waitgroup_singleflight

import (
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func getAFile(filename string) io.WriteCloser {
	f1, err := os.Create(filename)
	if err != nil {
		return nil
	}
	return f1
}

func BenchmarkChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := make(chan struct{})
		go func() {
			<-ch
		}()
		close(ch)
	}
}

func BenchmarkWaitGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			wg.Wait()
		}()
		wg.Done()
	}
}

func TestNormal(t *testing.T) {
	atomic.StoreInt64(&count, 0)
	atomic.StoreInt64(&cache, 0)
	eg := errgroup.Group{}

	for i := 0; i < 1000; i++ {
		eg.Go(func() error {
			v := queryCache()
			if v == 0 {
				v = queryDB()
				setCache(v)
			}
			return nil
		})
	}
	_ = eg.Wait()
	t.Log(count)
}

func TestSingleFlight(t *testing.T) {
	atomic.StoreInt64(&count, 0)
	atomic.StoreInt64(&cache, 0)
	eg := errgroup.Group{}
	single := singleflight.Group{}

	for i := 0; i < 1000; i++ {
		eg.Go(func() error {
			v := queryCache()
			if v == 0 {
				obj, _, _ := single.Do("cacheKey", func() (i interface{}, e error) {
					v = queryDB()
					setCache(v)
					return v, nil
				})
				v = obj.(int64)
			}
			return nil
		})
	}
	_ = eg.Wait()
	t.Log(count)
}

var count = int64(0)
var cache = int64(0)

func queryCache() int64 {
	return atomic.LoadInt64(&cache)
}

func setCache(v int64) {
	atomic.StoreInt64(&cache, v)
}

func queryDB() int64 {
	atomic.AddInt64(&count, 1)   // 统计次数
	time.Sleep(time.Millisecond) // 模拟耗时
	return 1
}
