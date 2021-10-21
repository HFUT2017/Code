package _0211021

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	resAddr := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			sum := int(num1[i]-'0')*int(num2[j]-'0') + resAddr[i+j+1]
			resAddr[i+j+1] = sum % 10
			resAddr[i+j] += sum / 10
		}
	}
	res := ""
	for i := 0; i < len(resAddr); i++ {
		if i == 0 && resAddr[0] == 0 {
			continue
		}
		res += string(resAddr[i] + '0')
	}
	return res
}
