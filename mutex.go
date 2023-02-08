package easy

import (
	"sync"
	"sync/atomic"
)

// ReentrantMutex (recursive mutex, recursive lock) is particular type of mutual exclusion (mutex) device
// that may be locked multiple times by the same goroutine, without causing a deadlock.
type ReentrantMutex struct {
	mutex sync.Mutex

	gid       uint64
	recursion int64
}

// NewReentrantMutex returns a new reentrant mutex.
func NewReentrantMutex() *ReentrantMutex {
	return &ReentrantMutex{}
}

// Lock locks the reentrant mutex.
func (rm *ReentrantMutex) Lock() {
	gid := Gid()
	if atomic.LoadUint64(&rm.gid) == gid {
		rm.recursion++
		return
	}

	rm.mutex.Lock()
	rm.gid = gid
	rm.recursion = 1
}

// Unlock unlocks the reentrant mutex.
func (rm *ReentrantMutex) Unlock() {
	gid := Gid()
	if atomic.LoadUint64(&rm.gid) != gid {
		panic("unlock before lock")
	}

	rm.recursion--
	if rm.recursion == 0 {
		rm.gid = 0
		rm.mutex.Unlock()
	}
}
