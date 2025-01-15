package main

func longestMountain(arr []int) int {
	n := len(arr)
	if len(arr) < 3 {
		return 0
	}
	res := 0
	up := make([]int, n)
	down := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			down[i] = down[i+1] + 1
		}
	}
	for i := 0; i < n; i++ {
		if i > 0 && arr[i] > arr[i-1] {
			up[i] = up[i-1] + 1
		}
		if up[i] > 0 && down[i] > 0 {
			res = max(res, up[i]+down[i]+1)
		}
	}
	return res
}
