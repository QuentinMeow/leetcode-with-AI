"""
Python interview cheatsheet.

Use this as the fast scan before an interview. The deeper explanations live in:

- language-specific-pattern/*.py: Python syntax, containers, libraries, OOP.
- algorithms/*.py: reusable algorithm templates by pattern.

This file is intentionally compact: adjacent lines often show equivalent ways
to do the same thing, plus the small caveats that cause interview bugs.
"""

# ruff: noqa: E501, F401, F841, B006
# pyright: reportUnusedImport=false, reportUnusedVariable=false

from __future__ import annotations

import heapq
import math
import re
import sys
from bisect import bisect_left, bisect_right, insort
from collections import Counter, OrderedDict, defaultdict, deque
from dataclasses import dataclass, field
from functools import cache, cmp_to_key, lru_cache
from heapq import heapify, heappop, heappush, nlargest, nsmallest
from itertools import accumulate, combinations, pairwise, permutations, product
from operator import attrgetter, itemgetter
from typing import Iterable


# =============================================================================
# 0. Version / Syntax Quick Map
# =============================================================================

"""
Python 3.8+:  walrus `:=`, `int.bit_count()`
Python 3.9+:  `list[int]`, `dict[str, int]`, `dict_a | dict_b`, `@cache`
Python 3.10+: `int | None`, `match/case`, `itertools.pairwise`

Older equivalents:
    list[int]        -> typing.List[int]
    dict[str, int]   -> typing.Dict[str, int]
    int | None       -> typing.Optional[int]
    @cache           -> @lru_cache(maxsize=None)
"""


# =============================================================================
# 1. Common Imports
# =============================================================================

"""
collections: Counter, OrderedDict, defaultdict, deque
heapq:       heappush, heappop, heapify, nlargest, nsmallest
bisect:      bisect_left, bisect_right, insort
functools:   cache, lru_cache, cmp_to_key
itertools:   accumulate, pairwise, product, combinations, permutations
math:        gcd, lcm, isqrt, comb, inf, ceil
dataclasses: dataclass, field
operator:    itemgetter, attrgetter
re:          findall, sub
"""


# =============================================================================
# 2. Containers: Init, Copy, Mutate
# =============================================================================


def list_array_matrix_patterns(n: int, rows: int, cols: int, nums: list[int]) -> None:
    # Empty list / dynamic array / stack.
    a: list[int] = []
    b = list()

    # Fixed-size lists.
    zeros = [0] * n
    falses = [False] * n
    none_slots = [None] * n
    one_to_n = list(range(1, n + 1))

    # Copying: all are shallow for a flat list.
    copy1 = nums[:]
    copy2 = nums.copy()
    copy3 = list(nums)

    # Matrix: make fresh inner lists.
    grid = [[0] * cols for _ in range(rows)]
    visited = [[False for _ in range(cols)] for _ in range(rows)]
    bad_grid = [[0] * cols] * rows  # AVOID: every row aliases the same list.

    # Stack: use end of list.
    stack: list[int] = []
    stack.append(1)
    top = stack[-1]
    popped = stack.pop()

    # Queue: use deque, not list.pop(0).
    q = deque([0])
    q.append(1)
    left = q.popleft()

    # In-place operations that return None.
    nums.sort()
    nums.reverse()

    # New-list operations.
    sorted_nums = sorted(nums)
    reversed_nums = nums[::-1]
    reversed_iter = reversed(nums)


def dict_defaultdict_counter_patterns(words: list[str], nums: list[int]) -> None:
    # Empty dict / hash map.
    d: dict[str, int] = {}
    e = dict()

    # Insert/update/read.
    d["a"] = 1
    d["a"] = d.get("a", 0) + 1
    exists = "a" in d
    value_or_default = d.get("missing", 0)
    removed = d.pop("a", None)

    # setdefault vs defaultdict: keep the alternatives adjacent.
    groups1: dict[str, list[str]] = {}
    for w in words:
        groups1.setdefault("".join(sorted(w)), []).append(w)

    # defaultdict(<factory>): when groups2[key] is missing, Python creates
    # groups2[key] = list(), then returns that list so append can run.
    groups2: defaultdict[str, list[str]] = defaultdict(list)
    for w in words:
        groups2["".join(sorted(w))].append(w)

    counts1: dict[int, int] = {}
    for x in nums:
        counts1[x] = counts1.get(x, 0) + 1

    # defaultdict(int): int() returns 0, so missing counts start from 0.
    counts2: defaultdict[int, int] = defaultdict(int)
    for x in nums:
        counts2[x] += 1

    # Counter(<iterable>): consume items and build a frequency map.
    # Unlike a normal dict, counts3[missing_key] reads as 0.
    counts3 = Counter(nums)
    top_three = counts3.most_common(3)
    counts3.update([10])
    counts3.subtract([10])
    counts3 += Counter()  # Drop zero/negative counts.

    # Iteration views.
    keys = list(d.keys())
    values = list(d.values())
    pairs = list(d.items())

    # Dict merge: right side wins.
    merged_modern = {"a": 1} | {"a": 9, "b": 2}
    merged_old = {"a": 1}
    merged_old.update({"a": 9, "b": 2})


