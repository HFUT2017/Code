package intership

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for index, value := range nums {
		if _, ok := hash[target-value]; ok {
			return []int{index, hash[target-value]}
		} else {
			hash[value] = index
		}
	}
	return nil
}
