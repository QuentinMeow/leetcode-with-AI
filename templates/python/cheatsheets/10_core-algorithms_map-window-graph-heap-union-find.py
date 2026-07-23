"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

import collections
import collections.abc
import functools
import heapq

# ====================================================================
# 10. Algorithm Skeletons
# ====================================================================

# Returns two indices whose values sum to target, or None. A map stores prior values. Time O(n),
# space O(n).
def two_sum_indices_using_map(
    nums: list[int], target: int
) -> tuple[int, int] | None:
    seen: dict[int, int] = {}
    for i, x in enumerate(nums):
        if target - x in seen:
            return seen[target - x], i
        seen[x] = i
    return None


# Requires ascending input and moves inward based on whether the sum is too small or large. Time
# O(n), space O(1).
def sorted_two_sum_indices_using_two_pointers(
    nums: list[int], target: int
) -> tuple[int, int] | None:
    left, right = 0, len(nums) - 1
    while left < right:
        total = nums[left] + nums[right]
        if total == target:
            return left, right
        if total < target:
            left += 1
        else:
            right -= 1
    return None




# Counts contiguous subarrays with at most k distinct values. The window shrinks until valid.
# Time O(n), space O(k).
# Requires: import collections
def count_subarrays_with_at_most_k_distinct_values(nums: list[int], k: int) -> int:
    if k < 0:
        return 0
    count: collections.defaultdict[int, int]
    count = collections.defaultdict(int)
    left = total = 0
    for right, x in enumerate(nums):
        count[x] += 1
        while len(count) > k:
            y = nums[left]
            count[y] -= 1
            if count[y] == 0:
                del count[y]
            left += 1
        total += right - left + 1
    return total


# Returns the first sorted index whose value is at least target, possibly len(nums). Time O(log
# n).
def first_index_at_least_target(nums: list[int], target: int) -> int:
    left, right = 0, len(nums)
    while left < right:
        mid = (left + right) // 2
        if nums[mid] < target:
            left = mid + 1
        else:
            right = mid
    return left


# Requires a monotonic predicate and returns its first true value in the search range. Time
# O(log range) predicate calls.
def smallest_feasible_value_using_binary_search(low: int, high: int, can) -> int:
    while low < high:
        mid = (low + high) // 2
        if is_feasible(mid):
            high = mid
        else:
            low = mid + 1
    return low




# Counts contiguous subarrays summing to target by matching earlier prefix sums. Time O(n),
# space O(n).
# Requires: import collections
def count_subarrays_with_target_sum_using_prefix_sums(nums: list[int], k: int) -> int:
    seen: collections.defaultdict[int, int] = collections.defaultdict(
        int
    )
    seen[0] = 1
    prefix = result = 0
    for x in nums:
        prefix += x
        result += seen[prefix - k]
        seen[prefix] += 1
    return result




# Returns the fewest edges from start to target in an unweighted graph, or -1. Time O(vertices +
# edges).
# Requires: import collections
def shortest_path_length_unweighted_breadth_first_search(
    graph: dict[int, list[int]], start: int, target: int
) -> int:
    q = collections.deque([(start, 0)])
    seen = {start}
    while q:
        node, dist = q.popleft()
        if node == target:
            return dist
        for neighbor in graph.get(node, []):
            if neighbor not in seen:
                seen.add(neighbor)
                q.append((neighbor, dist + 1))
    return -1




# Yields valid up, down, left, and right grid coordinates without diagonals.
# Requires: import collections.abc
def four_way_grid_neighbors(
    r: int, c: int, rows: int, cols: int
) -> collections.abc.Iterable[tuple[int, int]]:
    for dr, dc in ((1, 0), (-1, 0), (0, 1), (0, -1)):
        nr, nc = r + dr, c + dc
        if 0 <= nr < rows and 0 <= nc < cols:
            yield nr, nc


