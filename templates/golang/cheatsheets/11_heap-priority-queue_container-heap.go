// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"container/heap"
	"slices"
)

// ===================================================================
// 11. Heap and Priority Queue
// ===================================================================

/*
container/heap does not provide a concrete heap type. A type supplies five
methods: Len, Less, Swap, Push, and Pop. The package uses Len/Less/Swap to
restore heap order. Its heap.Pop first moves the root to the slice's end, then
calls this type's Pop method to remove that final element. That is why Pop below
removes the last slice item rather than index zero.
*/

// IntegerMinHeap is a min-heap. Reverse Less for a max-heap.
type IntegerMinHeap []int

// Len returns the number of stored values for heap.Interface.
func (h IntegerMinHeap) Len() int { return len(h) }

// Less gives smaller integers higher priority; reverse < for a max-heap.
func (h IntegerMinHeap) Less(i, j int) bool { return h[i] < h[j] }

// Swap exchanges positions while the heap package restores ordering.
func (h IntegerMinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push appends the concrete int passed through heap.Push's any parameter.
func (h *IntegerMinHeap) Push(value any) {
	*h = append(*h, value.(int))
}

// Pop removes the final item after the heap package has moved the root there.
func (h *IntegerMinHeap) Pop() any {
	old := *h
	n := len(old)
	value := old[n-1]
	*h = old[:n-1]
	return value
}

// heapOperationsExample demonstrates initialization, insertion, root peek, removal with
// a type assertion, priority repair, arbitrary removal, and clearing.
// Requires: import "container/heap"
// Requires: import "slices"
func heapOperationsExample(nums []int) {
	h := IntegerMinHeap(slices.Clone(nums))
	heap.Init(&h)     // O(n).
	heap.Push(&h, 42) // O(log n).
	smallest := h[0]  // Peek, O(1); requires non-empty heap.
	popped := heap.Pop(&h).(int)
	if len(h) > 0 {
		h[0]++          // Example priority change.
		heap.Fix(&h, 0) // Restore after changing h[0].
		heap.Remove(&h, 0)
	}
	clear(h)
	h = h[:0]
	_ = []any{smallest, popped}
}

// drainQueueByReslicing shows the concise queue form often used in short
// interview solutions. For long-lived queues, prefer the head-index form in
// queueWithHeadIndexExample and compact explicitly when the consumed prefix grows.
// Requires: import "slices"
func drainQueueByReslicing(values []int) []int {
	queue := slices.Clone(values)
	result := make([]int, 0, len(queue))
	for len(queue) > 0 {
		front := queue[0]
		queue[0] = 0 // Drop references early when T contains pointers.
		queue = queue[1:]
		result = append(result, front)
	}
	return result
}

// PriorityItem pairs a payload with an integer priority. PriorityItemHeap is a
// min-heap, so heap.Pop removes the item with the smallest priority first.
// If priorities never change and repeated insertion is unnecessary, sorting a
// slice once can be a simpler interview solution than implementing a heap.
type PriorityItem struct {
	value    int
	priority int
}

type PriorityItemHeap []PriorityItem

// These five methods implement heap.Interface for PriorityItem values.
func (h PriorityItemHeap) Len() int { return len(h) }

// Less defines which of two elements has higher heap priority.
func (h PriorityItemHeap) Less(i, j int) bool { return h[i].priority < h[j].priority }

// Swap exchanges two heap positions while container/heap restores its invariant.
func (h PriorityItemHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push appends a concrete value received through heap.Interface's any parameter.
func (h *PriorityItemHeap) Push(value any) {
	*h = append(*h, value.(PriorityItem))
}

// Pop removes the final slice item after container/heap has moved the root there.
func (h *PriorityItemHeap) Pop() any {
	old := *h
	last := old[len(old)-1]
	*h = old[:len(old)-1]
	return last
}

// kLargestValues keeps only k candidates in a min-heap. The root is the
// smallest retained value, so any extra item can be discarded in O(log k).
// Total time O(n log k), extra space O(k). Result order is unspecified.
// Requires: import "container/heap"
// Requires: import "slices"
func kLargestValues(values []int, k int) []int {
	if k <= 0 {
		return nil
	}
	priorityQueue := &IntegerMinHeap{}
	for _, value := range values {
		heap.Push(priorityQueue, value)
		if priorityQueue.Len() > k {
			heap.Pop(priorityQueue)
		}
	}
	return slices.Clone(*priorityQueue)
}
