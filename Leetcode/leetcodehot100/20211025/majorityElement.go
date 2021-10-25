package _0211025

func majorityElement(nums []int) int {
	count, num := 0, 0
	for _, value := range nums {
		if count == 0 {
			num = value
		}
		if value == num {
			count++
		} else {
			count--
		}
	}
	return num
}
