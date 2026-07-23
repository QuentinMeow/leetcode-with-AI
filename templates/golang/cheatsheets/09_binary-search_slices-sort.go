// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"slices"
	"sort"
)

// ===================================================================
// 9. Binary Search
// ===================================================================

// binarySearchLibrary returns the insertion index and exact-match flag from
// slices.BinarySearch; input must be ascending.
// Requires: import "slices"
func binarySearchLibrary(sortedNums []int, target int) (int, bool) {
	index, found := slices.BinarySearch(sortedNums, target)
	return index, found
}

// firstIndexAtLeastTargetUsingSortSearch returns the first index whose value is greater
// than or equal to target, possibly len(input). Input must be ascending.
// Requires: import "sort"
func firstIndexAtLeastTargetUsingSortSearch(sortedNums []int, target int) int {
	// First index i where sortedNums[i] >= target.
	return sort.Search(len(sortedNums), func(i int) bool {
		return sortedNums[i] >= target
	})
}

// firstIndexGreaterThanTargetUsingSortSearch returns the first index whose value is
// strictly greater than target, possibly len(input). Input must be ascending.
// Requires: import "sort"
func firstIndexGreaterThanTargetUsingSortSearch(sortedNums []int, target int) int {
	// First index i where sortedNums[i] > target.
	return sort.Search(len(sortedNums), func(i int) bool {
		return sortedNums[i] > target
	})
}

// firstIndexAtLeastTarget manually implements a half-open [left,right) lower-bound
// search and returns the first value greater than or equal to target.
func firstIndexAtLeastTarget(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		middle := left + (right-left)/2
		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return left
}

// Count in sorted data by subtracting the first index > target from the first
// index >= target. This is O(log n), unlike scanning after one match.
// Requires via helper: import "sort"
func countOccurrences(sortedNums []int, target int) int {
	return firstIndexGreaterThanTargetUsingSortSearch(sortedNums, target) -
		firstIndexAtLeastTargetUsingSortSearch(sortedNums, target)
}

// exactIndexBinarySearchClosedInterval searches [left, right], including both
// endpoints. It returns -1 when target is absent. Time O(log n), space O(1).
func exactIndexBinarySearchClosedInterval(sortedValues []int, target int) int {
	left, right := 0, len(sortedValues)-1
	for left <= right {
		middle := left + (right-left)/2
		if sortedValues[middle] == target {
			return middle
		}
		if sortedValues[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}

// legacySearchIntsExample returns the insertion index used by sort.SearchInts.
// Requires: import "sort"
func legacySearchIntsExample(sortedValues []int, target int) (index int, found bool) {
	index = sort.SearchInts(sortedValues, target)
	return index, index < len(sortedValues) && sortedValues[index] == target
}
