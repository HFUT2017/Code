package _0211024

func isPalindrome(head *ListNode) bool {
	nums := []int{}
	cur := head
	for cur != nil {
		nums = append(nums, cur.Val)
		cur = cur.Next
	}
	return judge(nums)
}

func judge(nums []int) bool {
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-i-1] {
			return false
		}
	}
	return true
}
