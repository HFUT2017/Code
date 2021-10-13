package main

import (
	"fmt"
)
func main() {
	n:=0
	fmt.Scan(&n)
	res:=[]int{}
	dp:=make([]bool,n)
	for i:=0;i<n;i++{
		dp[i]=true
	}
	for i:=2;i<n;i++{
		if dp[i]==false{
			continue
		}
		for j:=i+i;j<n;j+=i{
			dp[j]=false
		}
	}
	for i:=2;i<n;i++{
		if dp[i]==true{
			if i%10==1{
				res=append(res,i)
			}
		}
	}
	if len(res)==0{
		fmt.Printf("-1")
	}else{
		for i:=0;i<len(res);i++{
			fmt.Printf("%d",res[i])
			if i!=len(res)-1{
				fmt.Printf(" ")
			}
		}
	}
}


