"""
002 - Two pointers patterns

Use when the input is sorted, when you need to scan from both ends, or when
the answer depends on comparing/compacting elements in one linear pass.
"""


# Variant 1: opposite-direction pointers on a sorted array.
# Example problems: Two Sum II, 3Sum inner loop, valid palindrome.
# Time: O(n)
# Space: O(1)
def two_sum_sorted(nums: list[int], target: int) -> tuple[int, int] | None:
    left, right = 0, len(nums) - 1

    while left < right:
        total = nums[left] + nums[right]
        if total == target:
            return left, right
        if total < target:
            left += 1
        else:
            right -= 1

    return None


# Variant 2: skip duplicates after choosing a value.
# Example problems: 3Sum, 4Sum, unique pair/triplet generation.
# Time: O(n^2) for 3Sum after sorting
# Space: O(1) extra, excluding the output list.
def three_sum(nums: list[int]) -> list[list[int]]:
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


# Variant 3: same-direction read/write pointers for in-place compaction.
# Example problems: remove duplicates, move zeroes, partition array.
# Time: O(n)
# Space: O(1)
def move_zeroes(nums: list[int]) -> None:
    write = 0

    for read, x in enumerate(nums):
        if x != 0:
            nums[write], nums[read] = nums[read], nums[write]
            write += 1


# Variant 4: expand around center.
# Example problems: palindromic substrings, longest palindromic substring.
# Time: O(n^2)
# Space: O(1)
def count_palindromic_substrings(s: str) -> int:
    def expand(left: int, right: int) -> int:
        count = 0
        while left >= 0 and right < len(s) and s[left] == s[right]:
            count += 1
            left -= 1
            right += 1
        return count

    total = 0
    for center in range(len(s)):
        total += expand(center, center)
        total += expand(center, center + 1)

    return total
