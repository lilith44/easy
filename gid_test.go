package easy

import (
	"sync"
	"testing"
)

func TestGid(t *testing.T) {
	num := 10
	ch := make(chan uint64, num)
	defer close(ch)

	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()

			ch <- Gid()
		}()
	}
	wg.Wait()

	gidMapping := make(map[uint64]struct{}, len(ch))
	for i := 0; i < num; i++ {
		gid := <-ch
		if _, ok := gidMapping[gid]; ok {
			t.Fatal("two goroutines have same id. ")
		}

		gidMapping[gid] = struct{}{}
	}
}
