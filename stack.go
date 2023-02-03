package easy

import "sync"

// NormalStack a normal stack is a first-in, last-out data structure.
type NormalStack struct {
	sync.Mutex

	stack []any
}

// NewNormalStack creates a normal stack.
func NewNormalStack() *NormalStack {
	return new(NormalStack)
}

// Push pushes a value into the stack.
func (ns *NormalStack) Push(value any) {
	ns.Lock()
	defer ns.Unlock()

	ns.stack = append(ns.stack, value)
	return
}

// Pop pops and deletes the pop value.
func (ns *NormalStack) Pop() (any, bool) {
	ns.Lock()
	defer ns.Unlock()

	if ns.IsEmpty() {
		return nil, false
	}
	value := ns.stack[len(ns.stack)-1]
	ns.stack = ns.stack[:len(ns.stack)-1]
	return value, true
}

// Top returns the top value.
func (ns *NormalStack) Top() (any, bool) {
	ns.Lock()
	defer ns.Unlock()

	if ns.IsEmpty() {
		return nil, false
	}
	return ns.stack[len(ns.stack)-1], true
}

// Bottom returns the bottom value.
func (ns *NormalStack) Bottom() (any, bool) {
	ns.Lock()
	defer ns.Unlock()

	if ns.IsEmpty() {
		return nil, false
	}
	return ns.stack[0], true
}

// IsEmpty checks if the stack is empty.
func (ns *NormalStack) IsEmpty() bool {
	return len(ns.stack) == 0
}

// Len returns the depth of stack.
func (ns *NormalStack) Len() int {
	return len(ns.stack)
}

// SortedStack a sorted stack is a first-in, last-out data structure.
// To guarantee the sort order, it will remove the numbers those are less or greater(depended on asc) than the given number while pushing.
type SortedStack[n number] struct {
	sync.Mutex

	asc   bool
	stack []n
}

// NewSortedStack creates a new sorted stack.
func NewSortedStack[n number](asc bool) *SortedStack[n] {
	return &SortedStack[n]{
		asc: asc,
	}
}

// Push pushes a number into the stack, and removes the numbers those are less or greater(depended on asc) than the given number.
func (ss *SortedStack[number]) Push(n number) {
	ss.Lock()
	defer ss.Unlock()

	if ss.IsEmpty() {
		ss.stack = append(ss.stack, n)
		return
	}

	index := len(ss.stack) - 1
	for ; index >= 0; index-- {
		if ss.asc && ss.stack[index] <= n {
			break
		}
		if !ss.asc && ss.stack[index] >= n {
			break
		}
	}

	ss.stack = append(ss.stack[:index+1], n)
}

// Pop pops and deletes the popped number.
func (ss *SortedStack[number]) Pop() (n number, popped bool) {
	ss.Lock()
	defer ss.Unlock()

	if ss.IsEmpty() {
		return
	}
	value := ss.stack[len(ss.stack)-1]
	ss.stack = ss.stack[:len(ss.stack)-1]
	return value, true
}

// Top returns the top number.
func (ss *SortedStack[number]) Top() (n number, exist bool) {
	ss.Lock()
	defer ss.Unlock()

	if ss.IsEmpty() {
		return
	}
	return ss.stack[len(ss.stack)-1], true
}

// Bottom returns the bottom number.
func (ss *SortedStack[number]) Bottom() (n number, exist bool) {
	ss.Lock()
	defer ss.Unlock()

	if ss.IsEmpty() {
		return
	}
	return ss.stack[0], true
}

// IsEmpty checks if the stack is empty.
func (ss *SortedStack[number]) IsEmpty() bool {
	return len(ss.stack) == 0
}

// Len returns the depth of stack.
func (ss *SortedStack[number]) Len() int {
	return len(ss.stack)
}
