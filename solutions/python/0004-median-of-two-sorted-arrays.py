# LeetCode 4 — Median of Two Sorted Arrays
# https://leetcode.com/problems/median-of-two-sorted-arrays/
#
# Approaches in this file (all correct; pick one to submit as `Solution`):
#
# 1) Solution — linear two-pointer “virtual merge” up to the median.
#    Time O(m + n) worst case, space O(1). Easiest to explain in interviews.
#
# 2) SolutionBinarySearchPartition — binary search on where to cut the shorter array.
#    Time O(log(min(m, n))), space O(1). Matches the usual “hard” follow-up.
#
# 3) SolutionKthByElimination — recursively discard chunks from one array (k-th smallest).
#    Time O(log(m + n)) typical depth, space O(log(m + n)) call stack (iterable variant exists).

from typing import List, Tuple


class Solution:
    """
    Virtual merge: advance through nums1 and nums2 in sorted order only as far as needed.

    Let total = m + n. In 1-based “rank” among the merged order:
      - Odd length: median is the ((total + 1) // 2)-th smallest.
      - Even length: median is average of (total // 2)-th and (total // 2 + 1)-th smallest.

    After (total + 1) // 2 steps, `curr` is the lower middle element (the (total//2)-th smallest
    when total is even, or the sole median when odd). One more step yields the upper middle for even.
    """

    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        m, n = len(nums1), len(nums2)
        total = m + n
        steps = (total + 1) // 2

        i = j = 0
        curr = 0

        for _ in range(steps):
            i, j, curr = self._take_smaller(nums1, i, nums2, j)

        if total % 2 == 1:
            return float(curr)

        _, _, nxt = self._take_smaller(nums1, i, nums2, j)
        return (curr + nxt) / 2.0

    @staticmethod
    def _take_smaller(
        nums1: List[int], i: int, nums2: List[int], j: int
    ) -> Tuple[int, int, int]:
        """Advance one pointer and return (new_i, new_j, value_taken)."""
        if i >= len(nums1):
            return i, j + 1, nums2[j]
        if j >= len(nums2):
            return i + 1, j, nums1[i]
        if nums1[i] <= nums2[j]:
            return i + 1, j, nums1[i]
        return i, j + 1, nums2[j]


class SolutionBinarySearchPartition:
    """
    Partition the two arrays into a left half and right half of equal size (left gets the extra
    element when total length is odd), such that every value in the left half is <= every value in
    the right half. The median is then determined by the boundary values around the cut.

    Binary search the cut position `i` in the *shorter* array (after swapping if needed). The cut
    in the longer array is implied: `j = half - i`, where `half = (m + n + 1) // 2` is the size of
    the left partition.

    Invariants (0-based):
      - Left side uses nums1[0 : i] and nums2[0 : j]
      - Right side uses nums1[i : m] and nums2[j : n]

    A valid partition satisfies:
      - nums1[i-1] <= nums2[j]   (largest on left from nums1 does not exceed smallest on right from nums2)
      - nums2[j-1] <= nums1[i]   (symmetric)

    Use sentinels: if i == 0, left1 = -inf; if i == m, right1 = +inf (and similarly for j).
    """

    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1

        m, n = len(nums1), len(nums2)
        left, right = 0, m
        half = (m + n + 1) // 2

        while left <= right:
            i = (left + right) // 2
            j = half - i

            left1 = float("-inf") if i == 0 else nums1[i - 1]
            right1 = float("inf") if i == m else nums1[i]
            left2 = float("-inf") if j == 0 else nums2[j - 1]
            right2 = float("inf") if j == n else nums2[j]

            if left1 <= right2 and left2 <= right1:
                if (m + n) % 2 == 1:
                    return float(max(left1, left2))
                return (max(left1, left2) + min(right1, right2)) / 2.0

            # Too many elements from nums1 on the left — move partition left in nums1.
            if left1 > right2:
                right = i - 1
            else:
                left = i + 1

        raise RuntimeError("unreachable for valid sorted inputs")


class SolutionKthByElimination:
    """
    Generalization: median reduces to one or two k-th smallest queries in the merged order.

    find_kth(i, j, k): k-th smallest (1-based) in the union of nums1[i:] and nums2[j:].

    Compare the candidates at positions that would remove about k/2 elements from one side.
    Whichever side has the smaller “pivot” at that probe cannot contain the k-th smallest for
    this k, so discard that prefix and reduce k by the number of elements discarded.

    Always recurse on the side with fewer remaining elements first so probes stay in bounds.
    """

    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        total = len(nums1) + len(nums2)
        mid = total // 2
        if total % 2 == 1:
            return float(self._kth(nums1, 0, nums2, 0, mid + 1))
        return (self._kth(nums1, 0, nums2, 0, mid) + self._kth(nums1, 0, nums2, 0, mid + 1)) / 2.0

    @staticmethod
    def _kth(nums1: List[int], i: int, nums2: List[int], j: int, k: int) -> int:
        # Keep nums1[i:] the shorter (or equal) segment so the first discard is always safe.
        if len(nums1) - i > len(nums2) - j:
            return SolutionKthByElimination._kth(nums2, j, nums1, i, k)

        m = len(nums1) - i
        n = len(nums2) - j

        if m == 0:
            return nums2[j + k - 1]
        if k == 1:
            return min(nums1[i], nums2[j])

        # Discard up to k//2 elements from the shorter array, and the rest from the longer.
        take1 = min(k // 2, m)
        take2 = min(k - take1, n)
        if nums1[i + take1 - 1] < nums2[j + take2 - 1]:
            return SolutionKthByElimination._kth(nums1, i + take1, nums2, j, k - take1)
        return SolutionKthByElimination._kth(nums1, i, nums2, j + take2, k - take2)
