package _0211002

func exist(board [][]byte, word string) bool {
	dir:=[][]int{{0,1},{0,-1},{1,0},{-1,0}}
	used:=make([][]bool,len(board))
	for i:=0;i<len(board);i++{
		used[i]=make([]bool,len(board[i]))
	}
	var dfs func(i,j,k int) bool
	dfs=func(i,j,k int) bool{
		if i<0||j<0||i>=len(board)||j>=len(board[0])||used[i][j]==true||board[i][j]!=word[k]{
			return false
		}
		if k==len(word)-1{
			return true
		}
		used[i][j]=true
		defer func(){
			used[i][j]=false
		}()
		for _,value:=range dir{
			i_,j_:=i+value[0],j+value[1]
			if dfs(i_,j_,k+1){
				return true
			}
		}
		return false
	}
	for i:=0;i<len(board);i++{
		for j:=0;j<len(board[0]);j++{
			if dfs(i,j,0){
				return true
			}
		}
	}
	return false
}
