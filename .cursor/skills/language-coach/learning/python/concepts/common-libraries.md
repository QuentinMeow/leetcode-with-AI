---
concept_id: py-common-libraries
title: Python Common Libraries
lang: python
depth: standard
last_session: 2026-06-09
related_slugs:
  - python-specific-interview-patterns
  - interview-algorithm-templates
---

# Python Common Libraries

## At a glance

- Interview-heavy modules: `collections`, `heapq`, `bisect`, `itertools`, `functools`, and `math`.
- Daily-task modules worth recognizing: `re`, `json`, `csv`, `pathlib`, `datetime`, and `dataclasses`.
- Library choice should follow the invariant: queue behavior, sorted search, frequency map, memoized state, numeric helper, or data parsing.
- Standard-library functions save time, but each still has complexity and input-shape assumptions.

## What you already know

- Python interview prep benefits from reusable templates for common containers, iteration, sorting, strings, version syntax, and algorithm patterns.
- Current template files use small runnable functions plus a final interview-notes block.

## Gaps / not yet covered

- Practice choosing between a hand-written algorithm and a standard-library helper under interview constraints.
- Add more examples later for `operator`, `string`, `statistics`, `random`, `copy`, and `contextlib` if daily-task needs grow.
- Compare `heapq.nlargest` / `nsmallest` with full sorting and fixed-size heap patterns in real top-k problems.

## Detailed notes

### 2026-06-09

Added `templates/python/language-specific-pattern/common-libraries.py` as a reference for common Python libraries used in interviews and daily programming.

Core interview coverage:

- `collections.Counter` for frequencies, anagrams, and multiset comparison.
- `collections.defaultdict` for grouping and missing-key initialization.
- `collections.deque` for BFS queues and O(1) operations at both ends.
- `heapq` for min-heaps, priority queues, and top-k extraction.
- `bisect` for lower/upper-bound searches over sorted lists.
- `itertools.accumulate`, `pairwise`, `product`, `combinations`, `permutations`, and `groupby` for compact iteration patterns.
- `functools.cache`, `lru_cache`, and `cmp_to_key` for memoization and comparator adaptation.
- `math.gcd`, `lcm`, `isqrt`, `comb`, `inf`, and `ceil` for numeric helpers.

Daily-task coverage:

- `re` for parsing and normalizing text.
- `json`, `csv`, and `pathlib` for data and file handling.
- `datetime` for ISO date/timestamp parsing and date arithmetic.
- `dataclasses` for lightweight records with safe mutable defaults via `field(default_factory=...)`.

Important pitfalls to remember:

- `heapq` operates on a list and is a min-heap by default.
- `bisect` assumes sorted input.
- `groupby` groups only adjacent equal keys.
- `@cache` requires hashable arguments.
- Mutable dataclass defaults should use `default_factory`.

## Syntax and equivalents

- `@cache` is the short Python 3.9+ spelling for `@lru_cache(maxsize=None)`.
- `heappush(heap, -x)` is the common max-heap simulation for numeric priorities.
- `bisect_left(nums, x)` finds the first insertion point before existing `x` values; `bisect_right(nums, x)` finds the insertion point after them.
- `Path` objects compose paths with `/`, avoiding manual string concatenation.

## Interview angles

- Explain why `deque.popleft()` is O(1) but `list.pop(0)` is O(n).
- Know when top-k should use `heapq` versus sorting.
- Be ready to explain why memoized recursive DP state must be hashable.
- For sorted arrays, map "first index where condition holds" to `bisect_left` style lower-bound thinking.
- Mention standard-library helpers only when you can still describe the underlying data structure or complexity.

## Related concepts

- [python-specific-interview-patterns](python-specific-interview-patterns.md) — broader Python interview syntax and idioms.
- [interview-algorithm-templates](interview-algorithm-templates.md) — algorithm-pattern templates that often use these libraries.

## Addenda

### 2026-06-09

Expanded `common-libraries.py` with inline explanations and caveats for each imported module and example function. Emphasis added for sorted-input assumptions, lazy iterator materialization costs, heap tie-breaking, hashable memoization state, regex limitations, JSON/CSV type conversion, pathlib glob scope, timezone-naive datetimes, and dataclass mutable-default behavior.
