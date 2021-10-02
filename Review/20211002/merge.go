package _0211002

func merge(nums1 []int, m int, nums2 []int, n int) {
	index1, index2, tail := m-1, n-1, m+n-1
	for index1 >= 0 && index2 >= 0 {
		if index2 < 0 || nums1[index1] > nums2[index2] {
			nums1[tail] = nums1[index1]
			tail--
			index1--
		} else {
			nums1[tail] = nums2[index2]
			tail--
			index2--
		}
	}
}
