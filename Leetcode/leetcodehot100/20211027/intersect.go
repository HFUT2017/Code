package _0211027

func intersect(nums1 []int, nums2 []int) []int {
	hash := map[int]int{}
	res := []int{}
	for _, value := range nums1 {
		hash[value]++
	}
	for _, value := range nums2 {
		if hash[value] != 0 {
			res = append(res, value)
			hash[value]--
		}
	}
	return res
}

//有序
//func intersect(nums1 []int, nums2 []int) []int {
//	sort.Ints(nums1)
//	sort.Ints(nums2)
//	res := []int{}
//	l1, l2 := 0, 0
//	for l1 < len(nums1) && l2 < len(nums2) {
//		if nums1[l1] == nums2[l2] {
//			res = append(res, nums1[l1])
//			l1++
//			l2++
//		} else if nums1[l1] < nums2[l2] {
//			l1++
//		} else {
//			l2++
//		}
//	}
//	return res
//}
