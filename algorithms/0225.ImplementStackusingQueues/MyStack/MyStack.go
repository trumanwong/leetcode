package MyStack

type MyStack struct {
	Q []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	m := MyStack{[]int{}}
	return m
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.Q = append([]int{x}, this.Q...)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	res := this.Q[0]
	this.Q = this.Q[1:]
	return res
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Q[0]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.Q) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */