package _0211019

type hash [26]int

func groupAnagrams(strs []string) [][]string {
	hashMap := map[hash][]string{}
	res := [][]string{}
	for _, v := range strs {
		var t hash
		t.set(v)
		hashMap[t] = append(hashMap[t], v)
	}

	for _, v := range hashMap {
		res = append(res, v)
	}
	return res
}

func (h *hash) set(str string) {
	for _, v := range str {
		(*h)[v-'a']++
	}
}
