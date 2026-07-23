// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

// ===================================================================
// 12. Sets via Maps
// ===================================================================

/*
Go has no built-in set. Choose one of these map representations:

	map[T]bool       // simpler membership: if set[value] { ... }
	map[T]struct{}   // stores keys with a zero-size value; use comma-ok

A missing bool map entry reads as false. That makes membership concise, but a
stored false is indistinguishable from absence unless comma-ok is also used.
The struct form communicates that only keys matter.
*/

// setCreateReadUpdateDeleteExamples demonstrates a key-only set implemented by
// map[int]struct{}.
func setCreateReadUpdateDeleteExamples(nums []int) {
	set := make(map[int]struct{}, len(nums))
	for _, value := range nums {
		set[value] = struct{}{}
	}
	set[42] = struct{}{}
	_, exists := set[42]
	delete(set, 42)
	size := len(set)
	clear(set)
	_ = []any{exists, size}
}

// Set algebra follows the operation names directly. map[T]struct{} avoids a
// redundant bool value; use map[T]bool when the shorter membership check wins.
func setUnion(left, right map[int]struct{}) map[int]struct{} {
	result := make(map[int]struct{}, len(left)+len(right))
	for value := range left {
		result[value] = struct{}{}
	}
	for value := range right {
		result[value] = struct{}{}
	}
	return result
}

// setIntersection returns values present in both sets. Iterating the smaller set
// reduces membership checks.
func setIntersection(left, right map[int]struct{}) map[int]struct{} {
	if len(left) > len(right) {
		left, right = right, left
	}
	result := make(map[int]struct{})
	for value := range left {
		if _, exists := right[value]; exists {
			result[value] = struct{}{}
		}
	}
	return result
}

// setDifference returns values present in left but absent from right.
func setDifference(left, right map[int]struct{}) map[int]struct{} {
	result := make(map[int]struct{}, len(left))
	for value := range left {
		if _, exists := right[value]; !exists {
			result[value] = struct{}{}
		}
	}
	return result
}

// boolSetExamples demonstrates the concise bool representation.
func boolSetExamples(values []int, target int) (map[int]bool, bool) {
	set := make(map[int]bool, len(values))
	for _, value := range values {
		set[value] = true
	}
	containsTarget := set[target] // Missing keys read as false.
	delete(set, target)
	return set, containsTarget
}

// sliceToSet builds the struct representation from a slice in O(n) time.
func sliceToSet(values []int) map[int]struct{} {
	set := make(map[int]struct{}, len(values))
	for _, value := range values {
		set[value] = struct{}{}
	}
	return set
}
