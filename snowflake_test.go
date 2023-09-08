package easy

import (
	"sync"
	"testing"
)

func TestSnowflake_NextId(t *testing.T) {
	s1 := NewSnowflake(func() int64 {
		return 1
	})

	s2 := NewSnowflake(func() int64 {
		return 2
	})

	m := sync.Map{}
	num := 1000

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < num; i++ {
			id := s1.NextId()
			if _, ok := m.Load(id); ok {
				t.Errorf("same id is generated. ")
				t.Fail()
			}
			m.Store(id, "")
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < num; i++ {
			id := s2.NextId()
			if _, ok := m.Load(id); ok {
				t.Errorf("same id is generated. ")
				t.Fail()
			}
			m.Store(id, "")
		}
	}()

	wg.Wait()
}
