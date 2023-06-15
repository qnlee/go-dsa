package lc73_set_matrix_zeroes

/*
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
You must do it in place.

Ex1) Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
	 Output: [[1,0,1],[0,0,0],[1,0,1]]

Ex2) Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
	 Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

Constraints:
	m == matrix.length
	n == matrix[0].length
	1 <= m, n <= 200
	-231 <= matrix[i][j] <= 231 - 1
*/

func setZeroes(matrix [][]int) {
	// brute force- for i:rows, j:cols {if matrix[i][j]==0 then set all of r/c to 0}
	// optimize performance with memoization: have 2 maps, 1 for rows 1 for cols and mark which are 0s
	// optimize space: use one of the rows/cols to keep track of ^
	m := len(matrix)
	n := len(matrix[0])

	var rowIsZero = make(map[int]bool, m)
	var colIsZero = make(map[int]bool, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rowIsZero[i] = true
				colIsZero[j] = true
			}
		}
	}

	for i, isZero := range rowIsZero {
		if isZero {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j, isZero := range colIsZero {
		if isZero {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}

	return
}
