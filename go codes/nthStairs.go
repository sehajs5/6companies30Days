package main

import "math"

func waysToReachStair(k int) int {
	dp := make(map[int]map[int]map[bool]int)
	return solve4(k, dp, 1, 0, false)
}
func solve4(k int, dp map[int]map[int]map[bool]int, i int, jump int, down bool) int {
	if i > k+1 {
		return 0
	}

	if dp[i] == nil {
		dp[i] = make(map[int]map[bool]int)
	}

	if dp[i][jump] == nil {
		dp[i][jump] = make(map[bool]int)
	}

	if val, ok := dp[i][jump][down]; ok {
		return val
	}
	ans := 0
	if i == k {
		ans = 1
	}
	if i != 0 && !down {
		ans += solve4(k, dp, i-1, jump, true)
	}
	ans += solve4(k, dp, i+int(math.Pow(float64(2), float64(jump))), jump+1, false)

	dp[i][jump][down] = ans

	return ans

}
