package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int = 0

func aaCount(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println("counter =", counter)
	lock.Unlock()
}

func aamain() {

	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go aaCount(lock)
	}

	for {
		lock.Lock()

		c := counter

		lock.Unlock()

		runtime.Gosched()

		if c >= 10 {
			break
		}
	}
}
