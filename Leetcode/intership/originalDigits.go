package intership

import "strconv"

func OriginalDigits(s string) string {
	count := map[uint8]int{}
	res := ""
	hash := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}
	for index, str := range hash {
		for {
			flag := true
			for i := 0; i < len(str); i++ {
				if count[str[i]] <= 0 {
					for j := i - 1; j >= 0; j-- {
						count[str[j]]++
					}
					flag = false
					break
				} else {
					count[str[i]]--
				}
			}
			if flag == true {
				res += strconv.Itoa(index)
			} else {
				break
			}
		}
	}
	return res
}
