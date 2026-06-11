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


def run_assertion_tests():
    solution = Solution()

    assert solution.reverse(123) == 321
    assert solution.reverse(-123) == -321
    assert solution.reverse(120) == 21
    assert solution.reverse(1534236469) == 0
    assert solution.reverse(-2147483648) == 0


if __name__ == "__main__":
    solution = Solution()
    print(f"Test case 1 (expected: 321): {solution.reverse(123)}")
    print(f"Test case 2 (expected: -321): {solution.reverse(-123)}")
    print(f"Test case 3 (expected: 21): {solution.reverse(120)}")
    print(f"Test case 4 (expected: 0): {solution.reverse(1534236469)}")
    print(f"Test case 5 (expected: 0): {solution.reverse(-2147483648)}")

    run_assertion_tests()
