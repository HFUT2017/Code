package _0211027

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	hash := map[int]int{}
	res := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			hash[nums1[i]+nums2[j]]++
		}
	}

	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			res += hash[-nums3[i]-nums4[j]]
		}
	}
	return res
}
