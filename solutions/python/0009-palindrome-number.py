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
        
