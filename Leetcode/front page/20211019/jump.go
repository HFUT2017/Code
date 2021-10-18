package _0211019

func jump(nums []int) int {
	end:=0
	maxPosition:=0
	res:=0
	for i:=0;i<len(nums)-1;i++{
		if i+nums[i]>maxPosition{
			maxPosition=i+nums[i]
		}
		if i==end{
			end=maxPosition
			res++
		}
	}
	return res
}
