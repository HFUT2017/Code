package _0210916

func TranslateNum(num int) int {
	pre1, pre2 := 1, 1
	preNum := num % 10
	num /= 10
	for num != 0 {
		temp := num % 10
		if temp*10+preNum <= 25 && temp != 0 {
			pre1, pre2 = pre1+pre2, pre1
		} else {
			pre1, pre2 = pre1, pre1
		}
		preNum = temp
		num /= 10
	}
	return pre1

}
