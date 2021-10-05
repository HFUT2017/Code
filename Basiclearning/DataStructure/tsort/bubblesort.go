package tsort

func BubbleSort(nums []int) {
	var exchanged bool
	for i := len(nums) - 1; i >= 0; i-- {
		exchanged = false
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				exchanged = true
			}
		}
		if exchanged == false {
			return
		}
	}
}
