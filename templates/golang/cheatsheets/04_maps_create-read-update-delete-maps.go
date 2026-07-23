// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"maps"
)

// ===================================================================
// 4. Maps
// ===================================================================

// Map reads are safe on nil maps, but writes panic. The comma-ok form separates
// a missing key from a stored zero value. Map iteration order is unspecified.

// mapCreateReadUpdateDeleteExamples demonstrates initialization, comma-ok lookup,
// counting, grouping, cloning, deletion, clearing, and safe reads from a nil map.
// Requires via helper: import "slices"
// Requires: import "maps"
func mapCreateReadUpdateDeleteExamples(words []string, nums []int) {
	// CREATE
	var nilMap map[string]int
	counts := make(map[string]int)
	withData := map[string]int{"a": 1, "b": 2}

	// READ
	value := withData["a"]
	missingZero := withData["missing"]
	value, exists := withData["a"]
	size := len(withData)
	_, contains := withData["a"]

	// UPDATE / ADD
	counts["c"] = 3
	for _, word := range words {
		counts[word]++ // Missing key starts at int's zero value.
	}

	clone := maps.Clone(withData)
	for key, value := range counts {
		clone[key] = value
	}

	// Grouping.
	groups := make(map[string][]string)
	for _, word := range words {
		key := sortedRunesKey(word)
		groups[key] = append(groups[key], word)
	}

	// Frequency counting.
	frequency := make(map[int]int)
	for _, x := range nums {
		frequency[x]++
	}

	// DELETE
	delete(counts, "c") // No-op when absent.
	removed, ok := counts["x"]
	if ok {
		delete(counts, "x")
	}
	clear(counts)

	// Safe: reads from nil maps return zero. Unsafe: nilMap["x"] = 1.
	nilRead := nilMap["x"]

	_ = []any{
		value, missingZero, exists, size, contains, clone, groups,
		frequency, removed, nilRead,
	}
}

// mapIterationExamples shows all three useful range forms. Map order is not
// stable, so collect and sort keys when output order matters.
func mapIterationExamples(counts map[string]int) {
	for key, value := range counts {
		_, _ = key, value
	}
	for key := range counts {
		_ = key
	}
	for _, value := range counts {
		_ = value
	}
}
