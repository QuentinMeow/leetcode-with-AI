// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"slices"
)

// ===================================================================
// 3. Arrays and Slices: Create -> Read -> Update -> Delete
// ===================================================================

/*
Master memory table:

| Operation       | slice               | map                 | set          |
|-----------------|---------------------|---------------------|--------------|
| Empty           | []T{} / make([]T,0) | make(map[K]V)       | map[T]struct |
| Add one         | append              | m[k] = v            | s[x]=struct{}|
| Add many        | append(a, b...)     | loop assignments    | loop         |
| Read            | a[i]                | v, ok := m[k]       | _, ok := s[x]|
| Remove          | reslice/copy trick  | delete(m, k)        | delete(s, x) |
| Empty all       | clear / a = a[:0]   | clear(m)            | clear(s)     |

Key caveats:

- Arrays copy by value. Slices are small headers pointing to backing arrays.
- Map reads return the value type's zero value when a key is absent.
- Use the comma-ok form when zero and missing mean different things.
- append may reuse the backing array or allocate a new one.
- Map iteration order is deliberately unspecified.
- nil slices can be read, ranged over, and appended to.
- Writing to a nil map panics; reading from it is safe.
*/

// sliceCreateReadUpdateDeleteExamples demonstrates creation, reads, copies, insertion,
// deletion, clearing, matrix allocation, and stack operations. Indexed examples require
// at least three input values.
// Requires: import "slices"
func sliceCreateReadUpdateDeleteExamples(n, rows, cols int, nums []int) {
	// The indexed read examples below require len(nums) >= 3.
	x := 2

	// CREATE
	var nilSlice []int
	empty := []int{}
	withLength := make([]int, n)
	withCapacity := make([]int, 0, n)
	fromData := []int{1, 2, 3}
	_ = [3]int{1, 2, 3} // Array, not slice.

	// Matrix create: each row gets its own backing array.
	grid := make([][]int, rows)
	for r := range grid {
		grid[r] = make([]int, cols)
	}

	// READ
	first := nums[0] // Panics if empty.
	last := nums[len(nums)-1]
	middle := nums[1:3] // View sharing the backing array.
	size := len(nums)
	capacity := cap(nums)
	firstX := slices.Index(nums, x) // -1 if absent.
	containsX := slices.Contains(nums, x)

	// COPY
	copy1 := slices.Clone(nums)
	copy2 := append([]int(nil), nums...)
	copy3 := make([]int, len(nums))
	copy(copy3, nums)

	matrixCopy := make([][]int, len(grid))
	for r := range grid {
		matrixCopy[r] = slices.Clone(grid[r])
	}

	// UPDATE / ADD
	withCapacity = append(withCapacity, x)
	withCapacity = append(withCapacity, 4, 5)
	withCapacity = append(withCapacity, nums...)
	nums[0] = 99
	slices.Sort(nums)    // In-place ascending.
	slices.Reverse(nums) // In-place.

	// Insert x before index i.
	i := min(1, len(withCapacity))
	withCapacity = append(withCapacity, 0)
	copy(withCapacity[i+1:], withCapacity[i:])
	withCapacity[i] = x

	// DELETE while preserving order.
	if len(withCapacity) > 0 {
		i = 0
		copy(withCapacity[i:], withCapacity[i+1:])
		withCapacity = withCapacity[:len(withCapacity)-1]
	}

	// DELETE without preserving order, O(1).
	if len(fromData) > 0 {
		i = 0
		fromData[i] = fromData[len(fromData)-1]
		fromData = fromData[:len(fromData)-1]
	}

	clear(withLength)           // Zero elements; length is unchanged.
	withLength = withLength[:0] // Empty but retain capacity.
	nilSlice = nil              // Release reference to backing array.

	// Stack: append and trim the right end.
	stack := []int{1, 2}
	popped := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	_ = []any{
		nilSlice, empty, first, last, middle, size, capacity,
		firstX, containsX, copy1, copy2, copy3, matrixCopy,
		popped, stack,
	}
}

// asciiDigitValue converts one ASCII byte '0' through '9' to its integer value.
// It deliberately does not accept non-ASCII Unicode digits.
func asciiDigitValue(character byte) (int, bool) {
	if character < '0' || character > '9' {
		return 0, false
	}
	return int(character - '0'), true
}

// fillAndReverseSlice demonstrates index-based mutation without a library.
func fillAndReverseSlice(values []int, fillValue int) {
	for index := range values {
		values[index] = fillValue
	}
	for left, right := 0, len(values)-1; left < right; left, right = left+1, right-1 {
		values[left], values[right] = values[right], values[left]
	}
}
