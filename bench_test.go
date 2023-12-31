package main

import (
	"sync"
	"testing"
)

type Counter struct {
	A int
	B int
}

func inc(c *Counter) {
	c.A++
	c.B++
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(Counter)
	},
}

func BenchmarkWithoutPool(b *testing.B) {
	var s *Counter
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			s = new(Counter)
			b.StopTimer()
			inc(s)
			b.StartTimer()
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var s *Counter
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			s = pool.Get().(*Counter)
			b.StopTimer()
			inc(s)
			b.StartTimer()
			pool.Put(s)
		}
	}
}
