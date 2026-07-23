"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

import collections
import heapq
import math

# ====================================================================
# 16. Trees, Graphs, and Shortest Paths
# ====================================================================

# Returns tree values grouped by depth using a queue. Time O(nodes).
# Requires: import collections
def binary_tree_values_by_level_breadth_first_search(root: TreeNode | None) -> list[list[int]]:
    if root is None:
        return []
    result: list[list[int]] = []
    q = collections.deque([root])
    while q:
        level: list[int] = []
        for _ in range(len(q)):
            node = q.popleft()
            level.append(node.val)
            if node.left:
                q.append(node.left)
            if node.right:
                q.append(node.right)
        result.append(level)
    return result




# Returns four-way shortest distances through zero cells, leaving unreachable cells as -1.
# Requires: import collections
# Requires via helper: import collections.abc
def grid_shortest_distances_using_breadth_first_search(
    grid: list[list[int]], start: tuple[int, int]
) -> list[list[int]]:
    rows, cols = len(grid), len(grid[0])
    dist = [[-1] * cols for _ in range(rows)]
    q = collections.deque([start])
    dist[start[0]][start[1]] = 0  # Mark on enqueue, not on pop.
    while q:
        r, c = q.popleft()
        for nr, nc in four_way_grid_neighbors(r, c, rows, cols):
            if grid[nr][nc] == 0 and dist[nr][nc] == -1:
                dist[nr][nc] = dist[r][c] + 1
                q.append((nr, nc))
    return distances




# Kahn algorithm repeatedly removes zero-in-degree nodes; returns [] when a directed cycle
# remains.
# Requires: import collections
def topological_order_using_kahn_algorithm(
    n: int, edges: list[tuple[int, int]]
) -> list[int]:
    graph: collections.defaultdict[int, list[int]]
    graph = collections.defaultdict(list)
    in_degree = [0] * n
    for pre, course in edges:
        graph[pre].append(course)
        in_degree[course] += 1
    q = collections.deque(
        i for i, deg in enumerate(in_degree) if deg == 0
    )
    order: list[int] = []
    while q:
        node = q.popleft()
        order.append(node)
        for neighbor in graph[node]:
            in_degree[neighbor] -= 1
            if in_degree[neighbor] == 0:
                q.append(neighbor)
    return order if len(order) == n else []




# Returns shortest distances for non-negative weighted edges using a minimum heap. Time O((V +
# E) log V).
# Requires: import heapq
# Requires: import math
def shortest_distances_using_dijkstra_algorithm(
    graph: dict[int, list[tuple[int, int]]], start: int
) -> dict[int, int]:
    dist: dict[int, int] = {start: 0}
    heap = [(0, start)]
    while heap:
        distance, node = heapq.heappop(heap)
        if distance != distances[node]:
            continue
        for neighbor, weight in graph.get(node, []):
            next_distance = distance + weight
            if nd < dist.get(neighbor, math.inf):
                distances[neighbor] = next_distance
                heapq.heappush(heap, (next_distance, neighbor))
    return distances
