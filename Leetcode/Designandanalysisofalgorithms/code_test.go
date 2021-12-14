package Designandanalysisofalgorithms

import (
	"fmt"
	"testing"
)

func TestEuclidcomfactor(t *testing.T) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("i: %d j:%d comfactor:%d\n", i, j, Euclidcomfactor(i, j))
		}
	}
}

func TestMinMax(t *testing.T) {
	println(MinMax([]int{1, 2, 3}))
	println(MinMax([]int{3, 2, 1}))
	println(MinMax([]int{3, 3, 3}))
}

func TestMultiplicationAlgorithm(t *testing.T) {
	fmt.Printf("%v\n", MultiplicationAlgorithm([]int{1, 2, 3}, []int{4, 5, 6}))
}
