"""
Python sorting and comparison patterns for interviews.

Sorting is often the simplest way to expose order, intervals, greedy choices,
or duplicate structure. Python's `key=` function is the main idiom.
"""

from dataclasses import dataclass
from heapq import heappop, heappush


# -----------------------------------------------------------------------------
# sorted(...) vs list.sort(...)
# -----------------------------------------------------------------------------


def sort_copy(nums: list[int]) -> tuple[list[int], list[int]]:
    sorted_nums = sorted(nums)  # New list.
    nums.sort()  # In-place, returns None.
    return sorted_nums, nums


# -----------------------------------------------------------------------------
# key functions and tuple keys
# -----------------------------------------------------------------------------


def sort_intervals(intervals: list[list[int]]) -> list[list[int]]:
    # key= receives each item and returns the value Python should compare.
    return sorted(intervals, key=lambda interval: (interval[0], interval[1]))


def sort_words(words: list[str]) -> list[str]:
    # Primary key: length. Secondary key: lexicographic word.
    return sorted(words, key=lambda word: (len(word), word))


def sort_descending_scores(scores: list[tuple[str, int]]) -> list[tuple[str, int]]:
    # Score descending, name ascending.
    return sorted(scores, key=lambda item: (-item[1], item[0]))


# -----------------------------------------------------------------------------
# Stable sorting
# -----------------------------------------------------------------------------


def stable_sort_example(records: list[tuple[str, int, int]]) -> list[tuple[str, int, int]]:
    # records are (name, score, age). Python sort is stable.
    records = sorted(records, key=lambda record: record[2])  # age ascending
    records = sorted(records, key=lambda record: record[1], reverse=True)  # score descending
    return records


# The two-pass version works because equal-score records keep their prior
# age order. A tuple key is often simpler, but stability is useful to know.


# -----------------------------------------------------------------------------
# min/max with key
# -----------------------------------------------------------------------------


def best_word(words: list[str]) -> str:
    return max(words, key=lambda word: (len(word), word))


def smallest_interval(intervals: list[tuple[int, int]]) -> tuple[int, int]:
    return min(intervals, key=lambda interval: (interval[1] - interval[0], interval[0]))


# -----------------------------------------------------------------------------
# Heap tie breakers
# -----------------------------------------------------------------------------


def heap_with_tie_breaker(tasks: list[tuple[int, str]]) -> list[str]:
    heap: list[tuple[int, int, str]] = []

    for order, (priority, name) in enumerate(tasks):
        # Include order so ties do not require comparing task objects.
        heappush(heap, (priority, order, name))

    result: list[str] = []
    while heap:
        _, _, name = heappop(heap)
        result.append(name)

    return result


# -----------------------------------------------------------------------------
# Custom comparable classes: rare in LeetCode, but useful to recognize
# -----------------------------------------------------------------------------


@dataclass(order=True)
class RankedItem:
    # order=True generates comparison methods using fields in definition order.
    priority: int
    name: str


def sort_ranked_items(items: list[RankedItem]) -> list[RankedItem]:
    return sorted(items)


"""
Interview notes:

- `sorted(iterable)` returns a new list; `list.sort()` mutates and returns None.
- Prefer `key=` over custom comparators.
- Python sorts ascending by default; use negated numeric keys or `reverse=True`
  for descending order.
- Tuple keys compare lexicographically: first element, then second, and so on.
- Sorting is stable, which can simplify multi-key ordering.
- If heap entries tie on earlier tuple fields, Python compares later fields. Add
  an integer tie breaker when later fields may be non-comparable objects.

Concept explanations:

- `sorted(iterable)` works on any iterable and returns a new list. `list.sort()`
  only exists on lists, mutates the list in place, and returns `None` by design.
- Python 3 does not commonly use comparator functions for sorting. The idiom is
  `key=...`, where the key function is called once per element and the returned
  keys are compared.
- `lambda x: ...` creates a small anonymous function. In interview code, it is
  commonly used for `key=` because naming a one-line function would add noise.
- Tuple keys compare lexicographically. `(a, b, c)` sorts by `a`, breaks ties by
  `b`, then breaks remaining ties by `c`.
- Descending numeric fields are often represented by negating that field inside
  a tuple key, for example `(-score, name)`.
- `reverse=True` reverses the entire ordering. It is not the same as "only make
  the first tuple field descending" unless every field should reverse.
- Python sorting is stable: when two keys are equal, their original relative
  order is preserved. This supports multi-pass sorts and can simplify reasoning.
- `min(items, key=...)` and `max(items, key=...)` use the same key idea as
  sorting but avoid sorting the whole input when you only need one best item.
- Heap entries are compared as tuples. If priorities tie, Python compares the
  next tuple item. Add a monotonically increasing integer tie breaker before
  non-comparable payload objects.
- `@dataclass(order=True)` generates ordering methods, but only when the field
  order exactly matches the comparison you want. For one-off interview sorting,
  `key=` is usually more explicit.
"""
