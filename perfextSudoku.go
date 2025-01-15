package main

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		rows := make(map[byte]bool)
		cols := make(map[byte]bool)
		boxs := make(map[byte]bool)

		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if _, ok := rows[board[i][j]]; ok {
					return false
				}
				rows[board[i][j]] = true
			}
			if board[j][i] != '.' {
				if _, ok := cols[board[j][i]]; ok {
					return false
				}
				cols[board[j][i]] = true
			}
			m := (i/3)*3 + j/3
			n := (i%3)*3 + j%3
			if board[m][n] != '.' {
				if _, ok := boxs[board[m][n]]; ok {
					return false
				}
				boxs[board[m][n]] = true
			}
		}
	}
	return true
}
