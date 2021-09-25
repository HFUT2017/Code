package tsort

func BucketSort(nums []int, maxBucket int) []int {
	bucket := make([]int, maxBucket+1)
	for i := 0; i < len(nums); i++ {
		bucket[nums[i]]++
	}
	res := []int{}
	for i := 0; i < len(bucket); i++ {
		for j := 0; j < bucket[i]; j++ {
			res = append(res, i)
		}
	}
	return res
}
