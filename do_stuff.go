package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Let's go for a serial program, which is what we're used to. This'll run the DoComplicatedComputeThingy func sequentally.
func DoBigQuery(t int) (time.Duration, int) {

	start := time.Now()

	totalSum := 0
	for i := 0; i < t; i++ {
		// We're casting this to null because we don't actually need to use it here.
		returned := DoComplicatedComputeThingy(i, "serial")
		totalSum += returned
	}

	timeElapsed := time.Since(start)
	return timeElapsed, totalSum

}

// We'll implement concurrent programming here with goroutines, to show the difference in speeds.
func DoBigQueryInARoutine(t int) (time.Duration, int) {

	start := time.Now()

	// Make a channel. This'll be used to receive results back from the other threads (go routines) to the main process (this one).
	// I guess you can think of this like a pub/sub queue. We declare the type in it, and the size of the 'queue'.
	results := make(chan int, t)

	// Create a wait group. This is used to ensure that we don't move on in the program before all of our go routines are complete.
	var wg sync.WaitGroup

	// We add the number of things we're waiting to complete to the wait group.
	wg.Add(t)

	// Obvious loop is obvious.
	for i := 0; i < t; i++ {
		// These funcs are separate threads. It means it can run asynchronously in the background. 
		// "go x" is kinda like "run separate process of go"
		go func(i int) {

			// Defer keyword is used to be like "close this thread when finished. So even if it errors, it'll still close"
			// wg.Done() will iterate the waitgroup by -1.
			defer wg.Done()
			// The <- denotes "pass back our result back to the channel (line 33)"
			results <- DoComplicatedComputeThingy(i, "concurrent")

		}(i)

	}

	go func() {
		// This is the thing that waits for the wait group to be 0. Then it'll close the channel (the pubsub-esque thingy)
		wg.Wait()
		close(results)
	}()

	// We move on here. This is pretty standard

	totalSum := 0
	// range through channel and sum.
	for result := range results {
		totalSum += result
	}

	timeElapsed := time.Since(start)
	return timeElapsed, totalSum
}

func DoComplicatedComputeThingy(i int, s string) int {
	// Let's pretend this does something important that takes a sec.
	time.Sleep(time.Millisecond * 200)
	randNum := rand.IntN(1_000_000_000)
	val := i + 1*randNum
	fmt.Printf("%v query: %v \n", s, val)
	return val
}
