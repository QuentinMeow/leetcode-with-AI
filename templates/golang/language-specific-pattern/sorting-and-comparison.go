// Go sorting and comparison patterns for coding interviews.
//
// The sort package is widely available; Go 1.21's slices/cmp packages provide
// type-safe generic alternatives. Both families sort in place.
package languagepatterns

import (
	"cmp"
	"slices"
	"sort"
)

// SortCopyComparison clones before sorting so the caller's slice is unchanged.
func SortCopyComparison(nums []int) []int {
	result := slices.Clone(nums)
	slices.Sort(result)
	return result
}

// LegacySortComparison shows the pre-Go-1.21 helpers still common in interview
// templates and older environments.
func LegacySortComparison(nums []int, words []string) {
	sort.Ints(nums)
	sort.Strings(words)
}

// IntervalComparison is a small sortable record.
type IntervalComparison struct {
	Start int
	End   int
}

// SortIntervalsComparison sorts by start, then end. cmp.Compare avoids
// subtraction-based comparators, which can overflow.
func SortIntervalsComparison(intervals []IntervalComparison) {
	slices.SortFunc(intervals, func(left, right IntervalComparison) int {
		if byStart := cmp.Compare(left.Start, right.Start); byStart != 0 {
			return byStart
		}
		return cmp.Compare(left.End, right.End)
	})
}

// ScoreComparison demonstrates mixed direction: score descending, name
// ascending on ties.
type ScoreComparison struct {
	Name  string
	Score int
}

func SortScoresComparison(scores []ScoreComparison) {
	slices.SortFunc(scores, func(left, right ScoreComparison) int {
		if byScore := cmp.Compare(right.Score, left.Score); byScore != 0 {
			return byScore
		}
		return cmp.Compare(left.Name, right.Name)
	})
}

// StableRecordComparison carries an original ordering that should be retained
// when groups compare equal.
type StableRecordComparison struct {
	Group string
	Value int
}

func StableSortComparison(records []StableRecordComparison) {
	sort.SliceStable(records, func(left, right int) bool {
		return records[left].Group < records[right].Group
	})
}

// SortSliceComparison uses sort.Slice when a Go 1.18-1.20 environment lacks
// slices.SortFunc. The closure indexes the same slice being sorted.
func SortSliceComparison(intervals []IntervalComparison) {
	sort.Slice(intervals, func(left, right int) bool {
		if intervals[left].Start != intervals[right].Start {
			return intervals[left].Start < intervals[right].Start
		}
		return intervals[left].End < intervals[right].End
	})
}

// LowerBoundComparison returns the first index whose value is >= target.
// sort.Search assumes the predicate is false and then true (monotonic).
func LowerBoundComparison(sortedNums []int, target int) int {
	return sort.Search(len(sortedNums), func(index int) bool {
		return sortedNums[index] >= target
	})
}

// EqualRangeComparison returns the half-open range containing target.
func EqualRangeComparison(sortedNums []int, target int) (left, right int) {
	left = sort.Search(len(sortedNums), func(index int) bool {
		return sortedNums[index] >= target
	})
	right = sort.Search(len(sortedNums), func(index int) bool {
		return sortedNums[index] > target
	})
	return left, right
}

// BinarySearchComparison uses the Go 1.21 generic helper.
func BinarySearchComparison(sortedNums []int, target int) (index int, found bool) {
	return slices.BinarySearch(sortedNums, target)
}

// BestScoreComparison scans in O(n) rather than sorting in O(n log n) when
// only one best element is needed.
func BestScoreComparison(scores []ScoreComparison) (ScoreComparison, bool) {
	if len(scores) == 0 {
		return ScoreComparison{}, false
	}
	best := scores[0]
	for _, candidate := range scores[1:] {
		if candidate.Score > best.Score ||
			(candidate.Score == best.Score && candidate.Name < best.Name) {
			best = candidate
		}
	}
	return best, true
}

/*
Comparator contracts:

- sort.Slice asks for less(i, j) bool.
- slices.SortFunc asks for a comparator returning negative/zero/positive.
- A comparator must define a consistent strict ordering. Contradictory or
  state-changing comparison logic can produce incorrect results.
- Do not return left-right for arbitrary integers: subtraction can overflow.
  Use comparisons or cmp.Compare.

Practical choices:

- Integers/strings, Go 1.21: slices.Sort.
- Structs, Go 1.21: slices.SortFunc with cmp.Compare.
- Older environment: sort.Ints, sort.Strings, sort.Slice.
- Preserve equal-key order: sort.SliceStable or slices.SortStableFunc.
- Sorted search: sort.Search for custom monotonic predicates;
  slices.BinarySearch for an exact ordered value.
- Need one min/max only: scan instead of sorting.

All standard Go sorting helpers mutate the supplied slice. Clone first when the
input must remain unchanged.
*/
