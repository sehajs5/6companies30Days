package main

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	queue := [][2]int{}
	fresh := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
				grid[i][j] = 3
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}
	count := 0

	if fresh == 0 {
		return 0
	}
	if len(queue) == 0 {
		return -1
	}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			temp := queue[0]
			queue = queue[1:]
			row := temp[0]
			col := temp[1]
			if row+1 < rows && grid[row+1][col] == 1 {
				queue = append(queue, [2]int{row + 1, col})
				grid[row+1][col] = 3
			}

			if row-1 >= 0 && grid[row-1][col] == 1 {
				queue = append(queue, [2]int{row - 1, col})
				grid[row-1][col] = 3
			}

			if col+1 < cols && grid[row][col+1] == 1 {
				queue = append(queue, [2]int{row, col + 1})
				grid[row][col+1] = 3
			}

			if col-1 >= 0 && grid[row][col-1] == 1 {
				queue = append(queue, [2]int{row, col - 1})
				grid[row][col-1] = 3
			}
		}
		count++
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return count - 1
}
