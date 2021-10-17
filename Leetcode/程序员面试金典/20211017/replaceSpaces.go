package _0211017

import "strings"

func replaceSpaces(s string, length int) string {
	var res strings.Builder
	for i:=0;i<length;i++{
		if s[i]!=' '{
			res.WriteByte(s[i])
		}else{
			res.WriteString("%20")
		}
	}
	return res.String()
}
