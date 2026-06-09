"""
011 - Interval patterns

Use when inputs are ranges with start/end boundaries: meetings, calendars,
merge/insert intervals, overlaps, and sweep-line counting.
"""


# Variant 1: sort by start, then merge overlapping intervals.
# Example problems: merge intervals, summary ranges after sorting.
# Time: O(n log n)
# Space: O(n) for output.
def merge_intervals(intervals: list[list[int]]) -> list[list[int]]:
    if not intervals:
        return []

    intervals.sort(key=lambda x: x[0])
    merged = [intervals[0]]

    for start, end in intervals[1:]:
        last = merged[-1]
        if start <= last[1]:
            last[1] = max(last[1], end)
        else:
            merged.append([start, end])

    return merged


# Variant 2: insert interval into already sorted non-overlapping intervals.
# Example problems: insert interval, calendar insertion.
# Time: O(n)
# Space: O(n)
def insert_interval(intervals: list[list[int]], new_interval: list[int]) -> list[list[int]]:
    result: list[list[int]] = []
    i = 0
    n = len(intervals)

    while i < n and intervals[i][1] < new_interval[0]:
        result.append(intervals[i])
        i += 1

    while i < n and intervals[i][0] <= new_interval[1]:
        new_interval[0] = min(new_interval[0], intervals[i][0])
        new_interval[1] = max(new_interval[1], intervals[i][1])
        i += 1
    result.append(new_interval)

    result.extend(intervals[i:])
    return result


# Variant 3: overlap check after sorting.
# Example problems: meeting rooms, non-overlapping intervals.
# Time: O(n log n)
# Space: O(1) extra after sorting.
def can_attend_all_meetings(intervals: list[list[int]]) -> bool:
    intervals.sort(key=lambda x: x[0])

    for i in range(1, len(intervals)):
        if intervals[i][0] < intervals[i - 1][1]:
            return False

    return True


# Variant 4: sweep line for maximum simultaneous intervals.
# Example problems: meeting rooms II, car pooling, min platforms.
# Time: O(n log n)
# Space: O(n)
def max_overlapping_intervals(intervals: list[list[int]]) -> int:
    events: list[tuple[int, int]] = []
    for start, end in intervals:
        events.append((start, 1))
        events.append((end, -1))

    current = best = 0
    for _, delta in sorted(events):
        current += delta
        best = max(best, current)

    return best
