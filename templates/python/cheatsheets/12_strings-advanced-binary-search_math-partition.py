"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

import math

# ====================================================================
# 12. Strings and Advanced Binary Search
# ====================================================================

# Keeps alphanumeric characters, lowercases them, and removes punctuation and spaces.
def normalize_letters_and_digits_lowercase(s: str) -> str:
    return "".join(ch.lower() for ch in s if ch.isalnum())


# Checks a case-insensitive palindrome while skipping non-alphanumeric characters. Time O(n).
def is_palindrome_ignoring_non_letters_and_digits(s: str) -> bool:
    left, right = 0, len(s) - 1
    while left < right:
        while left < right and not s[left].isalnum():
            left += 1
        while left < right and not s[right].isalnum():
            right -= 1
        if s[left].lower() != s[right].lower():
            return False
        left += 1
        right -= 1
    return True


# Adds non-negative decimal integers represented as strings using right-to-left carry.
def add_decimal_strings(a: str, b: str) -> str:
    i, j, carry = len(a) - 1, len(b) - 1, 0
    result: list[str] = []
    while i >= 0 or j >= 0 or carry:
        x = ord(a[i]) - ord("0") if i >= 0 else 0
        y = ord(b[j]) - ord("0") if j >= 0 else 0
        carry, digit = divmod(x + y + carry, 10)
        result.append(chr(ord("0") + digit))
        i -= 1
        j -= 1
    return "".join(reversed(result))


# Compares dot-separated integer components and treats missing trailing components as zero.
def compare_version_numbers(version1: str, version2: str) -> int:
    a = [int(part) for part in version1.split(".")]
    b = [int(part) for part in version2.split(".")]
    n = max(len(a), len(b))
    for i in range(n):
        x = a[i] if i < len(a) else 0
        y = b[i] if i < len(b) else 0
        if x != y:
            return -1 if x < y else 1
    return 0


# Returns consecutive character runs as (character, count) pairs. Time O(n).
def run_length_encode(s: str) -> list[tuple[str, int]]:
    if not s:
        return []
    groups: list[tuple[str, int]] = []
    cur, count = s[0], 1
    for ch in s[1:]:
        if ch == cur:
            count += 1
        else:
            groups.append((cur, count))
            cur, count = ch, 1
    groups.append((cur, count))
    return groups


# Finds target in a distinct ascending array rotated at an unknown pivot. Time O(log n).
def search_rotated_sorted_array(nums: list[int], target: int) -> int:
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




# Binary-searches a partition where every left value is at most every right value. Time O(log
# min(m, n)).
# Requires: import math
def median_of_two_sorted_arrays_using_partition(a: list[int], b: list[int]) -> float:
    if len(a) > len(b):
        a, b = b, a
    m, n = len(a), len(b)
    half = (m + n + 1) // 2
    left, right = 0, m
    while left <= right:
        i = (left + right) // 2
        j = half - i
        left1 = -math.inf if i == 0 else a[i - 1]
        right1 = math.inf if i == m else a[i]
        left2 = -math.inf if j == 0 else b[j - 1]
        right2 = math.inf if j == n else b[j]
        if left1 <= right2 and left2 <= right1:
            if (m + n) % 2:
                return float(max(left1, left2))
            return (max(left1, left2) + min(right1, right2)) / 2
        if left1 > right2:
            right = i - 1
        else:
            left = i + 1
    raise ValueError("inputs must be sorted")
