"""
Python common-library patterns for interviews and daily tasks.

Use this file when the algorithm is clear but you need to remember which
standard-library tool gives the cleanest Python implementation.

Each section explains what the module is for and the caveats that matter in
interviews: input assumptions, runtime costs, mutability, and version details.
"""

# csv: read/write comma-separated data. Caveat: CSV values are strings unless
# you convert them, and real CSV needs the csv module rather than split(",").
import csv
# json: serialize Python data to text and parse it back. Caveat: JSON supports
# only JSON types; tuples become lists, and object keys are strings in JSON.
import json
# math: numeric helpers implemented carefully in C/Python. Caveat: prefer exact
# integer helpers such as isqrt over float sqrt when integer precision matters.
import math
# re: regular expressions for text search/extraction. Caveat: powerful but less
# readable than string methods for simple split/strip/replace tasks.
import re
# bisect: binary search insertion points in sorted lists. Caveat: it never
# checks that the list is sorted, and insertion into a list is still O(n).
from bisect import bisect_left, bisect_right, insort
# collections: specialized containers. Counter counts, defaultdict creates
# missing values, and deque is the queue type you want for O(1) popleft().
from collections import Counter, defaultdict, deque
# dataclasses: generate init/repr/comparison methods for records. Caveat: use
# default_factory for mutable defaults such as list/set/dict.
from dataclasses import dataclass, field
# datetime: dates, timestamps, and durations. Caveat: naive datetimes have no
# timezone; daily scripts should be explicit when timezone correctness matters.
from datetime import date, datetime, timedelta
# functools: decorators and helpers for function-level behavior. Caveat: cache
# keys must be hashable, and cmp_to_key is usually a fallback behind key=.
from functools import cache, cmp_to_key, lru_cache, reduce
# heapq: heap operations over a normal list. Caveat: Python gives a min-heap,
# not a heap class; use negated priorities or tuple keys for other orderings.
from heapq import heapify, heappop, heappush, nlargest, nsmallest
# itertools: lazy iterator building blocks. Caveat: converting product,
# permutations, or combinations to a list can explode in size.
from itertools import accumulate, combinations, groupby, pairwise, permutations, product
# operator: named function versions of common operators and accessors. Useful
# with sorted/groupby/reduce when a lambda would only index or add values.
from operator import add, itemgetter
# pathlib: object-oriented paths. Caveat: Path methods touch your local
# filesystem only when you call read/write/exists/glob/etc.
from pathlib import Path
# typing: type hints only; they guide readers and tooling, not runtime checks.
from typing import Iterable


# -----------------------------------------------------------------------------
# collections: Counter, defaultdict, deque
# -----------------------------------------------------------------------------


def counter_examples(words: list[str]) -> tuple[Counter[str], list[tuple[str, int]], bool]:
    # Counter is a dict subclass for frequencies. Missing keys read as 0, so it
    # works well for counting, anagrams, and multiset-style comparisons.
    counts = Counter(words)

    # most_common(k) returns (item, count) pairs ordered by count descending.
    # Caveat: ties keep first-seen order in modern Python, but interview logic
    # should define tie-breaking explicitly when output order matters.
    top_two = counts.most_common(2)

    # Counter equality ignores zero counts in modern Python, making this a
    # compact anagram check. It is still O(n) to build each Counter.
    same_multiset = Counter("listen") == Counter("silent")

    # Counter keeps zero/negative counts until normalized or deleted.
    # `subtract` can push counts below zero, unlike normal frequency maps.
    counts.subtract(["debug"])
    # Adding an empty Counter drops zero and negative counts.
    counts += Counter()

    return counts, top_two, same_multiset


def defaultdict_examples(pairs: list[tuple[str, int]]) -> tuple[dict[str, list[int]], dict[str, int]]:
    # defaultdict(factory) calls factory only when a missing key is accessed via
    # grouped[key]. It does not create entries for grouped.get(key).
    grouped: defaultdict[str, list[int]] = defaultdict(list)

    # int() returns 0, so defaultdict(int) is a compact frequency/sum map.
    totals: defaultdict[str, int] = defaultdict(int)

    for key, value in pairs:
        grouped[key].append(value)
        totals[key] += value

    # Convert to plain dict when callers should not inherit missing-key behavior.
    return dict(grouped), dict(totals)


def deque_examples(start: int, graph: dict[int, list[int]], window_size: int) -> list[int]:
    # deque is a double-ended queue. append/pop and appendleft/popleft are O(1),
    # unlike list.pop(0), which shifts all remaining elements and costs O(n).
    queue = deque([start])
    seen = {start}
    order: list[int] = []

    while queue:
        node = queue.popleft()
        order.append(node)

        for nei in graph.get(node, []):
            if nei not in seen:
                seen.add(nei)
                queue.append(nei)

    # maxlen makes a bounded deque: when it is full, appending on one side
    # automatically discards from the other side. Good for "last k" streams.
    recent = deque(order, maxlen=window_size)
    recent.append(start)
    return list(recent)


