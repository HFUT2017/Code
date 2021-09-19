package _0210919

func FinalValueAfterOperations(operations []string) int {
	x := 0
	str1 := "++X"
	str2 := "X++"
	str3 := "--X"
	str4 := "X--"
	for _, value := range operations {
		if value == str1 || value == str2 {
			x++
		} else if value == str3 || value == str4 {
			x--
		}
	}
	return x
}
