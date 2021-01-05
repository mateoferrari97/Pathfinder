package internal

import (
	"container/heap"
)

type item struct {
	parent            *vertex
	current           *vertex
	currentBestWeight int
	priority          float64
	index             int
}

type priorityQueue []*item

func newPriorityQueue() priorityQueue {
	return priorityQueue{}
}

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *priorityQueue) update(item *item, priority float64) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
