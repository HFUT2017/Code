package BDA2

func FindKElement(nums []int,k int) int{
	buildHeap(nums)
	length:=len(nums)-1
	for k!=1{
		nums[0],nums[length]=nums[length],nums[0]
		length--
		adjustHeap(nums,0,length)
		k--
	}
	return nums[0]
}

func buildHeap(nums []int){
	for i:=len(nums)/2;i>=0;i--{
		adjustHeap(nums,i,len(nums)-1)
	}
}

func adjustHeap(nums []int,index ,length int) {
	left,right,maxIndex:=2*index,2*index+1,index
	if left<=length&&nums[left]>nums[maxIndex]{
		maxIndex=left
	}
	if right<=length&&nums[right]>nums[maxIndex]{
		maxIndex=right
	}
	if maxIndex!=index{
		nums[maxIndex],nums[index]=nums[index],nums[maxIndex]
		adjustHeap(nums,maxIndex,length)
	}
}
