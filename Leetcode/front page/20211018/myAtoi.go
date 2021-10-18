package _0211018

import "math"

func myAtoi(s string) int {
	begin := 0
	flag := 1
	sum := 0
	for begin < len(s) && s[begin] == ' ' {
		begin++
	}
	if begin == len(s) {
		return 0
	}
	if s[begin] == '-' {
		flag = -1
		begin++
	} else if s[begin] == '+' {
		begin++
	}
	for i := begin; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if sum > math.MaxInt32/10 {
				if flag == -1 {
					return math.MinInt32
				} else {
					return math.MaxInt32
				}
			}
			sum = sum*10 + int(s[i]-'0')
		} else {
			break
		}
	}
	if sum*flag > math.MaxInt32 {
		return math.MaxInt32
	} else if sum*flag < math.MinInt32 {
		return math.MinInt32
	}
	return sum * flag

}
