"""
008 - DFS and backtracking patterns

Use when exploring all choices, connected components, recursion over trees,
or any search space where you choose, recurse, then undo.
"""


# Variant 1: recursive DFS over a graph/grid component.
# Example problems: number of islands, max area of island, connected components.
# Time: O(V + E), or O(rows * cols) for grids
# Space: O(V) recursion depth/visited set in the worst case.
def count_islands(grid: list[list[str]]) -> int:
    if not grid:
        return 0

    rows, cols = len(grid), len(grid[0])

    def dfs(r: int, c: int) -> None:
        if not (0 <= r < rows and 0 <= c < cols) or grid[r][c] != "1":
            return

        grid[r][c] = "0"
        dfs(r + 1, c)
        dfs(r - 1, c)
        dfs(r, c + 1)
        dfs(r, c - 1)

    islands = 0
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == "1":
                islands += 1
                dfs(r, c)

    return islands


# Variant 2: choose -> recurse -> undo.
# Example problems: subsets, combinations, permutations, combination sum.
# Time: O(n * 2^n)
# Space: O(n) recursion path, excluding output.
def subsets(nums: list[int]) -> list[list[int]]:
    result: list[list[int]] = []
    path: list[int] = []

    def backtrack(start: int) -> None:
        result.append(path.copy())

        for i in range(start, len(nums)):
            path.append(nums[i])
            backtrack(i + 1)
            path.pop()

    backtrack(0)
    return result


# Variant 3: permutations with a used array.
# Example problems: permutations, generate arrangements, assignment search.
# Time: O(n * n!)
# Space: O(n) recursion path, excluding output.
def permutations(nums: list[int]) -> list[list[int]]:
    result: list[list[int]] = []
    path: list[int] = []
    used = [False] * len(nums)

    def backtrack() -> None:
        if len(path) == len(nums):
            result.append(path.copy())
            return

        for i, x in enumerate(nums):
            if used[i]:
                continue
            used[i] = True
            path.append(x)
            backtrack()
            path.pop()
            used[i] = False

    backtrack()
    return result


# Variant 4: backtracking with pruning.
# Example problems: combination sum, word search, N-Queens.
# Time: problem-specific exponential, reduced by pruning
# Space: O(depth)
def combination_sum(candidates: list[int], target: int) -> list[list[int]]:
    candidates.sort()
    result: list[list[int]] = []
    path: list[int] = []

    def backtrack(start: int, remaining: int) -> None:
        if remaining == 0:
            result.append(path.copy())
            return

        for i in range(start, len(candidates)):
            x = candidates[i]
            if x > remaining:
                break
            path.append(x)
            backtrack(i, remaining - x)
            path.pop()

    backtrack(0, target)
    return result
