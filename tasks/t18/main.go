package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// So we need mutex for this.
// Let's lock it while  inc and get
// Mutexes in go are also providing order
// of lock-unlock queries
type ConcurrentCounter struct {
	value int
	rwm   *sync.Mutex
}

func NewConcurrentCounter() ConcurrentCounter {
	cc := new(ConcurrentCounter)
	cc.value = 0
	cc.rwm = new(sync.Mutex)
	return *cc
}

func (cc *ConcurrentCounter) Inc() {
	cc.rwm.Lock()
	defer cc.rwm.Unlock()
	cc.value++
}

func (cc *ConcurrentCounter) Get() int {
	cc.rwm.Lock()
	defer cc.rwm.Unlock()
	return cc.value
}

func main() {
	rand.Seed(time.Now().Unix())
	cc := NewConcurrentCounter()
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				cc.Inc()
			}
		}()
		// To see progress
		fmt.Println(cc.Get())
	}

	wg.Wait()
	fmt.Println(cc.Get())
}