def set_tuple_key_patterns(nums: list[int], grid: list[list[int]]) -> None:
    seen: set[int] = set()
    unique = set(nums)

    seen.add(1)
    seen.discard(1)  # No error if absent.
    # seen.remove(1)  # Raises KeyError if absent.

    union = {1, 2} | {2, 3}
    intersection = {1, 2} & {2, 3}
    difference = {1, 2} - {2, 3}

    # Tuple keys are hashable when all parts are hashable.
    visited: set[tuple[int, int]] = set()
    for r, row in enumerate(grid):
        for c, value in enumerate(row):
            if value:
                visited.add((r, c))


# =============================================================================
# 3. Copying / Binding / Mutability
# =============================================================================


def copy_and_aliasing(grid: list[list[int]], nums: list[int]) -> None:
    alias = nums  # Same list object.
    shallow1 = nums[:]
    shallow2 = nums.copy()
    shallow3 = list(nums)

    matrix_shallow = grid[:]  # New outer list, shared rows.
    matrix_copy = [row[:] for row in grid]  # Common matrix copy.

    path: list[int] = []
    result: list[list[int]] = []
    result.append(path.copy())  # Save a snapshot in backtracking.
    # result.append(path)       # AVOID: later path mutations affect saved row.

    # Mutable default args: avoid `def f(bucket=[]): ...`; use None sentinel.


# =============================================================================
# 4. Sorting / Comparison / Heap
# =============================================================================


def sorting_patterns(items: list[tuple[str, int]], intervals: list[list[int]]) -> None:
    nums = [3, 1, 2]

    new_sorted = sorted(nums)
    nums.sort()  # Mutates; returns None.

    desc_all = sorted(nums, reverse=True)

    # sorted(<list>, key=lambda <argument>: <representation sorting key>)
    # <argument>: each element extracted from the list/iterable.
    # <representation sorting key>: value Python compares for that element.
    by_start_end = sorted(intervals, key=lambda x: (x[0], x[1]))
    # x is one interval like [start, end]; (x[0], x[1]) sorts by start, then end.

    score_desc_name_asc = sorted(items, key=lambda item: (-item[1], item[0]))
    # item is one tuple like (name, score); (-score, name) means score desc, name asc.

    # Stable two-pass sort: sort secondary first, primary second.
    records = [("a", 90, 20), ("b", 90, 18)]
    records = sorted(records, key=lambda r: r[2])  # r is each record; r[2] is age.
    records = sorted(records, key=lambda r: r[1], reverse=True)  # r[1] is score.
    records = sorted(records, key=itemgetter(2))  # same as lambda r: r[2]
    records = sorted(records, key=itemgetter(1), reverse=True)  # same as lambda r: r[1]

    best = max(items, key=lambda item: (item[1], item[0]))
    worst = min(items, key=lambda item: item[1])

    # Comparator fallback: this is the -1/0/+1 style from Java/C/C++.
    # Python 3 sorting does not accept compare= directly; wrap it with cmp_to_key.
    def compare(a: int, b: int) -> int:
        return (a % 2) - (b % 2) or a - b

    sorted_with_cmp = sorted(nums, key=cmp_to_key(compare))


def custom_sorting_patterns(words: list[str], transactions: list[TransactionRecord]) -> None:
    # 1) Lambda key: most common.
    # w is each string in words; (len(w), w) is the value used for ordering.
    by_len_then_word = sorted(words, key=lambda w: (len(w), w))

    # 2) Named key function: easier to debug/explain when the key is not tiny.
    def word_key(word: str) -> tuple[int, str]:
        return len(word), word

    by_named_key = sorted(words, key=word_key)

    # 3) Comparator function: use only when "a before b?" logic is naturally pairwise.
    # This function receives TWO items and returns negative / zero / positive.
    def compare_words(a: str, b: str) -> int:
        if len(a) != len(b):
            return len(a) - len(b)
        return -1 if a < b else (1 if a > b else 0)

    by_comparator = sorted(words, key=cmp_to_key(compare_words))

    # 4) Custom object ordering: sorted(objs) calls obj.__lt__(other).
    by_time = sorted(transactions)
    by_time_key = sorted(transactions, key=attrgetter("time"))  # same as lambda t: t.time

    # 5) Common accessors: itemgetter(0) is like lambda item: item[0].
    pairs = [("b", 2), ("a", 3)]
    by_first = sorted(pairs, key=itemgetter(0))
    by_first_then_second = sorted(pairs, key=itemgetter(0, 1))


class TransactionRecord:
    def __init__(self, raw: str) -> None:
        name, time, amount, city = raw.split(",")
        self.name = name
        self.time = int(time)
        self.amount = int(amount)
        self.city = city
        self.raw = raw

    def __lt__(self, other: TransactionRecord) -> bool:
        return (self.time, self.amount, self.city) < (other.time, other.amount, other.city)

    def __repr__(self) -> str:
        return f"TransactionRecord({self.raw!r})"


