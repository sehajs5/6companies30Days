package main

import "sort"

func minMoves2(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	median := nums[n/2]
	if n%2 == 0 {
		median = (nums[n/2-1] + nums[n/2]) / 2
	}
	ans := 0
	for _, num := range nums {
		diff := median - num
		if diff < 0 {
			diff = -diff
		}
		ans += diff
	}
	return ans
}
