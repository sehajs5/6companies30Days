package main

func incremovableSubarrayCount(nums []int) int {
	count := 0
	n := len(nums)
	for i := 0; i < n; i++ {

		for j := i; j < n; j++ {

			ok := true
			prev := -1
			for k := 0; k < n; k++ {
				if k >= i && k <= j {
					continue
				} else {
					if nums[k] <= prev {
						ok = false
					} else {
						prev = nums[k]
					}
				}
			}
			if ok {
				count++
			}

		}
	}
	return count
}
