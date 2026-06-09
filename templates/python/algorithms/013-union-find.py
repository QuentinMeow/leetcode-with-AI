"""
013 - Union-find / disjoint set patterns

Use when the problem asks about dynamic connectivity, grouping components,
cycle detection in undirected graphs, or merging accounts/equivalence classes.
"""


class DSU:
    # Variant 1: path compression + union by size.
    # Example problems: number of connected components, redundant connection.
    # Amortized time: almost O(1) per find/union, commonly written as O(alpha(n)).
    # Space: O(n)
    def __init__(self, n: int) -> None:
        self.parent = list(range(n))
        self.size = [1] * n
        self.components = n

    def find(self, x: int) -> int:
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, a: int, b: int) -> bool:
        root_a = self.find(a)
        root_b = self.find(b)
        if root_a == root_b:
            return False

        if self.size[root_a] < self.size[root_b]:
            root_a, root_b = root_b, root_a

        self.parent[root_b] = root_a
        self.size[root_a] += self.size[root_b]
        self.components -= 1
        return True


# Variant 2: count connected components from edges.
# Example problems: number of connected components in an undirected graph.
# Time: O((n + e) * alpha(n))
# Space: O(n)
def count_components(n: int, edges: list[tuple[int, int]]) -> int:
    dsu = DSU(n)
    for a, b in edges:
        dsu.union(a, b)
    return dsu.components


# Variant 3: detect cycle / redundant edge.
# Example problems: redundant connection, validate tree.
# Time: O(e * alpha(n))
# Space: O(n)
def find_redundant_edge(edges: list[tuple[int, int]]) -> tuple[int, int] | None:
    dsu = DSU(max(max(a, b) for a, b in edges) + 1)

    for a, b in edges:
        if not dsu.union(a, b):
            return a, b

    return None


# Variant 4: union by arbitrary keys using id compression.
# Example problems: accounts merge, sentence similarity, equation satisfiability.
# Time: O((n + e) * alpha(n))
# Space: O(n)
def group_equivalent_items(pairs: list[tuple[str, str]]) -> dict[int, list[str]]:
    ids: dict[str, int] = {}

    def get_id(item: str) -> int:
        if item not in ids:
            ids[item] = len(ids)
        return ids[item]

    for a, b in pairs:
        get_id(a)
        get_id(b)

    dsu = DSU(len(ids))
    for a, b in pairs:
        dsu.union(ids[a], ids[b])

    groups: dict[int, list[str]] = {}
    for item, item_id in ids.items():
        root = dsu.find(item_id)
        groups.setdefault(root, []).append(item)

    return groups
