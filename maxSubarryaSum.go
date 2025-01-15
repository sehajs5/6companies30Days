package main

func maximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	if n < k {
		return 0
	}
	currentSum := 0
	mapp := make(map[int]int)
	maxSum := 0
	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]
		mapp[nums[i]]++

		if i >= k {
			left := nums[i-k]
			mapp[left]--
			currentSum -= left

			if mapp[left] == 0 {
				delete(mapp, left)
			}
		}

		if len(mapp) == k {
			maxSum = max(maxSum, currentSum)
		}
	}
	return int64(maxSum)
}
