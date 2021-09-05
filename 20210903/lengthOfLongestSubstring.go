package _0210903

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// LengthOfLongestSubstring 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
func LengthOfLongestSubstring(s string) int {
	indexMap := make(map[rune]int)
	start, current, max := 0, 0, 0
	for index, value := range s {
		if i, ok := indexMap[value]; ok {
			if i >= start {
				if current-start > max {
					max = current - start
				}
				start = i + 1
			}
		}
		indexMap[value] = index
		current++
	}
	if current-start > max {
		max = current - start
	}
	return max
}
