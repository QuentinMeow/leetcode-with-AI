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


def run_assertion_tests():
    solution = Solution()

    assert solution.validWordAbbreviation("internationalization", "i12iz4n")
    assert solution.validWordAbbreviation("a", "1")
    assert not solution.validWordAbbreviation("apple", "a2e")
    assert not solution.validWordAbbreviation("substitution", "s010n")
    assert not solution.validWordAbbreviation("word", "5")


if __name__ == "__main__":
    solution = Solution()
    print(f"Test case 1 (expected: True): {solution.validWordAbbreviation('internationalization', 'i12iz4n')}")
    print(f"Test case 2 (expected: True): {solution.validWordAbbreviation('a', '1')}")
    print(f"Test case 3 (expected: False): {solution.validWordAbbreviation('apple', 'a2e')}")
    print(f"Test case 4 (expected: False): {solution.validWordAbbreviation('substitution', 's010n')}")
    print(f"Test case 5 (expected: False): {solution.validWordAbbreviation('word', '5')}")

    run_assertion_tests()
