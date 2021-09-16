package _0210915

import "testing"

func TestPermutation(t *testing.T) {
	s := "abc"
	res := Permutation(s)
	for i := 0; i < len(res); i++ {
		println(res[i])
	}
}

func order(node *TreeNode) {
	if node == nil {
		return
	}
	println(node.Val)
	order(node.Left)
	order(node.Right)
}

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := BuildTree(preorder, inorder)
	order(root)
}

func TestMovingCount(t *testing.T) {
	println(MovingCount(1, 2, 1))
}
