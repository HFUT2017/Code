package _0211023

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}
func rotate(nums []int, k int) {
	k=k%len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
