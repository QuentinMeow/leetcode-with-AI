package algorithms

import (
	"container/heap"
	"sort"
)

/*
010 - Heap / priority queue patterns

Use when repeatedly needing the smallest/largest item, top k items, or merging
sorted streams. Go's container/heap is a min-heap when Less uses <.
*/

type intMinHeap []int

func (h intMinHeap) Len() int           { return len(h) }
func (h intMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intMinHeap) Push(value any)    { *h = append(*h, value.(int)) }
func (h *intMinHeap) Pop() any {
	old := *h
	last := len(old) - 1
	value := old[last]
	*h = old[:last]
	return value
}

// Variant 1: top k with a bounded min-heap.
// Example problems: kth largest element, top k scores, streaming kth largest.
// Time: O(n log k)
// Space: O(k)
func KLargest(nums []int, k int) []int {
	values := &intMinHeap{}
	heap.Init(values)
	for _, x := range nums {
		heap.Push(values, x)
		if values.Len() > k {
			heap.Pop(values)
		}
	}
	result := append([]int(nil), (*values)...)
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result
}

// Variant 2: max-heap by negating priority.
// Example problems: last stone weight, repeatedly take largest.
// Time: O(n log n)
// Space: O(n)
func RepeatedlyTakeLargest(nums []int) []int {
	values := make(intMinHeap, len(nums))
	for i, x := range nums {
		values[i] = -x
	}
	heap.Init(&values)
	order := make([]int, 0, len(nums))
	for values.Len() > 0 {
		order = append(order, -heap.Pop(&values).(int))
	}
	return order
}

type frequencyEntry struct {
	frequency int
	number    int
}

type frequencyMinHeap []frequencyEntry

func (h frequencyMinHeap) Len() int      { return len(h) }
func (h frequencyMinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h frequencyMinHeap) Less(i, j int) bool {
	if h[i].frequency == h[j].frequency {
		return h[i].number < h[j].number
	}
	return h[i].frequency < h[j].frequency
}
func (h *frequencyMinHeap) Push(value any) {
	*h = append(*h, value.(frequencyEntry))
}
func (h *frequencyMinHeap) Pop() any {
	old := *h
	last := len(old) - 1
	value := old[last]
	*h = old[:last]
	return value
}

// Variant 3: heap of structs for custom priority and deterministic tie-breaking.
// Example problems: top k frequent elements, task scheduler variations.
// Time: O(n log k)
// Space: O(n)
func TopKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, x := range nums {
		counts[x]++
	}
	values := &frequencyMinHeap{}
	heap.Init(values)
	for number, frequency := range counts {
		heap.Push(values, frequencyEntry{frequency, number})
		if values.Len() > k {
			heap.Pop(values)
		}
	}

	entries := append([]frequencyEntry(nil), (*values)...)
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].frequency == entries[j].frequency {
			return entries[i].number > entries[j].number
		}
		return entries[i].frequency > entries[j].frequency
	})
	result := make([]int, len(entries))
	for i, entry := range entries {
		result[i] = entry.number
	}
	return result
}

type mergeEntry struct {
	value        int
	arrayIndex   int
	elementIndex int
}

type mergeMinHeap []mergeEntry

func (h mergeMinHeap) Len() int      { return len(h) }
func (h mergeMinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h mergeMinHeap) Less(i, j int) bool {
	if h[i].value != h[j].value {
		return h[i].value < h[j].value
	}
	if h[i].arrayIndex != h[j].arrayIndex {
		return h[i].arrayIndex < h[j].arrayIndex
	}
	return h[i].elementIndex < h[j].elementIndex
}
func (h *mergeMinHeap) Push(value any) {
	*h = append(*h, value.(mergeEntry))
}
func (h *mergeMinHeap) Pop() any {
	old := *h
	last := len(old) - 1
	value := old[last]
	*h = old[:last]
	return value
}

// Variant 4: merge k sorted lists/arrays.
// Example problems: merge k sorted lists, kth smallest in sorted matrix.
// Time: O(total_items * log k)
// Space: O(k)
func MergeSortedArrays(arrays [][]int) []int {
	values := &mergeMinHeap{}
	heap.Init(values)
	for arrayIndex, array := range arrays {
		if len(array) > 0 {
			heap.Push(values, mergeEntry{array[0], arrayIndex, 0})
		}
	}

	result := make([]int, 0)
	for values.Len() > 0 {
		current := heap.Pop(values).(mergeEntry)
		result = append(result, current.value)
		nextIndex := current.elementIndex + 1
		if nextIndex < len(arrays[current.arrayIndex]) {
			heap.Push(values, mergeEntry{
				arrays[current.arrayIndex][nextIndex],
				current.arrayIndex,
				nextIndex,
			})
		}
	}
	return result
}
