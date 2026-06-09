"""
012 - Dynamic programming patterns

Use when a problem has overlapping subproblems and an optimal answer can be
built from smaller states. Start by naming the state and transition.
"""

from functools import cache


# Variant 1: 1D rolling DP.
# Example problems: climbing stairs, house robber, min cost climbing stairs.
# Time: O(n)
# Space: O(1)
def climb_stairs(n: int) -> int:
    if n <= 2:
        return n

    prev2, prev1 = 1, 2
    for _ in range(3, n + 1):
        prev2, prev1 = prev1, prev1 + prev2

    return prev1


# Variant 2: top-down memoization.
# Example problems: coin change, word break, decode ways, recursive choices.
# Time: O(amount * len(coins))
# Space: O(amount)
def coin_change(coins: list[int], amount: int) -> int:
    inf = amount + 1

    @cache
    def dp(remaining: int) -> int:
        if remaining == 0:
            return 0
        if remaining < 0:
            return inf

        return 1 + min(dp(remaining - coin) for coin in coins)

    answer = dp(amount)
    return -1 if answer >= inf else answer


# Variant 3: 2D table for two sequences.
# Example problems: longest common subsequence, edit distance, interleaving string.
# Time: O(m * n)
# Space: O(m * n)
def longest_common_subsequence(a: str, b: str) -> int:
    rows, cols = len(a), len(b)
    dp = [[0] * (cols + 1) for _ in range(rows + 1)]

    for i in range(rows - 1, -1, -1):
        for j in range(cols - 1, -1, -1):
            if a[i] == b[j]:
                dp[i][j] = 1 + dp[i + 1][j + 1]
            else:
                dp[i][j] = max(dp[i + 1][j], dp[i][j + 1])

    return dp[0][0]


# Variant 4: knapsack-style capacity DP.
# Example problems: partition equal subset sum, 0/1 knapsack, target sum variants.
# Time: O(n * target)
# Space: O(target)
def can_partition(nums: list[int]) -> bool:
    total = sum(nums)
    if total % 2:
        return False

    target = total // 2
    possible = [False] * (target + 1)
    possible[0] = True

    for x in nums:
        for current in range(target, x - 1, -1):
            possible[current] = possible[current] or possible[current - x]

    return possible[target]
