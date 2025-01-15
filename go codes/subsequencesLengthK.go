package main

import (
	"container/heap"
	"sort"
)

type Pair struct {
	val, indx int
}

func maxSubsequence(nums []int, k int) []int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for i, num := range nums {
		heap.Push(minHeap, Pair{num, i})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}
	result := make([]Pair, 0, k)
	for minHeap.Len() > 0 {
		result = append(result, heap.Pop(minHeap).(Pair))
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].indx < result[j].indx
	})
	ans := []int{}
	for i := 0; i < len(result); i++ {
		ans = append(ans, result[i].val)
	}
	return ans
}

type MinHeap []Pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}
