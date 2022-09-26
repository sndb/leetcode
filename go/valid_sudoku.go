package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		seen := map[byte]struct{}{}
		for j := 0; j < 9; j++ {
			if _, ok := seen[board[i][j]]; ok {
				return false
			}
			if board[i][j] != '.' {
				seen[board[i][j]] = struct{}{}
			}
		}
	}
	for i := 0; i < 9; i++ {
		seen := map[byte]struct{}{}
		for j := 0; j < 9; j++ {
			if _, ok := seen[board[j][i]]; ok {
				return false
			}
			if board[j][i] != '.' {
				seen[board[j][i]] = struct{}{}
			}
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			seen := map[byte]struct{}{}
			for k := i * 3; k < i*3+3; k++ {
				for l := j * 3; l < j*3+3; l++ {
					if _, ok := seen[board[k][l]]; ok {
						return false
					}
					if board[k][l] != '.' {
						seen[board[k][l]] = struct{}{}
					}
				}
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
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
	fmt.Println(isValidSudoku(board))
}