# Counts four-directionally connected groups of 1 cells and mutates visited cells to 0. Time
# O(rows * columns).
# Requires via helper: import collections.abc
def count_islands_using_depth_first_search(grid: list[list[str]]) -> int:
    rows, cols = len(grid), len(grid[0]) if grid else 0

    # visit_connected_land marks every land cell connected to this cell.
    def visit_connected_land(r: int, c: int) -> None:
        if not (0 <= r < rows and 0 <= c < cols) or grid[r][c] != "1":
            return
        grid[r][c] = "0"
        for nr, nc in four_way_grid_neighbors(r, c, rows, cols):
            visit_connected_land(nr, nc)

    count = 0
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == "1":
                count += 1
                visit_connected_land(r, c)
    return count


# Returns the power set. The path stores current choices and is copied before backtracking.
# Output O(n * 2^n).
def all_subsets_using_backtracking(nums: list[int]) -> list[list[int]]:
    result: list[list[int]] = []
    path: list[int] = []

    # backtrack chooses one candidate, explores it, then undoes that choice.
    def build_subsets(start: int) -> None:
        result.append(path.copy())
        for i in range(start, len(nums)):
            path.append(nums[i])
            build_subsets(i + 1)
            path.pop()

    build_subsets(0)
    return result


# Returns the next strictly greater index to the right, or -1. The stack stores unresolved
# indices. Time O(n).
def next_greater_index_using_monotonic_stack(nums: list[int]) -> list[int]:
    result = [-1] * len(nums)
    stack: list[int] = []  # indexes; decreasing values
    for i, x in enumerate(nums):
        while stack and nums[stack[-1]] < x:
            result[stack.pop()] = i
        stack.append(i)
    return result


# Sorts closed intervals and combines overlaps. Time O(n log n), output O(n).
def merge_intervals(intervals: list[list[int]]) -> list[list[int]]:
    intervals.sort(key=lambda x: x[0])
    merged: list[list[int]] = []
    for start, end in intervals:
        if not merged or start > merged[-1][1]:
            merged.append([start, end])
        else:
            merged[-1][1] = max(merged[-1][1], end)
    return merged




# Returns k most frequent values using a size-k minimum heap. Time O(n log k), space O(n).
# Requires: import collections
# Requires: import heapq
def top_k_frequent_values_using_heap(nums: list[int], k: int) -> list[int]:
    counts = collections.Counter(nums)
    return [
        num
        for _, num in heapq.nlargest(
            k, ((freq, num) for num, freq in counts.items())
        )
    ]




# Returns the fewest coins totaling amount, or -1, using memoized recursion.
# Requires: import functools
def minimum_coins_top_down_dynamic_programming(coins: tuple[int, ...], amount: int) -> int:
    inf = amount + 1

    # minimum_coins_remaining memoizes the best answer for each remaining amount.
    # Requires: import functools
    @functools.cache
    def minimum_coins_remaining(remaining: int) -> int:
        if remaining == 0:
            return 0
        if remaining < 0:
            return inf
        return 1 + min(minimum_coins_remaining(remaining - coin) for coin in coins)

    result = minimum_coins_remaining(amount)
    return -1 if result >= inf else result


# Disjoint-set union, also called union-find, tracks connected groups with path compression and
# size-balanced merging.
class DisjointSetUnion:
    # Initializes a new instance and establishes its invariants.
    def __init__(self, n: int) -> None:
        self.parent = list(range(n))
        self.size = [1] * n
        self.components = n

    # find returns the group representative and shortens future paths.
    def find(self, x: int) -> int:
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    # union merges two groups by size and reports whether a merge occurred.
    def union(self, a: int, b: int) -> bool:
        root_a, root_b = self.find(a), self.find(b)
        if root_a == root_b:
            return False
        if self.size[root_a] < self.size[root_b]:
            root_a, root_b = root_b, root_a
        self.parent[root_b] = root_a
        self.size[root_a] += self.size[root_b]
        self.components -= 1
        return True
