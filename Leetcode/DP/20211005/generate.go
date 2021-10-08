package _0211005

// Generate 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
func Generate(numRows int) [][]int {
	dp := [][]int{}
	for i := 0; i < numRows; i++ {
		temp := []int{}
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				temp = append(temp, 1)
			} else {
				temp = append(temp, (dp[i-1][j-1] + dp[i-1][j]))
			}
		}
		dp = append(dp, temp)
	}
	return dp
}
