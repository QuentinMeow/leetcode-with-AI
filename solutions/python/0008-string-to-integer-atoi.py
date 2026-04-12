# LeetCode 8 — String to Integer (atoi)
# https://leetcode.com/problems/string-to-integer-atoi/

class Solution:
    def myAtoi(self, s: str) -> int:
        started = False
        sign = 1
        result = 0
        for char in s:
            if char.isalpha() or char == '.' or (started and not char.isdigit()):
                return self.sanitizeResult(sign * result)
            
            # either not started yet, or started already with number calculation
            if char == '+':
                started = True
            elif char == '-':
                started = True
                sign = -1
            elif char.isdigit():
                started = True
                result = result * 10 + int(char)
            elif char == ' ' and not started:
                continue
            else:
                self.sanitizeResult(sign * result)
        return self.sanitizeResult(sign * result)

    def sanitizeResult(self, value) -> int:
        if value < -2**31:
            return -2**31
        if value > 2**31 - 1:
            return 2**31 - 1
        return value
        
