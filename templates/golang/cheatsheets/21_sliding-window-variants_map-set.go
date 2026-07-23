// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

// ===================================================================
// 21. Sliding Window Variants
// ===================================================================

// longestSubstringWithoutRepeatedASCIIBytes returns the longest substring containing no
// repeated bytes. It is intentionally ASCII-oriented; arbitrary Unicode text needs
// rune-based indexing. The set always equals bytes in the current window. Time O(n);
// space O(alphabet).
func longestSubstringWithoutRepeatedASCIIBytes(s string) int {
	// Byte-oriented version, appropriate when the problem guarantees ASCII.
	seen := make(map[byte]struct{})
	left, best := 0, 0
	for right := 0; right < len(s); right++ {
		for {
			if _, ok := seen[s[right]]; !ok {
				break
			}
			delete(seen, s[left])
			left++
		}
		seen[s[right]] = struct{}{}
		best = max(best, right-left+1)
	}
	return best
}

// minimumSubarrayLengthWithSumAtLeastTarget requires non-negative values so removing
// from the left can only decrease the sum. It returns 0 when no qualifying contiguous
// subarray exists. Time O(n); space O(1).
func minimumSubarrayLengthWithSumAtLeastTarget(nums []int, target int) int {
	left, total := 0, 0
	best := len(nums) + 1
	for right, value := range nums {
		total += value
		for total >= target {
			best = min(best, right-left+1)
			total -= nums[left]
			left++
		}
	}
	if best == len(nums)+1 {
		return 0
	}
	return best
}

// maximumSumFixedLengthWindow returns the largest sum among contiguous windows of
// exactly k elements, or 0 for invalid k. Add the entering value and remove the leaving
// value to update each window in O(1). Total time O(n).
func maximumSumFixedLengthWindow(nums []int, k int) int {
	if k <= 0 || k > len(nums) {
		return 0
	}
	window := 0
	for _, value := range nums[:k] {
		window += value
	}
	best := window
	for right := k; right < len(nums); right++ {
		window += nums[right] - nums[right-k]
		best = max(best, window)
	}
	return best
}

// countSubarraysWithExactlyKDistinctValues uses: exactly(k) = atMost(k) - atMost(k-1).
// Every subarray with exactly k distinct values appears in the first count and not the
// second. Time O(n); space O(k).
func countSubarraysWithExactlyKDistinctValues(nums []int, k int) int {
	return countSubarraysWithAtMostKDistinctValues(nums, k) -
		countSubarraysWithAtMostKDistinctValues(nums, k-1)
}
