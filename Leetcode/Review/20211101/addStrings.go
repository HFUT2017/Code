package _0211101

import "strconv"

func addStrings(num1 string, num2 string) string {
	l1, l2 := len(num1)-1, len(num2)-1
	cnt := 0
	ans := ""
	res := ""
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
		ans += strconv.Itoa(sum)
	}
	for i := len(ans) - 1; i >= 0; i-- {
		res += string(ans[i])
	}
	return res
}
