package easy

import (
	"container/heap"
	"sync"
)

// PriorityQueue a priority queue is a first-in, largest-out data structure.
type PriorityQueue struct {
	sync.Mutex

	queue *priorityQueue
}

// NewPriorityQueue creates a priority queue. Member having greater priority pops earlier if asc = true.
func NewPriorityQueue(asc bool) *PriorityQueue {
	return &PriorityQueue{
		queue: &priorityQueue{
			members: make([]*PriorityMember, 0),
			asc:     asc,
		},
	}
}

// Push pushes a member with priority into the queue.
func (pq *PriorityQueue) Push(value any, priority int) {
	pq.Lock()
	defer pq.Unlock()

	heap.Push(pq.queue, &PriorityMember{
		value:    value,
		priority: priority,
	})
}

// Pop pops a member from the queue.
func (pq *PriorityQueue) Pop() (any, bool) {
	pq.Lock()
	defer pq.Unlock()

	if pq.queue.Len() == 0 {
		return nil, false
	}

	member := heap.Pop(pq.queue).(*PriorityMember)
	return member.value, true
}

// Len gets the length of the queue.
func (pq *PriorityQueue) Len() int {
	pq.Lock()
	defer pq.Unlock()

	return pq.queue.Len()
}

// Peak returns the member waiting to be popped.
func (pq *PriorityQueue) Peak() (any, bool) {
	pq.Lock()
	defer pq.Unlock()

	if pq.queue.Len() == 0 {
		return nil, false
	}

	member := pq.queue.Peak().(*PriorityMember)
	return member.value, false
}

type PriorityMember struct {
	value    any
	priority int
}

type priorityQueue struct {
	asc     bool
	members []*PriorityMember
}

func (pq *priorityQueue) Len() int {
	return len(pq.members)
}

func (pq *priorityQueue) Less(i, j int) bool {
	if pq.asc {
		return pq.members[i].priority < pq.members[j].priority
	}
	return pq.members[i].priority > pq.members[j].priority
}

func (pq *priorityQueue) Swap(i, j int) {
	pq.members[i], pq.members[j] = pq.members[j], pq.members[i]
}

func (pq *priorityQueue) Pop() any {
	p := pq.members[len(pq.members)-1]
	pq.members = pq.members[: len(pq.members)-1 : len(pq.members)-1]
	return p
}

func (pq *priorityQueue) Push(member any) {
	pq.members = append(pq.members, member.(*PriorityMember))
}

func (pq *priorityQueue) Peak() any {
	if pq.asc {
		return pq.members[0]
	}
	return pq.members[len(pq.members)-1]
}
