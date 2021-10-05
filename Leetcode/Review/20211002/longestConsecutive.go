package _0211002

func longestConsecutive(nums []int) int {
	hash:=map[int]bool{}
	for i:=0;i<len(nums);i++{
		hash[nums[i]]=true
	}
	res:=0
	for i:=0;i<len(nums);i++{
		if !hash[nums[i]-1]{
			currentNum:=nums[i]
			currentLength:=1
			for{
				if hash[currentNum+1]{
					currentNum=currentNum+1
					currentLength=currentLength+1
				}else{
					break
				}
			}
			if currentLength>res{
				res=currentLength
			}
		}
	}
	return res
}
