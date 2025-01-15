package main

func numberOfSubarrays(nums []int, k int) int {
	oddNumCount := make(map[int]int)
	oddNumCount[0] = 1
	odd := 0
	count := 0
	for _, num := range nums {
		if num%2 != 0 {
			odd++
		}
		oddNumCount[odd]++
		if odd >= k {
			count += oddNumCount[odd-k]
		}
	}
	return count
}
