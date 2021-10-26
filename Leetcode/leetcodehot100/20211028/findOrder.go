package _0211028

func findOrder(numCourses int, prerequisites [][]int) []int {
	edge := make([][]int, numCourses)
	preCourseNum := make([]int, numCourses)
	queue := []int{}
	res := []int{}
	for _, preCourse := range prerequisites {
		edge[preCourse[1]] = append(edge[preCourse[1]], preCourse[0])
		preCourseNum[preCourse[0]]++
	}
	for i := 0; i < numCourses; i++ {
		if preCourseNum[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		nowCourse := queue[0]
		queue = queue[1:]
		res = append(res, nowCourse)
		for _, nextCourse := range edge[nowCourse] {
			preCourseNum[nextCourse]--
			if preCourseNum[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}
	if len(res) == numCourses {
		return res
	}
	return []int{}
}
