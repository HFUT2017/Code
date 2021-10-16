package _0211016

////排序
//func groupAnagrams(strs []string) [][]string {
//	hash := map[string][]string{}
//	res := [][]string{}
//	for i := 0; i < len(strs); i++ {
//		str := []byte(strs[i])
//		sort.Slice(str, func(i, j int) bool {
//			return str[i] < str[j]
//		})
//		str_ := string(str[:])
//		hash[str_] = append(hash[str_], strs[i])
//	}
//	for _, value := range hash {
//		res = append(res, value)
//	}
//	return res
//}

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
