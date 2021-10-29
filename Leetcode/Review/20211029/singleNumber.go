package _0211029

func singleNumber(nums []int) []int {
	a, b := 0, 0
	cnt := 0
	for _, value := range nums {
		cnt ^= value
	}
	p := cnt & -cnt
	for _, value := range nums {
		if value&p > 0 {
			a ^= value
		} else {
			b ^= value
		}
	}
	return []int{a, b}
}
