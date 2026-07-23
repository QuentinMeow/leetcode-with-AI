"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

import heapq

# ====================================================================
# 17. Stacks and Heap-Based Merging
# ====================================================================

# Checks that round, square, and curly brackets close in last-in-first-out order.
def has_valid_bracket_nesting(s: str) -> bool:
    close_to_open = {")": "(", "]": "[", "}": "{"}
    stack: list[str] = []
    for ch in s:
        if ch in close_to_open.values():
            stack.append(ch)
        elif ch in close_to_open:
            if not stack or stack.pop() != close_to_open[ch]:
                return False
    return not stack


# Stores each value beside the minimum at that depth for O(1) minimum queries.
class StackWithMinimum:
    # Initializes a new instance and establishes its invariants.
    def __init__(self) -> None:
        self.stack: list[tuple[int, int]] = []

    # push stores the value and the minimum at this new stack depth.
    def push(self, val: int) -> None:
        cur_min = (
            val if not self.stack else min(val, self.stack[-1][1])
        )
        self.stack.append((val, cur_min))

    # pop removes and returns the newest value.
    def pop(self) -> int:
        return self.stack.pop()[0]

    # top returns the newest value without removing it.
    def top(self) -> int:
        return self.stack[-1][0]

    # get_min returns the current minimum without scanning the stack.
    def get_min(self) -> int:
        return self.stack[-1][1]




# Performs a k-way merge by storing one current value from each sorted array. Time O(total log
# k).
# Requires: import heapq
def merge_sorted_arrays_using_min_heap(arrays: list[list[int]]) -> list[int]:
    heap: list[tuple[int, int, int]] = []
    for arr_i, arr in enumerate(arrays):
        if arr:
            heapq.heappush(heap, (arr[0], arr_i, 0))
    result: list[int] = []
    while heap:
        val, arr_i, elem_i = heapq.heappop(heap)
        result.append(val)
        next_node = elem_i + 1
        if next_node < len(arrays[arr_i]):
            heapq.heappush(heap, (arrays[arr_i][next_node], arr_i, next_node))
    return result