def bisect_patterns(sorted_nums: list[int], x: int) -> None:
    left = bisect_left(sorted_nums, x)  # First index with value >= x.
    right = bisect_right(sorted_nums, x)  # First index with value > x.
    count_x = right - left
    exists = left < len(sorted_nums) and sorted_nums[left] == x

    insert_at_left = bisect_left(sorted_nums, x)
    insert_at_right = bisect_right(sorted_nums, x)
    insort(sorted_nums, x)  # Finds position O(log n), inserts into list O(n).


def heap_patterns(nums: list[int], k: int, tasks: list[tuple[int, str]]) -> None:
    # Min-heap over a normal list.
    heap = nums[:]
    heapify(heap)
    smallest = heappop(heap)
    heappush(heap, 42)

    # Max-heap by negating numeric priorities.
    max_heap = [-x for x in nums]
    heapify(max_heap)
    largest = -heappop(max_heap)

    # Top-k one-off helpers.
    k_largest = nlargest(k, nums)
    k_smallest = nsmallest(k, nums)

    # Bounded min-heap for streaming kth largest.
    top: list[int] = []
    for x in nums:
        heappush(top, x)
        if len(top) > k:
            heappop(top)

    # Tuple heap: priority, tie breaker, payload.
    pq: list[tuple[int, int, str]] = []
    for order, (priority, name) in enumerate(tasks):
        heappush(pq, (priority, order, name))


# =============================================================================
# 5. Iteration / Comprehensions / Ranges
# =============================================================================


def iteration_patterns(nums: list[int], a: list[int], b: list[int]) -> None:
    for i, x in enumerate(nums):
        pass

    for i in range(len(nums)):
        pass

    for i in range(len(nums) - 1, -1, -1):
        pass

    pairs1 = list(pairwise(nums))  # Python 3.10+
    pairs2 = list(zip(nums, nums[1:]))  # Older equivalent; slice copies.

    zipped = list(zip(a, b))  # Stops at shortest input.
    dot = sum(x * y for x, y in zip(a, b))

    # [<output_expr> for <element> in <iterable> if <filter_condition>]
    # x is each number from nums; x * x is what gets stored.
    squares = [x * x for x in nums if x >= 0]

    # Braces with one expression make a set comprehension.
    unique_mods = {x % 10 for x in nums}

    # Braces with key: value make a dict comprehension.
    # enumerate(nums) yields (index, value), unpacked as i, x.
    index_by_value = {x: i for i, x in enumerate(nums)}

    # Generator expression: no list is built; sum/any/all pull values lazily.
    total_pos = sum(x for x in nums if x > 0)
    has_even = any(x % 2 == 0 for x in nums)
    all_positive = all(x > 0 for x in nums)

    # Star unpacking: middle becomes a list of the leftover values.
    first, *middle, last = nums if len(nums) >= 2 else [0, 0]
    a0, b0 = 1, 2
    a0, b0 = b0, a0


# =============================================================================
# 6. Strings / Numbers / Bits
# =============================================================================


def string_patterns(s: str, chars: list[str], nums: list[int]) -> None:
    words = s.split()  # Split on runs of whitespace.
    csv_parts = [part.strip() for part in s.split(",") if part.strip()]
    joined = " ".join(words)

    # Efficient build: list append + join.
    pieces: list[str] = []
    for ch in chars:
        pieces.append(ch)
    built = "".join(pieces)

    lower = s.lower()
    stripped = s.strip()
    replaced = s.replace("old", "new")
    starts = s.startswith("pre")
    ends = s.endswith("suf")

    idx = ord("c") - ord("a")
    ch = chr(ord("a") + idx)
    freq26 = [0] * 26
    for ch2 in s:
        if "a" <= ch2 <= "z":
            freq26[ord(ch2) - ord("a")] += 1

    ints_from_spaces = [int(part) for part in s.split()]
    ints_from_text = [int(m) for m in re.findall(r"-?\d+", s)]
    answer_line = ",".join(str(x) for x in nums)


