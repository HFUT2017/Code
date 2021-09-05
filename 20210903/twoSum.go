package _0210903

// TwoSum 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
func TwoSum(nums []int, target int) []int {
	hashtable := make(map[int]int, len(nums))
	for i, x := range nums {
		if p, ok := hashtable[target-x]; ok {
			return []int{p, i}
		}
		hashtable[x] = i
	}
	return nil
}
