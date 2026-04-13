# LeetCode 408 — Valid Word Abbreviation
# https://leetcode.com/problems/valid-word-abbreviation/

class Solution:
    def validWordAbbreviation(self, word: str, abbr: str) -> bool:
        i = j = 0
        while i < len(abbr):
            if j >= len(word):
                return False
            if abbr[i] == word[j]:
                j += 1
                i += 1
            elif not abbr[i].isdigit():
                return False
            elif abbr[i] == '0':
                return False
            else:
                num = 0
                while i < len(abbr) and abbr[i].isdigit():
                    num = num * 10 + int(abbr[i])
                    i += 1
                j += num
        return j == len(word)

