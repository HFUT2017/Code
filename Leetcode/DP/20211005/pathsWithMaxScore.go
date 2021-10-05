package _0211005

import "math"

// PathsWithMaxScore 1301. 最大得分的路径数目
func PathsWithMaxScore(board []string) []int {
	dp:=make([][]int,len(board))
	for i:=0;i<len(dp);i++{
		dp[i]=make([]int,len(board[i]))
	}
	ans:=1
	for i:=len(board)-1;i>=0;i--{
		for j:=len(board[i])-1;j>=0;j--{
			if board[i][j]=='X'{
				dp[i][j]=math.MinInt32
			}else if board[i][j]=='E'{
				ans*=getSameNum(dp[i+1][j],dp[i][j+1],dp[i+1][j+1])
				dp[i][j]=max(dp[i+1][j],max(dp[i][j+1],dp[i+1][j+1]))
			}else if board[i][j]=='S'{
				dp[i][j]=0
			}else if i==len(board)-1{
				if j+1<len(dp[i]){
					dp[i][j]=int(board[i][j]-'0')+dp[i][j+1]
				}else{
					dp[i][j]=int(board[i][j]-'0')
				}
			}else{
				if j+1<len(dp[i]){
					ans*=getSameNum(dp[i+1][j],dp[i][j+1],dp[i+1][j+1])
					dp[i][j]=max(dp[i+1][j],max(dp[i][j+1],dp[i+1][j+1]))+int(board[i][j]-'0')
				}else{
					dp[i][j]=dp[i+1][j]+int(board[i][j]-'0')
				}
			}
		}
	}
	if dp[0][0]<0{
		return []int{0,0}
	}else if dp[0][0]==0{
		return []int{0,1}
	}
	//return []int{dp[0][0],ans}
}

func getSameNum(a,b,c int) int{
	res:=0
	if a==b{
		res++
	}
	if a==c{
		res++
	}
	if b==c{
		res++
	}
	if res==3{
		return res
	}else if res==0{
		return 1
	}else{
		return res+1
	}
}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}
