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

type Stack[E any] struct {
	sync.Mutex

	stack    []E
	capacity int
}

func NewStack[E any](capacity int) *Stack[E] {
	if capacity <= 0 {
		panic("non-positive capacity. ")
	}

	return &Stack[E]{
		stack:    make([]E, 0, capacity),
		capacity: capacity,
	}
}

func (s *Stack[E]) Push(element E) error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if s.isFull() {
		return ErrFullStack
	}

	s.stack = append(s.stack, element)
	return nil
}

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

type SortedStack[E cmp.Ordered] struct {
	sync.Mutex

	stack []E
	asc   bool
}

func NewSortedStack[E cmp.Ordered](asc bool) *SortedStack[E] {
	return &SortedStack[E]{
		asc: asc,
	}
}

func (ss *SortedStack[E]) Push(element E) []E {
	ss.Lock()
	defer ss.Unlock()

	if ss.isEmpty() {
		ss.stack = append(ss.stack, element)
		return nil
	}

	if ss.asc && ss.stack[len(ss.stack)-1] <= element {
		ss.stack = append(ss.stack, element)
		return nil
	}

	if !ss.asc && ss.stack[0] >= element {
		ss.stack = append([]E{element}, ss.stack...)
		return nil
	}

	index, exist := slices.BinarySearch(ss.stack, element)
	if ss.asc {
		i := index
		if exist {
			for i = index; i < len(ss.stack); i++ {
				if ss.stack[i] != ss.stack[index] {
					break
				}
			}
		}
		popped := make([]E, len(ss.stack[i:]))
		copy(popped, ss.stack[i:])
		slices.Reverse(popped)

		ss.stack = append(ss.stack[:i], element)
		return popped
	}

	popped := make([]E, len(ss.stack[:index]))
	copy(popped, ss.stack[:index])
	ss.stack = append([]E{element}, ss.stack[index:]...)
	return popped
}

func (ss *SortedStack[E]) Pop() (E, error) {
	ss.Lock()
	defer ss.Unlock()

	var popped E
	if ss.isEmpty() {
		return popped, ErrEmptyStack
	}

	if ss.asc {
		popped = ss.stack[len(ss.stack)-1]
		ss.stack = ss.stack[:len(ss.stack)-1]
	} else {
		popped = ss.stack[0]
		ss.stack = ss.stack[1:]
	}
	return popped, nil
}

func (ss *SortedStack[E]) isEmpty() bool {
	return len(ss.stack) == 0
}
