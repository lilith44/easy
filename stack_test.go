package easy

import "testing"

func TestNormalStack_Push(t *testing.T) {
	suit := []struct {
		queue []int
		top   []int
		len   []int
	}{
		{
			queue: []int{1, 3, 5, 2, 6, 3},
			top:   []int{1, 3, 5, 2, 6, 3},
			len:   []int{1, 2, 3, 4, 5, 6},
		},
		{
			queue: []int{300, 100, 50, 25, 12, 0},
			top:   []int{300, 100, 50, 25, 12, 0},
			len:   []int{1, 2, 3, 4, 5, 6},
		},
		{
			queue: []int{0},
			top:   []int{0},
			len:   []int{1},
		},
	}

	for _, _case := range suit {
		stack := NewNormalStack()
		for i := range _case.queue {
			stack.Push(_case.queue[i])
			top, exist := stack.Top()
			if !exist || top != _case.top[i] {
				t.Fatal()
			}

			if stack.Len() != _case.len[i] {
				t.Fatal()
			}
		}
	}
}

func TestNormalStack_Pop(t *testing.T) {
	suit := []int{0, 5, 10}

	for _, _case := range suit {
		stack := NewNormalStack()
		for i := 0; i < _case; i++ {
			stack.Push(i)
		}

		for i := 0; i < _case*2; i++ {
			pop, exist := stack.Pop()
			if i < _case && (!exist || pop != _case-i-1) {
				t.Fatal()
			}

			if i >= _case && exist {
				t.Fatal()
			}
		}
	}
}

func TestNormalStack_Top(t *testing.T) {
	suit := []int{0, 5, 10}

	for _, _case := range suit {
		stack := NewNormalStack()
		for i := 0; i < _case; i++ {
			stack.Push(i)
		}

		top, exist := stack.Top()
		if _case == 0 && exist {
			t.Fatal()
		}
		if _case != 0 && (!exist || top != _case-1) {
			t.Fatal()
		}
	}
}

func TestNormalStack_Bottom(t *testing.T) {
	suit := []int{0, 5, 10}

	for _, _case := range suit {
		stack := NewNormalStack()
		for i := 0; i < _case; i++ {
			stack.Push(i)
		}

		bottom, exist := stack.Bottom()
		if _case == 0 && exist {
			t.Fatal()
		}
		if _case != 0 && (!exist || bottom != 0) {
			t.Fatal()
		}
	}
}

func TestNormalStack_IsEmpty(t *testing.T) {
	stack := NewNormalStack()
	if !stack.IsEmpty() {
		t.Fatal()
	}
	for i := 0; i < 100; i++ {
		stack.Push(i)
		if stack.IsEmpty() {
			t.Fatal()
		}
	}
}

func TestNormalStack_Len(t *testing.T) {
	stack := NewNormalStack()
	if stack.Len() != 0 {
		t.Fatal()
	}
	for i := 0; i < 100; i++ {
		stack.Push(i)
		if stack.Len() != i+1 {
			t.Fatal()
		}
	}
}

func TestSortedStack_Push(t *testing.T) {
	suit := []struct {
		queue []int
		top   []int
		len   []int
	}{
		{
			queue: []int{1, 3, 5, 2, 6, 3},
			top:   []int{1, 3, 5, 2, 6, 3},
			len:   []int{1, 2, 3, 2, 3, 3},
		},
		{
			queue: []int{300, 100, 50, 25, 12, 0},
			top:   []int{300, 100, 50, 25, 12, 0},
			len:   []int{1, 1, 1, 1, 1, 1},
		},
		{
			queue: []int{0},
			top:   []int{0},
			len:   []int{1},
		},
	}

	for _, _case := range suit {
		stack := NewSortedStack[int](true)
		for i := range _case.queue {
			stack.Push(_case.queue[i])
			top, exist := stack.Top()
			if !exist || top != _case.top[i] {
				t.Fatal()
			}

			if stack.Len() != _case.len[i] {
				t.Fatal()
			}
		}
	}
}

func TestSortedStack_Pop(t *testing.T) {
	suit := []int{0, 5, 10}

	for _, _case := range suit {
		stack := NewSortedStack[int](true)
		for i := 0; i < _case; i++ {
			stack.Push(i)
		}

		for i := 0; i < _case*2; i++ {
			pop, exist := stack.Pop()
			if i < _case && (!exist || pop != _case-i-1) {
				t.Fatal()
			}

			if i >= _case && exist {
				t.Fatal()
			}
		}
	}
}

func TestSortedStack_Top(t *testing.T) {
	suit := []int{0, 5, 10}

	for _, _case := range suit {
		stack := NewSortedStack[int](true)
		for i := 0; i < _case; i++ {
			stack.Push(i)
		}

		top, exist := stack.Top()
		if _case == 0 && exist {
			t.Fatal()
		}
		if _case != 0 && (!exist || top != _case-1) {
			t.Fatal()
		}
	}
}

func TestSortedStack_Bottom(t *testing.T) {
	suit := []int{0, 5, 10}

	for _, _case := range suit {
		stack := NewSortedStack[int](true)
		for i := 0; i < _case; i++ {
			stack.Push(i)
		}

		bottom, exist := stack.Bottom()
		if _case == 0 && exist {
			t.Fatal()
		}
		if _case != 0 && (!exist || bottom != 0) {
			t.Fatal()
		}
	}
}

func TestSortedStack_IsEmpty(t *testing.T) {
	stack := NewSortedStack[int](true)
	if !stack.IsEmpty() {
		t.Fatal()
	}
	for i := 0; i < 100; i++ {
		stack.Push(i)
		if stack.IsEmpty() {
			t.Fatal()
		}
	}
}

func TestSortedStack_Len(t *testing.T) {
	stack := NewSortedStack[int](true)
	if stack.Len() != 0 {
		t.Fatal()
	}
	for i := 0; i < 100; i++ {
		stack.Push(i)
		if stack.Len() != i+1 {
			t.Fatal()
		}
	}
}
