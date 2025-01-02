package main

func imageSmoother(img [][]int) [][]int {
	rows := len(img)
	cols := len(img[0])
	res := make([][]int, rows)
	for i := 0; i < rows; i++ {
		res[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res[i][j] = imgSmooth(img, i, j)
		}
	}
	return res
}
func imgSmooth(img [][]int, x int, y int) int {
	rows := len(img)
	cols := len(img[0])
	sum := 0
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			nx := x + i
			ny := y + j
			if nx < 0 || nx >= rows || ny < 0 || ny >= cols {
				continue
			}
			sum += img[nx][ny]
			count++
		}
	}
	return sum / count
}
