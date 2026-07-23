"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

import collections

# ====================================================================
# 15. Least-Recently-Used Cache Design
# ====================================================================

# CacheNode stores one cache entry plus links in the recency list.
class CacheNode:
    # Initializes a new instance and establishes its invariants.
    def __init__(self, key: int = 0, val: int = 0) -> None:
        self.key = key
        self.val = val
        self.prev: CacheNode | None = None
        self.next: CacheNode | None = None


# Least-recently-used cache using a map plus doubly linked recency list for O(1) operations.
class LeastRecentlyUsedCacheDoublyLinkedList:
    # Initializes a new instance and establishes its invariants.
    def __init__(self, capacity: int) -> None:
        self.capacity = capacity
        self.nodes: dict[int, CacheNode] = {}
        self.head = CacheNode()
        self.tail = CacheNode()
        self.head.next = self.tail
        self.tail.prev = self.head

    # Internal helper for remove.
    def _remove(self, node: CacheNode) -> None:
        prev, next_node = node.prev, node.next
        assert prev is not None and next_node is not None
        prev.next = next_node
        next_node.prev = prev

    # Internal helper for add to back.
    def _add_to_back(self, node: CacheNode) -> None:
        prev = self.tail.prev
        assert prev is not None
        node.prev, node.next = prev, self.tail
        prev.next = self.tail.prev = node

    # get returns a cached value and marks its entry most recently used.
    def get(self, key: int) -> int:
        if key not in self.nodes:
            return -1
        node = self.nodes[key]
        self._remove(node)
        self._add_to_back(node)
        return node.val

    # put inserts or updates a value and evicts the least recent entry when full.
    def put(self, key: int, value: int) -> None:
        if key in self.nodes:
            self._remove(self.nodes[key])
        node = CacheNode(key, value)
        self.nodes[key] = node
        self._add_to_back(node)
        if len(self.nodes) > self.capacity:
            victim = self.head.next
            assert victim is not None and victim is not self.tail
            self._remove(victim)
            del self.nodes[victim.key]




# Least-recently-used cache using OrderedDict recency operations in O(1).
# Requires: import collections
class LeastRecentlyUsedCacheOrderedDictionary:
    # Initializes a new instance and establishes its invariants.
    # Requires: import collections
    def __init__(self, capacity: int) -> None:
        self.capacity = capacity
        self.data: collections.OrderedDict[int, int]
        self.data = collections.OrderedDict()

    # get returns a cached value and marks its entry most recently used.
    def get(self, key: int) -> int:
        if key not in self.data:
            return -1
        self.data.move_to_end(key)
        return self.data[key]

    # put inserts or updates a value and evicts the least recent entry when full.
    def put(self, key: int, value: int) -> None:
        if key in self.data:
            self.data.move_to_end(key)
        self.data[key] = value
        if len(self.data) > self.capacity:
            self.data.popitem(last=False)
