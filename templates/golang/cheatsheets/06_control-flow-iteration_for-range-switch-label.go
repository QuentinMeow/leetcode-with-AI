// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"slices"
)

// ===================================================================
// 6. Control Flow and Iteration
// ===================================================================

// Go uses for for every loop shape: `for condition {}` replaces while, and
// `for {}` is an intentional infinite loop exited with break or return. Switch
// cases stop automatically, and only explicit bool expressions can be conditions.

// findMatrixValue demonstrates the one Go control-flow tool that is easy to
// forget in interviews: a label can leave nested loops in one step.
func findMatrixValue(matrix [][]int, target int) (row, col int, found bool) {
	row, col = -1, -1
search:
	for r := range matrix {
		for c, value := range matrix[r] {
			if value == target {
				row, col, found = r, c, true
				break search
			}
		}
	}
	return row, col, found
}

// iterationPatterns compares index/value ranges, reverse iteration, nested matrix
// loops, range-copy mutation, closure capture, and deterministic map-key ordering.
// Requires: import "slices"
func iterationPatterns(nums []int, matrix [][]int) {
	for index, value := range nums {
		_, _ = index, value
	}

	for index := range nums {
		_ = index
	}

	for index := len(nums) - 1; index >= 0; index-- {
		_ = nums[index]
	}

	for row := range matrix {
		for col := range matrix[row] {
			_ = matrix[row][col]
		}
	}

	// range copies each element into value. Mutate through nums[i].
	for i := range nums {
		nums[i] *= 2
	}

	// Before Go 1.22, closures could accidentally capture the reused range
	// variable. Rebinding remains clear and portable interview code.
	functions := make([]func() int, 0, len(nums))
	for _, value := range nums {
		value := value
		functions = append(functions, func() int { return value })
	}

	// Maps are unordered. Sort keys for deterministic output.
	counts := map[string]int{"b": 2, "a": 1}
	keys := make([]string, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	_ = []any{functions, keys}
}

// controlFlowExamples collects loop and switch forms that are easy to forget.
func controlFlowExamples(value, limit int) string {
	for index := 0; index < limit; index += 2 { // Classic loop with step 2.
		if index == value {
			break
		}
	}
	for value < 0 { // Go's while-style loop.
		value++
	}
	if parsed := value * 2; parsed > limit { // Initializer is scoped to if/else.
		return "large"
	}
	switch { // A switch without a value replaces a long if/else chain.
	case value < 0:
		return "negative"
	case value == 0:
		return "zero"
	default:
		return "positive"
	}
}
