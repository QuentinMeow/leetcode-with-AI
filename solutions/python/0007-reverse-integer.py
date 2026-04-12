# LeetCode 7 — Reverse Integer
# https://leetcode.com/problems/reverse-integer/

class Solution:
    def reverse(self, x: int) -> int:
        INT_MAX = 2**31 - 1 # 2147483647
        INT_MIN = -2**31 # -2147483648
        if x == INT_MIN:
            # -2147483648, which will overflow
            return 0

        # Now abs value must be within 2^31 - 1
        sign = 1
        if (x < 0):
            sign = -1
        
        ans, remainder = 0, abs(x)
        while (remainder):
            digit = remainder % 10
            if ans > INT_MAX // 10 or (ans == INT_MAX // 10 and digit > 7):
                return 0
            ans = ans * 10 + digit
            remainder //= 10
        
        return ans * sign
        
