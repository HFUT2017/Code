package _0210918

import "testing"

func TestExist(t *testing.T) {
	board := [][]byte{{'A', 'B', 'C', 'D'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	println(Exist(board, word))
}
