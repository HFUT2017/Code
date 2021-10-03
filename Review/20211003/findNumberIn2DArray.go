package _0211003

func findNumberIn2DArray(matrix [][]int, target int) bool {
	i:=len(matrix)-1
	if i<0{
		return false
	}
	j:=0
	for i>=0&&j<len(matrix[0]){
		if matrix[i][j]==target{
			return true
		}else if matrix[i][j]<target{
			j++
		}else{
			i--
		}
	}
	return false
}
