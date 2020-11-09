/*
	leetcode tag:155 title:最小栈
*/

package leetcode155

import "math"

type MinStack struct {
	stack      []int
	minElement []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:      []int{},
		minElement: []int{math.MaxInt32}, // 用栈来保存新top元素加入时当前最小值
	}
}

// Push : add a new element in stack
func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.minElement[len(this.minElement)-1]
	xx := top
	if x < top {
		xx = x
	}
	this.minElement = append(this.minElement, xx)
}

// Pop : delete top element
func (this *MinStack) Pop() {
	if this == nil {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.minElement = this.minElement[:len(this.minElement)-1]
}

// Top : get the stack top element
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

// GetMin : Get the smallest in stack
func (this *MinStack) GetMin() int {
	return this.minElement[len(this.minElement)-1]
}
