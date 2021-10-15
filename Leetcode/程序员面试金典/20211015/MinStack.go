package _0211015

import "container/list"

type MinStack struct {
	stack1 *list.List
	stack2 *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *MinStack) Push(x int) {
	this.stack1.PushBack(x)
	if this.stack2.Len() == 0 || x <= this.stack2.Back().Value.(int) {
		this.stack2.PushBack(x)
	}
}

func (this *MinStack) Pop() {
	if this.stack1.Back().Value.(int) == this.stack2.Back().Value.(int) {
		this.stack2.Remove(this.stack2.Back())
	}
	this.stack1.Remove(this.stack1.Back())
}

func (this *MinStack) Top() int {
	if this.stack1.Len() == 0 {
		return 0
	}
	return this.stack1.Back().Value.(int)
}

func (this *MinStack) GetMin() int {
	if this.stack2.Len() == 0 {
		return 0
	}
	return this.stack2.Back().Value.(int)
}
