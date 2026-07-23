// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"slices"
)

// ===================================================================
// 25. Depth-First Search and Backtracking Variants
// ===================================================================

// allPermutationsUsingBacktracking returns every ordering of nums. used[index] records
// which positions are already in the current path; choose, recurse, and unchoose
// restores the invariant. Time/output O(n*n!).
// Requires: import "slices"
func allPermutationsUsingBacktracking(nums []int) [][]int {
	answer := make([][]int, 0)
	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			answer = append(answer, slices.Clone(path))
			return
		}
		for index, value := range nums {
			if used[index] {
				continue
			}
			used[index] = true
			path = append(path, value)
			backtrack()
			path = path[:len(path)-1]
			used[index] = false
		}
	}

	backtrack()
	return answer
}

// combinationsSummingToTargetWithReuse returns combinations totaling target where each positive candidate may
// be reused. Sorting permits early stopping; recursing with the same index allows
// reuse, while start prevents reordered duplicates. It mutates candidates.
// Requires: import "slices"
func combinationsSummingToTargetWithReuse(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	answer := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(int, int)
	backtrack = func(start, remaining int) {
		if remaining == 0 {
			answer = append(answer, slices.Clone(path))
			return
		}
		for index := start; index < len(candidates); index++ {
			value := candidates[index]
			if value > remaining {
				break
			}
			path = append(path, value)
			backtrack(index, remaining-value)
			path = path[:len(path)-1]
		}
	}

	backtrack(0, target)
	return answer
}
