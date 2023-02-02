package easy

import (
	"reflect"
	"runtime/debug"
	"sync"
)

const defaultPoolCap = 20

// Fill invokes f() and stores the output into slice result for each element in slice origin concurrently.
func Fill(origin any, result any, f func(any) any, poolCaps ...int) {
	c := defaultPoolCap
	if len(poolCaps) != 0 && poolCaps[0] > 0 {
		c = poolCaps[0]
	}

	o := reflect.ValueOf(origin)
	r := reflect.ValueOf(result)

	wg := sync.WaitGroup{}
	wg.Add(o.Len())
	ch := make(chan struct{}, c)
	defer close(ch)

	for i := 0; i < o.Len(); i++ {
		go func(i int) {
			defer func() {
				if r := recover(); r != nil {
					debug.PrintStack()
				}

				wg.Done()
			}()

			ch <- struct{}{}
			defer func() { <-ch }()

			r.Index(i).Set(reflect.ValueOf(f(o.Index(i).Interface())))
		}(i)
	}
	wg.Wait()
}