def numeric_bit_patterns(a: int, b: int, n: int, mask: int) -> None:
    true_division = a / b
    floor_division = a // b  # Floors toward -inf.
    trunc_toward_zero = int(a / b)
    ceil_div_positive = -(-a // b)
    remainder = a % b  # If b > 0, result is in [0, b).

    inf = float("inf")
    neg_inf = float("-inf")
    also_inf = math.inf

    gcd = math.gcd(a, b)
    lcm = math.lcm(a, b)
    root_floor = math.isqrt(n)
    choose_two = math.comb(n, 2) if n >= 2 else 0

    one_bit = 1 << 3
    mask |= 1 << 3  # Set bit.
    has_bit = (mask & (1 << 3)) != 0
    mask &= ~(1 << 3)  # Clear bit.
    mask ^= 1 << 3  # Toggle bit.
    bits = mask.bit_count()


# =============================================================================
# 7. Functions / Scope / Decorators
# =============================================================================


def function_call_syntax(required: int, *args: int, **kwargs: int) -> None:
    collected_args = args  # tuple
    collected_kwargs = kwargs  # dict

    values = [1, 2, 3]
    options = {"required": 1}
    # function_call_syntax(*values)
    # function_call_syntax(**options)


def nested_helper_and_nonlocal(root: TreeNode | None) -> int:
    count = 0

    def dfs(node: TreeNode | None) -> None:
        # nonlocal means "when assigning count, use the count in the outer
        # function scope." Without it, count += 1 would create/read a local name.
        nonlocal count
        if node is None:
            return
        count += 1
        dfs(node.left)
        dfs(node.right)

    dfs(root)
    return count


@cache
def cached_dp(i: int, remaining: int) -> int:
    # @cache is decorator syntax: cached_dp = cache(cached_dp).
    # It memoizes by the argument tuple (i, remaining), so args must be hashable.
    if remaining == 0:
        return 1
    if i == 0 or remaining < 0:
        return 0
    return cached_dp(i - 1, remaining) + cached_dp(i - 1, remaining - 1)


@lru_cache(maxsize=None)
def cached_dp_old_spelling(i: int) -> int:
    return i if i <= 1 else cached_dp_old_spelling(i - 1) + cached_dp_old_spelling(i - 2)


# =============================================================================
# 8. Classes / Nodes / Dataclasses
# =============================================================================


class ListNode:
    def __init__(self, val: int = 0, next: ListNode | None = None) -> None:
        self.val = val
        self.next = next


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


@dataclass
class Interval:
    # @dataclass reads annotated fields and generates __init__, __repr__, __eq__.
    start: int
    end: int


@dataclass(frozen=True)
class Point:
    # frozen=True makes value objects immutable/hashable if fields are hashable.
    row: int
    col: int


@dataclass(order=True)
class Task:
    # order=True makes comparisons use fields in definition order.
    priority: int
    name: str
    tags: list[str] = field(default_factory=list)


@dataclass
class TrieNode:
    children: dict[str, TrieNode] = field(default_factory=dict)
    is_word: bool = False


class Trie:
    def __init__(self) -> None:
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        node = self.root
        for ch in word:
            node = node.children.setdefault(ch, TrieNode())
        node.is_word = True

    def search(self, word: str) -> bool:
        node = self.root
        for ch in word:
            if ch not in node.children:
                return False
            node = node.children[ch]
        return node.is_word


# Object notes:
#   `is` checks identity; `==` checks value equality.
#   Mutable class variables are shared by every instance; use `self.x` in __init__.
#   `field(default_factory=list)` gives every dataclass instance a fresh list.


# =============================================================================
# 9. Algorithm Skeletons
# =============================================================================


def two_sum_hash(nums: list[int], target: int) -> tuple[int, int] | None:
    seen: dict[int, int] = {}
    for i, x in enumerate(nums):
        if target - x in seen:
            return seen[target - x], i
        seen[x] = i
    return None


def two_pointers_sorted(nums: list[int], target: int) -> tuple[int, int] | None:
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


def sliding_window_at_most_k_distinct(nums: list[int], k: int) -> int:
    if k < 0:
        return 0
    count: defaultdict[int, int] = defaultdict(int)
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


def lower_bound(nums: list[int], target: int) -> int:
    left, right = 0, len(nums)
    while left < right:
        mid = (left + right) // 2
        if nums[mid] < target:
            left = mid + 1
        else:
            right = mid
    return left


def first_feasible(low: int, high: int, can) -> int:
    while low < high:
        mid = (low + high) // 2
        if can(mid):
            high = mid
        else:
            low = mid + 1
    return low


def prefix_sum_subarray_count(nums: list[int], k: int) -> int:
    seen: defaultdict[int, int] = defaultdict(int)
    seen[0] = 1
    prefix = ans = 0
    for x in nums:
        prefix += x
        ans += seen[prefix - k]
        seen[prefix] += 1
    return ans


def bfs_graph(graph: dict[int, list[int]], start: int, target: int) -> int:
    q = deque([(start, 0)])
    seen = {start}
    while q:
        node, dist = q.popleft()
        if node == target:
            return dist
        for nei in graph.get(node, []):
            if nei not in seen:
                seen.add(nei)
                q.append((nei, dist + 1))
    return -1


def grid_neighbors(r: int, c: int, rows: int, cols: int) -> Iterable[tuple[int, int]]:
    for dr, dc in ((1, 0), (-1, 0), (0, 1), (0, -1)):
        nr, nc = r + dr, c + dc
        if 0 <= nr < rows and 0 <= nc < cols:
            yield nr, nc


def dfs_island(grid: list[list[str]]) -> int:
    rows, cols = len(grid), len(grid[0]) if grid else 0

    def dfs(r: int, c: int) -> None:
        if not (0 <= r < rows and 0 <= c < cols) or grid[r][c] != "1":
            return
        grid[r][c] = "0"
        for nr, nc in grid_neighbors(r, c, rows, cols):
            dfs(nr, nc)

    count = 0
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == "1":
                count += 1
                dfs(r, c)
    return count


def backtracking_subsets(nums: list[int]) -> list[list[int]]:
    ans: list[list[int]] = []
    path: list[int] = []

    def backtrack(start: int) -> None:
        ans.append(path.copy())
        for i in range(start, len(nums)):
            path.append(nums[i])
            backtrack(i + 1)
            path.pop()

    backtrack(0)
    return ans


def monotonic_next_greater(nums: list[int]) -> list[int]:
    ans = [-1] * len(nums)
    stack: list[int] = []  # indexes; decreasing values
    for i, x in enumerate(nums):
        while stack and nums[stack[-1]] < x:
            ans[stack.pop()] = i
        stack.append(i)
    return ans


def merge_intervals(intervals: list[list[int]]) -> list[list[int]]:
    intervals.sort(key=lambda x: x[0])
    merged: list[list[int]] = []
    for start, end in intervals:
        if not merged or start > merged[-1][1]:
            merged.append([start, end])
        else:
            merged[-1][1] = max(merged[-1][1], end)
    return merged


def top_k_frequent(nums: list[int], k: int) -> list[int]:
    counts = Counter(nums)
    return [num for _, num in nlargest(k, ((freq, num) for num, freq in counts.items()))]


def coin_change_top_down(coins: tuple[int, ...], amount: int) -> int:
    inf = amount + 1

    @cache
    def dp(rem: int) -> int:
        if rem == 0:
            return 0
        if rem < 0:
            return inf
        return 1 + min(dp(rem - coin) for coin in coins)

    ans = dp(amount)
    return -1 if ans >= inf else ans


class DSU:
    def __init__(self, n: int) -> None:
        self.parent = list(range(n))
        self.size = [1] * n
        self.components = n

    def find(self, x: int) -> int:
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, a: int, b: int) -> bool:
        ra, rb = self.find(a), self.find(b)
        if ra == rb:
            return False
        if self.size[ra] < self.size[rb]:
            ra, rb = rb, ra
        self.parent[rb] = ra
        self.size[ra] += self.size[rb]
        self.components -= 1
        return True


