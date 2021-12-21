package intership

func getDirections(root *TreeNode, startValue int, destValue int) string {
	lTemp := ""
	rTemp := ""
	lRes := ""
	rRes := ""
	var getCommonRoot func(root *TreeNode, left, right int) *TreeNode
	getCommonRoot = func(root *TreeNode, left, right int) *TreeNode {
		if root == nil {
			return nil
		}
		if root.Val == left {
			return root
		}
		if root.Val == right {
			return root
		}
		leftNode := getCommonRoot(root.Left, left, right)
		rightNode := getCommonRoot(root.Right, left, right)
		if leftNode == nil {
			return rightNode
		}
		if rightNode == nil {
			return leftNode
		}

		return root
	}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if lRes != "" && rRes != "" {
			return
		}
		if root.Val == startValue {
			lRes = lTemp
		}
		if root.Val == destValue {
			rRes = rTemp
		}
		lTemp += "U"
		rTemp += "L"
		dfs(root.Left)
		lTemp = lTemp[:len(lTemp)-1]
		rTemp = rTemp[:len(rTemp)-1]
		lTemp += "U"
		rTemp += "R"
		dfs(root.Right)
		lTemp = lTemp[:len(lTemp)-1]
		rTemp = rTemp[:len(rTemp)-1]
	}
	commonRoot := getCommonRoot(root, startValue, destValue)
	dfs(commonRoot)
	return lRes + rRes
}
