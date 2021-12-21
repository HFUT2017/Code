package amazon

func LadderLength(beginWord string, endWord string, wordList []string) int {
	hash := map[string]bool{}
	for _, word := range wordList {
		hash[word] = true
	}
	if len(wordList) == 0 || hash[endWord] == false {
		return 0
	}
	count := 1
	visited := map[string]bool{beginWord: true, endWord: true}
	beginVisited := map[string]bool{beginWord: true}
	endVisited := map[string]bool{endWord: true}
	for len(beginVisited) > 0 && len(endVisited) > 0 {
		if len(beginVisited) > len(endVisited) {
			beginVisited, endVisited = endVisited, beginVisited
		}
		nextLevelVisited := map[string]bool{}
		for word, _ := range beginVisited {
			nextLevelList := expand(word, hash)
			for _, next := range nextLevelList {
				if endVisited[next] {
					return count + 1
				}
				if visited[next] == false {
					visited[next] = true
					nextLevelVisited[next] = true
				}
			}
		}
		beginVisited = nextLevelVisited
		count += 1
	}
	return 0
}

func expand(word string, hash map[string]bool) []string {
	str := []string{}
	letter := []byte(word)
	for index, _ := range letter {
		ori := letter[index]
		for c := byte('a'); c <= byte('z'); c++ {
			if c == ori {
				continue
			}
			letter[index] = c
			if hash[string(letter)] == true {
				str = append(str, string(letter))
			}
		}
		letter[index] = ori
	}
	return str
}
