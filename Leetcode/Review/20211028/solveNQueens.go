package _0211028

import "strings"

func solveNQueens(n int) [][]string {
	res := [][]string{}
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	check := func(row, col int) bool {
		for i := 0; i < n; i++ {
			if board[i][col] == "Q" {
				return false
			}
		}

		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == "Q" {
				return false
			}
		}

		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == "Q" {
				return false
			}
		}
		return true
	}
	var backTrack func(row int)
	backTrack = func(row int) {
		if row == n {
			temp := make([]string, n)
			for i := 0; i < n; i++ {
				temp[i] = strings.Join(board[i], "")
			}
			res = append(res, temp)
		}
		for col := 0; col < n; col++ {
			if check(row, col) == false {
				continue
			}
			board[row][col] = "Q"
			backTrack(row + 1)
			board[row][col] = "."
		}
	}
	backTrack(0)
	return res
}
