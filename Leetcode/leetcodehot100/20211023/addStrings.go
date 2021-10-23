package _0211023

import "strconv"

func addStrings(num1 string, num2 string) string {
	res := ""
	l1, l2 := len(num1)-1, len(num2)-1
	cnt := 0
	for l1 >= 0 || l2 >= 0 {
		n1, n2 := 0, 0
		if l1 >= 0 {
			n1 = int(num1[l1] - '0')
			l1--
		}
		if l2 >= 0 {
			n2 = int(num2[l2] - '0')
			l2--
		}
		sum := (n1 + n2 + cnt) % 10
		cnt = (n1 + n2 + cnt) / 10
		res += strconv.Itoa(sum)
	}
	if cnt != 0 {
		res += strconv.Itoa(cnt)
	}
	ans := ""
	for i := len(res) - 1; i >= 0; i-- {
		ans += string(res[i])
	}
	return ans
}
