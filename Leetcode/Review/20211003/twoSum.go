package _0211003

func twoSum(nums []int, target int) []int {
	hash:=map[int]int{}
	for i:=0;i<len(nums);i++{
		if _,ok:=hash[target-nums[i]];ok{
			return []int{hash[target-nums[i]],nums[i]}
		}
		hash[nums[i]]=nums[i]
	}
	return []int{}
}
