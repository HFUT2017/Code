package _0211003

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A==nil&&B==nil{
		return true
	}
	if A==nil||B==nil{
		return false
	}
	return isSubStructure(A.Left,B)||isSubStructure(A.Right,B)||helper(A,B)
}

func helper(A,B *TreeNode) bool{
	if A==nil&&B!=nil{
		return false
	}
	if B==nil{
		return true
	}
	if A.Val!=B.Val{
		return false
	}
	return helper(A.Left,B.Left)&&helper(A.Right,B.Right)
}
