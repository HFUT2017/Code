package _0211004

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func levelOrder(root *TreeNode) []int {
	if root==nil{
		return []int{}
	}
	queue:=[]*TreeNode{root}
	res:=[]int{}
	index,size:=0,1
	for index<size{
		n:=size-index
		for i:=0;i<n;i++{
			node:=queue[index+i]
			res=append(res,node.Val)
			if node.Left!=nil{
				queue=append(queue,node.Left)
				size++
			}
			if node.Right!=nil{
				queue=append(queue,node.Right)
				size++
			}
		}
		index+=n
	}
	return res
}
