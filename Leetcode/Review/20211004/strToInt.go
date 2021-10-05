package _0211004

import (
	"math"
	"strings"
)

func strToInt(str string) int {
	str = strings.TrimSpace(str)
	res:=0
	sign:=1
	for i:=0;i<len(str);i++{
		if str[i]>='0'&&str[i]<='9'{
			res=res*10+int(str[i]-'0')
		}else if str[i]=='-'&&i==0{
			sign=-1
		}else if str[i]=='+'&&i==0{
			sign=1
		}else{
			break
		}
		if res>math.MaxInt32{
			if sign==-1{
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	return sign*res
}
