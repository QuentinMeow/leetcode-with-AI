# LeetCode 723 — Candy Crush
# https://leetcode.com/problems/candy-crush/
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


