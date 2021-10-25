package _0211025

func canFinish(numCourses int, prerequisites [][]int) bool {
	hash := map[int][]int{}
	dp := make([]bool, numCourses)
	used := make([]bool, numCourses)
	for _, value := range prerequisites {
		hash[value[0]] = append(hash[value[0]], value[1])
	}
	var helper func(course int) bool
	helper = func(course int) bool {
		if used[course] == true {
			return dp[course]
		}
		used[course] = true
		for _, precourse := range hash[course] {
			if dp[precourse] == true {
				continue
			}
			if helper(precourse) == false {
				return false
			}
		}
		dp[course] = true
		return true
	}
	for i := 0; i <= numCourses-1; i++ {
		if helper(i) == false {
			return false
		}
	}
	return true
}
