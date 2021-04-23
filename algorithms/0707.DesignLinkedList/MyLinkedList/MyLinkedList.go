package MyLinkedList

type MyLinkedList struct {
	nums []int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	nums := make([]int, 0)
	return MyLinkedList{nums: nums}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= len(this.nums) {
		return -1
	}
	return this.nums[index]
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.nums = append([]int{val}, this.nums...)
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.nums = append(this.nums, val)
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > len(this.nums) {
		return
	}
	if index == len(this.nums) {
		this.nums = append(this.nums, val)
		return
	}
	if index < 0 {
		this.nums = append([]int{val}, this.nums...)
		return
	}
	temp := make([]int, len(this.nums[index:]))
	copy(temp, this.nums[index:])
	this.nums = append(append(this.nums[:index], val), temp...)
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= len(this.nums) {
		return
	}
	this.nums = append(this.nums[:index], this.nums[index+1:]...)
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