# -----------------------------------------------------------------------------
# heapq: priority queues and top-k
# -----------------------------------------------------------------------------


def heapq_examples(nums: list[int], k: int) -> tuple[list[int], list[int], list[int]]:
    # heapq mutates a normal list into heap order. heap[0] is the smallest item,
    # but the rest of the list is not globally sorted.
    heap = nums[:]
    # heapify is O(n), faster than pushing n items one by one at O(n log n).
    heapify(heap)

    # Each heappop is O(log n). Popping k times is useful when k is small.
    smallest = [heappop(heap) for _ in range(min(k, len(heap)))]

    # Python's heap is a min-heap. Negating numbers simulates a max-heap, but
    # remember to negate again when popping.
    max_heap: list[int] = []
    for x in nums:
        heappush(max_heap, -x)
    largest_by_negation = [-heappop(max_heap) for _ in range(min(k, len(max_heap)))]

    # nlargest/nsmallest are convenient for one-off top-k. For repeated updates,
    # maintain your own heap. For large k near n, sorting may be simpler.
    largest_direct = nlargest(k, nums)
    return smallest, largest_by_negation, largest_direct


def priority_queue_with_tie_breaker(tasks: list[tuple[int, str]]) -> list[str]:
    # Heap entries are compared lexicographically: first field, then second, etc.
    # Include a monotonically increasing order so equal priorities do not force
    # Python to compare non-comparable task objects.
    heap: list[tuple[int, int, str]] = []

    for order, (priority, name) in enumerate(tasks):
        heappush(heap, (priority, order, name))

    result: list[str] = []
    while heap:
        _, _, name = heappop(heap)
        result.append(name)

    return result


def top_k_frequent_words(words: list[str], k: int) -> list[str]:
    counts = Counter(words)
    # Highest frequency first, lexicographically smaller word first on ties.
    # nsmallest works because the key starts with negative frequency.
    return nsmallest(k, counts, key=lambda word: (-counts[word], word))


# -----------------------------------------------------------------------------
# bisect: binary search over sorted lists
# -----------------------------------------------------------------------------


def bisect_examples(sorted_nums: list[int], target: int) -> tuple[int, int, int]:
    # bisect_left: first index where target could be inserted before equals.
    left = bisect_left(sorted_nums, target)

    # bisect_right: first index after existing equals. The half-open slice
    # sorted_nums[left:right] contains exactly the target values.
    right = bisect_right(sorted_nums, target)
    count = right - left
    return left, right, count


def insert_sorted(sorted_nums: list[int], value: int) -> list[int]:
    result = sorted_nums[:]
    # insort finds the insertion point in O(log n), then inserts into the list in
    # O(n) because elements after the position must shift.
    insort(result, value)
    return result


def lower_bound_answer_space(nums: list[int], need: int) -> int:
    # accumulate lazily yields running totals. Converting to list gives a prefix
    # array. Caveat: bisect is valid here only if prefix is monotonic, which
    # requires non-negative nums for ordinary prefix sums.
    prefix = list(accumulate(nums))
    return bisect_left(prefix, need)


# -----------------------------------------------------------------------------
# itertools: compact iteration building blocks
# -----------------------------------------------------------------------------


def itertools_examples(nums: list[int]) -> tuple[list[int], list[int], list[tuple[int, int]]]:
    # accumulate(nums) yields prefix sums by default. It is lazy, so wrap with
    # list only when you actually need all values stored.
    prefix_sums = list(accumulate(nums))

    # pairwise is Python 3.10+. It yields adjacent pairs: (nums[0], nums[1]),
    # then (nums[1], nums[2]), and so on.
    adjacent_diffs = [b - a for a, b in pairwise(nums)]

    # combinations(nums, 2) yields unordered index-distinct pairs. The number of
    # pairs is O(n^2), so this is only for small search spaces.
    pairs = list(combinations(nums, 2))
    return prefix_sums, adjacent_diffs, pairs


def cartesian_grid(rows: int, cols: int) -> list[tuple[int, int]]:
    # product(A, B) yields every ordered pair from A x B. For grids this replaces
    # nested loops, but materializing the list costs rows * cols space.
    return list(product(range(rows), range(cols)))


def permutation_examples(chars: str) -> list[str]:
    # permutations(chars) yields ordered arrangements. Count is n!, so avoid
    # generating all permutations unless n is tiny or the problem requires it.
    return ["".join(p) for p in permutations(chars)]


