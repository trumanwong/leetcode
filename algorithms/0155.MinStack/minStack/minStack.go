package minStack

import "sort"

type MinStack struct {
	Nums []int
	Min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{Nums: []int{}, Min: 1 << 31 - 1}
}


func (this *MinStack) Push(x int)  {
	this.Nums = append(this.Nums, x)
	if this.Min > x {
		this.Min = x
	}
}


func (this *MinStack) Pop()  {
	last := this.Nums[len(this.Nums) - 1]
	this.Nums = this.Nums[:len(this.Nums) - 1]
	if last == this.Min {
		temp := make([]int, len(this.Nums))
		copy(temp, this.Nums)
		sort.Ints(temp)
		if len(this.Nums) > 0 {
			this.Min = temp[0]
		} else {
			this.Min = 1 << 31 - 1
		}
	}
}


func (this *MinStack) Top() int {
	return this.Nums[len(this.Nums) - 1]
}


func (this *MinStack) GetMin() int {
	return this.Min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */