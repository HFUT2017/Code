package _0211019

func searchRange(nums []int, target int) []int {
	mid:=binarySearch(nums,target)
	if mid==-1{
		return []int{-1,-1}
	}
	left,right:=mid,mid
	for left>=0&&nums[left]==nums[mid]{
		left--
	}
	for right<len(nums)&&nums[right]==nums[mid]{
		right++
	}
	return []int{left+1,right-1}
}

func binarySearch(nums []int,target int) int{
	left,right:=0,len(nums)-1
	for left<=right{
		mid:=(left+right)/2
		if nums[mid]==target{
			return mid
		}else if nums[mid]>target{
			right=mid-1
		}else{
			left=mid+1
		}
	}
	return -1
}
