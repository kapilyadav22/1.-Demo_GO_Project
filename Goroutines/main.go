package main

import (
	"fmt"
	"time"
)

func HeyKapil() {
	fmt.Println("Hey This is Kapil")
	time.Sleep(2200 * time.Millisecond)
	fmt.Println("Woke up after sleep")
}

func HelloRahul() {
	fmt.Println("Hello Rahul")
}

func main() {
	// fmt.Println("Lets Learn GoRoutines")
	// go HeyKapil()
	// HelloRahul()
	// time.Sleep(2300 * time.Millisecond)
	SyncWait()
}

/*
Concurrency vs Parallelism
-> Concurrency : When two or more Control flows (threads) of execution share one or more CPUs.
   * In such cases, the CPU scheduler is responsible for deciding when each thread gets to execute and on which CPU
   for example: even if there is only one CPU, but two or more threads share the CPU, then its considered concurrent execution.

-> Parallelism :
	* Is a subset of concurrency
	* its when two or more threads execute at the same real time on two or more CPUs
	* for example, three threads executing on three different CPUs simultaneously.

	Note: We use the term "thread" above loosely to refer to either threads or processes.

-> GoRoutines in Golang:
 -> GoRoutines are a key feature of the GO programming language that allow you to run functions concurrently or in parallel,
 with other parts of our program.
 -> Imagine we have a task that can be broken down into smaller sub-tasks that can be executed independently. Instead of waiting for each sub-task
    to finish before starting the next one,
-> we can use goroutines to execute them concurrently. This can greatly improve the performance of our program, especially when dealing with tasks that involve I/O
	operations or other types of blocking operations.
-> use go keyword to put that function in concurrent space.
-> if both functions will use go subroutine, then task schedular will decide which function to execute first


-> We can wait in main function for some time, to check whether all the functions executed or not
*/
