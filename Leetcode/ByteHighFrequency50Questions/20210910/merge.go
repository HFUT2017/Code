package _0210910

import "sort"

// Merge 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间
func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0]<intervals[j][0]
	})

	res:=[][]int{}
	prev:=intervals[0]

	for i:=1;i<len(intervals);i++{
		cur :=intervals[i]
		if prev[1]<cur[0]{
			res=append(res,prev)
			prev=cur
		}else {
			prev[1]=max(prev[1],cur[1])
		}
	}
	res=append(res,prev)
	return res
}
func max(a, b int) int {
	if a > b { return a }
	return b
}
