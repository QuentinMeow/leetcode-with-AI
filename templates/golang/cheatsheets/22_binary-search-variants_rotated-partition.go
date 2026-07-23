// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

// ===================================================================
// 22. Binary Search Variants
// ===================================================================

// searchRotatedSortedArray finds target in an ascending array rotated at an unknown
// pivot. At least one half around middle is still sorted; range checks decide which
// half can contain target. Returns -1 if absent. Time O(log n), assuming distinct
// values.
func searchRotatedSortedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		}
		if nums[left] <= nums[middle] {
			if nums[left] <= target && target < nums[middle] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else if nums[middle] < target && target <= nums[right] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}

// medianOfTwoSortedArraysUsingPartition binary-searches a cut in the shorter array and
// derives the matching cut in the other. A valid partition has every left value <=
// every right value and places half the combined elements on the left. Time O(log
// min(m,n)); space O(1).
func medianOfTwoSortedArraysUsingPartition(a, b []int) float64 {
	if len(a) > len(b) {
		return medianOfTwoSortedArraysUsingPartition(b, a)
	}
	if len(a)+len(b) == 0 {
		panic("at least one input must be non-empty")
	}

	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	m, n := len(a), len(b)
	half := (m + n + 1) / 2
	left, right := 0, m

	for left <= right {
		aCut := left + (right-left)/2
		bCut := half - aCut

		aLeft, aRight := minInt, maxInt
		bLeft, bRight := minInt, maxInt
		if aCut > 0 {
			aLeft = a[aCut-1]
		}
		if aCut < m {
			aRight = a[aCut]
		}
		if bCut > 0 {
			bLeft = b[bCut-1]
		}
		if bCut < n {
			bRight = b[bCut]
		}

		if aLeft <= bRight && bLeft <= aRight {
			if (m+n)%2 == 1 {
				return float64(max(aLeft, bLeft))
			}
			leftMax := float64(max(aLeft, bLeft))
			rightMin := float64(min(aRight, bRight))
			return (leftMax + rightMin) / 2
		}
		if aLeft > bRight {
			right = aCut - 1
		} else {
			left = aCut + 1
		}
	}
	panic("inputs must be sorted")
}
