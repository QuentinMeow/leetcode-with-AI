# LeetCode 6 — Zigzag Conversion
# https://leetcode.com/problems/zigzag-conversion/

class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s
        
        n = len(s)
        ans = []
        charsInSection = 2 * numRows - 2

        for row in range(numRows):
            i = row
            while (i < n):
                ans.append(s[i])
                if row != 0 and row != numRows - 1:
                    diff =  2 * numRows - row * 2 - 2
                    second_i = i + diff
                    if second_i < n:
                        ans.append(s[second_i])
                i += charsInSection
        
        return "".join(ans)
        
