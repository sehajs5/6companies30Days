package main

import "sort"

func wiggleSort(arr []int) {
	n := len(arr)
	ans := make([]int, n)
	sort.Ints(arr)
	i := 1
	j := n - 1
	for i < n {
		ans[i] = arr[j]
		i += 2
		j--
	}
	i = 0
	for i < n {
		ans[i] = arr[j]
		i += 2
		j--
	}
	for k := 0; k < n; k++ {
		arr[k] = ans[k]
	}
	return
}
