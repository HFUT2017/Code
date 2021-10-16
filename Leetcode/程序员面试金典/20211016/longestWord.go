package _0211016

func longestWord(words []string) string {
	hash := make(map[string]bool)
	for _, value := range words {
		hash[value] = true
	}
	var dfs func(str string) bool
	dfs = func(str string) bool {
		if hash[str] == true {
			return true
		}
		for i := 0; i < len(str); i++ {
			if hash[str[:i]] == false {
				continue
			}
			if dfs(str[i:]) {
				hash[str[i:]] = true
				return true
			}

		}
		return false
	}
	res := ""
	for _, value := range words {
		delete(hash, value)
		if dfs(value) {
			res = compare(res, value)
		}
		hash[value] = true
	}
	return res
}

func compare(word1, word2 string) string {
	if len(word1) < len(word2) {
		return word2
	} else if len(word1) > len(word2) {
		return word1
	}

	for i := 0; i < len(word1); i++ {
		if word1[i] < word2[i] {
			return word1
		} else if word1[i] > word2[i] {
			return word2
		}
	}
	return word1
}
