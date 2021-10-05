package _0210919

import (
	"testing"
)

func TestDetectSquares(t *testing.T) {
	f := Constructor()
	nums := [][]int{{3, 10}, {11, 2}, {3, 2}, {11, 2}}
	for i := 0; i < len(nums); i++ {
		f.Add(nums[i])
	}
	//for i := 0; i < len(nums); i++ {
	//	temp := [2]int{}
	//	temp[0] = nums[i][0]
	//	temp[1] = nums[i][1]
	//	println(f.Num[temp])
	//}

	//xNums := f.Hashy[10]
	//yNums := f.Hash[11]
	//for i := 0; i < len(xNums); i++ {
	//	println(xNums[i])
	//}
	//println("------------------")
	//for i := 0; i < len(yNums); i++ {
	//	println(yNums[i])
	//}
	println(f.Count([]int{11,10}))
}
