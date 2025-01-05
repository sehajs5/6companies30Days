package main

import "math"

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	n := 26

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < len(matrix[i]); j++ {
			if i == j {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = math.MaxInt / 2
			}
		}
	}

	for i, c := range cost {
		from := int(original[i] - 'a')
		to := int(changed[i] - 'a')
		if c < matrix[from][to] {
			matrix[from][to] = c
		}
	}

	floydWarshall2(matrix, n)
	total := int64(0)
	for i := 0; i < len(source); i++ {
		s := source[i] - 'a'
		t := target[i] - 'a'
		if target[i] != source[i] {
			if matrix[s][t] == math.MaxInt/2 {
				return -1
			}
			total += int64(matrix[s][t])

		}
	}
	return total
}

func floydWarshall2(matrix [][]int, n int) {
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
