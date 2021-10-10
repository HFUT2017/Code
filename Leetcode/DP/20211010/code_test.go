package _0211010

import "testing"

func TestCanPartition(t *testing.T) {
	//nums := []int{1, 5, 11, 5}
	//println(CanPartition(nums))
}

func TestLongestCommonSubsequence(t *testing.T) {
	test1 := "abcde"
	text2 := "ace"
	println(LongestCommonSubsequence(test1, text2))
}

func TestFindLength(t *testing.T) {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}
	println(FindLength(nums2, nums1))
}
