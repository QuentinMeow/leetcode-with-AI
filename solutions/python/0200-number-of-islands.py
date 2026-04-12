# LeetCode 200 — Number of Islands
# https://leetcode.com/problems/number-of-islands/

from collections import deque

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
            print(f"r: {r}, c: {c}")
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
        

        
