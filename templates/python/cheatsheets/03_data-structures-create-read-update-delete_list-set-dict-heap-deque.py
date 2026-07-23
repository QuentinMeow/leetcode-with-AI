"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

import collections
import heapq
import queue

# ====================================================================
# 3. Python Data Structures - CRUD Cheat Sheet
# ====================================================================

"""
Python Data Structures - CRUD Cheat Sheet

Big naming logic:

| Concept          | Pattern                       |
|------------------|-------------------------------|
| Add one item     | append / add / heapq.heappush |
| Add many items   | extend / update               |
| Remove + return  | pop / popleft / heapq.heappop |
| Remove by value  | remove / discard / del        |
| Remove all       | clear                         |
| Contains?        | in                            |
| Soft read/delete | get / discard / pop(default)  |

Why the names stick:

- `pop*` means "take it out and give it back".
- `remove` / `discard` delete a known value.
- `del` / `pop(key)` delete a known dict key.
- `clear` empties the container.
- `in` is the universal membership test.

Master memory table:

Core containers:

| Operation       | list   | set            | dict          |
|-----------------|--------|----------------|---------------|
| Empty           | []     | set()          | {}            |
| Add one         | append | add            | d[k] = v      |
| Add many        | extend | update         | update        |
| Read            | lst[i] | x in s         | d[k] / get    |
| Remove + return | pop    | pop            | pop / popitem |
| Remove by value | remove | discard/remove | del / pop(k)  |
| Empty all       | clear  | clear          | clear         |

Priority and double-ended containers:

| Operation       | heapq          | collections.deque |
|-----------------|----------------|-------------------|
| Empty           | []             | collections.deque |
| Add one         | heapq.heappush | append/appendleft |
| Add many        | heapq.heapify  | extend/extendleft |
| Read            | heap[0]        | q[0] / q[-1]      |
| Remove + return | heapq.heappop  | pop / popleft     |
| Remove by value | rebuild        | remove            |
| Empty all       | clear          | clear             |

Mental model: `pop*` returns what it removes; `remove`, `discard`,
`del`, `clear`, `sort`, and `reverse` mutate and return None.
"""




# Shows list creation, indexing, copying, stack use, mutation, deletion, and their main costs.
def list_create_read_update_delete_examples(
    n: int, rows: int, cols: int, nums: list[int]
) -> None:
    x = 2

    # CREATE
    lst: list[int] = []  # Empty dynamic array / stack.
    from_iterable = list(nums)
    zeros = [0] * n
    one_to_n = list(range(1, n + 1))

    # Matrix create: make fresh inner lists.
    grid = [[0] * cols for _ in range(rows)]
    visited = [[False for _ in range(cols)] for _ in range(rows)]
    bad_grid = [[0] * cols] * rows  # AVOID: every row aliases.

    # READ
    first = nums[0]  # IndexError if empty.
    last = nums[-1]
    middle = nums[1:3]  # Slice returns a new shallow list.
    size = len(nums)
    contains_x = x in nums
    first_x = nums.index(x) if x in nums else -1
    count_x = nums.count(x)

    # COPY: all are shallow for a flat list.
    copy1 = nums[:]
    copy2 = nums.copy()
    copy3 = list(nums)
    matrix_copy = [row[:] for row in grid]

    # UPDATE / ADD
    lst.append(x)  # Add one to the right end.
    lst.extend([4, 5])  # Add many to the right end.
    lst.insert(1, x)  # Insert before index 1.
    nums[0] = 99
    nums.sort()  # In-place; returns None.
    nums.reverse()  # In-place; returns None.
    sorted_nums = sorted(nums)  # New list.
    reversed_nums = nums[::-1]  # New list.

    # DELETE
    removed_last = lst.pop()  # Remove and return last item.
    removed_at_index = lst.pop(0)  # O(n): shifts everything left.
    if x in lst:
        lst.remove(x)  # Remove first matching value; returns None.
    del lst[:1]  # Delete by index/slice.
    lst.clear()

    # Stack pattern: right end only, O(1).
    stack: list[int] = []
    stack.append(1)
    top = stack[-1]
    popped = stack.pop()


# Shows set creation, membership, algebra, updates, deletion, and hashable tuple keys.
def set_create_read_update_delete_examples(nums: list[int], grid: list[list[int]]) -> None:
    x = 2

    # CREATE
    s: set[int] = set()  # Empty set. `{}` is an empty dict.
    from_iterable = set(nums)
    with_data = {1, 2, 3}

    # READ
    exists = x in with_data  # Average O(1); sets have no indexing.
    size = len(with_data)

    # UPDATE / ADD
    s.add(x)  # Add one.
    s.update([3, 4, 5])  # Add many from any iterable.
    union = {1, 2} | {2, 3}
    intersection = {1, 2} & {2, 3}
    difference = {1, 2} - {2, 3}
    symmetric_difference = {1, 2} ^ {2, 3}

    # DELETE
    s.discard(x)  # Soft: no error if absent.
    # s.remove(x)  # Strict: KeyError if absent.
    arbitrary = s.pop() if s else None  # Remove and return any item.
    s.clear()

    # Tuple keys are hashable when all parts are hashable.
    visited: set[tuple[int, int]] = set()
    for r, row in enumerate(grid):
        for c, value in enumerate(row):
            if value:
                visited.add((r, c))


