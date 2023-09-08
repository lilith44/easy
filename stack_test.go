package easy

import (
	"reflect"
	"sync"
	"testing"
)

var stackPushTests = []struct {
	element int
	want    []int
}{
	{
		element: 1,
		want:    []int{1},
	},
	{
		element: 15,
		want:    []int{1, 15},
	},
	{
		element: 17,
		want:    []int{1, 15, 17},
	},
	{
		element: 22,
		want:    []int{1, 15, 17, 22},
	},
	{
		element: 35,
		want:    []int{1, 15, 17, 22, 35},
	},
}

func TestStack_Push(t *testing.T) {
	stack := NewStack[int](len(stackPushTests))
	for _, test := range stackPushTests {
		if err := stack.Push(test.element); err != nil {
			t.Errorf("error occurs while Push: %s, want no error", err)
			t.Fail()
		}
	}

	if err := stack.Push(100); err == nil {
		t.Errorf("no error occurs while Push, want error")
		t.Fail()
	}
}

func TestStack_Push_Concurrent(t *testing.T) {
	num := 100
	stack := NewStack[int](num)
	ch := make(chan error, num)
	wg := sync.WaitGroup{}
	wg.Add(num * 2)
	for i := 0; i < num*2; i++ {
		go func(i int) {
			defer wg.Done()
			err := stack.Push(i)
			if err != nil {
				ch <- err
			}
		}(i)
	}
	wg.Wait()

	if len(ch) != num || len(stack.stack) != num {
		t.Errorf("element missed")
	}
}

func TestStack_Pop(t *testing.T) {
	num := 100
	stack := NewStack[int](num)
	for i := 0; i < num; i++ {
		_ = stack.Push(i)
	}

	for i := 99; i >= 0; i-- {
		popped, err := stack.Pop()
		if err != nil {
			t.Errorf("error occurs while Pop: %s, want no error", err)
			t.Fail()
		}
		if popped != i {
			t.Errorf("wrong popped value, got: %v, want %v", popped, i)
		}
	}

	if _, err := stack.Pop(); err == nil {
		t.Errorf("no error occurs while Pop, want error")
		t.Fail()
	}
}

func TestStack_Pop_Concurrent(t *testing.T) {
	num := 100
	stack := NewStack[int](num)
	for i := 0; i < num; i++ {
		_ = stack.Push(i)
	}

	ch := make(chan error, num)
	wg := sync.WaitGroup{}
	wg.Add(num * 2)
	for i := 0; i < num*2; i++ {
		go func(i int) {
			defer wg.Done()
			_, err := stack.Pop()
			if err != nil {
				ch <- err
			}
		}(i)
	}
	wg.Wait()

	if len(ch) != num || len(stack.stack) != 0 {
		t.Errorf("element missed")
	}
}

var increaseSortedStackPushTests = []struct {
	element int
	want    []int
	popped  []int
}{
	{
		element: 10,
		want:    []int{10},
		popped:  nil,
	},
	{
		element: 20,
		want:    []int{10, 20},
		popped:  nil,
	},
	{
		element: 15,
		want:    []int{10, 15},
		popped:  []int{20},
	},
	{
		element: 25,
		want:    []int{10, 15, 25},
		popped:  nil,
	},
	{
		element: 35,
		want:    []int{10, 15, 25, 35},
		popped:  nil,
	},
	{
		element: 35,
		want:    []int{10, 15, 25, 35, 35},
		popped:  nil,
	},
	{
		element: 35,
		want:    []int{10, 15, 25, 35, 35, 35},
		popped:  nil,
	},
	{
		element: 25,
		want:    []int{10, 15, 25, 25},
		popped:  []int{35, 35, 35},
	},
	{
		element: 1,
		want:    []int{1},
		popped:  []int{25, 25, 15, 10},
	},
}

var decreaseSortedStackPushTests = []struct {
	element int
	want    []int
	popped  []int
}{
	{
		element: 100,
		want:    []int{100},
		popped:  nil,
	},
	{
		element: 120,
		want:    []int{120},
		popped:  []int{100},
	},
	{
		element: 90,
		want:    []int{90, 120},
		popped:  nil,
	},
	{
		element: 70,
		want:    []int{70, 90, 120},
		popped:  nil,
	},
	{
		element: 70,
		want:    []int{70, 70, 90, 120},
		popped:  nil,
	},
	{
		element: 70,
		want:    []int{70, 70, 70, 90, 120},
		popped:  nil,
	},
	{
		element: 60,
		want:    []int{60, 70, 70, 70, 90, 120},
		popped:  nil,
	},
	{
		element: 70,
		want:    []int{70, 70, 70, 70, 90, 120},
		popped:  []int{60},
	},
	{
		element: 150,
		want:    []int{150},
		popped:  []int{70, 70, 70, 70, 90, 120},
	},
}

func TestSortedStack_Push(t *testing.T) {
	stack1 := NewSortedStack[int](true)
	for _, test := range increaseSortedStackPushTests {
		popped := stack1.Push(test.element)
		if !reflect.DeepEqual(stack1.stack, test.want) {
			t.Errorf("Push(%v), ss.stack got %v, want %v", test.element, stack1.stack, test.want)
		}
		if !reflect.DeepEqual(popped, test.popped) {
			t.Errorf("Push(%v), popped %v, want %v", test.element, popped, test.popped)
		}
	}

	stack2 := NewSortedStack[int](false)
	for _, test := range decreaseSortedStackPushTests {
		popped := stack2.Push(test.element)
		if !reflect.DeepEqual(stack2.stack, test.want) {
			t.Errorf("Push(%v), got %v, want %v", test.element, stack2.stack, test.want)
		}
		if !reflect.DeepEqual(popped, test.popped) {
			t.Errorf("Push(%v), popped %v, want %v", test.element, popped, test.popped)
		}
	}
}

var increaseSortedStackPopTests = []struct {
	push []int
	pop  []int
}{
	{
		push: []int{1, 2, 3, 4, 5},
		pop:  []int{5, 4, 3, 2, 1},
	},
	{
		push: []int{1, 2, 3, 4, 3},
		pop:  []int{3, 3, 2, 1},
	},
}

var decreaseSortedStackPopTests = []struct {
	push []int
	pop  []int
}{
	{
		push: []int{1, 2, 3, 4, 5},
		pop:  []int{5},
	},
	{
		push: []int{5, 4, 3, 2, 1},
		pop:  []int{1, 2, 3, 4, 5},
	},
}

func TestSortedStack_Pop(t *testing.T) {
	stack1 := NewSortedStack[int](true)
	for _, test := range increaseSortedStackPopTests {
		for _, element := range test.push {
			stack1.Push(element)
		}

		for _, element := range test.pop {
			popped, err := stack1.Pop()
			if err != nil {
				t.Errorf("error occurs while Pop: %s, want no error", err)
				t.Fail()
			}
			if !reflect.DeepEqual(popped, element) {
				t.Errorf("Pop(), popped %v, want %v", popped, element)
			}
		}
	}

	stack2 := NewSortedStack[int](false)
	for _, test := range decreaseSortedStackPopTests {
		for _, element := range test.push {
			stack2.Push(element)
		}

		for _, element := range test.pop {
			popped, err := stack2.Pop()
			if err != nil {
				t.Errorf("error occurs while Pop: %s, want no error", err)
				t.Fail()
			}
			if !reflect.DeepEqual(popped, element) {
				t.Errorf("Pop(), popped %v, want %v", popped, element)
			}
		}
	}
}
