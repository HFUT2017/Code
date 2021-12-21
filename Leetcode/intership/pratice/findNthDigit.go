package intership

import (
	"math"
	"strconv"
)

//func findNthDigit(n int) int {
//	digit := 1
//	start := 1
//	count := digit * 9 * start
//	for n > count {
//		n -= count
//		digit++
//		start *= 10
//		count = digit * 9 * start
//	}
//	num := start + (n - 1) / digit
//	remain := (n - 1) % digit
//	str := strconv.Itoa(num)
//	return int(str[remain] - '0')
//}

func getDNum(cnt int) int {
	c := float64(cnt)
	d := float64(10)
	if cnt == 1 {
		return 9
	}
	return int(math.Pow(d, c) - math.Pow(d, c-1))
}

func FindNthDigit(n int) int {
	cnt := 1
	num := cnt * getDNum(cnt)
	for n-num > 0 {
		n -= num
		cnt += 1
		num = cnt * getDNum(cnt)
	}
	begin := int(math.Pow(float64(10), float64(cnt))) - getDNum(cnt)
	if n%cnt == 0 {
		return int(strconv.Itoa(begin + n/cnt - 1)[cnt-1] - '0')
	} else {
		return int(strconv.Itoa(begin + n/cnt)[n%cnt-1] - '0')
	}
}
