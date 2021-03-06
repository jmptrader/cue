// Copyright (c) 2016 Bob Ziuchkovski
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cue

import (
	"testing"
	"time"
)

type noopCollector struct{}

func (c *noopCollector) Collect(event *Event) error { return nil }

func BenchmarkUncollected(b *testing.B) {
	defer resetCue()
	defer b.StopTimer()

	log := NewLogger("test")
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		log.Debug("test")
	}
}

func BenchmarkParallelUncollected(b *testing.B) {
	defer resetCue()
	defer b.StopTimer()

	log := NewLogger("test")
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Debug("test")
		}
	})
}

func BenchmarkSyncNoopCollector(b *testing.B) {
	defer resetCue()
	defer b.StopTimer()

	c := &noopCollector{}
	Collect(DEBUG, c)

	log := NewLogger("test")
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		log.Debug("test")
	}
}

func BenchmarkParallelSyncNoopCollector(b *testing.B) {
	defer resetCue()
	defer b.StopTimer()

	c := &noopCollector{}
	Collect(DEBUG, c)

	log := NewLogger("test")
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Debug("test")
		}
	})
}

func BenchmarkAsyncNoopCollector(b *testing.B) {
	defer resetCue()
	defer b.StopTimer()

	c := &noopCollector{}
	CollectAsync(DEBUG, b.N, c)

	log := NewLogger("test")
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		log.Debug("test")
	}

	err := Close(time.Minute)
	if err != nil {
		panic(err)
	}
}

func BenchmarkParallelAsyncNoopCollector(b *testing.B) {
	defer resetCue()
	defer b.StopTimer()

	c := &noopCollector{}
	CollectAsync(DEBUG, b.N, c)

	log := NewLogger("test")
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Debug("test")
		}
	})

	err := Close(time.Minute)
	if err != nil {
		panic(err)
	}
}
