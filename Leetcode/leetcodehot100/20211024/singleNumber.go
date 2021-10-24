package _0211024

func singleNumber(nums []int) int {
	a := 0
	for _, v := range nums {
		a ^= v
	}
	return a
}
