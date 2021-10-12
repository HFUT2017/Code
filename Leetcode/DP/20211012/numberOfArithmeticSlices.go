package _0211012

func numberOfArithmeticSlices(nums []int) int {
	if len(nums)==1{
		return 0
	}
	d,t:=nums[0]-nums[1],0
	res:=0
	for i:=2;i<len(nums);i++{
		if nums[i-1]-nums[i]==d{
			t++
		}else{
			d,t=nums[i-1]-nums[i],0
		}
		res+=t
	}
	return res
}
