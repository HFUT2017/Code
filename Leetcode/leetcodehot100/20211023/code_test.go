package _0211023

import (
	"strconv"
	"testing"
)

func TestString2Num(t *testing.T) {
	test := "001"
	res, _ := strconv.Atoi(test)
	println(res)
}
