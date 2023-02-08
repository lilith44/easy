package easy

import (
	"sync"
	"sync/atomic"
)

// The reentrant mutex (recursive mutex, recursive lock) is particular type of mutual exclusion (mutex) device
// that may be locked multiple times by the same goroutine, without causing a deadlock.
type reentrantMutex struct {
	mutex sync.Mutex

	gid       uint64
	recursion int64
}

// NewReentrantMutex returns a new reentrant mutex.
func NewReentrantMutex() sync.Locker {
	return &reentrantMutex{}
}

// Lock locks the reentrant mutex.
func (rm *reentrantMutex) Lock() {
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
func (rm *reentrantMutex) Unlock() {
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
