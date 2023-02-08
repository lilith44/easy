package easy

import (
	"testing"
)

func TestNewReentrantMutex(t *testing.T) {
	mutex := NewReentrantMutex()
	m := make(map[int]struct{})
	per := 50
	go func() {
		for i := 0; i < per; i++ {
			mutex.Lock()
			m[i] = struct{}{}
			mutex.Unlock()
		}
	}()

	go func() {
		for i := per; i < per*2; i++ {
			mutex.Lock()
			m[i] = struct{}{}
			mutex.Unlock()
		}
	}()

	go func() {
		for i := per * 2; i < per*3; i++ {
			mutex.Lock()
			m[i] = struct{}{}
			mutex.Unlock()
		}
	}()
}
