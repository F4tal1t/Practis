package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i, delay int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %v started \n", i)
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Printf("Worker %v finished \n", i)

}

func main() {
	var wg sync.WaitGroup

	//for i := 0; i < 3; i++ {
	//	wg.Add(1)
	//	go worker(i+1, &wg)
	//}

	wg.Add(5)
	go worker(1, 3, &wg)
	go worker(2, 2, &wg)
	go worker(3, 3, &wg)
	go worker(4, 4, &wg)
	go worker(5, 5, &wg)

	wg.Wait()
	fmt.Println("kllkdvnlsk")
}
