// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"container/heap"
)

// ===================================================================
// 27. Heap Applications
// ===================================================================

// SortedArrayCursor identifies one current value and where its successor lives.
type SortedArrayCursor struct {
	value      int
	arrayIndex int
	valueIndex int
}

// SortedArrayCursorMinHeap exposes the smallest current value across all arrays.
type SortedArrayCursorMinHeap []SortedArrayCursor

// Len reports the number of elements required by heap.Interface.
func (h SortedArrayCursorMinHeap) Len() int { return len(h) }

// Less defines which of two elements has higher heap priority.
func (h SortedArrayCursorMinHeap) Less(i, j int) bool { return h[i].value < h[j].value }

// Swap exchanges two heap positions while container/heap restores its invariant.
func (h SortedArrayCursorMinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push appends a concrete value received through heap.Interface's any parameter.
func (h *SortedArrayCursorMinHeap) Push(value any) {
	*h = append(*h, value.(SortedArrayCursor))
}

// Pop removes the final slice item after container/heap has moved the root there.
func (h *SortedArrayCursorMinHeap) Pop() any {
	old := *h
	value := old[len(old)-1]
	*h = old[:len(old)-1]
	return value
}

// mergeSortedArraysUsingMinHeap performs a k-way merge. The min-heap stores the next
// unused value from each sorted input; after removing one cursor, its successor from
// the same array is inserted. Time O(totalValues log k); space O(k).
// Requires: import "container/heap"
func mergeSortedArraysUsingMinHeap(arrays [][]int) []int {
	priorityQueue := &SortedArrayCursorMinHeap{}
	totalLength := 0
	for arrayIndex, values := range arrays {
		totalLength += len(values)
		if len(values) > 0 {
			heap.Push(priorityQueue, SortedArrayCursor{
				value:      values[0],
				arrayIndex: arrayIndex,
			})
		}
	}

	answer := make([]int, 0, totalLength)
	for priorityQueue.Len() > 0 {
		current := heap.Pop(priorityQueue).(SortedArrayCursor)
		answer = append(answer, current.value)
		nextIndex := current.valueIndex + 1
		values := arrays[current.arrayIndex]
		if nextIndex < len(values) {
			heap.Push(priorityQueue, SortedArrayCursor{
				value:      values[nextIndex],
				arrayIndex: current.arrayIndex,
				valueIndex: nextIndex,
			})
		}
	}
	return answer
}
