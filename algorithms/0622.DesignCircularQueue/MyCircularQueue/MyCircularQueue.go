package MyCircularQueue

type MyCircularQueue struct {
	queue  []int
	length int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	queue := make([]int, 0)
	return MyCircularQueue{queue: queue, length: k}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.queue) >= this.length {
		return false
	}
	this.queue = append(this.queue, value)
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if len(this.queue) == 0 {
		return false
	}
	this.queue = this.queue[1:]
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[0]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[len(this.queue)-1]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if len(this.queue) == 0 {
		return true
	}
	return false
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if len(this.queue) == this.length {
		return true
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
