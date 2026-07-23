// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"cmp"
	"slices"
	"sort"
)

// ===================================================================
// 8. Sorting
// ===================================================================

// Sort a copy unless mutation is intended. Prefer cmp.Compare over subtraction
// in comparators, and state which bound your binary search returns.

type ScoreRecord struct {
	name  string
	score int
	age   int
}

// sortingPatterns demonstrates copied ascending/descending sorts, multi-key comparison,
// and stable sorting. The record slice is sorted in place.
// Requires: import "cmp"
// Requires: import "slices"
// Requires: import "sort"
func sortingPatterns(nums []int, words []string, records []ScoreRecord) {
	sortedNums := slices.Clone(nums)
	slices.Sort(sortedNums)
	slices.SortFunc(sortedNums, func(a, b int) int {
		return cmp.Compare(b, a) // Descending without subtraction overflow.
	})

	sortedWords := slices.Clone(words)
	slices.SortFunc(sortedWords, func(a, b string) int {
		if len(a) != len(b) {
			return cmp.Compare(len(a), len(b))
		}
		return cmp.Compare(a, b)
	})

	// score descending, then name ascending.
	sort.Slice(records, func(i, j int) bool {
		if records[i].score != records[j].score {
			return records[i].score > records[j].score
		}
		return records[i].name < records[j].name
	})

	// Stable sort preserves earlier order for equal keys.
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].age < records[j].age
	})

	_ = []any{sortedNums, sortedWords}
}

// legacySortPackageExamples shows APIs common in code written before Go 1.21.
// New Go code can usually use slices.Sort and slices.IsSorted instead.
// Requires: import "sort"
func legacySortPackageExamples(values []int, words []string) bool {
	sort.Ints(values)
	sort.Strings(words)
	intValuesAreSorted := sort.IntsAreSorted(values)
	wordsAreSorted := sort.StringsAreSorted(words)
	valuesAreSortedByComparator := sort.SliceIsSorted(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	return intValuesAreSorted && wordsAreSorted && valuesAreSortedByComparator
}
