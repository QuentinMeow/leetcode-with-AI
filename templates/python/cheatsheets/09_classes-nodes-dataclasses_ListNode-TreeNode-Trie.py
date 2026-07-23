"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

import dataclasses

# ====================================================================
# 9. Classes / Nodes / Dataclasses
# ====================================================================

# Singly linked-list node matching common interview problem definitions.
class ListNode:
    # Initializes a new instance and establishes its invariants.
    def __init__(
        self, val: int = 0, next: "ListNode | None" = None
    ) -> None:
        self.val = val
        self.next = next


# Binary-tree node matching common interview problem definitions.
class TreeNode:
    # Initializes a new instance and establishes its invariants.
    def __init__(
        self,
        val: int = 0,
        left: "TreeNode | None" = None,
        right: "TreeNode | None" = None,
    ) -> None:
        self.val = val
        self.left = left
        self.right = right




# Interval is a closed numeric range from start through end.
# Requires: import dataclasses
@dataclasses.dataclass
class Interval:
    # @dataclasses.dataclass reads fields and generates __init__,
    # __repr__, __eq__.
    start: int
    end: int




# Point is an immutable grid coordinate usable as a set or dictionary key.
# Requires: import dataclasses
@dataclasses.dataclass(frozen=True)
class Point:
    # frozen=True makes value objects immutable/hashable if
    # fields are hashable.
    row: int
    col: int




# Task compares by priority and then name because dataclass order follows field order.
# Requires: import dataclasses
@dataclasses.dataclass(order=True)
class Task:
    # order=True makes comparisons use fields in definition order.
    priority: int
    name: str
    tags: list[str] = dataclasses.field(default_factory=list)




# TrieNode stores outgoing character edges and whether a complete word ends here.
# Requires: import dataclasses
@dataclasses.dataclass
# Requires via helper: import dataclasses
class TrieNode:
    children: dict[str, "TrieNode"] = dataclasses.field(
        default_factory=dict
    )
    is_word: bool = False


# Trie, also called a prefix tree, stores words by shared character prefixes.
class Trie:
    # Initializes a new instance and establishes its invariants.
    def __init__(self) -> None:
        self.root = TrieNode()

    # insert creates missing prefix nodes and marks the final node as a word.
    def insert(self, word: str) -> None:
        node = self.root
        for ch in word:
            node = node.children.setdefault(ch, TrieNode())
        node.is_word = True

    # search requires every character edge and an end-of-word marker.
    def search(self, word: str) -> bool:
        node = self.root
        for ch in word:
            if ch not in node.children:
                return False
            node = node.children[ch]
        return node.is_word


# Object notes:
#   `is` checks identity; `==` checks value equality.
#   Mutable class variables are shared by every instance; use
#   `self.x` in __init__.
#   `dataclasses.field(default_factory=list)` gives every dataclass
#   instance a fresh list.
