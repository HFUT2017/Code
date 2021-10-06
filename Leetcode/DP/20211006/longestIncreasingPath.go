package _0211006

//func longestIncreasingPath(matrix [][]int) int {
//	res:=0
//	var dfs func(i,j int,temp int,length int)
//	dir:=[][]int{{-1,0},{1,0},{0,-1},{0,1}}
//	used:=make([][]bool,len(matrix))
//	for i:=0;i<len(used);i++{
//		used[i]=make([]bool,len(matrix[i]))
//	}
//	dfs=func(i,j int,temp int,length int){
//		if i<0||j<0||i>=len(matrix)||j>=len(matrix[0])||used[i][j]==true||matrix[i][j]<=temp{
//			return
//		}
//		used[i][j]=true
//		length=length+1
//		if length>res{
//			res=length
//		}
//		for t:=0;t<len(dir);t++{
//			dfs(i+dir[t][0],j+dir[t][1],matrix[i][j],length)
//		}
//		used[i][j]=false
//		length=length-1
//	}
//	for i:=0;i<len(matrix);i++{
//		for j:=0;j<len(matrix[i]);j++{
//			dfs(i,j,-1,0)
//		}
//	}
//	return res
//}

//正确解法:记忆化回溯
func longestIncreasingPath(matrix [][]int) int {
	dir:=[][]int{{-1,0},{1,0},{0,-1},{0,1}}
	if len(matrix)==0||len(matrix[0])==0{
		return 0
	}
	mem:=make([][]int,len(matrix))
	var dfs func(matrix [][]int,i,j int,mem [][]int) int
	dfs=func(matrix [][]int,i,j int,mem [][]int) int{
		if mem[i][j]!=0{
			return mem[i][j]
		}
		mem[i][j]++
		for _,value:=range dir{
			newI,newJ:=i+value[0],j+value[1]
			if newI>=0&&newJ>=0&&newI<len(matrix)&&newJ<len(matrix[0])&&matrix[newI][newJ]>matrix[i][j]{
				mem[i][j]=max(mem[i][j],dfs(matrix,newI,newJ,mem)+1)
			}
		}
		return mem[i][j]
	}
	for i:=0;i<len(mem);i++{
		mem[i]=make([]int,len(matrix[i]))
	}
	res:=0
	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[i]);j++{
			res=max(res,dfs(matrix,i,j,mem))
		}
	}
	return res
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}
