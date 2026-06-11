"""
014 - Trie and structure-construction patterns

Use when a problem needs prefix lookup, word dictionaries, or a setup step that
turns raw input into the tree/graph shape the real algorithm expects.
"""

from __future__ import annotations

from collections import deque


class TrieNode:
    def __init__(self) -> None:
        self.children: dict[str, TrieNode] = {}
        self.is_word = False


# Variant 1: trie with dictionary-backed children.
# Example problems: implement trie, search suggestions, replace words.
# Time: O(L) per insert/search, where L is the word/prefix length.
# Space: O(total characters inserted)
class Trie:
    def __init__(self) -> None:
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.is_word = True

    def _walk(self, prefix: str) -> TrieNode | None:
        node = self.root
        for ch in prefix:
            if ch not in node.children:
                return None
            node = node.children[ch]
        return node

    def search(self, word: str) -> bool:
        node = self._walk(word)
        return node is not None and node.is_word

    def starts_with(self, prefix: str) -> bool:
        return self._walk(prefix) is not None


# Variant 2: trie with wildcard search.
# Example problems: design add and search words data structure.
# Time: O(L) without wildcards; worst-case O(number of trie nodes) with many "."
# Space: O(total characters inserted)
class WordDictionary:
    def __init__(self) -> None:
        self.root = TrieNode()

    def add_word(self, word: str) -> None:
        node = self.root
        for ch in word:
            node.children.setdefault(ch, TrieNode())
            node = node.children[ch]
        node.is_word = True

    def search(self, pattern: str) -> bool:
        def dfs(i: int, node: TrieNode) -> bool:
            if i == len(pattern):
                return node.is_word

            ch = pattern[i]
            if ch == ".":
                return any(dfs(i + 1, child) for child in node.children.values())

            if ch not in node.children:
                return False
            return dfs(i + 1, node.children[ch])

        return dfs(0, self.root)


class TreeNode:
    def __init__(
        self,
        val: int = 0,
        left: TreeNode | None = None,
        right: TreeNode | None = None,
    ) -> None:
        self.val = val
        self.left = left
        self.right = right


# Variant 3: build a binary tree from LeetCode-style level-order values.
# Example problems: local testing helpers, deserialize binary tree inputs.
# Time: O(n)
# Space: O(n)
def build_binary_tree_from_level_order(values: list[int | None]) -> TreeNode | None:
    if not values or values[0] is None:
        return None

    root = TreeNode(values[0])
    queue = deque([root])
    i = 1

    while queue and i < len(values):
        node = queue.popleft()

        if i < len(values) and values[i] is not None:
            node.left = TreeNode(values[i])
            queue.append(node.left)
        i += 1

        if i < len(values) and values[i] is not None:
            node.right = TreeNode(values[i])
            queue.append(node.right)
        i += 1

    return root


# Variant 4: rebuild binary tree from preorder and inorder traversals.
# Example problems: construct binary tree from preorder/inorder.
# Time: O(n)
# Space: O(n)
def build_tree_from_preorder_inorder(
    preorder: list[int],
    inorder: list[int],
) -> TreeNode | None:
    inorder_index = {value: i for i, value in enumerate(inorder)}
    preorder_i = 0

    def build(left: int, right: int) -> TreeNode | None:
        nonlocal preorder_i
        if left > right:
            return None

        root_value = preorder[preorder_i]
        preorder_i += 1

        root = TreeNode(root_value)
        split = inorder_index[root_value]
        root.left = build(left, split - 1)
        root.right = build(split + 1, right)
        return root

    return build(0, len(inorder) - 1)


# Variant 5: orient undirected edges into a rooted tree.
# Example problems: tree DP, subtree sizes, collect coins in a tree.
# Time: O(n + e)
# Space: O(n + e)
def build_rooted_tree(
    n: int,
    edges: list[tuple[int, int]],
    root: int = 0,
) -> list[list[int]]:
    graph = [[] for _ in range(n)]
    for a, b in edges:
        graph[a].append(b)
        graph[b].append(a)

    children = [[] for _ in range(n)]
    stack = [(root, -1)]

    while stack:
        node, parent = stack.pop()
        for nei in graph[node]:
            if nei == parent:
                continue
            children[node].append(nei)
            stack.append((nei, node))

    return children
