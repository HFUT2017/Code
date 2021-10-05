package _0211004

func levelOrder3(root *TreeNode) [][]int {
	if root==nil{
		return [][]int{}
	}
	queue:=[]*TreeNode{root}
	res:=[][]int{}
	index,size:=0,1
	for index<size{
		temp:=[]int{}
		n:=size-index
		for i:=0;i<n;i++{
			node:=queue[i+index]
			temp=append(temp,node.Val)
			if node.Left!=nil{
				queue=append(queue,node.Left)
				size++
			}
			if node.Right!=nil{
				queue=append(queue,node.Right)
				size++
			}
		}
		res=append(res,temp)
		index+=n
	}
	return res
}