def group_sorted_items(items: list[tuple[str, int]]) -> dict[str, list[int]]:
    result: dict[str, list[int]] = {}

    # groupby only groups adjacent equal keys, so sort first unless the input is
    # already grouped by the same key. itemgetter(0) is like lambda item: item[0].
    for key, group in groupby(sorted(items, key=itemgetter(0)), key=itemgetter(0)):
        # Each group is an iterator that is consumed once. Store it immediately
        # if you need the values after the loop advances.
        result[key] = [value for _, value in group]

    return result


# -----------------------------------------------------------------------------
# functools and operator: memoization, comparators, reductions
# -----------------------------------------------------------------------------


@cache
def fibonacci(n: int) -> int:
    # @cache memoizes return values by argument tuple. It is ideal for top-down
    # DP with immutable state such as ints, strings, and tuples.
    # Caveat: cached recursive functions keep state across calls in the same
    # process; call fibonacci.cache_clear() if stale state matters in a script.
    if n <= 1:
        return n
    return fibonacci(n - 1) + fibonacci(n - 2)


@lru_cache(maxsize=None)
def count_paths(r: int, c: int) -> int:
    # @lru_cache(maxsize=None) is the older spelling of @cache. With a finite
    # maxsize it evicts least-recently-used entries, which is rarely needed for
    # interview DP but useful in long-running services.
    if r == 0 or c == 0:
        return 1
    return count_paths(r - 1, c) + count_paths(r, c - 1)


def sort_with_comparator(nums: list[int]) -> list[int]:
    def compare(a: int, b: int) -> int:
        # Example ordering: even numbers first, then normal ascending order.
        # Comparator contract: return negative if a < b, zero if equal, positive
        # if a > b. Prefer key= when possible because it is simpler and faster.
        if a % 2 != b % 2:
            return -1 if a % 2 == 0 else 1
        return a - b

    # cmp_to_key adapts an old-style comparator for sorted(..., key=...).
    return sorted(nums, key=cmp_to_key(compare))


def reduce_examples(nums: Iterable[int]) -> int:
    # reduce repeatedly combines values: (((0 + a) + b) + c) ...
    # In Python, sum(nums) is clearer for addition; reduce is more useful when
    # the operation is not already a built-in aggregate.
    return reduce(add, nums, 0)


# -----------------------------------------------------------------------------
# math: numeric helpers
# -----------------------------------------------------------------------------


def math_examples(a: int, b: int, n: int) -> tuple[int, int, int, int, float, int]:
    # gcd/lcm are exact integer helpers. lcm is Python 3.9+.
    gcd_value = math.gcd(a, b)
    lcm_value = math.lcm(a, b)

    # isqrt(n) returns floor(sqrt(n)) using integer arithmetic, avoiding float
    # precision bugs for large n.
    root_floor = math.isqrt(n)

    # comb(n, k) computes "n choose k" exactly. It raises ValueError for negative
    # n or k, so guard inputs when they may be invalid.
    combinations_count = math.comb(n, 2) if n >= 2 else 0

    # math.inf is a readable sentinel for "larger than any finite number."
    infinity = math.inf

    # This float-based ceiling is fine for ordinary-sized values. For huge
    # positive integers, prefer exact integer math: -(-a // b).
    rounded_up = math.ceil(a / b)
    return gcd_value, lcm_value, root_floor, combinations_count, infinity, rounded_up


def prime_check(n: int) -> bool:
    if n < 2:
        return False
    # If n has a factor larger than sqrt(n), it must also have one smaller than
    # sqrt(n). isqrt keeps the loop bound exact.
    for d in range(2, math.isqrt(n) + 1):
        if n % d == 0:
            return False
    return True


# -----------------------------------------------------------------------------
# string and re: text parsing
# -----------------------------------------------------------------------------


def regex_parse_ints(text: str) -> list[int]:
    # findall returns every non-overlapping match. This pattern captures optional
    # minus signs followed by digits, so it parses "-12" as one token.
    # Caveat: it does not parse decimals, plus signs, or scientific notation.
    return [int(match) for match in re.findall(r"-?\d+", text)]


def normalize_words(text: str) -> list[str]:
    # Lowercase first, then keep alphabetic runs. This is useful for quick word
    # counts, but the pattern is ASCII-only; use a Unicode-aware pattern or
    # str.isalpha logic if non-English text matters.
    return re.findall(r"[a-z]+", text.lower())


def replace_repeated_spaces(text: str) -> str:
    # \s matches spaces, tabs, and newlines. Use literal " +" if only normal
    # spaces should collapse.
    return re.sub(r"\s+", " ", text).strip()


