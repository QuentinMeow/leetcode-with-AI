"""
010 - Heap / priority queue patterns

Use when repeatedly needing the smallest/largest item, top k items, or merging
sorted streams. Python's heapq is a min-heap; use negative values for max-heap.
"""

import heapq
from collections import Counter


# Variant 1: top k with a bounded min-heap.
# Example problems: kth largest element, top k scores, streaming kth largest.
# Time: O(n log k)
# Space: O(k)
def k_largest(nums: list[int], k: int) -> list[int]:
    heap: list[int] = []

    for x in nums:
        heapq.heappush(heap, x)
        if len(heap) > k:
            heapq.heappop(heap)

    return sorted(heap, reverse=True)


# Variant 2: max-heap by negating priority.
# Example problems: last stone weight, repeatedly take largest.
# Time: O(n log n)
# Space: O(n)
def repeatedly_take_largest(nums: list[int]) -> list[int]:
    heap = [-x for x in nums]
    heapq.heapify(heap)
    order: list[int] = []

    while heap:
        order.append(-heapq.heappop(heap))

    return order


# Variant 3: heap of tuples for custom priority and deterministic tie-breaking.
# Example problems: top k frequent elements, task scheduler variations.
# Time: O(n log k)
# Space: O(n)
def top_k_frequent(nums: list[int], k: int) -> list[int]:
    counts = Counter(nums)
    heap: list[tuple[int, int]] = []

    for num, freq in counts.items():
        heapq.heappush(heap, (freq, num))
        if len(heap) > k:
            heapq.heappop(heap)

    return [num for _, num in sorted(heap, reverse=True)]


# Variant 4: merge k sorted lists/arrays.
# Example problems: merge k sorted lists, kth smallest in sorted matrix.
# Time: O(total_items * log k)
# Space: O(k)
def merge_sorted_arrays(arrays: list[list[int]]) -> list[int]:
    heap: list[tuple[int, int, int]] = []
    result: list[int] = []

    for array_index, arr in enumerate(arrays):
        if arr:
            heapq.heappush(heap, (arr[0], array_index, 0))

    while heap:
        value, array_index, element_index = heapq.heappop(heap)
        result.append(value)
        next_index = element_index + 1
        if next_index < len(arrays[array_index]):
            heapq.heappush(
                heap,
                (arrays[array_index][next_index], array_index, next_index),
            )

    return result
