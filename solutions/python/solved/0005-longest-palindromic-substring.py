# LeetCode 5 — Longest Palindromic Substring
# https://leetcode.com/problems/longest-palindromic-substring/

class Solution:
    def longestPalindrome(self, s: str) -> str:
        n = len(s)
        dp = [[False] * n for _ in range(n)]
        ans = [0, 0]

        for i in range(n):
            dp[i][i] = True

        # center of even length
        for i in range(n - 1):
            if s[i] == s[i + 1]:
                dp[i][i + 1] = True
                ans = [i, i + 1]
        
        for diff in range(2, n):
            for i in range(n - diff):
                j = i + diff
                if dp[i + 1][j - 1] and s[i] == s[j]:
                    dp[i][j] = True
                    if diff > ans[1] - ans[0]:
                        ans = [i, j]

        return s[ans[0]: ans[1] + 1]


def run_assertion_tests():
    solution = Solution()

    assert solution.longestPalindrome("babad") in {"bab", "aba"}
    assert solution.longestPalindrome("cbbd") == "bb"
    assert solution.longestPalindrome("a") == "a"
    assert solution.longestPalindrome("ac") in {"a", "c"}


if __name__ == "__main__":
    solution = Solution()
    print(f"Test case 1 (expected: 'bab' or 'aba'): {solution.longestPalindrome('babad')}")
    print(f"Test case 2 (expected: 'bb'): {solution.longestPalindrome('cbbd')}")
    print(f"Test case 3 (expected: 'a'): {solution.longestPalindrome('a')}")
    print(f"Test case 4 (expected: 'a' or 'c'): {solution.longestPalindrome('ac')}")

    run_assertion_tests()
