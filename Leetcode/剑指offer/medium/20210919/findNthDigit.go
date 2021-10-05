package _0210919

import "strconv"

func FindNthDigit(n int) int {
	digit, start, count := 1, 1, 9
	for n > count {
		n -= count
		digit++
		start *= 10
		count = 9 * digit * start
	}
	num := start + (n-1)/digit
	return int(strconv.Itoa(num)[(n-1)%digit] - '0')
}
