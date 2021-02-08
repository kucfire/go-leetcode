package leetcode155

/** initialize your data structure here. */
func ConstructorReview() MinStack {
	return MinStack{
		stack:      []int{},
		minElement: []int{},
	}
}

func (this *MinStack) PushReview(x int) {
	this.stack = append(this.stack, x)
	top := this.minElement[len(this.minElement)-1]
	xx := top
	if x < top {
		xx = x
	}
	this.minElement = append(this.minElement, xx)

}

func (this *MinStack) PopReview() {
	if this == nil {
		return
	}
	this.minElement = this.minElement[:len(this.minElement)-1]
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) TopReview() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMinReview() int {
	return this.minElement[len(this.minElement)-1]
}
