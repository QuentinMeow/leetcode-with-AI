"""
003 - Sliding window patterns

Use when the problem asks about a contiguous subarray/substring and the window
can move monotonically from left to right.
"""

from collections import defaultdict


# Variant 1: variable-size window with a validity condition.
# Example problems: longest substring without repeating characters, max window
# after at most k replacements, at most k distinct characters.
# Time: O(n)
# Space: O(k), where k is the number of distinct values in the window.
def longest_substring_without_repeating(s: str) -> int:
    count: defaultdict[str, int] = defaultdict(int)
    left = 0
    best = 0

    for right, ch in enumerate(s):
        count[ch] += 1

        while count[ch] > 1:
            count[s[left]] -= 1
            left += 1

        best = max(best, right - left + 1)

    return best


# Variant 2: fixed-size window.
# Example problems: maximum average subarray, fixed-length anagram checks.
# Time: O(n)
# Space: O(1)
def max_sum_fixed_window(nums: list[int], k: int) -> int:
    if k <= 0 or k > len(nums):
        raise ValueError("k must be between 1 and len(nums)")

    window_sum = sum(nums[:k])
    best = window_sum

    for right in range(k, len(nums)):
        window_sum += nums[right] - nums[right - k]
        best = max(best, window_sum)

    return best


# Variant 3: minimum window that satisfies a requirement.
# Example problems: minimum size subarray sum, minimum window substring.
# Time: O(n)
# Space: O(1) for numeric threshold problems.
def min_subarray_len_at_least_target(nums: list[int], target: int) -> int:
    left = 0
    window_sum = 0
    best = float("inf")

    for right, x in enumerate(nums):
        window_sum += x

        while window_sum >= target:
            best = min(best, right - left + 1)
            window_sum -= nums[left]
            left += 1

    return 0 if best == float("inf") else int(best)


# Variant 4: count subarrays with at most k distinct values.
# Use exactly-k = at_most(k) - at_most(k - 1).
# Time: O(n)
# Space: O(k)
def count_subarrays_at_most_k_distinct(nums: list[int], k: int) -> int:
    if k < 0:
        return 0

    count: defaultdict[int, int] = defaultdict(int)
    left = 0
    total = 0

    for right, x in enumerate(nums):
        count[x] += 1

        while len(count) > k:
            y = nums[left]
            count[y] -= 1
            if count[y] == 0:
                del count[y]
            left += 1

        total += right - left + 1

    return total
