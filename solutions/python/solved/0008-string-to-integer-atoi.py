# LeetCode 8 — String to Integer (atoi)
# https://leetcode.com/problems/string-to-integer-atoi/

class Solution:
    def myAtoi(self, s: str) -> int:
        i = 0
        while i < len(s) and s[i] == " ":
            i += 1

        sign = 1
        if i < len(s) and s[i] in "+-":
            sign = -1 if s[i] == "-" else 1
            i += 1

        result = 0
        while i < len(s) and s[i].isdigit():
            result = result * 10 + int(s[i])
            i += 1

        return self.sanitizeResult(sign * result)

    def sanitizeResult(self, value: int) -> int:
        if value < -2**31:
            return -2**31
        if value > 2**31 - 1:
            return 2**31 - 1
        return value


def run_assertion_tests():
    solution = Solution()

    assert solution.myAtoi("42") == 42
    assert solution.myAtoi("   -042") == -42
    assert solution.myAtoi("1337c0d3") == 1337
    assert solution.myAtoi("words and 987") == 0
    assert solution.myAtoi("+-12") == 0
    assert solution.myAtoi("91283472332") == 2147483647


if __name__ == "__main__":
    solution = Solution()
    print(f"Test case 1 (expected: 42): {solution.myAtoi('42')}")
    print(f"Test case 2 (expected: -42): {solution.myAtoi('   -042')}")
    print(f"Test case 3 (expected: 1337): {solution.myAtoi('1337c0d3')}")
    print(f"Test case 4 (expected: 0): {solution.myAtoi('words and 987')}")
    print(f"Test case 5 (expected: 0): {solution.myAtoi('+-12')}")
    print(f"Test case 6 (expected: 2147483647): {solution.myAtoi('91283472332')}")

    run_assertion_tests()
