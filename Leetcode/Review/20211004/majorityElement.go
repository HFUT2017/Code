package _0211004

func majorityElement(nums []int) int {
	cnt:=0
	res:=0
	for i:=0;i<len(nums);i++{
		if nums[i]==res{
			cnt++
		}else{
			if cnt==0{
				cnt++
				res=nums[i]
			}else{
				cnt--
			}
		}
	}
	return res
}
