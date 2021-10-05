package _0211003

func getLeastNumbers(nums []int, k int) []int {
	res:=[]int{}
	buildHeap(nums)
	length:=len(nums)-1
	for k!=0{
		res=append(res,nums[0])
		nums[0],nums[length]=nums[length],nums[0]
		length--
		k--
		adjust(nums,0,length)
	}
	return res
}

func buildHeap(nums []int){
	for i:=len(nums)/2;i>=0;i--{
		adjust(nums,i,len(nums)-1)
	}
}

func adjust(nums []int,index,length int){
	left,right,minIndex:=2*index,2*index+1,index
	if left<=length &&nums[left]<nums[minIndex]{
		minIndex=left
	}
	if right<=length &&nums[right]<nums[minIndex]{
		minIndex=right
	}
	if minIndex!=index{
		nums[index],nums[minIndex]=nums[minIndex],nums[index]
		adjust(nums,minIndex,length)
	}
}
