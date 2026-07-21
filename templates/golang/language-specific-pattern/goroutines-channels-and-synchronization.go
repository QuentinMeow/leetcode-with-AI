// Go goroutine, channel, context, mutex, and WaitGroup patterns.
//
// Concurrency is not an automatic speedup for interview algorithms. Use it when
// the problem has independent work or explicitly tests coordination.
package languagepatterns

import (
	"context"
	"sync"
)

type squareJobPattern struct {
	index int
	value int
}

type squareResultPattern struct {
	index int
	value int
}

// ConcurrentSquaresPattern runs a bounded worker pool and preserves input
// order by carrying indexes. The sender closes jobs; a coordinator closes
// results after every worker exits.
func ConcurrentSquaresPattern(ctx context.Context, nums []int, workers int) ([]int, error) {
	if workers < 1 {
		workers = 1
	}

	jobs := make(chan squareJobPattern)
	results := make(chan squareResultPattern)

	var waitGroup sync.WaitGroup
	waitGroup.Add(workers)
	for worker := 0; worker < workers; worker++ {
		go func() {
			defer waitGroup.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case job, open := <-jobs:
					if !open {
						return
					}
					result := squareResultPattern{
						index: job.index,
						value: job.value * job.value,
					}
					select {
					case results <- result:
					case <-ctx.Done():
						return
					}
				}
			}
		}()
	}

	go func() {
		defer close(jobs)
		for index, value := range nums {
			select {
			case jobs <- squareJobPattern{index: index, value: value}:
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	squares := make([]int, len(nums))
	completed := 0
	for result := range results {
		squares[result.index] = result.value
		completed++
	}
	if completed != len(nums) {
		return nil, ctx.Err()
	}
	return squares, nil
}

// GenerateWithContextPattern returns a receive-only channel. The select lets
// the goroutine stop if the consumer cancels instead of draining every value.
func GenerateWithContextPattern(ctx context.Context, values []int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for _, value := range values {
			select {
			case output <- value:
			case <-ctx.Done():
				return
			}
		}
	}()
	return output
}

// MutexCounterPattern protects a map invariant. Do not copy this value after it
// has been used because it contains a sync.Mutex.
type MutexCounterPattern struct {
	mu     sync.Mutex
	counts map[string]int
}

func NewMutexCounterPattern() *MutexCounterPattern {
	return &MutexCounterPattern{counts: make(map[string]int)}
}

func (counter *MutexCounterPattern) Add(word string) {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	counter.counts[word]++
}

func (counter *MutexCounterPattern) Snapshot() map[string]int {
	counter.mu.Lock()
	defer counter.mu.Unlock()

	result := make(map[string]int, len(counter.counts))
	for word, count := range counter.counts {
		result[word] = count
	}
	return result
}

// CountWordsConcurrentPattern demonstrates WaitGroup plus a mutex-protected
// shared result. Each goroutine builds locally to reduce lock contention.
func CountWordsConcurrentPattern(groups [][]string) map[string]int {
	counts := make(map[string]int)
	var mutex sync.Mutex
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(groups))
	for _, group := range groups {
		group := group // Portable closure capture for Go 1.21.
		go func() {
			defer waitGroup.Done()

			local := make(map[string]int)
			for _, word := range group {
				local[word]++
			}

			mutex.Lock()
			defer mutex.Unlock()
			for word, count := range local {
				counts[word] += count
			}
		}()
	}

	waitGroup.Wait()
	return counts
}

// OnceValuePattern runs initialization exactly once even when many goroutines
// call it concurrently.
func OnceValuePattern(callers int) []int {
	if callers <= 0 {
		return nil
	}
	var once sync.Once
	value := 0
	results := make([]int, callers)

	var waitGroup sync.WaitGroup
	waitGroup.Add(callers)
	for index := 0; index < callers; index++ {
		index := index
		go func() {
			defer waitGroup.Done()
			once.Do(func() { value = 42 })
			results[index] = value
		}()
	}
	waitGroup.Wait()
	return results
}

/*
Channel rules:

- `make(chan T)` is unbuffered: send and receive rendezvous.
- `make(chan T, n)` buffers up to n values before sends block.
- The sending side normally owns close. Closing announces that no more values
  will be sent; it is not required merely to reclaim a channel.
- Receiving from a closed channel yields buffered values, then the zero value
  with ok == false. `for value := range ch` stops after close and drain.
- Sending on a closed channel or closing an already closed channel panics.
- A nil channel blocks forever in send/receive and can disable a select case.

Synchronization rules:

- WaitGroup tracks goroutine completion: Add before launch, Done in a defer,
  Wait in the coordinator. Do not call Add concurrently with a Wait that may
  observe a zero count.
- Mutex protects shared memory. Keep critical sections small and always unlock;
  defer is convenient after a successful Lock.
- Channels communicate ownership/events; mutexes protect shared state. Pick the
  model that makes the invariant easiest to explain.
- Run `go test -race` in a module/package when concurrency code has tests.

Context rules:

- context carries cancellation/deadlines across call boundaries. Put ctx first,
  do not store it in a struct, and call a returned cancel function.
- A goroutine must have a defined exit path. Blocked sends, forgotten closes,
  and consumers that stop early are common leak sources.
*/
