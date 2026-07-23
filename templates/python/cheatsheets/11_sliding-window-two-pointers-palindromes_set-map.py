"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

import math

# ====================================================================
# 11. Sliding Windows, Two Pointers, and Palindromes
# ====================================================================

# Returns the longest substring with unique characters using a sliding set window. Time O(n).
def longest_substring_without_repeated_characters(s: str) -> int:
    seen: set[str] = set()
    left = best = 0
    for right, ch in enumerate(s):
        while ch in seen:
            seen.remove(s[left])
            left += 1
        seen.add(ch)
        best = max(best, right - left + 1)
    return best




# Requires non-negative values and returns the shortest window reaching target, or 0. Time O(n).
# Requires: import math
def minimum_subarray_length_with_sum_at_least_target(nums: list[int], target: int) -> int:
    left = total = 0
    best = math.inf
    for right, x in enumerate(nums):
        total += x
        while total >= target:
            best = min(best, right - left + 1)
            total -= nums[left]
            left += 1
    return 0 if best == math.inf else int(best)


# Returns the largest sum of exactly k contiguous values by adding one and removing one per
# step. Time O(n).
def maximum_sum_fixed_length_window(nums: list[int], k: int) -> int:
    window = sum(nums[:k])
    best = window
    for right in range(k, len(nums)):
        window += nums[right] - nums[right - k]
        best = max(best, window)
    return best




# Uses exactly(k) = at_most(k) - at_most(k-1). Time O(n), space O(k).
# Requires via helper: import collections
def count_subarrays_with_exactly_k_distinct_values(nums: list[int], k: int) -> int:
    return count_subarrays_with_at_most_k_distinct_values(
        nums, k
    ) - count_subarrays_with_at_most_k_distinct_values(nums, k - 1)


# Returns unique triples summing to zero after sorting and scanning with two pointers. Time
# O(n^2).
def unique_triplets_summing_to_zero(nums: list[int]) -> list[list[int]]:
    nums.sort()
    result: list[list[int]] = []
    for i, x in enumerate(nums):
        if i > 0 and x == nums[i - 1]:
            continue
        left, right = i + 1, len(nums) - 1
        while left < right:
            total = x + nums[left] + nums[right]
            if total == 0:
                result.append([x, nums[left], nums[right]])
                left += 1
                right -= 1
                while left < right and nums[left] == nums[left - 1]:
                    left += 1
                while left < right and nums[right] == nums[right + 1]:
                    right -= 1
            elif total < 0:
                left += 1
            else:
                right -= 1
    return result


# Moves zeros to the end in place while preserving non-zero order. Time O(n), space O(1).
def move_zeroes(nums: list[int]) -> None:
    write = 0
    for read, x in enumerate(nums):
        if x != 0:
            nums[write], nums[read] = nums[read], nums[write]
            write += 1


# Returns the largest water-container area; moving the shorter wall is the only move that can
# improve height. Time O(n).
def maximum_water_container_area_using_two_pointers(height: list[int]) -> int:
    left, right = 0, len(height) - 1
    best = 0
    while left < right:
        best = max(
            best, (right - left) * min(height[left], height[right])
        )
        if height[left] < height[right]:
            left += 1
        else:
            right -= 1
    return best


# Counts palindromic substrings by expanding around every character and gap. Time O(n^2).
def count_palindromic_substrings_by_expanding_centers(s: str) -> int:
    # expand counts palindromes sharing the supplied center.
    def count_from_center(left: int, right: int) -> int:
        count = 0
        while 0 <= left and right < len(s) and s[left] == s[right]:
            count += 1
            left -= 1
            right += 1
        return count

    return sum(count_from_center(i, i) + count_from_center(i, i + 1) for i in range(len(s)))
