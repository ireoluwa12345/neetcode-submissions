func isValidSudoku(board [][]byte) bool {
	var rows [9][256]bool
	var cols [9][256]bool
	var squares [9][256]bool

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			val := board[r][c]

			if val == '.' {
				continue
			}

			squareKey := (r/3)*3 + (c/3)

			if rows[r][val] || cols[c][val] || squares[squareKey][val] {
				return false
			}

			rows[r][val] = true
			cols[c][val] = true
			squares[squareKey][val] = true
		}
	}

	return true
}