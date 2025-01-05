package main

import "math"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = math.MaxInt / 2
			}
		}
	}

	for _, edge := range edges {
		matrix[edge[0]][edge[1]] = edge[2]
		matrix[edge[1]][edge[0]] = edge[2]
	}

	floydWarshall(matrix, n)
	ans := 0
	prev := math.MaxInt
	for i := n - 1; i >= 0; i-- {
		samp := 0
		for j := 0; j < n; j++ {
			if matrix[i][j] <= distanceThreshold && i != j {
				samp++
			}
		}
		if prev > samp {
			ans = i
			prev = samp
		}
	}
	return ans
}

func floydWarshall(matrix [][]int, n int) {
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if matrix[i][k] < math.MaxInt && matrix[k][j] < math.MaxInt {
					matrix[i][j] = min(matrix[i][j], matrix[k][j]+matrix[i][k])
				}
			}
		}
	}
}
