package _0211004

func verifyPostorder(postorder []int) bool {
	for i:=len(postorder)-1;i>=0;i--{
		j:=i-1
		for j>=0&&postorder[j]>postorder[i]{
			j--
		}
		for j>=0&&postorder[j]<postorder[i]{
			j--
		}
		if j>=0{
			return false
		}
	}
	return true
}
