package main

type Node struct {
}

func maxRotateFunction(nums []int) int {
	n := len(nums)
	sum := 0
	f0 := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		f0 += i * nums[i]
	}
	mapp := make(map[int]int)
	mapp[0] = f0
	maxi := f0
	for i := 1; i < n; i++ {
		mapp[i] = mapp[i-1] + sum - n*nums[n-i]
		maxi = max(mapp[i], maxi)
	}
	return maxi
}
