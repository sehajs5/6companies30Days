package main

import "sort"

func findPairs(nums []int, k int) int {
	sort.Ints(nums)

	mapp := make(map[int]int)
	for _, num := range nums {
		mapp[num]++
	}
	pairs := 0
	if k == 0 {
		for _, val := range mapp {
			if val > 1 {
				pairs++
			}
		}
		return pairs
	} else {
		for i := 0; i < len(nums); i++ {
			if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
				if mapp[nums[i]+k] > 0 {
					pairs++
				}
			}
		}
	}
	return pairs
}
