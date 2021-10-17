package _0211017

// 回溯
//func findLadders(beginWord string, endWord string, wordList []string) []string {
//	if len(wordList) == 0 {
//		return nil
//	}
//	return backTrack(beginWord, endWord, wordList, 0, []string{}, map[string]bool{})
//}
//
//func backTrack(beginWord string, endWord string, wordList []string, index int, temp []string, mem map[string]bool) []string {
//	if mem[beginWord] == true {
//		return nil
//	}
//	if beginWord == endWord {
//		return temp
//	}
//	if index >= len(wordList) {
//		return nil
//	}
//	for i := index; i < len(wordList); i++ {
//		if check(beginWord, wordList[i]) == false {
//			continue
//		}
//		t := wordList[i]
//		wordList[i], wordList[index] = wordList[index], wordList[i]
//		newTmp := make([]string, len(temp))
//		copy(newTmp, temp)
//		newTmp = append(newTmp, t)
//		result := dfs(t, endWord, wordList, index+1, newTmp, mem)
//		if result != nil {
//			return result
//		}
//		wordList[i], wordList[index] = wordList[index], wordList[i]
//		mem[t] = true
//	}
//	return nil
//}
//
//func check(word1, word2 string) bool {
//	diff := 0
//	for i := 0; i < len(word1); i++ {
//		if word1[i] != word2[i] {
//			diff++
//		}
//	}
//	if diff != 1 {
//		return false
//	}
//	return true
//}

func findLadders(beginWord string, endWord string, wordList []string) []string {
	m := map[string]bool{}
	for _, w := range wordList {
		m[w] = true
	}
	if m[beginWord] {
		m[beginWord] = false
	}
	if !m[endWord] {
		return nil
	}
	return dfs_(m, []string{}, []byte(beginWord), []byte(endWord))
}

func dfs_(m map[string]bool, path []string, beginWord []byte, endWord []byte) []string {
	path = append(path, string(beginWord))
	if string(beginWord) == string(endWord) {
		return path
	}

	for i, b := range beginWord {
		var j byte
		for j = 'a'; j <= 'z'; j++ {
			if j == b {
				continue
			}
			beginWord[i] = j
			w := string(beginWord)
			if !m[w] {
				beginWord[i] = b
				continue
			}
			m[w] = false
			res := dfs_(m, path, beginWord, endWord)
			if res != nil {
				return res
			}
			m[w] = false
			beginWord[i] = b
		}
	}
	return nil
}
