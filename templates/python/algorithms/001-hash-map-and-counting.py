"""
001 - Hash map and counting patterns

Use when the problem asks for pairs, frequencies, first/last seen positions,
membership checks, grouping, or "have we seen this before?" logic.
"""

from collections import Counter, defaultdict
from typing import Hashable, Iterable


# Variant 1: complement lookup, the most common pair pattern.
# Example problems: Two Sum, pair with target difference, pair existence.
# Time: O(n)
# Space: O(n)
def find_pair_indices(nums: list[int], target: int) -> tuple[int, int] | None:
    seen: dict[int, int] = {}

    for i, x in enumerate(nums):
        need = target - x
        if need in seen:
            return seen[need], i
        seen[x] = i

    return None


# Variant 2: frequency counting.
# Example problems: anagrams, majority checks, least/most frequent values.
# Time: O(n)
# Space: O(k), where k is the number of distinct values.
def frequency_table(items: Iterable[Hashable]) -> Counter:
    return Counter(items)


# Variant 3: group by derived key.
# Example problems: group anagrams, bucket by normalized form, classify strings.
# Time: O(n * m log m) for sorted-string keys, where m is average word length.
# Space: O(n * m)
def group_anagrams(words: list[str]) -> list[list[str]]:
    groups: defaultdict[tuple[str, ...], list[str]] = defaultdict(list)

    for word in words:
        key = tuple(sorted(word))
        groups[key].append(word)

    return list(groups.values())


# Variant 4: first seen index for longest distance / subarray transforms.
# Example problems: contiguous array, equal prefix states, first repeated state.
# Time: O(n)
# Space: O(k), where k is the number of distinct states.
def longest_span_with_same_state(states: list[int]) -> int:
    first_seen: dict[int, int] = {}
    best = 0

    for i, state in enumerate(states):
        if state in first_seen:
            best = max(best, i - first_seen[state])
        else:
            first_seen[state] = i

    return best