# =============================================================================
# 10. High-Frequency Add-ons From Solutions / Hot Patterns
# =============================================================================


def set_based_unique_window(s: str) -> int:
    seen: set[str] = set()
    left = best = 0
    for right, ch in enumerate(s):
        while ch in seen:
            seen.remove(s[left])
            left += 1
        seen.add(ch)
        best = max(best, right - left + 1)
    return best


def min_subarray_len_at_least(nums: list[int], target: int) -> int:
    left = total = 0
    best = math.inf
    for right, x in enumerate(nums):
        total += x
        while total >= target:
            best = min(best, right - left + 1)
            total -= nums[left]
            left += 1
    return 0 if best == math.inf else int(best)


def fixed_window_max_sum(nums: list[int], k: int) -> int:
    window = sum(nums[:k])
    best = window
    for right in range(k, len(nums)):
        window += nums[right] - nums[right - k]
        best = max(best, window)
    return best


def exactly_k_distinct(nums: list[int], k: int) -> int:
    return sliding_window_at_most_k_distinct(nums, k) - sliding_window_at_most_k_distinct(nums, k - 1)


def three_sum(nums: list[int]) -> list[list[int]]:
    nums.sort()
    ans: list[list[int]] = []
    for i, x in enumerate(nums):
        if i > 0 and x == nums[i - 1]:
            continue
        left, right = i + 1, len(nums) - 1
        while left < right:
            total = x + nums[left] + nums[right]
            if total == 0:
                ans.append([x, nums[left], nums[right]])
                left += 1
                right -= 1
                while left < right and nums[left] == nums[left - 1]:
                    left += 1
                while left < right and nums[right] == nums[right + 1]:
                    right -= 1
            elif total < 0:
                left += 1
            else:
                right -= 1
    return ans


def move_zeroes(nums: list[int]) -> None:
    write = 0
    for read, x in enumerate(nums):
        if x != 0:
            nums[write], nums[read] = nums[read], nums[write]
            write += 1


def max_area_container(height: list[int]) -> int:
    left, right = 0, len(height) - 1
    best = 0
    while left < right:
        best = max(best, (right - left) * min(height[left], height[right]))
        if height[left] < height[right]:
            left += 1
        else:
            right -= 1
    return best


def count_palindromes_expand(s: str) -> int:
    def expand(left: int, right: int) -> int:
        count = 0
        while 0 <= left and right < len(s) and s[left] == s[right]:
            count += 1
            left -= 1
            right += 1
        return count

    return sum(expand(i, i) + expand(i, i + 1) for i in range(len(s)))


def search_rotated(nums: list[int], target: int) -> int:
    left, right = 0, len(nums) - 1
    while left <= right:
        mid = (left + right) // 2
        if nums[mid] == target:
            return mid
        if nums[left] <= nums[mid]:
            if nums[left] <= target < nums[mid]:
                right = mid - 1
            else:
                left = mid + 1
        else:
            if nums[mid] < target <= nums[right]:
                left = mid + 1
            else:
                right = mid - 1
    return -1


