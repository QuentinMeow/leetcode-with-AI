"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

import collections
import math

# ====================================================================
# 13. Hashing, Prefix Sums, and Intervals
# ====================================================================

# Groups lowercase words by their 26-letter frequency tuple. Time O(total characters).
# Requires: import collections
def group_anagrams_by_letter_counts(words: list[str]) -> list[list[str]]:
    groups: collections.defaultdict[tuple[int, ...], list[str]]
    groups = collections.defaultdict(list)
    for word in words:
        count = [0] * 26
        for ch in word:
            count[ord(ch) - ord("a")] += 1
        groups[tuple(count)].append(word)
    return list(groups.values())


# Starts only where predecessor is absent, so each set value is visited once. Expected time
# O(n).
def longest_consecutive_sequence_length(nums: list[int]) -> int:
    values = set(nums)
    best = 0
    for x in values:
        if x - 1 not in values:
            y = x
            while y in values:
                y += 1
            best = max(best, y - x)
    return best


# Stores cumulative sums with a leading zero for constant-time inclusive range queries.
class OneDimensionalPrefixSum:
    # Initializes a new instance and establishes its invariants.
    def __init__(self, nums: list[int]) -> None:
        self.prefix = [0]
        for x in nums:
            self.prefix.append(self.prefix[-1] + x)

    # sum_range returns the inclusive left-through-right sum in O(1).
    def sum_range(self, left: int, right: int) -> int:
        return self.prefix[right + 1] - self.prefix[left]


# Stores cumulative rectangle sums with a zero border for constant-time region queries.
class TwoDimensionalPrefixSum:
    # Initializes a new instance and establishes its invariants.
    def __init__(self, matrix: list[list[int]]) -> None:
        rows = len(matrix)
        cols = len(matrix[0]) if matrix else 0
        self.prefix = [[0] * (cols + 1) for _ in range(rows + 1)]
        for r in range(rows):
            for c in range(cols):
                self.prefix[r + 1][c + 1] = (
                    matrix[r][c]
                    + self.prefix[r][c + 1]
                    + self.prefix[r + 1][c]
                    - self.prefix[r][c]
                )

    # sum_region returns an inclusive rectangle sum in O(1).
    def sum_region(self, r1: int, c1: int, r2: int, c2: int) -> int:
        return (
            self.prefix[r2 + 1][c2 + 1]
            - self.prefix[r1][c2 + 1]
            - self.prefix[r2 + 1][c1]
            + self.prefix[r1][c1]
        )


# Returns the maximum overlap of half-open meeting intervals. Time O(n log n).
def maximum_concurrent_meetings_using_two_pointers(intervals: list[list[int]]) -> int:
    starts = sorted(start for start, _ in intervals)
    ends = sorted(end for _, end in intervals)
    rooms = best = 0
    s = e = 0
    while s < len(starts):
        if starts[s] < ends[e]:
            rooms += 1
            best = max(best, rooms)
            s += 1
        else:
            rooms -= 1
            e += 1
    return best


# Inserts into sorted non-overlapping closed intervals and merges every overlap. Time O(n).
def insert_interval(
    intervals: list[list[int]], new_interval: list[int]
) -> list[list[int]]:
    result: list[list[int]] = []
    i = 0
    while i < len(intervals) and intervals[i][1] < new_interval[0]:
        result.append(intervals[i])
        i += 1
    while i < len(intervals) and intervals[i][0] <= new_interval[1]:
        new_interval[0] = min(new_interval[0], intervals[i][0])
        new_interval[1] = max(new_interval[1], intervals[i][1])
        i += 1
    return result + [new_interval] + intervals[i:]




# Greedily keeps earliest-ending intervals and returns minimum removals. Time O(n log n).
# Requires: import math
def erase_overlap_intervals(intervals: list[list[int]]) -> int:
    intervals.sort(key=lambda x: x[1])
    removed = 0
    prev_end = -math.inf
    for start, end in intervals:
        if start >= prev_end:
            prev_end = end
        else:
            removed += 1
    return removed
