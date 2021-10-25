package _0211026

import (
	"strconv"
	"strings"
)

type Codec struct {

}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	sb:=&strings.Builder{}
	var dfs func(root *TreeNode)
	dfs= func(root *TreeNode) {
		if root==nil{
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

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	sp:=strings.Split(data,",")
	var build func()*TreeNode
	build= func() *TreeNode {
		if sp[0]=="null"{
			sp=sp[1:]
			return nil
		}
		val,_:=strconv.Atoi(sp[0])
		sp=sp[1:]
		return &TreeNode{Val: val,Left: build(),Right: build()}
	}
	return build()
}
