package main

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	heights := []int{}
	for _, e := range envelopes {
		heights = append(heights, e[1])
	}
	return lis(heights)
}

func lis(heights []int) int {
	dp := []int{}
	for _, num := range heights {
		pos := solve(num, dp)

		if pos == len(dp) {
			dp = append(dp, num)
		} else {
			dp[pos] = num
		}
	}
	return len(dp)
}
func solve(num int, dp []int) int {
	left := 0
	right := len(dp)
	for left < right {
		mid := left + (right-left)/2
		if dp[mid] < num {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
