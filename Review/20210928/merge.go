package _0210928

func merge(nums1 []int, m int, nums2 []int, n int) {
	index1, index2, tail := m-1, n-1, m+n-1
	for index1 >= 0 || index2 >= 0 {
		cur := 0
		if index1 < 0 {
			cur = nums2[index2]
			index2--
		} else if index2 < 0 {
			cur = nums1[index1]
			index1--
		} else if nums1[index1] > nums2[index2] {
			cur = nums1[index1]
			index1--
		} else {
			cur = nums2[index2]
			index2--
		}
		nums1[tail] = cur
		tail--
	}
}
