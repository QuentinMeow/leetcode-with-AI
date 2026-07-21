// Go data-structure patterns for coding interviews.
//
// Go's core interview containers are arrays, slices, and maps. Stacks and
// queues are normally slices; sets are normally maps with empty struct values.
package languagepatterns

// SlicePatternsData demonstrates creation, copying, appending, and deletion.
func SlicePatternsData(nums []int) (copied, inserted, deleted []int) {
	// make([]int, length, capacity): length elements are immediately indexable.
	buffer := make([]int, 0, len(nums)+1)
	buffer = append(buffer, nums...)

	// copy duplicates elements. Plain assignment would copy only the slice
	// header, leaving both slices backed by the same array.
	copied = make([]int, len(nums))
	copy(copied, nums)

	// Insert before index 1 without relying on a particular spare capacity.
	inserted = append([]int(nil), nums...)
	index := min(1, len(inserted))
	inserted = append(inserted, 0)
	copy(inserted[index+1:], inserted[index:])
	inserted[index] = 99

	// Delete index 0 while preserving order. Clear the now-unused tail slot
	// when element references should become collectible promptly.
	deleted = append([]int(nil), buffer...)
	if len(deleted) > 0 {
		copy(deleted, deleted[1:])
		deleted[len(deleted)-1] = 0
		deleted = deleted[:len(deleted)-1]
	}

	return copied, inserted, deleted
}

// MatrixData allocates each row separately. This is the conventional shape for
// a rectangular interview grid.
func MatrixData(rows, cols int) [][]int {
	grid := make([][]int, rows)
	for row := range grid {
		grid[row] = make([]int, cols)
	}
	return grid
}

// SliceViewData demonstrates that slicing does not copy. Changes through view
// are visible in nums until an append moves view to a different backing array.
func SliceViewData(nums []int, start, end int) []int {
	view := nums[start:end]
	if len(view) > 0 {
		view[0] = -1
	}
	return view
}

// StackData uses the right end of a slice, where push/pop are amortized O(1).
func StackData(values []int) []int {
	stack := make([]int, 0, len(values))
	for _, value := range values {
		stack = append(stack, value)
	}

	result := make([]int, 0, len(stack))
	for len(stack) > 0 {
		last := len(stack) - 1
		result = append(result, stack[last])
		stack = stack[:last]
	}
	return result
}

// QueueData uses a head index so removing the front does not shift O(n)
// elements. For a long-lived queue, compact occasionally so old entries and a
// large backing array are not retained forever.
func QueueData(values []int) []int {
	queue := append([]int(nil), values...)
	order := make([]int, 0, len(queue))
	for head := 0; head < len(queue); head++ {
		order = append(order, queue[head])
	}
	return order
}

// BFSData shows the standard slice queue plus a map-backed visited set.
func BFSData(start int, graph map[int][]int) []int {
	queue := []int{start}
	seen := map[int]struct{}{start: {}}
	order := make([]int, 0, len(graph))

	for head := 0; head < len(queue); head++ {
		node := queue[head]
		order = append(order, node)
		for _, neighbor := range graph[node] {
			if _, exists := seen[neighbor]; exists {
				continue
			}
			seen[neighbor] = struct{}{}
			queue = append(queue, neighbor)
		}
	}
	return order
}

// FrequencyData demonstrates map zero values and the comma-ok lookup.
func FrequencyData(words []string, query string) (map[string]int, int, bool) {
	counts := make(map[string]int, len(words))
	for _, word := range words {
		counts[word]++ // A missing int value reads as 0.
	}

	count, present := counts[query]
	return counts, count, present
}

// GroupData relies on append working with a nil slice returned by a missing map
// key. No separate "initialize group" branch is needed.
func GroupData(words []string) map[byte][]string {
	groups := make(map[byte][]string)
	for _, word := range words {
		if word == "" {
			continue
		}
		key := word[0]
		groups[key] = append(groups[key], word)
	}
	return groups
}

// SetData uses map[int]struct{} because struct{} occupies zero bytes as a map
// value. map[int]bool can be convenient when false has useful meaning.
func SetData(nums []int) (map[int]struct{}, bool) {
	seen := make(map[int]struct{}, len(nums))
	hasDuplicate := false
	for _, value := range nums {
		if _, exists := seen[value]; exists {
			hasDuplicate = true
		}
		seen[value] = struct{}{}
	}
	return seen, hasDuplicate
}

// ArrayValueData shows that an array's length is part of its type and assigning
// an array copies every element.
func ArrayValueData(input [3]int) ([3]int, [3]int) {
	copied := input
	copied[0] = -1
	return input, copied
}

/*
Container choice guide:

- Random access, stack, dense DP table: []T.
- FIFO queue: []T with a head index for interview-sized workloads.
- Key to value, frequency, grouping: map[K]V.
- Membership/visited: map[T]struct{}.
- Fixed small value whose length is meaningful: [N]T.

Important semantics:

- A slice is a descriptor (pointer, length, capacity), not an owning list.
  Assignment shares its backing array; copy or append to a fresh slice to clone.
- append returns the resulting slice and may allocate. Always retain its return
  value. A callee that appends cannot update the caller's slice header unless it
  returns the slice or receives *[]T.
- nil slices are safe to len, range, and append. Reading a nil map is safe, but
  assigning a key in a nil map panics.
- Map keys must be comparable. Slices, maps, and functions cannot be keys.
- Map iteration order is unspecified and can differ between iterations.
- delete on a missing key is safe. clear(map) removes all entries in Go 1.21.
*/
