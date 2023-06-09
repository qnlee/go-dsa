package matrix_test

import (
	"go-dsa/leetcode/lc_medium/matrix"
	"testing"
)

var (
	validPartial      [][]byte
	invalidRowPartial [][]byte
	invalidColPartial [][]byte
	invalidSqrPartial [][]byte
	validFilled       [][]byte
	invalidFilled     [][]byte
)

type testCase struct {
	board    [][]byte
	expected bool
	errCase  string
}

func init() {
	// Valid Sudoku board (partially filled)
	validPartial = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	// Invalid Sudoku board (partially filled) with duplicate number in row
	invalidRowPartial = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '3', '.', '.', '2', '.', '.', '.', '6'}, // Duplicate '3' in row
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	// Invalid Sudoku board (partially filled) with duplicate number in column
	invalidColPartial = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'5', '.', '.', '.', '8', '.', '.', '7', '9'}, // Duplicate '5' in column
	}

	// Invalid Sudoku board (partially filled) with duplicate number in square
	invalidSqrPartial = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '6'}, // Duplicate '6' in square
	}

	// Invalid Sudoku board (complete)
	invalidFilled = [][]byte{
		{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
		{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
		{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
		{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
		{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
		{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
		{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
		{'2', '3', '7', '4', '1', '9', '6', '3', '5'},
		{'3', '4', '5', '2', '8', '6', '1', '7', '9'}, // Duplicate '3' at board[7][1]
	}

	// Valid Sudoku board (complete)
	validFilled = [][]byte{
		{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
		{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
		{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
		{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
		{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
		{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
		{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
		{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
		{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
	}
}

func TestIsValidSudoku_FilledBoards(t *testing.T) {
	testCases := []testCase{
		{validFilled, true, "valid"},
		{invalidFilled, false, "invalid"},
	}
	for _, tc := range testCases {
		actual := matrix.IsValidSudoku(tc.board)
		if actual != tc.expected {
			t.Errorf("Expected %v on [%s], but got %v", tc.expected, tc.errCase, actual)
		}
	}
}

func TestIsValidSudoku_PartiallyFilledBoards(t *testing.T) {
	testCases := []testCase{
		{validPartial, true, "valid"},
		{invalidColPartial, false, "invalid (column)"},
		{invalidRowPartial, false, "invalid (row)"},
		{invalidSqrPartial, false, "invalid (square)"},
	}
	for _, tc := range testCases {
		actual := matrix.IsValidSudoku(tc.board)
		if actual != tc.expected {
			t.Errorf("Expected %v on [%s], but got %v", tc.expected, tc.errCase, actual)
		}
	}
}

func TestIsValidSudokuAlt_PartiallyFilledBoards(t *testing.T) {
	testCases := []testCase{
		{validPartial, true, "valid"},
		{invalidColPartial, false, "invalid (column)"},
		{invalidRowPartial, false, "invalid (row)"},
		{invalidSqrPartial, false, "invalid (square)"},
	}
	for _, tc := range testCases {
		actual := matrix.IsValidSudokuAlt(tc.board)
		if actual != tc.expected {
			t.Errorf("Expected %v on [%s], but got %v", tc.expected, tc.errCase, actual)
		}
	}
}

func TestIsValidSudokuAlt_FilledBoards(t *testing.T) {
	testCases := []testCase{
		{validFilled, true, "valid"},
		{invalidFilled, false, "invalid"},
	}
	for _, tc := range testCases {
		actual := matrix.IsValidSudoku(tc.board)
		if actual != tc.expected {
			t.Errorf("Expected %v on [%s], but got %v", tc.expected, tc.errCase, actual)
		}
	}
}

func BenchmarkIsValidSudoku(b *testing.B) {
	for n := 0; n < b.N; n++ {
		matrix.IsValidSudoku(validFilled)
	}
}

func BenchmarkIsValidSudokuAlt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		matrix.IsValidSudokuAlt(validFilled)
	}
}
