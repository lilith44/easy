package easy

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestNewPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue(true)

	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		x := int(rand.Int63n(math.MaxInt16))

		pq.Push(x, x)
	}

	var (
		x  int
		ok = true
	)

	a, _ := pq.Peak()
	b, _ := pq.Peak()

	if a != b {
		t.Fatal()
	}

	for ok {
		var y interface{}
		y, ok = pq.Pop()

		if ok && x > y.(int) {
			t.Fatal()
		}

		if !ok {
			break
		}

		x = y.(int)
	}

	if _, ok = pq.Peak(); ok {
		t.Fatal()
	}

	pq = NewPriorityQueue(false)

	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		x := int(rand.Int63n(math.MaxInt16))

		pq.Push(x, x)
	}

	ok = true

	a, _ = pq.Peak()
	b, _ = pq.Peak()
	if a != b {
		t.Fatal()
	}

	for ok {
		var y interface{}
		y, ok = pq.Pop()

		if ok && x < y.(int) {
			t.Fatal()
		}

		if !ok {
			break
		}

		x = y.(int)
	}
}
