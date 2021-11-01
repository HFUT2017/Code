package _0211101

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for index, value := range v1 {
		if index == len(v2) {
			for i := index; i < len(v1); i++ {
				num, _ := strconv.Atoi(v1[i])
				if num != 0 {
					return 1
				}
			}
			return 0
		}
		num1, _ := strconv.Atoi(value)
		num2, _ := strconv.Atoi(v2[index])
		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}
	}
	for i := len(v1); i < len(v2); i++ {
		num, _ := strconv.Atoi(v2[i])
		if num != 0 {
			return -1
		}
	}
	return 0
}
