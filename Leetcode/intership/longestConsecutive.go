package intership

func find(parent []int, x int) int {
	for parent[x] != x {
		x = parent[x]
	}
	return parent[x]
}

func longestConsecutive(nums []int) int {
	hash := map[int]int{}
	n := len(nums)
	parent := make([]int, n)
	count := make([]int, n)
	maxNum := 0
	for index, _ := range nums {
		if _, ok := hash[nums[index]]; ok {
			continue
		}
		hash[nums[index]] = index
		parent[index] = index
		count[index] = 1
		if value, ok := hash[nums[index]-1]; ok {
			head := find(parent, value)
			parent[head] = index
			count[index] = count[head] + count[index]
		}
		if value, ok := hash[nums[index]+1]; ok {
			head := find(parent, value)
			parent[head] = index
			count[index] = count[head] + count[index]
		}
		if count[index] > maxNum {
			maxNum = count[index]
		}
	}
	return maxNum
}
