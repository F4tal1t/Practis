// Livelock
// eg , 2 people approaching each other moving at each other , they move to provide space for each other at the narrow coridor.
// left right left rightr both dont go anywhere

// Current Code not working maybe cuz 9 year old book reference
// Nvm it worked
package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	cadence := sync.NewCond(&sync.Mutex{}) // Onesly i donno whats this function will come back later
	go func() {
		for range time.Tick(2 * time.Millisecond) {
			cadence.Broadcast() // Broadcast will wake up all the goroutines waiting on this condition variable
		}
	}()
	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}
	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		// Implementation for trying a direction
		fmt.Fprintf(out, "%v", dirName)
		atomic.AddInt32(dir, 1)         // Atomically increment the direction counter
		takeStep()                      // Wait for the cadence to allow the next step
		if atomic.LoadInt32(dir) == 1 { // Check if this goroutine is the only one trying this direction
			fmt.Fprintf(out, " . Success! ")
			return true
		}
		takeStep()               // Wait for the cadence to allow the next step
		atomic.AddInt32(dir, -1) // Atomically decrement the direction counter
		return false
	}

	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool {
		return tryDir("left", &left, out)
	}
	tryRight := func(out *bytes.Buffer) bool {
		return tryDir("right", &right, out)
	}
	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() { fmt.Println(out.String()) }()
		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, " . Flips the table ;.;")

	}
	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "BRUh")
	peopleInHallway.Wait()
}
