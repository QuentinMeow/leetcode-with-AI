# LeetCode 723 — Candy Crush
# https://leetcode.com/problems/candy-crush/
from typing import List

class Solution:
    def candyCrush(self, board: List[List[int]]) -> List[List[int]]:
        rows, cols = len(board), len(board[0])

        def find():
            crushable_set = set()
            # check vertical middle element
            for i in range(1, rows - 1):
                for j in range(cols):
                    if board[i][j] != 0 and board[i - 1][j] == board[i][j] == board[i + 1][j]:
                        crushable_set.add((i - 1, j))
                        crushable_set.add((i, j))
                        crushable_set.add((i + 1, j))

            # check horizontal middle element
            for i in range(rows):
                for j in range(1, cols - 1):
                    if board[i][j] != 0 and board[i][j - 1] == board[i][j] == board[i][j + 1]:
                        crushable_set.add((i, j - 1))
                        crushable_set.add((i, j))
                        crushable_set.add((i, j + 1))
            return crushable_set

        def crush(crushable_set):
            for i, j in crushable_set:
                board[i][j] = 0
            return

        def drop():
            for j in range(cols):
                all_candies = []
                i = rows - 1
                while i >= 0:
                    if board[i][j] != 0:
                        all_candies.append(board[i][j])
                    i -= 1

                for i in range(len(all_candies)):
                    board[rows - i - 1][j] = all_candies[i]

                for i in range(rows - len(all_candies)):
                    board[i][j] = 0 
            return

        crushable_set = find()
        while crushable_set:
            crush(crushable_set)
            drop()
            crushable_set = find()
        return board


def run_assertion_tests():
    solution = Solution()

    board = [
        [110, 5, 112, 113, 114],
        [210, 211, 5, 213, 214],
        [310, 311, 3, 313, 314],
        [410, 411, 412, 5, 414],
        [5, 1, 512, 3, 3],
        [610, 4, 1, 613, 614],
        [710, 1, 2, 713, 714],
        [810, 1, 2, 1, 1],
        [1, 1, 2, 2, 2],
        [4, 1, 4, 4, 1014],
    ]
    expected = [
        [0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0],
        [110, 0, 0, 0, 114],
        [210, 0, 0, 0, 214],
        [310, 0, 0, 113, 314],
        [410, 0, 0, 213, 414],
        [610, 211, 112, 313, 614],
        [710, 311, 412, 613, 714],
        [810, 411, 512, 713, 1014],
    ]
    assert solution.candyCrush(board) == expected

    board = [[1, 2, 3], [4, 5, 6]]
    assert solution.candyCrush(board) == [[1, 2, 3], [4, 5, 6]]


if __name__ == "__main__":
    solution = Solution()

    board = [
        [110, 5, 112, 113, 114],
        [210, 211, 5, 213, 214],
        [310, 311, 3, 313, 314],
        [410, 411, 412, 5, 414],
        [5, 1, 512, 3, 3],
        [610, 4, 1, 613, 614],
        [710, 1, 2, 713, 714],
        [810, 1, 2, 1, 1],
        [1, 1, 2, 2, 2],
        [4, 1, 4, 4, 1014],
    ]
    print(f"Test case 1 (expected: stable crushed board): {solution.candyCrush(board)}")

    board = [[1, 2, 3], [4, 5, 6]]
    print(f"Test case 2 (expected: unchanged board): {solution.candyCrush(board)}")

    run_assertion_tests()
