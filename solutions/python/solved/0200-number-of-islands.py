# LeetCode 200 — Number of Islands
# https://leetcode.com/problems/number-of-islands/

from collections import deque

from typing import List

class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        if not grid:
            return 0
        ans = 0
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if grid[i][j] == "1":
                    self.bfs(i, j, grid)
                    ans += 1

        return ans

    def dfs(self, r, c, grid):
        # out of boundary
        if r < 0 or c < 0 or r >= len(grid) or c >= len(grid[0]):
            return

        # water
        if grid[r][c] != "1":
            return

        # is island, keep searching
        grid[r][c] = "-1"
        self.dfs(r - 1, c, grid)
        self.dfs(r + 1, c, grid)
        self.dfs(r, c - 1, grid)
        self.dfs(r, c + 1, grid)

    def bfs(self, r, c, grid):
        grid[r][c] = "-1"
        queue = deque()
        queue.append((r, c))
        while queue:
            r, c = queue.popleft()
            if r - 1 >= 0 and grid[r - 1][c] == "1":
                queue.append((r - 1, c))
                grid[r - 1][c] = "-1"
            if r + 1 < len(grid) and grid[r + 1][c] == "1":
                queue.append((r + 1, c))
                grid[r + 1][c] = "-1"
            if c - 1 >= 0 and grid[r][c - 1] == "1":
                queue.append((r, c - 1))
                grid[r][c - 1] = "-1"
            if c + 1 < len(grid[0]) and grid[r][c + 1] == "1":
                queue.append((r, c + 1))
                grid[r][c + 1] = "-1"


def run_assertion_tests():
    solution = Solution()

    grid = [
        ["1", "1", "1", "1", "0"],
        ["1", "1", "0", "1", "0"],
        ["1", "1", "0", "0", "0"],
        ["0", "0", "0", "0", "0"],
    ]
    assert solution.numIslands(grid) == 1

    grid = [
        ["1", "1", "0", "0", "0"],
        ["1", "1", "0", "0", "0"],
        ["0", "0", "1", "0", "0"],
        ["0", "0", "0", "1", "1"],
    ]
    assert solution.numIslands(grid) == 3

    assert solution.numIslands([]) == 0


if __name__ == "__main__":
    solution = Solution()

    grid = [
        ["1", "1", "1", "1", "0"],
        ["1", "1", "0", "1", "0"],
        ["1", "1", "0", "0", "0"],
        ["0", "0", "0", "0", "0"],
    ]
    print(f"Test case 1 (expected: 1): {solution.numIslands(grid)}")

    grid = [
        ["1", "1", "0", "0", "0"],
        ["1", "1", "0", "0", "0"],
        ["0", "0", "1", "0", "0"],
        ["0", "0", "0", "1", "1"],
    ]
    print(f"Test case 2 (expected: 3): {solution.numIslands(grid)}")
    print(f"Test case 3 (expected: 0): {solution.numIslands([])}")

    run_assertion_tests()
