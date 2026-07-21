package algorithms

/*
003 - Sliding window patterns

Use when the problem asks about a contiguous subarray/substring and the window
can move monotonically from left to right.
*/

// Variant 1: variable-size window with a validity condition.
// Example problems: longest substring without repeating characters, max window
// after at most k replacements, at most k distinct characters.
// Time: O(n)
// Space: O(k), where k is the number of distinct values in the window.
func LongestSubstringWithoutRepeating(s string) int {
	count := make(map[byte]int)
	left, best := 0, 0
	for right := 0; right < len(s); right++ {
		ch := s[right]
		count[ch]++
		for count[ch] > 1 {
			count[s[left]]--
			left++
		}
		best = max(best, right-left+1)
	}
	return best
}

// Variant 2: fixed-size window.
// Example problems: maximum average subarray, fixed-length anagram checks.
// Time: O(n)
// Space: O(1)
func MaxSumFixedWindow(nums []int, k int) (int, bool) {
	if k <= 0 || k > len(nums) {
		return 0, false
	}

	windowSum := 0
	for _, x := range nums[:k] {
		windowSum += x
	}
	best := windowSum
	for right := k; right < len(nums); right++ {
		windowSum += nums[right] - nums[right-k]
		best = max(best, windowSum)
	}
	return best, true
}

// Variant 3: minimum window that satisfies a requirement.
// Example problems: minimum size subarray sum, minimum window substring.
// Time: O(n)
// Space: O(1) for numeric threshold problems.
func MinSubarrayLenAtLeastTarget(nums []int, target int) int {
	left, windowSum := 0, 0
	best := len(nums) + 1
	for right, x := range nums {
		windowSum += x
		for windowSum >= target {
			best = min(best, right-left+1)
			windowSum -= nums[left]
			left++
		}
	}
	if best == len(nums)+1 {
		return 0
	}
	return best
}

// Variant 4: count subarrays with at most k distinct values.
// Use exactly-k = at_most(k) - at_most(k - 1).
// Time: O(n)
// Space: O(k)
func CountSubarraysAtMostKDistinct(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	count := make(map[int]int)
	left, total := 0, 0
	for right, x := range nums {
		count[x]++
		for len(count) > k {
			y := nums[left]
			count[y]--
			if count[y] == 0 {
				delete(count, y)
			}
			left++
		}
		total += right - left + 1
	}
	return total
}
