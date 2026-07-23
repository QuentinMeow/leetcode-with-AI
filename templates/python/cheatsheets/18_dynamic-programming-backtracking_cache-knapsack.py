"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

# ====================================================================
# 18. Dynamic Programming and Backtracking
# ====================================================================

# Kadane algorithm keeps the best sum ending at each position. Time O(n), space O(1).
def maximum_subarray_sum_using_kadane_algorithm(nums: list[int]) -> int:
    best = best_ending_here = nums[0]
    for x in nums[1:]:
        best_ending_here = max(x, best_ending_here + x)
        best = max(best, best_ending_here)
    return best


# Counts one- or two-step sequences while retaining only the previous two dynamic-programming
# states.
def count_ways_to_climb_stairs_using_rolling_state(n: int) -> int:
    if n <= 2:
        return n
    prev2, prev1 = 1, 2
    for _ in range(3, n + 1):
        prev2, prev1 = prev1, prev1 + prev2
    return prev1


# Returns the longest sequence obtainable without reordering either string. Time and space O(m *
# n).
def longest_common_subsequence_length(a: str, b: str) -> int:
    dp = [[0] * (len(b) + 1) for _ in range(len(a) + 1)]
    for i in range(len(a) - 1, -1, -1):
        for j in range(len(b) - 1, -1, -1):
            dp[i][j] = (
                1 + dp[i + 1][j + 1]
                if a[i] == b[j]
                else max(dp[i + 1][j], dp[i][j + 1])
            )
    return dp[0][0]


# Uses zero-or-one knapsack state to test whether values split into equal sums.
def can_partition_into_equal_sum_subsets(nums: list[int]) -> bool:
    total = sum(nums)
    if total % 2:
        return False
    target = total // 2
    possible = [False] * (target + 1)
    possible[0] = True
    for x in nums:
        for candidate_sum in range(target, x - 1, -1):
            possible[candidate_sum] = (
                possible[candidate_sum]
                or possible[candidate_sum - x]
            )
    return possible[target]


# Returns every ordering by choosing unused positions and undoing each choice. Output O(n * n!).
def all_permutations_using_backtracking(nums: list[int]) -> list[list[int]]:
    result: list[list[int]] = []
    path: list[int] = []
    used = [False] * len(nums)

    # backtrack chooses one candidate, explores it, then undoes that choice.
    def build_permutations() -> None:
        if len(path) == len(nums):
            result.append(path.copy())
            return
        for i, x in enumerate(nums):
            if used[i]:
                continue
            used[i] = True
            path.append(x)
            build_permutations()
            path.pop()
            used[i] = False

    build_permutations()
    return result


# Returns target-sum combinations where positive candidates may be reused and ordering is
# ignored.
def combinations_summing_to_target_with_reuse(
    candidates: list[int], target: int
) -> list[list[int]]:
    candidates.sort()
    result: list[list[int]] = []
    path: list[int] = []

    # backtrack chooses one candidate, explores it, then undoes that choice.
    def build_combinations(start: int, remaining: int) -> None:
        if remaining == 0:
            result.append(path.copy())
            return
        for i in range(start, len(candidates)):
            x = candidates[i]
            if x > remaining:
                break
            path.append(x)
            build_combinations(i, remaining - x)
            path.pop()

    build_combinations(0, target)
    return result
