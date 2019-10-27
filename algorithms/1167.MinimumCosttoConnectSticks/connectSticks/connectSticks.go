package connectSticks

import (
	"container/heap"
)

type Item struct {
	value    interface{} // The value of the item; arbitrary.
	priority int         // The priority of the item in the queue.
	index    int
}

// A PriorityQueue implements heap. Interface and holds Items
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value interface{}, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func ConnectSticks(sticks []int) int {
	pq := make(PriorityQueue, len(sticks))
	for i, v := range sticks {
		pq[i] = &Item{
			value:    v,
			priority: v,
			index:    i,
		}
	}
	heap.Init(&pq)
	ret := 0
	for pq.Len() > 1 {
		p := heap.Pop(&pq).(*Item).value.(int)
		q := heap.Pop(&pq).(*Item).value.(int)
		ret = ret + p + q
		item := &Item{
			value:    p + q,
			priority: p + q,
		}
		heap.Push(&pq, item)
	}
	return ret
}