# Shows dictionary creation, strict and soft reads, grouping, counting, merging, and deletion.
# Requires: import collections
def dictionary_create_read_update_delete_examples(words: list[str], nums: list[int]) -> None:
    # CREATE
    d: dict[str, int] = {}
    with_data = {"a": 1, "b": 2}
    from_keywords = dict(a=1, b=2)
    from_pairs = dict(zip(["a", "b"], [1, 2]))

    # READ
    strict = with_data["a"]  # KeyError if missing.
    soft_none = with_data.get("missing")  # None if missing.
    soft_default = with_data.get("missing", 0)
    exists = "a" in with_data  # Tests keys, not values.
    size = len(with_data)
    keys = list(with_data.keys())
    values = list(with_data.values())
    pairs = list(with_data.items())

    # UPDATE / ADD
    d["c"] = 3  # Set or overwrite one key.
    d.update({"d": 4, "e": 5})  # Merge many; right side wins.
    merged_modern = {"a": 1} | {"a": 9, "b": 2}
    merged_old = {"a": 1}
    merged_old.update({"a": 9, "b": 2})

    # Grouping: setdefault is a dict method.
    groups1: dict[str, list[str]] = {}
    for w in words:
        groups1.setdefault("".join(sorted(w)), []).append(w)

    # defaultdict(<factory>): missing key creates factory().
    groups2: collections.defaultdict[str, list[str]]
    groups2 = collections.defaultdict(list)
    for w in words:
        groups2["".join(sorted(w))].append(w)

    counts1: dict[int, int] = {}
    for x in nums:
        counts1[x] = counts1.get(x, 0) + 1

    counts2: collections.defaultdict[int, int]
    counts2 = collections.defaultdict(int)
    for x in nums:
        counts2[x] += 1

    # Counter(<iterable>): frequency dict with 0 for missing keys.
    counts3 = collections.Counter(nums)
    top_three = counts3.most_common(3)
    counts3.update([10])
    counts3.subtract([10])
    counts3 += collections.Counter()  # Drop zero/negative counts.

    # DELETE
    removed_value = d.pop("c")  # KeyError if missing.
    removed_or_default = d.pop("missing", 0)
    last_pair = d.popitem()  # Remove and return last inserted pair.
    del with_data["a"]
    d.clear()


# Shows heap creation, minimum access, insertion, removal, replacement, and max-heap negation.
# Requires: import heapq
def heap_create_read_update_delete_examples(nums: list[int]) -> None:
    # CREATE
    heap: list[int] = nums[:]
    heapq.heapify(heap)  # Turn a list into a min-heap in-place, O(n).
    built_by_push: list[int] = []
    for x in nums:
        heapq.heappush(built_by_push, x)  # O(log n) each.

    # READ
    smallest = heap[0] if heap else None  # Peek; do not remove.
    size = len(heap)

    # UPDATE / ADD
    heapq.heappush(heap, 42)
    max_heap = [-x for x in nums]  # Python heapq is min-heap only.
    heapq.heapify(max_heap)
    largest = -heapq.heappop(max_heap)

    # DELETE
    removed_smallest = heapq.heappop(heap) if heap else None
    # No efficient remove-by-value: rebuild, or use lazy deletion
    # with a separate "deleted" counter in advanced problems.
    heap.clear()


# Shows double-ended queue operations for constant-time work at either end.
# Requires: import collections
def double_ended_queue_create_read_update_delete_examples(nums: list[int]) -> None:
    x = 2

    # CREATE
    q: collections.deque[int] = collections.deque()
    with_data = collections.deque(nums)
    bounded = collections.deque(maxlen=3)  # Auto-drops old items.

    # READ
    left = with_data[0] if with_data else None
    right = with_data[-1] if with_data else None
    size = len(with_data)
    contains_x = x in with_data

    # UPDATE / ADD
    q.append(1)  # Add right.
    q.appendleft(0)  # Add left.
    q.extend([2, 3])  # Add many to right.
    q.extendleft([-1, -2])  # Adds left one-by-one; order reverses.
    q.rotate(1)  # Move rightmost item to the left.

    # DELETE
    removed_right = q.pop()
    removed_left = q.popleft()
    if x in q:
        q.remove(x)  # Remove first matching value.
    q.clear()

    # Usage patterns:
    fifo = collections.deque([1])
    fifo.append(2)
    next_in_line = fifo.popleft()  # Queue: FIFO.

    stack = collections.deque([1])
    stack.append(2)
    top = stack.pop()  # Stack: LIFO.


# Shows synchronized FIFO, LIFO, and priority queues intended for communication between threads.
# Requires: import queue
def thread_safe_queue_create_read_update_delete_examples() -> None:
    # CREATE: for multi-threading; prefer deque/heapq in algorithms.
    fifo = queue.Queue()
    stack = queue.LifoQueue()
    priority = queue.PriorityQueue()

    # UPDATE / ADD
    fifo.put(1)
    stack.put(1)
    priority.put((0, "task"))

    # READ
    fifo_empty = fifo.empty()
    fifo_size = fifo.qsize()  # Approximate under concurrency.

    # DELETE
    fifo_item = fifo.get()
    stack_item = stack.get()
    priority_item = priority.get()