def median_two_sorted_partition(a: list[int], b: list[int]) -> float:
    if len(a) > len(b):
        a, b = b, a
    m, n = len(a), len(b)
    half = (m + n + 1) // 2
    left, right = 0, m
    while left <= right:
        i = (left + right) // 2
        j = half - i
        left1 = -math.inf if i == 0 else a[i - 1]
        right1 = math.inf if i == m else a[i]
        left2 = -math.inf if j == 0 else b[j - 1]
        right2 = math.inf if j == n else b[j]
        if left1 <= right2 and left2 <= right1:
            if (m + n) % 2:
                return float(max(left1, left2))
            return (max(left1, left2) + min(right1, right2)) / 2
        if left1 > right2:
            right = i - 1
        else:
            left = i + 1
    raise ValueError("inputs must be sorted")


def group_anagrams_count_key(words: list[str]) -> list[list[str]]:
    groups: defaultdict[tuple[int, ...], list[str]] = defaultdict(list)
    for word in words:
        count = [0] * 26
        for ch in word:
            count[ord(ch) - ord("a")] += 1
        groups[tuple(count)].append(word)
    return list(groups.values())


def longest_consecutive(nums: list[int]) -> int:
    values = set(nums)
    best = 0
    for x in values:
        if x - 1 not in values:
            y = x
            while y in values:
                y += 1
            best = max(best, y - x)
    return best


class PrefixSum1D:
    def __init__(self, nums: list[int]) -> None:
        self.prefix = [0]
        for x in nums:
            self.prefix.append(self.prefix[-1] + x)

    def sum_range(self, left: int, right: int) -> int:
        return self.prefix[right + 1] - self.prefix[left]


class PrefixSum2D:
    def __init__(self, matrix: list[list[int]]) -> None:
        rows = len(matrix)
        cols = len(matrix[0]) if matrix else 0
        self.prefix = [[0] * (cols + 1) for _ in range(rows + 1)]
        for r in range(rows):
            for c in range(cols):
                self.prefix[r + 1][c + 1] = (
                    matrix[r][c]
                    + self.prefix[r][c + 1]
                    + self.prefix[r + 1][c]
                    - self.prefix[r][c]
                )

    def sum_region(self, r1: int, c1: int, r2: int, c2: int) -> int:
        return (
            self.prefix[r2 + 1][c2 + 1]
            - self.prefix[r1][c2 + 1]
            - self.prefix[r2 + 1][c1]
            + self.prefix[r1][c1]
        )


def meeting_rooms_two_pointer(intervals: list[list[int]]) -> int:
    starts = sorted(start for start, _ in intervals)
    ends = sorted(end for _, end in intervals)
    rooms = best = 0
    s = e = 0
    while s < len(starts):
        if starts[s] < ends[e]:
            rooms += 1
            best = max(best, rooms)
            s += 1
        else:
            rooms -= 1
            e += 1
    return best


def insert_interval(intervals: list[list[int]], new_interval: list[int]) -> list[list[int]]:
    ans: list[list[int]] = []
    i = 0
    while i < len(intervals) and intervals[i][1] < new_interval[0]:
        ans.append(intervals[i])
        i += 1
    while i < len(intervals) and intervals[i][0] <= new_interval[1]:
        new_interval[0] = min(new_interval[0], intervals[i][0])
        new_interval[1] = max(new_interval[1], intervals[i][1])
        i += 1
    return ans + [new_interval] + intervals[i:]


def erase_overlap_intervals(intervals: list[list[int]]) -> int:
    intervals.sort(key=lambda x: x[1])
    removed = 0
    prev_end = -math.inf
    for start, end in intervals:
        if start >= prev_end:
            prev_end = end
        else:
            removed += 1
    return removed


def merge_two_lists(a: ListNode | None, b: ListNode | None) -> ListNode | None:
    dummy = tail = ListNode()
    while a and b:
        if a.val <= b.val:
            tail.next, a = a, a.next
        else:
            tail.next, b = b, b.next
        tail = tail.next
    tail.next = a or b
    return dummy.next


def reverse_list(head: ListNode | None) -> ListNode | None:
    prev = None
    cur = head
    while cur:
        nxt = cur.next
        cur.next = prev
        prev = cur
        cur = nxt
    return prev


def has_cycle(head: ListNode | None) -> bool:
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow is fast:
            return True
    return False


def remove_nth_from_end(head: ListNode | None, n: int) -> ListNode | None:
    dummy = ListNode(0, head)
    fast = slow = dummy
    for _ in range(n):
        fast = fast.next
    while fast and fast.next:
        fast = fast.next
        slow = slow.next
    if slow.next:
        slow.next = slow.next.next
    return dummy.next


class CacheNode:
    def __init__(self, key: int = 0, val: int = 0) -> None:
        self.key = key
        self.val = val
        self.prev: CacheNode | None = None
        self.next: CacheNode | None = None


