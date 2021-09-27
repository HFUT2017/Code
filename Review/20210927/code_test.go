package _0210927

import "testing"

func TestNumIslands(t *testing.T) {
	grid:=[][]byte{{'1','1','1','1','0'},{'1','1','0','1','0'},{'1','1','0','0','0'},{'0','0','0','0','0'}}
	println(NumIslands(grid))
}
