package _0211003

import "strconv"

func translateNum(num int) int {
	s:=strconv.Itoa(num)
	dp:=make([]int,len(s))
	for i:=0;i<len(s);i++{
		if i==0{
			dp[i]=1
			continue
		}
		if int(s[i-1]-'0')!=0&&int(s[i-1]-'0')*10+int(s[i]-'0')<=25{
			if i>=2{
				dp[i]=dp[i-1]+dp[i-2]
			}else{
				dp[i]=2
			}
		}else{
			dp[i]=dp[i-1]
		}
	}
	return dp[len(s)-1]
}
