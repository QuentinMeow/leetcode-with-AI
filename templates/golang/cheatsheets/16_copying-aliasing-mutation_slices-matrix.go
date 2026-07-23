// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"slices"
)

// ===================================================================
// 16. Copying, Aliasing, and Mutation
// ===================================================================

// copyAndAliasing contrasts shared slice headers, independent one-dimensional copies,
// shallow matrix copies, deep row copies, and saved backtracking snapshots.
// Requires: import "slices"
func copyAndAliasing(grid [][]int, nums []int) {
	alias := nums // Shared backing array.
	shallow1 := slices.Clone(nums)
	shallow2 := append([]int(nil), nums...)

	matrixShallow := slices.Clone(grid) // Rows still shared.
	matrixCopy := make([][]int, len(grid))
	for r := range grid {
		matrixCopy[r] = slices.Clone(grid[r])
	}

	path := []int{1, 2}
	result := make([][]int, 0)
	result = append(result, slices.Clone(path)) // Save snapshot.

	_ = []any{
		alias, shallow1, shallow2, matrixShallow, matrixCopy, result,
	}
}

// appendCaveat returns the new slice header because append may allocate a different
// backing array.
func appendCaveat(nums []int) []int {
	// The returned slice must be assigned by the caller because append may
	// allocate a new backing array and change the slice header.
	return append(nums, 42)
}

// mutateElements doubles elements through their indices, so aliases sharing the backing
// array observe the writes.
func mutateElements(nums []int) {
	// Element writes are visible through every slice sharing this array.
	for i := range nums {
		nums[i] *= 2
	}
}

// rebindLocally demonstrates that replacing a local slice header does not replace the
// caller's header.
func rebindLocally(nums []int) {
	// This changes only the local slice header, not the caller's header.
	nums = append(nums, 42)
	_ = nums
}
