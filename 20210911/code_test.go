package _0210911

import "testing"

func TestMinimumSwitchingTimes(t *testing.T) {
	source:=[][]int{{1,2,3},{3,4,5}}
	target:=[][]int{{1,3,5},{2,3,4}}
	println(MinimumSwitchingTimes(source,target))
}

//func TestMaxmiumScore(t *testing.T) {
//	cards:=[]int{1,2,8,9}
//	cnt:=3
//	println(MaxmiumScore(cards,cnt))
//}