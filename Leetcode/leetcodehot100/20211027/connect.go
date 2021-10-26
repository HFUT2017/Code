package _0211027

type NodeConnect struct {
	Val   int
	Left  *NodeConnect
	Right *NodeConnect
	Next  *NodeConnect
}

func connect(root *NodeConnect) *NodeConnect {
	if root == nil {
		return nil
	}
	queue := []*NodeConnect{root}
	index, size := 0, 1
	for index < size {
		n := size - index
		for i := 0; i < n-1; i++ {
			queue[index+i].Next = queue[index+i+1]
		}
		for i := 0; i < n; i++ {
			if queue[index+i].Left != nil {
				queue = append(queue, queue[index+i].Left)
				size++
			}
			if queue[index+i].Right != nil {
				queue = append(queue, queue[index+i].Right)
				size++
			}
		}
		index += n
	}
	return root
}
