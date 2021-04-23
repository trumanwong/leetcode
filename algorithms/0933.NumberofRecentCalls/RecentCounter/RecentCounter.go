package RecentCounter

type RecentCounter struct {
	Queue []int
}

func Constructor() RecentCounter {
	queue := make([]int, 0, 10000)
	return RecentCounter{Queue: queue}
}

func (this *RecentCounter) Ping(t int) int {
	this.Queue = append(this.Queue, t)
	for this.Queue[0] < t-3000 {
		this.Queue = this.Queue[1:]
	}
	return len(this.Queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
