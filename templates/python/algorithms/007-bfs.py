"""
007 - Breadth-first search patterns

Use for shortest paths in unweighted graphs, level-order traversal, grid
distance, and problems where "minimum number of steps" is the signal.
"""

from collections import deque


# Variant 1: queue BFS on a graph.
# Example problems: shortest path in unweighted graph, word ladder state graph.
# Time: O(V + E)
# Space: O(V)
def shortest_path_length(graph: dict[int, list[int]], start: int, target: int) -> int:
    queue = deque([(start, 0)])
    seen = {start}

    while queue:
        node, dist = queue.popleft()
        if node == target:
            return dist

        for nei in graph[node]:
            if nei not in seen:
                seen.add(nei)
                queue.append((nei, dist + 1))

    return -1


# Variant 2: level-order traversal.
# Example problems: binary tree level order, right side view, zigzag traversal.
# Time: O(n)
# Space: O(w), where w is max tree width.
def level_order(root) -> list[list[int]]:
    if not root:
        return []

    result: list[list[int]] = []
    queue = deque([root])

    while queue:
        level: list[int] = []
        for _ in range(len(queue)):
            node = queue.popleft()
            level.append(node.val)
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        result.append(level)

    return result


# Variant 3: grid BFS with four directions.
# Example problems: rotting oranges, walls and gates, shortest path in binary matrix.
# Time: O(rows * cols)
# Space: O(rows * cols)
def shortest_grid_path(grid: list[list[int]], start: tuple[int, int], target: tuple[int, int]) -> int:
    rows, cols = len(grid), len(grid[0])
    queue = deque([(start[0], start[1], 0)])
    seen = {start}
    directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]

    while queue:
        r, c, dist = queue.popleft()
        if (r, c) == target:
            return dist

        for dr, dc in directions:
            nr, nc = r + dr, c + dc
            if (
                0 <= nr < rows
                and 0 <= nc < cols
                and grid[nr][nc] == 0
                and (nr, nc) not in seen
            ):
                seen.add((nr, nc))
                queue.append((nr, nc, dist + 1))

    return -1


# Variant 4: multi-source BFS.
# Example problems: rotting oranges, nearest zero, map of highest peak.
# Time: O(rows * cols)
# Space: O(rows * cols)
def distance_from_sources(grid: list[list[int]], sources: list[tuple[int, int]]) -> list[list[int]]:
    rows, cols = len(grid), len(grid[0])
    dist = [[-1] * cols for _ in range(rows)]
    queue = deque()

    for r, c in sources:
        dist[r][c] = 0
        queue.append((r, c))

    while queue:
        r, c = queue.popleft()
        for dr, dc in ((1, 0), (-1, 0), (0, 1), (0, -1)):
            nr, nc = r + dr, c + dc
            if 0 <= nr < rows and 0 <= nc < cols and dist[nr][nc] == -1:
                dist[nr][nc] = dist[r][c] + 1
                queue.append((nr, nc))

    return dist
