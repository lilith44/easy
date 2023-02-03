package easy

import (
	"reflect"
	"sync"
	"testing"
)

func TestBitmap_Contain(t *testing.T) {
	total := 100
	b := NewBitmap()

	for i := 0; i < total; i++ {
		b.Store(i)
	}

	for i := 0; i < total*2; i++ {
		if i < total && !b.Contain(i) {
			t.Fatal()
		}
		if i >= total && b.Contain(i) {
			t.Fatal()
		}
	}

}

func TestBitmap_Store(t *testing.T) {
	total := 100
	b := NewBitmap()

	wg := sync.WaitGroup{}
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func(i int) {
			defer wg.Done()
			b.Store(i)
		}(i)
	}
	wg.Wait()

	keys := b.Values()
	if len(keys) != total {
		t.Fatal()
	}
	for i := range keys {
		if keys[i] != i {
			t.Fatal()
		}
	}
}

func TestBitmap_StoreMulti(t *testing.T) {
	total := 1000
	b := NewBitmap()

	wg := sync.WaitGroup{}
	wg.Add(total / 10)
	for i := 0; i < total/10; i++ {
		go func(i int) {
			defer wg.Done()
			keys := make([]int, 0, 10)
			for k := i * 10; k < (i+1)*10; k++ {
				keys = append(keys, k)
			}
			b.StoreMulti(keys...)
		}(i)
	}
	wg.Wait()

	keys := b.Values()
	if len(keys) != total {
		t.Fatal()
	}
	for i := range keys {
		if keys[i] != i {
			t.Fatal()
		}
	}
}

func TestBitmap_StoreNX(t *testing.T) {
	key, concurrent := 150, 100
	b := NewBitmap()
	ch := make(chan bool, concurrent)

	wg := sync.WaitGroup{}
	wg.Add(concurrent)
	for i := 0; i < concurrent; i++ {
		go func(i int) {
			defer wg.Done()
			ch <- b.StoreNX(key)
		}(i)
	}
	wg.Wait()

	type check struct {
		sync.Mutex
		m map[bool]int
	}

	m := &check{m: map[bool]int{}}

CHECK:
	for {
		select {
		case c := <-ch:
			m.Lock()
			m.m[c]++
			m.Unlock()
		default:
			break CHECK
		}
	}

	if m.m[true] != 1 || m.m[false] != concurrent-1 || !b.Contain(key) {
		t.Fatal()
	}
}

func TestBitmap_Remove(t *testing.T) {
	total := 100
	b := NewBitmap()

	for i := 0; i < total; i++ {
		b.Store(i)
	}

	for i := 0; i < total*2; i++ {
		b.Remove(i)
	}

	if len(b.Values()) != 0 {
		t.Fatal()
	}
}

func TestBitmap_RemoveMulti(t *testing.T) {
	total := 100
	b := NewBitmap()

	keys := make([]int, 0, total)
	for i := 0; i < total; i++ {
		b.Store(i)
		keys = append(keys, i, i+total)
	}

	b.RemoveMulti(keys...)

	if len(b.Values()) != 0 {
		t.Fatal()
	}
}

func TestBitmap_And(t *testing.T) {
	suit := []struct {
		bitmap1 []int
		bitmap2 []int
		result  []int
	}{
		{
			[]int{},
			[]int{},
			[]int{},
		},
		{
			[]int{},
			[]int{0, 1},
			[]int{},
		},
		{
			[]int{0, 1},
			[]int{},
			[]int{},
		},
		{
			[]int{1, 3, 5, 7, 9},
			[]int{0, 2, 4, 6, 8, 10},
			[]int{},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{0, 2, 4, 6, 8, 10},
			[]int{2, 4, 6, 8, 10},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, _case := range suit {
		b, c := NewBitmap(), NewBitmap()
		b.StoreMulti(_case.bitmap1...)
		c.StoreMulti(_case.bitmap2...)

		b.And(c)
		if !reflect.DeepEqual(b.Values(), _case.result) {
			t.Fatal()
		}

		if !reflect.DeepEqual(c.Values(), _case.bitmap2) {
			t.Fatal()
		}
	}
}

func TestBitmap_Or(t *testing.T) {
	suit := []struct {
		bitmap1 []int
		bitmap2 []int
		result  []int
	}{
		{
			[]int{},
			[]int{},
			[]int{},
		},
		{
			[]int{},
			[]int{0, 1},
			[]int{0, 1},
		},
		{
			[]int{0, 1},
			[]int{},
			[]int{0, 1},
		},
		{
			[]int{1, 3, 5, 7, 9},
			[]int{0, 2, 4, 6, 8, 10},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{0, 2, 4, 6, 8, 10},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100},
		},
	}

	for _, _case := range suit {
		b, c := NewBitmap(), NewBitmap()
		b.StoreMulti(_case.bitmap1...)
		c.StoreMulti(_case.bitmap2...)

		b.Or(c)
		if !reflect.DeepEqual(b.Values(), _case.result) {
			t.Fatal()
		}

		if !reflect.DeepEqual(c.Values(), _case.bitmap2) {
			t.Fatal()
		}
	}
}

func TestBitmap_FromDB(t *testing.T) {
	suit := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 100},
		{},
		{0},
	}

	for _, _case := range suit {
		b := NewBitmap()
		b.StoreMulti(_case...)

		bytes, _ := b.ToDB()
		if err := b.FromDB(bytes); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b.Values(), _case) {
			t.Fatal()
		}
	}
}

