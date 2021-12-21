package amazon

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for index, value := range nums {
		if index2, ok := hash[target-value]; ok {
			return []int{index2, index}
		}
		hash[value] = index
	}
	return []int{-1, -1}
}