# -----------------------------------------------------------------------------
# json, csv, pathlib: everyday file/data tasks
# -----------------------------------------------------------------------------


def json_round_trip(data: dict[str, object]) -> dict[str, object]:
    # dumps converts Python data to a JSON string; loads parses JSON text back.
    # sort_keys makes output deterministic for tests/diffs.
    text = json.dumps(data, sort_keys=True)
    # Caveat: JSON round trips can change types, e.g. tuple -> list.
    return json.loads(text)


def read_json_file(path: Path) -> dict[str, object]:
    # Path.read_text is concise for small files. For huge files, stream from an
    # open file object instead of reading everything into memory first.
    return json.loads(path.read_text(encoding="utf-8"))


def read_csv_rows(path: Path) -> list[dict[str, str]]:
    # newline="" lets csv handle newlines correctly across platforms.
    with path.open(newline="", encoding="utf-8") as f:
        # DictReader uses the header row as keys. Values are strings; convert
        # numbers/dates yourself.
        return list(csv.DictReader(f))


def pathlib_examples(root: Path) -> tuple[Path, str, list[str]]:
    # The / operator joins paths without manual slash handling.
    config_path = root / "config" / "settings.json"

    # suffix is the final extension, such as ".json". For "archive.tar.gz",
    # suffix is ".gz" and suffixes is [".tar", ".gz"].
    suffix = config_path.suffix

    # glob("*.py") searches one directory level. Use rglob("*.py") or
    # glob("**/*.py") for recursive search.
    python_files = [p.name for p in root.glob("*.py")]
    return config_path, suffix, python_files


# -----------------------------------------------------------------------------
# datetime: timestamps, durations, parsing
# -----------------------------------------------------------------------------


def datetime_examples(day: str) -> tuple[date, date, int]:
    # fromisoformat parses "YYYY-MM-DD" for dates. It is stricter than many
    # human date formats, which is good for machine-readable data.
    parsed = date.fromisoformat(day)

    # timedelta represents a duration. Date arithmetic returns a new date.
    tomorrow = parsed + timedelta(days=1)

    # Subtracting two dates gives a timedelta; .days is the whole-day difference.
    days_since_epoch = (parsed - date(1970, 1, 1)).days
    return parsed, tomorrow, days_since_epoch


def parse_timestamp(timestamp: str) -> datetime:
    # fromisoformat handles many ISO-8601 timestamp strings. Caveat: timezone
    # handling depends on the string; a result without tzinfo is "naive".
    return datetime.fromisoformat(timestamp)


# -----------------------------------------------------------------------------
# dataclasses: lightweight records for daily Python code
# -----------------------------------------------------------------------------


@dataclass(order=True)
class Task:
    # order=True generates comparison methods using fields in definition order:
    # priority first, then name, then tags unless a field disables comparison.
    priority: int
    name: str

    # default_factory creates a fresh list per instance. Using tags: list[str] =
    # [] would share the same list across instances and is rejected by dataclass.
    tags: list[str] = field(default_factory=list)


def dataclass_examples(tasks: list[Task]) -> list[Task]:
    # sorted(tasks) works because Task has order=True. Caveat: if later fields
    # are not comparable, mark them compare=False with field(..., compare=False).
    return sorted(tasks)


"""
Interview library shortlist:

- `collections.Counter`: frequencies, anagrams, multiset comparison.
- `collections.defaultdict`: grouping and missing-key initialization.
- `collections.deque`: BFS queues and O(1) pops from both ends.
- `heapq`: min-heaps, priority queues, top-k, scheduling.
- `bisect`: lower/upper bound in sorted arrays or monotonic prefix arrays.
- `itertools.accumulate`: prefix sums; `pairwise`: adjacent pairs; `product`:
  grid coordinates; `combinations` / `permutations`: generate small search sets.
- `functools.cache` / `lru_cache`: top-down DP memoization with hashable state.
- `math.gcd`, `math.lcm`, `math.isqrt`, `math.inf`: common numeric helpers.
- `re`: parsing tokens from text; useful daily, less common in pure LeetCode.
- `json`, `csv`, `pathlib`, `datetime`: daily scripting and data handling.

Pitfalls:

- Standard library tools are helpers, not magic. Know the complexity of the
  operation you are using.
- `heapq` is a module over a list, not a heap class; it is a min-heap by default.
- `groupby` only groups adjacent items; sort first when needed.
- `bisect` assumes the list is already sorted and does not verify it.
- `@cache` arguments must be hashable; convert lists to tuples for memo state.
- `field(default_factory=list)` avoids sharing one mutable default across every
  dataclass instance.
"""
