package _0211004

func exist(board [][]byte, word string) bool {
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
		if dfs(i+1,j,k+1){
			return true
		}
		if dfs(i,j+1,k+1){
			return true
		}
		if dfs(i-1,j,k+1){
			return true
		}
		if dfs(i,j-1,k+1){
			return true
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
