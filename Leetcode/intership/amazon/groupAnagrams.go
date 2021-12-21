package amazon

import "sort"

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	hash := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		hash[string(s)] = append(hash[string(s)], str)
	}
	for _, value := range hash {
		res = append(res, value)
	}
	return res
}
