package test

import (
	"math/rand"
	"testing"
	"time"
)

func TestArray(t *testing.T) {
	nums1 := [3]int{}
	nums2 := [3]int{}
	//nums3 := [4]int{}
	nums4 := nums1
	println(nums1 == nums2)
	//println(nums1 == nums3)
	println(nums1 == nums4)
}

func TestRand(t *testing.T) {
	rand.Seed(time.Now().Unix())
	var b [10]int
	for i := 0; i < len(b); i++ {
		b[i] = rand.Intn(1000)
	}
	res := 0
	for _, value := range b {
		res += value
	}
	println(res)
}

func TestCopy(t *testing.T) {
	array := [3]int{1, 2, 3}
	array2 := array
	println(&array)
	println(&array2)
	slice := []int{1, 2, 3}
	slice2 := slice
	println(slice)
	println(slice2)
}
