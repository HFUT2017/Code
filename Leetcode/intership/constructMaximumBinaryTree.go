package intership

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := 0
	t := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > t {
			t=nums[i]
			index = i
		}
	}
	root := &TreeNode{Val: nums[index]}
	root.Left = ConstructMaximumBinaryTree(nums[:index])
	root.Right = ConstructMaximumBinaryTree(nums[index+1:])
	return root
}
