// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"context"
	"sync"
	"time"
)

// ===================================================================
// 18. Goroutines, Channels, Context, and Synchronization
// ===================================================================

// squareWorkerPool demonstrates bounded concurrency: workers receive indexed jobs,
// preserve output order through result indices, and stop when context is canceled or
// jobs close. The WaitGroup closes results only after every worker exits.
// Requires: import "context"
// Requires: import "sync"
func squareWorkerPool(ctx context.Context, nums []int, workers int) []int {
	if workers <= 0 {
		return nil
	}

	type job struct {
		index int
		value int
	}
	type result struct {
		index int
		value int
	}

	jobs := make(chan job, len(nums))
	results := make(chan result, len(nums))
	var wg sync.WaitGroup

	worker := func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case current, ok := <-jobs:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case results <- result{
					index: current.index,
					value: current.value * current.value,
				}:
				}
			}
		}
	}

	wg.Add(workers)
	for workerIndex := 0; workerIndex < workers; workerIndex++ {
		go worker()
	}

	go func() {
		defer close(jobs)
		for index, value := range nums {
			select {
			case <-ctx.Done():
				return
			case jobs <- job{index: index, value: value}:
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	squared := make([]int, len(nums))
	for result := range results {
		squared[result.index] = result.value
	}
	return squared
}

// mutexCounterPattern starts workers that share one counter. The mutex makes each
// read-modify-write increment exclusive; the WaitGroup waits for completion. Without
// the lock, increments race and can be lost.
// Requires: import "sync"
func mutexCounterPattern(workers int) int {
	if workers <= 0 {
		return 0
	}
	var mu sync.Mutex
	var wg sync.WaitGroup
	count := 0

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
	wg.Wait()
	return count
}

// bufferedChannelPattern sends all values into a sufficiently buffered channel, closes
// it from the sending side, and ranges until drained.
func bufferedChannelPattern(values []int) []int {
	channel := make(chan int, len(values))
	for _, value := range values {
		channel <- value
	}
	close(channel)

	result := make([]int, 0, len(values))
	for value := range channel {
		result = append(result, value)
	}
	return result
}

// timeoutPattern derives a context with a deadline and races normal completion against
// cancellation. Always call cancel to release timer resources even when work finishes
// first.
// Requires: import "context"
// Requires: import "time"
func timeoutPattern(parent context.Context) error {
	ctx, cancel := context.WithTimeout(parent, 100*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(10 * time.Millisecond):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

/*
Concurrency memory rules:

- Start a goroutine only when you know how it stops.
- The sender normally closes a channel; receivers should not.
- Sending to or closing a closed channel panics.
- Receiving from a closed channel returns buffered values, then zero,false.
- A nil channel blocks forever in send/receive cases.
- Run `go test -race ./...` for concurrent production code.
*/
