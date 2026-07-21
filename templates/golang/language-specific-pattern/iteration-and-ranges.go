// Go iteration and range semantics for coding interviews.
//
// range is concise, but its index/value meanings depend on the operand. The
// range value is a copy; index the original slice when mutation or identity is
// required.
package languagepatterns

// IndexAndValueRange returns the first matching index.
func IndexAndValueRange(nums []int, target int) int {
	for index, value := range nums {
		if value == target {
			return index
		}
	}
	return -1
}

// MutateByIndexRange updates slice elements. Writing `value *= 2` in a
// `for _, value := range nums` loop would modify only the copied loop value.
func MutateByIndexRange(nums []int) {
	for index := range nums {
		nums[index] *= 2
	}
}

// ElementPointersRange takes addresses of actual slice elements, not the
// address of the reused range value under Go 1.21 semantics.
func ElementPointersRange(nums []int) []*int {
	pointers := make([]*int, 0, len(nums))
	for index := range nums {
		pointers = append(pointers, &nums[index])
	}
	return pointers
}

// ReverseIndexRange uses a classic loop because range always moves forward.
func ReverseIndexRange(nums []int) []int {
	reversed := make([]int, 0, len(nums))
	for index := len(nums) - 1; index >= 0; index-- {
		reversed = append(reversed, nums[index])
	}
	return reversed
}

// AdjacentPairsRange starts at 1 so index-1 is always valid.
func AdjacentPairsRange(nums []int) [][2]int {
	pairs := make([][2]int, 0, max(0, len(nums)-1))
	for index := 1; index < len(nums); index++ {
		pairs = append(pairs, [2]int{nums[index-1], nums[index]})
	}
	return pairs
}

// ParallelSlicesRange checks the length explicitly. Go has no zip built-in.
func ParallelSlicesRange(left, right []int) (int, bool) {
	if len(left) != len(right) {
		return 0, false
	}

	dotProduct := 0
	for index, value := range left {
		dotProduct += value * right[index]
	}
	return dotProduct, true
}

// MatrixRange uses the row's actual length, which also supports ragged grids.
func MatrixRange(grid [][]int) int {
	total := 0
	for row := range grid {
		for col := range grid[row] {
			total += grid[row][col]
		}
	}
	return total
}

// MapRange demonstrates key/value iteration. Never use the resulting order as
// part of an answer unless keys are collected and sorted first.
func MapRange(counts map[string]int) int {
	total := 0
	for _, count := range counts {
		total += count
	}
	return total
}

// RemoveInPlaceRange uses a write index instead of deleting while indexes are
// moving. It keeps the original order and reuses the backing array.
func RemoveInPlaceRange(nums []int, remove int) []int {
	write := 0
	for _, value := range nums {
		if value == remove {
			continue
		}
		nums[write] = value
		write++
	}
	clear(nums[write:])
	return nums[:write]
}

// CaptureValuesRange explicitly copies the loop value before creating each
// closure. This is the portable spelling across Go 1.21 and Go 1.22 semantics.
func CaptureValuesRange(nums []int) []func() int {
	functions := make([]func() int, 0, len(nums))
	for _, value := range nums {
		value := value
		functions = append(functions, func() int { return value })
	}
	return functions
}

/*
Range operand meanings:

- slice/array: index, copied element value.
- string: byte offset, decoded rune. Invalid UTF-8 produces utf8.RuneError.
- map: key, value in unspecified order.
- channel: received values until the channel is closed.

Common one-value forms:

	for index := range nums { ... }  // index
	for key := range counts { ... }  // key
	for value := range channel { ... } // received value

Go 1.21 closure warning:

- Before Go 1.22, freshly declared range variables are reused across
  iterations. A closure may therefore observe the final value. Copy inside the
  loop (`value := value`) before launching a goroutine or saving a closure.
- This file uses the explicit copy even though Go 1.22 changed the language
  semantics, because the target is Go 1.21 and the spelling is portable.

Mutation guidance:

- Mutating existing slice elements through nums[index] is safe during range.
- Appending to or reslicing the same slice during range makes reasoning about
  the original length/backing array difficult. Prefer a separate output slice
  or a controlled read/write index.
*/
