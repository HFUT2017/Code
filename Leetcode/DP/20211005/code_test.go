package _0211005

import "testing"

func TestGenerate(t *testing.T) {
	Generate(5)
}

func TestMinimumTotal(t *testing.T) {
	traiangle:=[][]int{{2},{3,4},{6,5,7},{4,1,8,3}}
	println(MinimumTotal(traiangle))
}

func TestMinFallingPathSum(t *testing.T) {
	traiangle:=[][]int{{17,82},{1,-44}}
	MinFallingPathSum(traiangle)
}

func TestMinFallingPathSum2(t *testing.T) {
	traiangle:=[][]int{{1,2,3},{4,5,6},{7,8,9}}
	println(MinFallingPathSum2(traiangle))
}

func TestPathsWithMaxScore(t *testing.T) {
	str:=[]string{"EX","XS"}
	PathsWithMaxScore(str)
}