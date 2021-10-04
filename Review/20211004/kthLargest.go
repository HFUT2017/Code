package _0211004

func kthLargest(root *TreeNode, k int) int {
	var inorder func(root *TreeNode)
	res:=[]int{}
	inorder=func(root *TreeNode){
		if root==nil||k==0{
			return
		}
		inorder(root.Left)
		res=append(res,root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return res[len(res)-k]
}
