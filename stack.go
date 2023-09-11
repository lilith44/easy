package easy

import (
	"cmp"
	"errors"
	"slices"
	"sync"
)

var (
	ErrFullStack  = errors.New("full stack. ")
	ErrEmptyStack = errors.New("empty stack. ")
)

// Stack is an abstract data type that serves as a collection of elements, with two main operations Push and Pop.
type Stack[E any] struct {
	sync.Mutex

	stack    []E
	capacity int
}

// NewStack creates a stack with capacity.
func NewStack[E any](capacity int) *Stack[E] {
	if capacity <= 0 {
		panic("non-positive capacity. ")
	}

	return &Stack[E]{
		stack:    make([]E, 0, capacity),
		capacity: capacity,
	}
}

// Push adds an element to the stack.
// Notice that it returns an error if the stack is full.
func (s *Stack[E]) Push(element E) error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if s.isFull() {
		return ErrFullStack
	}

	s.stack = append(s.stack, element)
	return nil
}

// Pop removes and returns the most recently added element that was not yet removed.
// Notice that it returns an error if the stack is empty.
func (s *Stack[E]) Pop() (E, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	var popped E
	if s.isEmpty() {
		return popped, ErrEmptyStack
	}

	popped = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return popped, nil
}

func (s *Stack[E]) isFull() bool {
	return len(s.stack) == s.capacity
}

func (s *Stack[E]) isEmpty() bool {
	return len(s.stack) == 0
}

// MonotoneStack is a special kind of stack, with elements stored in an increasing or decreasing order.
type MonotoneStack[E cmp.Ordered] struct {
	sync.Mutex

	stack []E
	asc   bool
}

// NewMonotoneStack creates a monotone stack in an increasing(if asc is true) or decreasing(if asc is false) order.
func NewMonotoneStack[E cmp.Ordered](asc bool) *MonotoneStack[E] {
	return &MonotoneStack[E]{
		asc: asc,
	}
}

// Push adds an element to the stack. To keep the elements in an increasing or decreasing order, elements which are greater
// or less than the element to push will be removed.
func (ms *MonotoneStack[E]) Push(element E) []E {
	ms.Lock()
	defer ms.Unlock()

	if ms.isEmpty() {
		ms.stack = append(ms.stack, element)
		return nil
	}

	if ms.asc && ms.stack[len(ms.stack)-1] <= element {
		ms.stack = append(ms.stack, element)
		return nil
	}

	if !ms.asc && ms.stack[0] >= element {
		ms.stack = append([]E{element}, ms.stack...)
		return nil
	}

	index, exist := slices.BinarySearch(ms.stack, element)
	if ms.asc {
		i := index
		if exist {
			for i = index; i < len(ms.stack); i++ {
				if ms.stack[i] != ms.stack[index] {
					break
				}
			}
		}
		popped := make([]E, len(ms.stack[i:]))
		copy(popped, ms.stack[i:])
		slices.Reverse(popped)

		ms.stack = append(ms.stack[:i], element)
		return popped
	}

	popped := make([]E, len(ms.stack[:index]))
	copy(popped, ms.stack[:index])
	ms.stack = append([]E{element}, ms.stack[index:]...)
	return popped
}

// Pop removes and returns the most recently added element that was not yet removed.
// Notice that it returns an error if the stack is empty.
func (ms *MonotoneStack[E]) Pop() (E, error) {
	ms.Lock()
	defer ms.Unlock()

	var popped E
	if ms.isEmpty() {
		return popped, ErrEmptyStack
	}

	if ms.asc {
		popped = ms.stack[len(ms.stack)-1]
		ms.stack = ms.stack[:len(ms.stack)-1]
	} else {
		popped = ms.stack[0]
		ms.stack = ms.stack[1:]
	}
	return popped, nil
}

func (ms *MonotoneStack[E]) isEmpty() bool {
	return len(ms.stack) == 0
}
