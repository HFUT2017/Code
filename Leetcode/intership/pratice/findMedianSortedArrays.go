package intership

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l%2 == 0 {
		return float64(findKElement(nums1, nums2, l/2)+findKElement(nums1, nums2, l/2+1)) / 2.0
	} else {
		return float64(findKElement(nums1, nums2, l/2+1))
	}
}

func findKElement(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for index1 < len(nums1) || index2 < len(nums2) {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		} else if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		diff := k / 2
		newIndex1 := min(index1+diff, len(nums1)) - 1
		newIndex2 := min(index2+diff, len(nums2)) - 1
		if nums1[newIndex1] < nums2[newIndex2] {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}
