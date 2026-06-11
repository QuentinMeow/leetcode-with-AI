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


def run_assertion_tests():
    solution = Solution()

    assert solution.convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR"
    assert solution.convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI"
    assert solution.convert("ABC", 1) == "ABC"


if __name__ == "__main__":
    solution = Solution()
    print(f"Test case 1 (expected: 'PAHNAPLSIIGYIR'): {solution.convert('PAYPALISHIRING', 3)}")
    print(f"Test case 2 (expected: 'PINALSIGYAHRPI'): {solution.convert('PAYPALISHIRING', 4)}")
    print(f"Test case 3 (expected: 'ABC'): {solution.convert('ABC', 1)}")

    run_assertion_tests()
