"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

import bisect
import functools
import heapq
import operator

# ====================================================================
# 5. Sorting / Comparison / Heap
# ====================================================================

# Compares in-place sort, sorted copies, key functions, reverse order, and comparator adapters.
# Requires: import functools
# Requires: import operator
def sorting_patterns(
    items: list[tuple[str, int]], intervals: list[list[int]]
) -> None:
    nums = [3, 1, 2]

    new_sorted = sorted(nums)
    nums.sort()  # Mutates; returns None.

    desc_all = sorted(nums, reverse=True)

    # sorted(<list>, key=lambda <argument>: <representation
    # sorting key>)
    # <argument>: each element extracted from the list/iterable.
    # <representation sorting key>: value Python compares for
    # that element.
    by_start_end = sorted(intervals, key=lambda x: (x[0], x[1]))
    # x is one interval like [start, end]; (x[0], x[1]) sorts by
    # start, then end.

    score_desc_name_asc = sorted(
        items, key=lambda item: (-item[1], item[0])
    )
    # item is one tuple like (name, score); (-score, name) means
    # score desc, name asc.

    # Stable two-pass sort: sort secondary first, primary second.
    records = [("a", 90, 20), ("b", 90, 18)]
    records = sorted(
        records, key=lambda r: r[2]
    )  # r is each record; r[2] is age.
    records = sorted(
        records, key=lambda r: r[1], reverse=True
    )  # r[1] is score.
    records = sorted(
        records, key=operator.itemgetter(2)
    )  # same as lambda r: r[2]
    records = sorted(
        records, key=operator.itemgetter(1), reverse=True
    )  # same as lambda r: r[1]

    best = max(items, key=lambda item: (item[1], item[0]))
    worst = min(items, key=lambda item: item[1])

    # Comparator fallback: this is the -1/0/+1 style from Java/C/C++.
    # Python 3 sorting does not accept compare= directly; wrap
    # it with cmp_to_key.
    # compare returns negative, zero, or positive for ascending pairwise ordering.
    def compare(a: int, b: int) -> int:
        return (a % 2) - (b % 2) or a - b

    sorted_with_cmp = sorted(nums, key=functools.cmp_to_key(compare))




# Shows named key extractors, tuple keys, attribute keys, and legacy pairwise comparators.
# Requires: import functools
# Requires: import operator
def custom_sorting_patterns(
    words: list[str], transactions: list["TransactionRecord"]
) -> None:
    # 1) Lambda key: most common.
    # w is each string in words; (len(w), w) is the value used
    # for ordering.
    by_len_then_word = sorted(words, key=lambda w: (len(w), w))

    # 2) Named key function: easier to debug/explain when the
    # key is not tiny.
    # word_key sorts by increasing length and then alphabetical spelling.
    def word_key(word: str) -> tuple[int, str]:
        return len(word), word

    by_named_key = sorted(words, key=word_key)

    # 3) Comparator function: use only when "a before b?" logic
    # is naturally pairwise.
    # This function receives TWO items and returns negative /
    # zero / positive.
    # compare_words expresses the same ordering as word_key through pairwise comparison.
    def compare_words(a: str, b: str) -> int:
        if len(a) != len(b):
            return len(a) - len(b)
        return -1 if a < b else (1 if a > b else 0)

    by_comparator = sorted(
        words, key=functools.cmp_to_key(compare_words)
    )

    # 4) Custom object ordering: sorted(objs) calls obj.__lt__(other).
    by_time = sorted(transactions)
    by_time_key = sorted(
        transactions, key=operator.attrgetter("time")
    )  # same as lambda t: t.time

    # 5) Common accessors: itemgetter(0) is like lambda item: item[0].
    pairs = [("b", 2), ("a", 3)]
    by_first = sorted(pairs, key=operator.itemgetter(0))
    by_first_then_second = sorted(
        pairs, key=operator.itemgetter(0, 1)
    )


# TransactionRecord supplies __lt__ so ordinary sorting compares amount and then name.
class TransactionRecord:
    # Initializes a new instance and establishes its invariants.
    def __init__(self, raw: str) -> None:
        name, time, amount, city = raw.split(",")
        self.name = name
        self.time = int(time)
        self.amount = int(amount)
        self.city = city
        self.raw = raw

    # Internal helper for lt.
    def __lt__(self, other: "TransactionRecord") -> bool:
        return (self.time, self.amount, self.city) < (
            other.time,
            other.amount,
            other.city,
        )

    # Internal helper for repr.
    def __repr__(self) -> str:
        return f"TransactionRecord({self.raw!r})"




# Shows binary-search insertion boundaries and maintaining an ascending list.
# Requires: import bisect
def bisect_patterns(sorted_nums: list[int], x: int) -> None:
    left = bisect.bisect_left(sorted_nums, x)  # First index >= x.
    right = bisect.bisect_right(
        sorted_nums, x
    )  # First index with value > x.
    count_x = right - left
    exists = left < len(sorted_nums) and sorted_nums[left] == x

    insert_at_left = bisect.bisect_left(sorted_nums, x)
    insert_at_right = bisect.bisect_right(sorted_nums, x)
    bisect.insort(
        sorted_nums, x
    )  # Finds position O(log n), inserts into list O(n).




# Shows minimum heaps, negated maximum heaps, tuple priorities, replacement, and merging.
# Requires: import heapq
def heap_patterns(
    nums: list[int], k: int, tasks: list[tuple[int, str]]
) -> None:
    # Basic heap CRUD lives in `heap_create_read_update_delete_examples`.
    # This section keeps the interview variants.

    # Top-k one-off helpers.
    k_largest = heapq.nlargest(k, nums)
    k_smallest = heapq.nsmallest(k, nums)

    # Bounded min-heap for streaming kth largest.
    top: list[int] = []
    for x in nums:
        heapq.heappush(top, x)
        if len(top) > k:
            heapq.heappop(top)

    # Tuple heap: priority, tie breaker, payload.
    priority_queue: list[tuple[int, int, str]] = []
    for order, (priority, name) in enumerate(tasks):
        heapq.heappush(priority_queue, (priority, order, name))
