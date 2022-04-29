package main

import (
	"fmt"
)

//https://leetcode.com/problems/valid-sudoku/solution/
func main() {
	sudoku := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println(isValidSudoku(sudoku))
}

func isValidSudoku(board [][]byte) bool {
	rowMap := make(map[int]map[string]bool, len(board))
	columnMap := make(map[int]map[string]bool, len(board))
	boxMap := make(map[int]map[string]bool, len(board))
	for i := 0; i < 9; i++ {
		rowMap[i] = make(map[string]bool, len(board))
		columnMap[i] = make(map[string]bool, len(board))
		boxMap[i] = make(map[string]bool, len(board))
	}
	for row := 0; row < len(board); row++ {
		for column := 0; column < len(board); column++ {
			val := board[row][column]
			if string(val) == "." {
				continue
			}
			//check row
			_, okRow := rowMap[row][string(val)]
			if okRow {
				return false
			}
			rowMap[row][string(val)] = true
			//check column
			_, okCol := columnMap[column][string(val)]
			if okCol {
				return false
			}
			columnMap[column][string(val)] = true

			//check boxes
			index := (row/3)*3 + column/3
			_, okBox := boxMap[index][string(val)]
			if okBox {
				return false
			}
			boxMap[index][string(val)] = true
		}
	}
	return true
}
