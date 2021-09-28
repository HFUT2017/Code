package _0210928

func longestConsecutive(nums []int) int {
	hash := make(map[int]bool, len(nums))
	res := 0
	for i := 0; i < len(nums); i++ {
		hash[nums[i]] = true
	}
	for i := 0; i < len(nums); i++ {
		if !hash[nums[i]-1] {
			currentNum := nums[i]
			currentLength := 1
			for {
				if hash[currentNum+1] {
					currentNum = currentNum + 1
					currentLength = currentLength + 1
				} else {
					break
				}
			}
			if currentLength > res {
				res = currentLength
			}

		}
	}
	return res
}
