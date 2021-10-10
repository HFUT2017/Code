package _0211010

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	res := []int{}
	hash1 := map[int]bool{}
	hash2 := map[int]bool{}
	hash3 := map[int]bool{}
	hash4 := map[int]bool{}
	for i := 0; i < len(nums1); i++ {
		hash1[nums1[i]] = true
	}
	for i := 0; i < len(nums2); i++ {
		hash2[nums2[i]] = true
	}
	for i := 0; i < len(nums3); i++ {
		hash3[nums3[i]] = true
	}

	for i := 0; i < len(nums1); i++ {
		if hash2[nums1[i]] || hash3[nums1[i]] {
			if _,ok:=hash4[nums1[i]];!ok {
				res = append(res, nums1[i])
			}
			hash4[nums1[i]] = true
		}
	}

	for i := 0; i < len(nums2); i++ {
		if hash3[nums2[i]] {
			if _,ok:=hash4[nums2[i]];!ok {
				res = append(res, nums2[i])
			}
			hash4[nums2[i]] = true
		}
	}

	return res
}
