"""
Python data-structure patterns for coding interviews.

Use this file when you know the algorithm but need to recall the Pythonic way
to initialize, update, and choose common containers.
"""

from collections import Counter, defaultdict, deque
from heapq import heappop, heappush
from typing import Iterable


# -----------------------------------------------------------------------------
# list: dynamic array, stack, matrix rows
# -----------------------------------------------------------------------------


def list_patterns(nums: list[int]) -> tuple[list[int], list[int], list[list[int]]]:
    # Empty and fixed-size lists.
    values: list[int] = []
    zeros = [0] * len(nums)

    for x in nums:
        values.append(x)  # Amortized O(1)

    # Stack usage: append/pop from the end are O(1).
    stack: list[int] = []
    for x in nums:
        stack.append(x)
    while stack:
        stack.pop()

    # Do not build a matrix with [[0] * cols] * rows: rows share one inner list.
    rows, cols = 2, 3
    matrix = [[0] * cols for _ in range(rows)]

    return values, zeros, matrix


def list_slicing_cost(nums: list[int]) -> list[int]:
    # Slicing copies. Useful, but it is O(k) time and space for slice length k.
    middle = nums[1:-1]
    reversed_copy = nums[::-1]

    # In-place reverse avoids an extra list.
    nums.reverse()

    return middle + reversed_copy


# -----------------------------------------------------------------------------
# dict: hash map, first seen, grouping, counting
# -----------------------------------------------------------------------------


def dict_patterns(words: list[str]) -> tuple[dict[str, int], dict[str, list[str]]]:
    first_index: dict[str, int] = {}
    groups: dict[str, list[str]] = {}

    for i, word in enumerate(words):
        if word not in first_index:
            first_index[word] = i

        key = "".join(sorted(word))
        if key not in groups:
            groups[key] = []
        groups[key].append(word)

    return first_index, groups


def defaultdict_patterns(words: list[str]) -> dict[str, list[str]]:
    # defaultdict calls the factory only when a missing key is read through
    # __getitem__, such as groups[key]. It does not call the factory for get().
    groups: defaultdict[str, list[str]] = defaultdict(list)

    for word in words:
        key = "".join(sorted(word))
        groups[key].append(word)

    # Return as a regular dict when callers do not need missing-key behavior.
    return dict(groups)


def counter_patterns(items: Iterable[str]) -> Counter[str]:
    # Counter is a dict subclass specialized for frequencies. Missing keys read
    # as 0, which is why counts[ch] += 1 works without initialization.
    counts = Counter(items)

    # Most common operations.
    counts.update(["extra"])
    counts.subtract(["extra"])

    # Counter keeps zero/negative counts until deleted or normalized.
    counts += Counter()

    return counts


# -----------------------------------------------------------------------------
# set: membership, deduplication, visited states
# -----------------------------------------------------------------------------


def set_patterns(nums: list[int]) -> tuple[bool, set[int]]:
    seen: set[int] = set()

    for x in nums:
        if x in seen:
            return True, seen
        seen.add(x)

    unique = set(nums)
    return False, unique


# -----------------------------------------------------------------------------
# tuple: immutable compound keys
# -----------------------------------------------------------------------------


def tuple_keys(grid: list[list[int]]) -> set[tuple[int, int]]:
    visited: set[tuple[int, int]] = set()

    for r, row in enumerate(grid):
        for c, value in enumerate(row):
            if value:
                visited.add((r, c))

    return visited


# -----------------------------------------------------------------------------
# deque: real queue for BFS / sliding windows
# -----------------------------------------------------------------------------


def queue_patterns(start: int, graph: dict[int, list[int]]) -> list[int]:
    # deque is a double-ended queue. Both append/pop on either end are O(1),
    # making it the standard Python choice for BFS queues.
    queue = deque([start])
    seen = {start}
    order: list[int] = []

    while queue:
        node = queue.popleft()  # O(1); list.pop(0) would be O(n).
        order.append(node)

        for nei in graph.get(node, []):
            if nei not in seen:
                seen.add(nei)
                queue.append(nei)

    return order


# -----------------------------------------------------------------------------
# heapq: min-heap; negate priorities for max-heap behavior
# -----------------------------------------------------------------------------


def heap_patterns(nums: list[int], k: int) -> list[int]:
    # heapq functions operate on a normal list and maintain the heap invariant:
    # heap[0] is always the smallest item, but the whole list is not sorted.
    heap: list[int] = []

    for x in nums:
        heappush(heap, x)

    smallest = [heappop(heap) for _ in range(min(k, len(heap)))]

    max_heap: list[int] = []
    for x in nums:
        heappush(max_heap, -x)
    largest = [-heappop(max_heap) for _ in range(min(k, len(max_heap)))]

    return smallest + largest


# -----------------------------------------------------------------------------
# Common interview choice guide
# -----------------------------------------------------------------------------


"""
Choose quickly:

- Need order by insertion and random access: list
- Need LIFO stack: list with append/pop
- Need FIFO queue: collections.deque
- Need membership / visited: set
- Need key -> value lookup: dict
- Need missing list/int values: defaultdict(list/int)
- Need frequencies and multiset math: Counter
- Need repeated min/max extraction: heapq
- Need coordinate/state key: tuple

Version notes:

- Python 3.9+: built-in generics like list[int], dict[str, int].
- Python 3.10+: union syntax like int | None.
- Older interviews/environments may expect typing.List and typing.Optional.

Concept explanations:

- `list` is Python's dynamic array. It is good for random access, append/pop at
  the end, stacks, and dense DP tables. Inserting or removing near the front is
  O(n) because elements must shift.
- `[0] * n` is safe for immutable values like integers. `[[0] * cols] * rows`
  is unsafe because each row references the same inner list.
- `dict` is Python's built-in hash map. Keys must be hashable, so integers,
  strings, tuples of hashable values, and frozen dataclasses work; lists and
  dicts do not.
- `dict.setdefault(key, default)` can initialize a missing key inline, but
  `defaultdict` is usually clearer when every missing key uses the same default.
- `defaultdict(list)` means "when a missing key is accessed, create a new list
  for that key." Common factories are `list`, `int`, `set`, and `deque`.
- `Counter(iterable)` builds a frequency map. It is great for anagrams and
  multiset-style comparisons, but a plain `dict` can be clearer when you need
  custom state transitions.
- `set` is a hash table with only keys. Use it for `seen`, `visited`,
  membership checks, and deduplication.
- `tuple` is immutable and hashable if all elements are hashable. This is why
  grid coordinates are usually stored as `(r, c)` rather than `[r, c]`.
- `deque` is from `collections`, not the language core. It is the right queue
  abstraction for BFS because `popleft()` is O(1).
- `heapq` is a module, not a heap class. You pass a list to `heappush` and
  `heappop`. Python only provides a min-heap directly; use negative priorities
  or reversed comparison fields for max-heap behavior.
"""
