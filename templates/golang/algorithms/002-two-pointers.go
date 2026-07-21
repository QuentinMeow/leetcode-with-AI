package algorithms

import "sort"

/*
002 - Two pointers patterns

Use when the input is sorted, when you need to scan from both ends, or when
the answer depends on comparing/compacting elements in one linear pass.
*/

// Variant 1: opposite-direction pointers on a sorted array.
// Example problems: Two Sum II, 3Sum inner loop, valid palindrome.
// Time: O(n)
// Space: O(1)
func TwoSumSorted(nums []int, target int) ([2]int, bool) {
	left, right := 0, len(nums)-1
	for left < right {
		total := nums[left] + nums[right]
		if total == target {
			return [2]int{left, right}, true
		}
		if total < target {
			left++
		} else {
			right--
		}
	}
	return [2]int{}, false
}

// Variant 2: skip duplicates after choosing a value.
// Example problems: 3Sum, 4Sum, unique pair/triplet generation.
// Time: O(n^2) for 3Sum after sorting
// Space: O(1) extra, excluding the output list.
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i, x := range nums {
		if i > 0 && x == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			total := x + nums[left] + nums[right]
			switch {
			case total == 0:
				result = append(result, []int{x, nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			case total < 0:
				left++
			default:
				right--
			}
		}
	}
	return result
}

// Variant 3: same-direction read/write pointers for in-place compaction.
// Example problems: remove duplicates, move zeroes, partition array.
// Time: O(n)
// Space: O(1)
func MoveZeroes(nums []int) {
	write := 0
	for read, x := range nums {
		if x != 0 {
			nums[write], nums[read] = nums[read], nums[write]
			write++
		}
	}
}

// Variant 4: expand around center.
// Example problems: palindromic substrings, longest palindromic substring.
// Time: O(n^2)
// Space: O(1)
func CountPalindromicSubstrings(s string) int {
	expand := func(left, right int) int {
		count := 0
		for left >= 0 && right < len(s) && s[left] == s[right] {
			count++
			left--
			right++
		}
		return count
	}

	total := 0
	for center := range s {
		total += expand(center, center)
		total += expand(center, center+1)
	}
	return total
}
