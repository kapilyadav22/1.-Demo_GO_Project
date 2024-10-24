package main

import (
	"fmt"
	"sync"
)

func Worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", i)
	fmt.Printf("Worker %d ended\n", i)
}

func SyncWait() {
	// fmt.Println("Let's Learn sync.WaitGroup")

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1) //Increment the WaitGroup Counter
		go Worker(i, &wg)
	}
	//we are waiting for all the workers to complete their execution
	wg.Wait()
	fmt.Println("Workers task completed")
}

/*
-> sync.WaitGroup is a synchronization primitive in Go that is used to wait for a collection of goroutines to finish their execution.
-> It allows you to coordinate the execution of multiple goroutines and ensure that they all complete before continuing with the rest of the program.

How sync.WaitGroup works :
1. We create a sync.WaitGroup variable to keep track of the number of goroutines we want to wait for
2. For each goroutine we start, we increament the WaitGroup counter using the Add method.
3. Inside each goroutine, we call Done on the WaitGroup to signal that the goroutine has finished its work.
4. Finally, we call Wait on the WaitGroup to block the main goroutine until all other goroutines have called Done.
*/
