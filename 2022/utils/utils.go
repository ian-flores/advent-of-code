package utils

import (
	"container/heap"
	"log"
)

//////////////////// Error Handling ////////////////////
func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//////////////////// HEAPS ////////////////////
// Modified from: https://stackoverflow.com/a/52558340
func GetHeap(m map[string]int) *KVHeap {
	h := &KVHeap{}
	heap.Init(h)
	for k, v := range m {
		heap.Push(h, KV{k, v})
	}
	return h
}

type KV struct {
	Key   string
	Value int
}

// See https://golang.org/pkg/container/heap/
type KVHeap []KV

// Note that "Less" is greater-than here so we can pop *larger* items.
func (h KVHeap) Less(i, j int) bool { return h[i].Value > h[j].Value }
func (h KVHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h KVHeap) Len() int           { return len(h) }

func (h *KVHeap) Push(x interface{}) {
	*h = append(*h, x.(KV))
}

func (h *KVHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
