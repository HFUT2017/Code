package _0211020

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length%2 == 1 {
		return float64(getKelement(nums1, nums2, length/2+1))
	} else {
		return float64(getKelement(nums1, nums2, length/2)+getKelement(nums1, nums2, length/2+1)) / 2.0
	}
}

func getKelement(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		diff := k / 2
		newIndex1 := min(len(nums1), index1+diff) - 1
		newIndex2 := min(len(nums2), index2+diff) - 1
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
