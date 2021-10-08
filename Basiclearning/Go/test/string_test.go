package test

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	str := "123456字符串测试"
	println("字符串byte形式:")
	for i := 0; i < len(str); i++ { //byte ascii
		fmt.Printf("%v(%c)", str[i], str[i])
	}
	println()
	println("字符串rune形式:")
	for _, value := range str { //rune int32 unicode
		fmt.Printf("%v(%c)", value, value)
	}
	println()
}
