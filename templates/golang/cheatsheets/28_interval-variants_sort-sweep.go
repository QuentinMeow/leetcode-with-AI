// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"cmp"
	"slices"
	"sort"
)

// ===================================================================
// 28. Interval Variants
// ===================================================================

// maximumConcurrentMeetingsUsingTwoPointers returns the most simultaneously active
// half-open intervals [start,end). Sorted starts add a room; an end at the same time
// frees a room before that start. Time O(n log n); space O(n).
// Requires: import "slices"
func maximumConcurrentMeetingsUsingTwoPointers(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	starts := make([]int, 0, len(intervals))
	ends := make([]int, 0, len(intervals))
	for _, interval := range intervals {
		if interval.start >= interval.end {
			continue
		}
		starts = append(starts, interval.start)
		ends = append(ends, interval.end)
	}
	if len(starts) == 0 {
		return 0
	}
	slices.Sort(starts)
	slices.Sort(ends)

	rooms, best := 0, 0
	startIndex, endIndex := 0, 0
	for startIndex < len(starts) {
		if starts[startIndex] < ends[endIndex] {
			rooms++
			best = max(best, rooms)
			startIndex++
		} else {
			rooms--
			endIndex++
		}
	}
	return best
}

// insertInterval assumes non-overlapping closed intervals [start,end] sorted by
// start. Touching at an endpoint counts as overlap. It appends intervals before
// the new one, merges every overlap, then appends the untouched suffix.
// Time O(n); output space O(n).
func insertInterval(
	intervals []Interval,
	newInterval Interval,
) []Interval {
	answer := make([]Interval, 0, len(intervals)+1)
	index := 0
	for index < len(intervals) &&
		intervals[index].end < newInterval.start {
		answer = append(answer, intervals[index])
		index++
	}
	for index < len(intervals) &&
		intervals[index].start <= newInterval.end {
		newInterval.start = min(newInterval.start, intervals[index].start)
		newInterval.end = max(newInterval.end, intervals[index].end)
		index++
	}
	answer = append(answer, newInterval)
	return append(answer, intervals[index:]...)
}

// eraseOverlapIntervals returns the minimum removals needed to make intervals
// non-overlapping. Keeping the interval with the earliest end leaves the most room for
// future intervals, which is the greedy invariant. Time O(n log n).
// Requires: import "sort"
func eraseOverlapIntervals(intervals []Interval) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].end < intervals[j].end
	})

	removed, previousEnd := 0, 0
	hasPrevious := false
	for _, interval := range intervals {
		if !hasPrevious || interval.start >= previousEnd {
			previousEnd = interval.end
			hasPrevious = true
		} else {
			removed++
		}
	}
	return removed
}

// maxOverlappingIntervals is a sweep-line algorithm. Start events add one active
// interval and end events remove one; sorting end before start on ties gives half-open
// [start,end) behavior. Time O(n log n).
// Requires: import "cmp"
// Requires: import "slices"
func maxOverlappingIntervals(intervals []Interval) int {
	type event struct {
		time  int
		delta int
	}
	events := make([]event, 0, 2*len(intervals))
	for _, interval := range intervals {
		if interval.start >= interval.end {
			continue
		}
		events = append(events,
			event{time: interval.start, delta: 1},
			event{time: interval.end, delta: -1},
		)
	}
	slices.SortFunc(events, func(a, b event) int {
		if a.time != b.time {
			return cmp.Compare(a.time, b.time)
		}
		return cmp.Compare(a.delta, b.delta) // End before start on ties.
	})

	active, best := 0, 0
	for _, current := range events {
		active += current.delta
		best = max(best, active)
	}
	return best
}
