# LeetCode 9 — Palindrome Number
# https://leetcode.com/problems/palindrome-number/

class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        if x < 10:
            return True
        if x % 10 == 0:
            return False

        leftside = x
        reversed_rightside = 0
        while leftside > 0:
            # even palindrome
            if leftside == reversed_rightside:
                return True
            # odd palindrome
            if leftside // 10 > 0 and leftside // 10 == reversed_rightside:
                return True
            # not yet done, continue
            reversed_rightside = reversed_rightside * 10 + leftside % 10
            leftside //= 10
        
        return False


def run_assertion_tests():
    solution = Solution()

    assert solution.isPalindrome(121)
    assert not solution.isPalindrome(-121)
    assert not solution.isPalindrome(10)
    assert solution.isPalindrome(12321)
    assert solution.isPalindrome(0)


if __name__ == "__main__":
    solution = Solution()
    print(f"Test case 1 (expected: True): {solution.isPalindrome(121)}")
    print(f"Test case 2 (expected: False): {solution.isPalindrome(-121)}")
    print(f"Test case 3 (expected: False): {solution.isPalindrome(10)}")
    print(f"Test case 4 (expected: True): {solution.isPalindrome(12321)}")
    print(f"Test case 5 (expected: True): {solution.isPalindrome(0)}")

    run_assertion_tests()
