// Starvation
// concurrent process dont get enough resources to start da work

// polite and greedy goroutines
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var sharedLock sync.Mutex

const runtime = 1 * time.Second

func main() {
	greedymf := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) < runtime; {
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("greedymf: %d\n", count)
	}
	politemf := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) < runtime; {
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			count++
		}
		fmt.Printf("politemf: %d\n", count)
	}
	wg.Add(2)
	go greedymf()
	go politemf()
	wg.Wait()
}
