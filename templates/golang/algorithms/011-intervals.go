package algorithms

import "sort"

/*
011 - Interval patterns

Use when inputs are ranges with start/end boundaries: meetings, calendars,
merge/insert intervals, overlaps, and sweep-line counting.
*/

// Variant 1: sort by start, then merge overlapping intervals.
// Example problems: merge intervals, summary ranges after sorting.
// Time: O(n log n)
// Space: O(n) for output.
func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		last := merged[len(merged)-1]
		if interval[0] <= last[1] {
			last[1] = max(last[1], interval[1])
		} else {
			merged = append(merged, interval)
		}
	}
	return merged
}

// Variant 2: insert interval into already sorted non-overlapping intervals.
// Example problems: insert interval, calendar insertion.
// Time: O(n)
// Space: O(n)
func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0, len(intervals)+1)
	index := 0
	for index < len(intervals) && intervals[index][1] < newInterval[0] {
		result = append(result, intervals[index])
		index++
	}
	for index < len(intervals) && intervals[index][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[index][0])
		newInterval[1] = max(newInterval[1], intervals[index][1])
		index++
	}
	result = append(result, newInterval)
	result = append(result, intervals[index:]...)
	return result
}

// Variant 3: overlap check after sorting.
// Example problems: meeting rooms, non-overlapping intervals.
// Time: O(n log n)
// Space: O(1) extra after sorting.
func CanAttendAllMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}

// Variant 4: sweep line for maximum simultaneous intervals.
// Example problems: meeting rooms II, car pooling, min platforms.
// Time: O(n log n)
// Space: O(n)
func MaxOverlappingIntervals(intervals [][]int) int {
	type event struct {
		position int
		delta    int
	}
	events := make([]event, 0, len(intervals)*2)
	for _, interval := range intervals {
		events = append(events, event{interval[0], 1}, event{interval[1], -1})
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i].position == events[j].position {
			return events[i].delta < events[j].delta
		}
		return events[i].position < events[j].position
	})

	current, best := 0, 0
	for _, event := range events {
		current += event.delta
		best = max(best, current)
	}
	return best
}
