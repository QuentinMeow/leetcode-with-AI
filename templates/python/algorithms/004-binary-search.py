"""
004 - Binary search patterns

Use when the input is sorted or when the answer can be checked with a monotonic
predicate: false false false true true true.
"""

from collections.abc import Callable


# Variant 1: lower bound, the safest default template.
# Finds the first index i where nums[i] >= target.
# Example problems: search insert position, first occurrence, bisect_left.
# Time: O(log n)
# Space: O(1)
def lower_bound(nums: list[int], target: int) -> int:
    left, right = 0, len(nums)

    while left < right:
        mid = (left + right) // 2
        if nums[mid] < target:
            left = mid + 1
        else:
            right = mid

    return left


# Variant 2: exact search in a sorted array.
# Example problems: binary search, lookup in sorted list.
# Time: O(log n)
# Space: O(1)
def binary_search(nums: list[int], target: int) -> int:
    index = lower_bound(nums, target)
    return index if index < len(nums) and nums[index] == target else -1


# Variant 3: answer-space binary search with a monotonic feasibility function.
# Example problems: Koko Eating Bananas, ship packages, split array largest sum.
# Time: O(log(range) * cost(can))
# Space: depends on can().
def first_feasible(low: int, high: int, can: Callable[[int], bool]) -> int:
    while low < high:
        mid = (low + high) // 2
        if can(mid):
            high = mid
        else:
            low = mid + 1

    return low


# Variant 4: binary search on a rotated sorted array.
# Example problems: search in rotated sorted array, find minimum.
# Time: O(log n)
# Space: O(1)
def search_rotated(nums: list[int], target: int) -> int:
    left, right = 0, len(nums) - 1

    while left <= right:
        mid = (left + right) // 2
        if nums[mid] == target:
            return mid

        if nums[left] <= nums[mid]:
            if nums[left] <= target < nums[mid]:
                right = mid - 1
            else:
                left = mid + 1
        else:
            if nums[mid] < target <= nums[right]:
                left = mid + 1
            else:
                right = mid - 1

    return -1
