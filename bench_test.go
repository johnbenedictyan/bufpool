package bufpool_test

import (
	"encoding/json"
	"sync"
	"testing"

	"github.com/vmihailenco/bufpool"
)

func BenchmarkNoPool(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var buf bufpool.Buffer
			benchOne(b, &buf)
		}
	})
}

type Record struct {
	A string
	B int
}

var data = make([]Record, 1000)

func benchOne(b *testing.B, buf *bufpool.Buffer) {
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		b.Fatal(err)
	}
}

var syncPool = sync.Pool{
	New: func() interface{} {
		return new(bufpool.Buffer)
	},
}

func BenchmarkSyncPool(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := syncPool.Get().(*bufpool.Buffer)
			buf.Reset(nil)

			benchOne(b, buf)

			syncPool.Put(buf)
		}
	})
}

var bufPool bufpool.Pool

func BenchmarkBufPool(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := bufPool.Get()

			benchOne(b, buf)

			bufPool.Put(buf)
		}
	})
}
