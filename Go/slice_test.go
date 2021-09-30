package Go

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	parent := []int{0, 0, 3, 4}
	sub := parent[0:4]
	parent = append(parent, 5)
	parent[1] = 2
	sub[0] = 1
	for i := 0; i < len(parent); i++ {
		fmt.Printf("%d ", parent[i])
	}
	println()
	for i := 0; i < len(sub); i++ {
		fmt.Printf("%d ", sub[i])
	}
}

func TestSliceInit(t *testing.T) {
	//1
	var slice1 []int
	if slice1 == nil {
		println("slice1 is nil")
	}
	slice2 := []int{}
	if slice2 != nil {
		println("slice2 is empty")
	}
	slice3 := make([]int, 0)
	if slice3 == nil {
		println("slice3 is nil")
	} else {
		println("slice3 is empty")
	}
	//2
	arr := [6]int{0, 1, 2, 3, 4, 5}
	slice4 := arr[1:4]
	for i := 0; i < len(slice4); i++ {
		fmt.Printf("%d ", slice4[i])
		if i == len(slice4)-1 {
			println()
		}
	}
	arr[1] = 2
	for i := 0; i < len(slice4); i++ {
		fmt.Printf("%d ", slice4[i])
		if i == len(slice4)-1 {
			println()
		}
	}
	//3
	slice5 := []int{}
	println(slice5)
	slice6 := []int{}
	println(slice6)
	slice5 = append(slice5, 1)
	slice6 = append(slice6, 2)
	for i := 0; i < len(slice6); i++ {
		println(slice6[i])
	}
	//4
	array := [4]int{0, 1, 2, 3}
	slice7 := array[2:4]
	slice7 = append(slice7, 4)
	for i := 0; i < len(slice7); i++ {
		fmt.Printf("%d ", slice7[i])
		if i == len(slice7)-1 {
			println()
		}
	}
	array[2] = 3
	for i := 0; i < len(slice7); i++ {
		fmt.Printf("%d ", slice7[i])
		if i == len(slice7)-1 {
			println()
		}
	}

}