class LRUCacheDLL:
    def __init__(self, capacity: int) -> None:
        self.capacity = capacity
        self.nodes: dict[int, CacheNode] = {}
        self.head = CacheNode()
        self.tail = CacheNode()
        self.head.next = self.tail
        self.tail.prev = self.head

    def _remove(self, node: CacheNode) -> None:
        prev, nxt = node.prev, node.next
        assert prev is not None and nxt is not None
        prev.next = nxt
        nxt.prev = prev

    def _add_to_back(self, node: CacheNode) -> None:
        prev = self.tail.prev
        assert prev is not None
        node.prev, node.next = prev, self.tail
        prev.next = self.tail.prev = node

    def get(self, key: int) -> int:
        if key not in self.nodes:
            return -1
        node = self.nodes[key]
        self._remove(node)
        self._add_to_back(node)
        return node.val

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


class LRUCacheOrderedDict:
    def __init__(self, capacity: int) -> None:
        self.capacity = capacity
        self.data: OrderedDict[int, int] = OrderedDict()

    def get(self, key: int) -> int:
        if key not in self.data:
            return -1
        self.data.move_to_end(key)
        return self.data[key]

    def put(self, key: int, value: int) -> None:
        if key in self.data:
            self.data.move_to_end(key)
        self.data[key] = value
        if len(self.data) > self.capacity:
            self.data.popitem(last=False)


def level_order(root: TreeNode | None) -> list[list[int]]:
    if root is None:
        return []
    ans: list[list[int]] = []
    q = deque([root])
    while q:
        level: list[int] = []
        for _ in range(len(q)):
            node = q.popleft()
            level.append(node.val)
            if node.left:
                q.append(node.left)
            if node.right:
                q.append(node.right)
        ans.append(level)
    return ans


def grid_bfs_distance(grid: list[list[int]], start: tuple[int, int]) -> list[list[int]]:
    rows, cols = len(grid), len(grid[0])
    dist = [[-1] * cols for _ in range(rows)]
    q = deque([start])
    dist[start[0]][start[1]] = 0  # Mark on enqueue, not on pop.
    while q:
        r, c = q.popleft()
        for nr, nc in grid_neighbors(r, c, rows, cols):
            if grid[nr][nc] == 0 and dist[nr][nc] == -1:
                dist[nr][nc] = dist[r][c] + 1
                q.append((nr, nc))
    return dist


def topological_sort_kahn(n: int, edges: list[tuple[int, int]]) -> list[int]:
    graph: defaultdict[int, list[int]] = defaultdict(list)
    indeg = [0] * n
    for pre, course in edges:
        graph[pre].append(course)
        indeg[course] += 1
    q = deque(i for i, deg in enumerate(indeg) if deg == 0)
    order: list[int] = []
    while q:
        node = q.popleft()
        order.append(node)
        for nei in graph[node]:
            indeg[nei] -= 1
            if indeg[nei] == 0:
                q.append(nei)
    return order if len(order) == n else []


def dijkstra(graph: dict[int, list[tuple[int, int]]], start: int) -> dict[int, int]:
    dist: dict[int, int] = {start: 0}
    heap = [(0, start)]
    while heap:
        d, node = heappop(heap)
        if d != dist[node]:
            continue
        for nei, weight in graph.get(node, []):
            nd = d + weight
            if nd < dist.get(nei, math.inf):
                dist[nei] = nd
                heappush(heap, (nd, nei))
    return dist


def valid_parentheses(s: str) -> bool:
    close_to_open = {")": "(", "]": "[", "}": "{"}
    stack: list[str] = []
    for ch in s:
        if ch in close_to_open.values():
            stack.append(ch)
        elif ch in close_to_open:
            if not stack or stack.pop() != close_to_open[ch]:
                return False
    return not stack


class MinStack:
    def __init__(self) -> None:
        self.stack: list[tuple[int, int]] = []

    def push(self, val: int) -> None:
        cur_min = val if not self.stack else min(val, self.stack[-1][1])
        self.stack.append((val, cur_min))

    def pop(self) -> int:
        return self.stack.pop()[0]

    def top(self) -> int:
        return self.stack[-1][0]

    def get_min(self) -> int:
        return self.stack[-1][1]


def merge_k_sorted_arrays(arrays: list[list[int]]) -> list[int]:
    heap: list[tuple[int, int, int]] = []
    for arr_i, arr in enumerate(arrays):
        if arr:
            heappush(heap, (arr[0], arr_i, 0))
    ans: list[int] = []
    while heap:
        val, arr_i, elem_i = heappop(heap)
        ans.append(val)
        nxt = elem_i + 1
        if nxt < len(arrays[arr_i]):
            heappush(heap, (arrays[arr_i][nxt], arr_i, nxt))
    return ans


def kadane_max_subarray(nums: list[int]) -> int:
    best = cur = nums[0]
    for x in nums[1:]:
        cur = max(x, cur + x)
        best = max(best, cur)
    return best


