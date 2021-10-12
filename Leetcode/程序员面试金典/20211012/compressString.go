package _0211012

import (
	"strconv"
	"strings"
)

func compressString(S string) string {
	if S == "" {
		return ""
	}
	var sb strings.Builder
	cur := S[0]
	ant := 0
	for i := 0; i < len(S); i++ {
		if cur == S[i] {
			ant++
		} else {
			sb.WriteByte(cur)
			sb.WriteString(strconv.Itoa(ant))
			cur = S[i]
			ant = 1
		}
	}
	sb.WriteByte(cur)
	sb.WriteString(strconv.Itoa(ant))
	if sb.Len() >= len(S) {
		return S
	}
	return sb.String()
}
