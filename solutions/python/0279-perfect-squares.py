# LeetCode 279 — Perfect Squares
# https://leetcode.com/problems/perfect-squares/

class Solution:
    # greedy
    def numSquares(self, n: int) -> int:
        def canWork(n, count):
            if count == 1:
                return n in square_nums
            for k in square_nums:
                if canWork(n - k, count - 1):
                    return True
            return False

        square_nums = set([i * i for i in range(1, int(n**0.5) + 1)])
        for count in range(1, n + 1):
            if canWork(n, count):
                return count

    # original idea, dp
    def helper(self, n: int, dp: dict[int, int]) -> int:
        if n in dp:
            return dp[n]
        i = 1
        minimum = math.inf
        while i * i <= n:
            res = self.helper(n - i * i, dp) + 1
            if res < minimum:
                minimum = res
            i += 1

        dp[n] = minimum
        return minimum

        
