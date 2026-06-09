"""
005 - Prefix sum patterns

Use when repeated range sums are needed, or when subarray problems can be
converted into differences between prefix states.
"""

from collections import defaultdict


# Variant 1: prefix sum with a hash map of seen prefix counts.
# Example problems: subarray sum equals k, path sum with prefix counts.
# Time: O(n)
# Space: O(n)
def count_subarrays_sum_k(nums: list[int], k: int) -> int:
    seen = defaultdict(int)
    seen[0] = 1
    prefix = 0
    total = 0

    for x in nums:
        prefix += x
        total += seen[prefix - k]
        seen[prefix] += 1

    return total


# Variant 2: prefix array for immutable range sum queries.
# Example problems: range sum query, sum between i and j many times.
# Build time: O(n)
# Query time: O(1)
# Space: O(n)
class PrefixSum:
    def __init__(self, nums: list[int]) -> None:
        self.prefix = [0]
        for x in nums:
            self.prefix.append(self.prefix[-1] + x)

    def range_sum(self, left: int, right: int) -> int:
        """Inclusive [left, right]."""
        return self.prefix[right + 1] - self.prefix[left]


# Variant 3: 2D prefix sum.
# Example problems: range sum query 2D, matrix block sum.
# Build time: O(m * n)
# Query time: O(1)
# Space: O(m * n)
class PrefixSum2D:
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

    def region_sum(self, r1: int, c1: int, r2: int, c2: int) -> int:
        """Inclusive top-left (r1, c1) to bottom-right (r2, c2)."""
        return (
            self.prefix[r2 + 1][c2 + 1]
            - self.prefix[r1][c2 + 1]
            - self.prefix[r2 + 1][c1]
            + self.prefix[r1][c1]
        )


# Variant 4: difference array for many range updates.
# Example problems: range addition, car pooling, meeting capacity.
# Time: O(n + q)
# Space: O(n)
def apply_range_additions(length: int, updates: list[tuple[int, int, int]]) -> list[int]:
    diff = [0] * (length + 1)

    for left, right, delta in updates:
        diff[left] += delta
        if right + 1 < len(diff):
            diff[right + 1] -= delta

    result = [0] * length
    running = 0
    for i in range(length):
        running += diff[i]
        result[i] = running

    return result
