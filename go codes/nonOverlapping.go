package main

import (
	"math/rand"
	"sort"
	"time"
)

type Solution struct {
	rects      [][]int
	prefixSums []int
}

func Constructor(rects [][]int) Solution {
	n := len(rects)
	prefixSums := make([]int, n+1)
	for i, rect := range rects {
		widthPoints := rect[2] - rect[0] + 1
		heightPoints := rect[3] - rect[1] + 1
		area := widthPoints * heightPoints
		prefixSums[i+1] = prefixSums[i] + area
	}

	rand.Seed(time.Now().UnixNano())

	return Solution{rects: rects, prefixSums: prefixSums}

}

func binarySearch(nums []int, k int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
func (s *Solution) Pick() []int {
	n := len(s.prefixSums)
	target := rand.Intn(s.prefixSums[n-1]) + 1
	indx := binarySearch(s.prefixSums, target)

	rect := s.rects[indx-1]
	x := rect[0] + rand.Intn(rect[2]-rect[0]+1)
	y := rect[1] + rand.Intn(rect[3]-rect[1]+1)

	return []int{x, y}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
