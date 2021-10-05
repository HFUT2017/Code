package _0211004

func search(nums []int, target int) int {
	var binarySearch func(left,right int) int
	binarySearch=func(left,right int) int{
		if left>right{
			return -1
		}
		mid:=(left+right)/2
		if nums[mid]==target{
			return mid
		}else if nums[mid]>target{
			return binarySearch(left,mid-1)
		}else{
			return binarySearch(mid+1,right)
		}
	}
	mid:=binarySearch(0,len(nums)-1)
	if mid==-1{
		return 0
	}
	left,right:=mid,mid
	for left>=0&&nums[left]==target{
		left--
	}
	for right<len(nums)&&nums[right]==target{
		right++
	}
	return right-left-1
}
