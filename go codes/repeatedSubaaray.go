package main

func findLength(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	dp := make(map[int]map[int]int)
	for i := 1; i <= m; i++ {
		dp[i] = make(map[int]int)
	}
	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			ans = max(ans, dp[i][j])
		}
	}
	return ans
}
