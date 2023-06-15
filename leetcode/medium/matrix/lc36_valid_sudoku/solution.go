package lc36_valid_sudoku

/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

	Each row must contain the digits 1-9 without repetition.
	Each column must contain the digits 1-9 without repetition.
	Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

Note: A Sudoku board (partially filled) could be valid but is not necessarily solvable. Only the filled cells need to
be validated according to the mentioned rules.

Example 1:
Input: board =  [["5","3",".",".","7",".",".",".","."]

	,["6",".",".","1","9","5",".",".","."]
	,[".","9","8",".",".",".",".","6","."]
	,["8",".",".",".","6",".",".",".","3"]
	,["4",".",".","8",".","3",".",".","1"]
	,["7",".",".",".","2",".",".",".","6"]
	,[".","6",".",".",".",".","2","8","."]
	,[".",".",".","4","1","9",".",".","5"]
	,[".",".",".",".","8",".",".","7","9"]]

Output: true

Constraints:

	board.length == 9
	board[i].length == 9
	board[i][j] is a digit 1-9 or '.'.
*/

func isValidSudoku(board [][]byte) bool {
	rows := make([][]bool, 9)
	cols := make([][]bool, 9)
	sqrs := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make([]bool, 9)
		cols[i] = make([]bool, 9)
		sqrs[i] = make([]bool, 9)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			val := int(board[r][c] - '1') // subtract '1' to keep val in index range [0,8]
			// check if val already in row
			if rows[r][val] == true {
				return false
			}
			rows[r][val] = true
			// check if val already in col
			if cols[c][val] == true {
				return false
			}
			cols[c][val] = true
			// check if val already in square
			sq := r/3*3 + c/3 // index of sub-square, starting top left (sq=0) -> bottom right (sq=8)
			if sqrs[sq][val] == true {
				return false
			}
			sqrs[sq][val] = true
		}
	}

	return true
}

// IsValidSudokuAlt - this solution reduces space slightly but is significantly less performant
func isValidSudokuAlt(board [][]byte) bool {
	counts := make([]map[byte]bool, 27)
	for i := 0; i < 27; i++ {
		counts[i] = make(map[byte]bool)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			val := board[r][c]
			rowIdx := r
			colIdx := 9 + c
			sqrIdx := 18 + r/3*3 + c/3
			if counts[rowIdx][val] || counts[colIdx][val] || counts[sqrIdx][val] {
				return false
			}
			counts[rowIdx][val] = true
			counts[colIdx][val] = true
			counts[sqrIdx][val] = true
		}
	}

	return true
}