func TestBitmap_ToDB(t *testing.T) {
	suit := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 100},
		{},
		{0},
	}

	for _, _case := range suit {
		b := NewBitmap()
		b.StoreMulti(_case...)

		bytes, _ := b.ToDB()
		if err := b.FromDB(bytes); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b.Values(), _case) {
			t.Fatal()
		}
	}
}

func TestBitmap_Clear(t *testing.T) {
	suit := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 100},
		{},
		{0},
	}

	for _, _case := range suit {
		b := NewBitmap()
		b.StoreMulti(_case...)

		b.Clear()
		if len(b.Values()) != 0 {
			t.Fatal()
		}
	}
}

func BenchmarkBitmap_Contain(b *testing.B) {
	total := 100
	bitmap := NewBitmap()
	for i := 0; i < total; i++ {
		bitmap.Store(i)
	}

	for i := 0; i < b.N; i++ {
		bitmap.Contain(i)
	}
}

func BenchmarkBitmap_Store(b *testing.B) {
	bitmap := NewBitmap()
	for i := 0; i < b.N; i++ {
		bitmap.Store(i)
	}
}

func BenchmarkBitmap_StoreMulti(b *testing.B) {
	var cases [][]int
	for i := 0; i < 10; i++ {
		cases = append(cases, make([]int, 0, 10))
		for j := i * 10; j < (i+1)*10; j++ {
			cases[i] = append(cases[i], j)
		}
	}

	bitmap := NewBitmap()
	for i := 0; i < b.N; i++ {
		bitmap.StoreMulti(cases[i%10]...)
	}
}

func BenchmarkBitmap_StoreNX(b *testing.B) {
	bitmap := NewBitmap()
	for i := 0; i < b.N; i++ {
		bitmap.StoreNX(i % 10)
	}
}

func BenchmarkBitmap_Remove(b *testing.B) {
	bitmap := NewBitmap()
	for i := 0; i < b.N; i++ {
		bitmap.Remove(i % 10)
	}
}

func BenchmarkBitmap_RemoveMulti(b *testing.B) {
	var cases [][]int
	for i := 0; i < 10; i++ {
		cases = append(cases, make([]int, 0, 10))
		for j := i * 10; j < (i+1)*10; j++ {
			cases[i] = append(cases[i], j)
		}
	}

	bitmap := NewBitmap()
	for i := 0; i < b.N; i++ {
		bitmap.RemoveMulti(cases[i%10]...)
	}
}

func BenchmarkBitmap_Values(b *testing.B) {
	bimap := NewBitmapWithContainers([]uint64{123, 3123, 4, 51, 5, 6, 61, 616})
	for i := 0; i < b.N; i++ {
		bimap.Values()
	}
}

func BenchmarkBitmap_And(b *testing.B) {
	a, c := NewBitmap(), NewBitmap()
	a.StoreMulti(1, 3, 5, 7, 9, 111, 333, 555, 777, 999)
	c.StoreMulti(2, 4, 6, 8, 0, 111, 333, 555, 777, 999)

	for i := 0; i < b.N; i++ {
		a.And(c)
	}
}

func BenchmarkBitmap_Or(b *testing.B) {
	a, c := NewBitmap(), NewBitmap()
	a.StoreMulti(1, 3, 5, 7, 9, 111, 333, 555, 777, 999)
	c.StoreMulti(2, 4, 6, 8, 0, 111, 333, 555, 777, 999)

	for i := 0; i < b.N; i++ {
		a.Or(c)
	}
}

func BenchmarkBitmap_FromDB(b *testing.B) {
	data, _ := NewBitmapWithContainers([]uint64{12345, 77228, 11111, 333213, 4124, 55161254125125}).ToDB()
	for i := 0; i < b.N; i++ {
		bitmap := NewBitmap()
		_ = bitmap.FromDB(data)
	}
}

func BenchmarkBitmap_ToDB(b *testing.B) {
	bitmap := NewBitmapWithContainers([]uint64{12345, 77228, 11111, 333213, 4124, 55161254125125})
	for i := 0; i < b.N; i++ {
		_, _ = bitmap.ToDB()
	}
}
