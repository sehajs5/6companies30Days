package main

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make(map[int]map[int][]int)
	prev := -1
	return solve5(nums, 0, prev, dp)
}
func solve5(nums []int, i int, prev int, dp map[int]map[int][]int) []int {
	if len(nums) == i {
		return []int{}
	}
	if dp[i] == nil {
		dp[i] = make(map[int][]int)
	}
	if val, found := dp[i][prev]; found {
		return val
	}
	take := []int{}
	if prev == -1 || nums[i]%nums[prev] == 0 {
		take = solve5(nums, i+1, i, dp)
		take = append(take, nums[i])
	}
	notTake := []int{}
	notTake = solve5(nums, i+1, prev, dp)

	if len(notTake) > len(take) {
		dp[i][prev] = notTake
		return notTake
	} else {
		dp[i][prev] = take
		return take
	}
}
