package amazon

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(root.Val))
		sb.WriteByte(',')
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return sb.String()
}

func (this *Codec) deserialize(data string) *TreeNode {
	str := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if str[0] == "null" {
			str = str[1:]
			return nil
		}
		val, _ := strconv.Atoi(str[0])
		str = str[1:]
		return &TreeNode{Val: val, Left: build(), Right: build()}
	}
	return build()
}
