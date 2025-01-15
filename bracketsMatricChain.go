package main

import (
	"math"
)

func matrixChain(arr []int) string {
	n := len(arr)
	dp := make([][]int, n)
	way := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		way[i] = make([]int, n)
	}

	solveHelp(arr, dp, 1, n-1, way)

	result := []byte{}

	build(1, n-1, &result, way)
	return string(result)
}

func solveHelp(arr []int, dp [][]int, i int, j int, way [][]int) int {
	if i == j {
		return 0
	}

	if dp[i][j] != 0 {
		return dp[i][j]
	}
	res := math.MaxInt
	for k := i; k < j; k++ {
		a := solveHelp(arr, dp, i, k, way) + solveHelp(arr, dp, k+1, j, way) + (arr[i-1] * arr[k] * arr[j])

		if res > a {
			res = a
			way[i][j] = k
		}
	}
	dp[i][j] = res
	return res
}

func build(i int, j int, result *[]byte, way [][]int) {
	if i == j {
		*result = append(*result, 'A'+byte(i-1))
		return
	}

	*result = append(*result, '(')
	build(i, way[i][j], result, way)
	build(way[i][j]+1, j, result, way)
	*result = append(*result, ')')
}
