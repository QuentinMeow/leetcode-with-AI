# LeetCode 3 — Longest Substring Without Repeating Characters
# https://leetcode.com/problems/longest-substring-without-repeating-characters/

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        left = 0
        seen = set()
        max_length = 0
        for right in range(len(s)):
            while s[right] in seen:
                seen.remove(s[left])
                left += 1
            seen.add(s[right])
            max_length = max(max_length, right - left + 1)
        
        return max_length


def run_assertion_tests():
    solution = Solution()

    assert solution.lengthOfLongestSubstring("abcabcbb") == 3
    assert solution.lengthOfLongestSubstring("bbbbb") == 1
    assert solution.lengthOfLongestSubstring("pwwkew") == 3
    assert solution.lengthOfLongestSubstring("") == 0


if __name__ == "__main__":
    solution = Solution()
    print(f"Test case 1 (expected: 3): {solution.lengthOfLongestSubstring('abcabcbb')}")
    print(f"Test case 2 (expected: 1): {solution.lengthOfLongestSubstring('bbbbb')}")
    print(f"Test case 3 (expected: 3): {solution.lengthOfLongestSubstring('pwwkew')}")
    print(f"Test case 4 (expected: 0): {solution.lengthOfLongestSubstring('')}")

    run_assertion_tests()
