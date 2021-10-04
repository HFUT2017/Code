package _0211003

func movingCount(m int, n int, k int) int {
	res:=0
	used:=make([][]bool,m)
	for i:=0;i<m;i++{
		used[i]=make([]bool,n)
	}
	var dfs func(i,j int)
	dfs=func(i,j int){
		if i<0||j<0||i>=m||j>=n||used[i][j]==true||helper_(i)+helper_(j)>k{
			return
		}
		res+=1
		used[i][j]=true
		dfs(i+1,j)
		dfs(i-1,j)
		dfs(i,j+1)
		dfs(i,j-1)
	}
	dfs(0,0)
	return res
}

func helper_(n int) int{
	sum:=0
	for n!=0{
		sum+=n%10
		n/=10
	}
	return sum
}
