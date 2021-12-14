package Designandanalysisofalgorithms

func multiply(x, y string) string {
	if x == "0" || y == "0" {
		return "0"
	}
	resArr := make([]int, len(x)+len(y))
	res := ""
	for i := len(x) - 1; i >= 0; i-- {
		for j := len(y) - 1; j >= 0; j-- {
			resArr[i+j+1] += int(x[i]-'0') * int(y[j]-'0')
			resArr[i+j] += resArr[i+j+1] / 10
			resArr[i+j+1] = resArr[i+j+1] % 10
		}
	}
	for index, value := range resArr {
		if index == 0 && value == 0 {
			continue
		}
		res += string(value + '0')
	}
	return res
}
