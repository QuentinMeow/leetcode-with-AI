package algorithms

/*
004 - Binary search patterns

Use when the input is sorted or when the answer can be checked with a monotonic
predicate: false false false true true true.
*/

// Variant 1: lower bound, the safest default template.
// Finds the first index i where nums[i] >= target.
// Example problems: search insert position, first occurrence, bisect_left.
// Time: O(log n)
// Space: O(1)
func LowerBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// Variant 2: exact search in a sorted array.
// Example problems: binary search, lookup in sorted list.
// Time: O(log n)
// Space: O(1)
func BinarySearch(nums []int, target int) int {
	index := LowerBound(nums, target)
	if index < len(nums) && nums[index] == target {
		return index
	}
	return -1
}

// Variant 3: answer-space binary search with a monotonic feasibility function.
// Example problems: Koko Eating Bananas, ship packages, split array largest sum.
// Time: O(log(range) * cost(can))
// Space: depends on can().
func FirstFeasible(low, high int, can func(int) bool) int {
	for low < high {
		mid := low + (high-low)/2
		if can(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

// Variant 4: binary search on a rotated sorted array.
// Example problems: search in rotated sorted array, find minimum.
// Time: O(log n)
// Space: O(1)
func SearchRotated(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < target && target <= nums[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
