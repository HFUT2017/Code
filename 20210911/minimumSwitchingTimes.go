package _0210911

// MinimumSwitchingTimes 给定两个大小均为 N*M 的二维数组 source 和 target 表示无人机方阵表演的两种颜色图案，由于无人机切换灯光颜色的耗能很大，请返回从 source 到 target 最少需要多少架无人机切换灯光颜色。
func MinimumSwitchingTimes(source [][]int, target [][]int) int {
	hashtable := make(map[int]int, 0)
	res := 0
	for i := 0; i < len(source); i++ {
		for j := 0; j < len(source[i]); j++ {
			if _, ok := hashtable[source[i][j]]; ok {
				hashtable[source[i][j]]++
			} else {
				hashtable[source[i][j]] = 1
			}
		}
	}
	for i := 0; i < len(target); i++ {
		for j := 0; j < len(target[i]); j++ {
			if _, ok := hashtable[target[i][j]]; ok {
				if(hashtable[target[i][j]]!=0){
					hashtable[target[i][j]]--
				}else{
					res++
				}
			} else {
				res++
			}
		}
	}
	return res
}