def climb_stairs_rolling(n: int) -> int:
    if n <= 2:
        return n
    prev2, prev1 = 1, 2
    for _ in range(3, n + 1):
        prev2, prev1 = prev1, prev1 + prev2
    return prev1


def lcs_2d(a: str, b: str) -> int:
    dp = [[0] * (len(b) + 1) for _ in range(len(a) + 1)]
    for i in range(len(a) - 1, -1, -1):
        for j in range(len(b) - 1, -1, -1):
            dp[i][j] = 1 + dp[i + 1][j + 1] if a[i] == b[j] else max(dp[i + 1][j], dp[i][j + 1])
    return dp[0][0]


def can_partition_01_knapsack(nums: list[int]) -> bool:
    total = sum(nums)
    if total % 2:
        return False
    target = total // 2
    possible = [False] * (target + 1)
    possible[0] = True
    for x in nums:
        for cur in range(target, x - 1, -1):
            possible[cur] = possible[cur] or possible[cur - x]
    return possible[target]


def permutations_used(nums: list[int]) -> list[list[int]]:
    ans: list[list[int]] = []
    path: list[int] = []
    used = [False] * len(nums)

    def backtrack() -> None:
        if len(path) == len(nums):
            ans.append(path.copy())
            return
        for i, x in enumerate(nums):
            if used[i]:
                continue
            used[i] = True
            path.append(x)
            backtrack()
            path.pop()
            used[i] = False

    backtrack()
    return ans


def combination_sum(candidates: list[int], target: int) -> list[list[int]]:
    candidates.sort()
    ans: list[list[int]] = []
    path: list[int] = []

    def backtrack(start: int, rem: int) -> None:
        if rem == 0:
            ans.append(path.copy())
            return
        for i in range(start, len(candidates)):
            x = candidates[i]
            if x > rem:
                break
            path.append(x)
            backtrack(i, rem - x)
            path.pop()

    backtrack(0, target)
    return ans


def single_number_xor(nums: list[int]) -> int:
    ans = 0
    for x in nums:
        ans ^= x
    return ans


def reverse_integer_32(x: int) -> int:
    sign = -1 if x < 0 else 1
    x = abs(x)
    limit = 2**31 - 1 if sign > 0 else 2**31
    ans = 0
    while x:
        digit = x % 10
        x //= 10
        if ans > limit // 10 or (ans == limit // 10 and digit > limit % 10):
            return 0
        ans = ans * 10 + digit
    return sign * ans


def atoi_clamped(s: str) -> int:
    i = 0
    while i < len(s) and s[i] == " ":
        i += 1
    sign = 1
    if i < len(s) and s[i] in "+-":
        sign = -1 if s[i] == "-" else 1
        i += 1
    ans = 0
    while i < len(s) and s[i].isdigit():
        ans = ans * 10 + int(s[i])
        i += 1
    return max(-(2**31), min(2**31 - 1, sign * ans))


def valid_word_abbreviation(word: str, abbr: str) -> bool:
    i = j = 0
    while i < len(abbr):
        if j >= len(word):
            return False
        if abbr[i] == word[j]:
            i += 1
            j += 1
        elif not abbr[i].isdigit() or abbr[i] == "0":
            return False
        else:
            skip = 0
            while i < len(abbr) and abbr[i].isdigit():
                skip = skip * 10 + int(abbr[i])
                i += 1
            j += skip
    return j == len(word)


def parse_transaction_record(raw: str) -> tuple[str, int, int, str]:
    name, time, amount, city = raw.split(",")
    return name, int(time), int(amount), city


def apply_gravity(board: list[list[int]]) -> None:
    rows, cols = len(board), len(board[0])
    for c in range(cols):
        write = rows - 1
        for r in range(rows - 1, -1, -1):
            if board[r][c] != 0:
                board[write][c] = board[r][c]
                write -= 1
        for r in range(write, -1, -1):
            board[r][c] = 0


def find_three_in_a_row(board: list[list[int]]) -> set[tuple[int, int]]:
    rows, cols = len(board), len(board[0])
    crush: set[tuple[int, int]] = set()
    for r in range(1, rows - 1):
        for c in range(cols):
            if board[r][c] and board[r - 1][c] == board[r][c] == board[r + 1][c]:
                crush.update({(r - 1, c), (r, c), (r + 1, c)})
    for r in range(rows):
        for c in range(1, cols - 1):
            if board[r][c] and board[r][c - 1] == board[r][c] == board[r][c + 1]:
                crush.update({(r, c - 1), (r, c), (r, c + 1)})
    return crush


def recursion_limit_for_deep_dfs() -> None:
    sys.setrecursionlimit(10**6)


# =============================================================================
# 11. Main Guard / Local Script Pattern
# =============================================================================


def solve(nums: list[int]) -> int:
    return sum(nums)


def main() -> None:
    sample = [1, 2, 3]
    print(solve(sample))


if __name__ == "__main__":
    # LeetCode normally calls `Solution` methods directly, so do not include this
    # in submissions unless you are building a local script or template.
    main()
