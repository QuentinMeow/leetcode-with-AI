// Go 1.21 version-specific syntax and standard-library additions.
//
// Prefer the simple core-language spelling when an interview environment has
// not stated its Go version. This repository targets Go 1.21.
package languagepatterns

import (
	"cmp"
	"errors"
	"maps"
	"slices"
)

// BuiltinsGo121 demonstrates min, max, and clear, all added in Go 1.21.
func BuiltinsGo121(nums []int, counts map[string]int) (int, int, []int) {
	low, high := 0, 0
	if len(nums) > 0 {
		low, high = nums[0], nums[0]
		for _, value := range nums[1:] {
			low = min(low, value)
			high = max(high, value)
		}
	}

	cleared := slices.Clone(nums)
	clear(cleared) // Keeps len/cap but sets every element to its zero value.
	clear(counts)  // Deletes every map entry.
	return low, high, cleared
}

// CloneAndSortGo121 uses the generic slices and maps packages added in Go 1.21.
func CloneAndSortGo121(nums []int, counts map[string]int) ([]int, map[string]int) {
	numsCopy := slices.Clone(nums)
	slices.Sort(numsCopy)
	return numsCopy, maps.Clone(counts)
}

// CompactGo121 removes adjacent duplicates after sorting. slices.Compact
// modifies the slice's backing array and returns the new logical length.
func CompactGo121(nums []int) []int {
	result := slices.Clone(nums)
	slices.Sort(result)
	return slices.Compact(result)
}

// -----------------------------------------------------------------------------
// Advanced: explicit generic functions (Go 1.18+)
// -----------------------------------------------------------------------------

// ClampVersioned works for the ordered predeclared types represented by
// cmp.Ordered. Generics are useful for genuinely reusable helpers, but concrete
// int/string helpers are usually easier to explain in a short interview.
func ClampVersioned[T cmp.Ordered](value, low, high T) T {
	if value < low {
		return low
	}
	if value > high {
		return high
	}
	return value
}

// UniqueVersioned demonstrates a comparable constraint. Map keys must be
// comparable, so slices, maps, and functions cannot be T here.
func UniqueVersioned[T comparable](values []T) []T {
	seen := make(map[T]struct{}, len(values))
	result := make([]T, 0, len(values))
	for _, value := range values {
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

// -----------------------------------------------------------------------------
// Advanced: combining independent errors (Go 1.20+)
// -----------------------------------------------------------------------------

var (
	ErrMissingVersioned = errors.New("missing value")
	ErrRangeVersioned   = errors.New("value outside range")
)

// ValidateVersioned joins independent validation failures. errors.Is can find
// any wrapped member. Most algorithm problems need one error at most, so
// errors.Join is more relevant to API/design interviews than LeetCode.
func ValidateVersioned(present bool, value int) error {
	var failures []error
	if !present {
		failures = append(failures, ErrMissingVersioned)
	}
	if value < 0 || value > 100 {
		failures = append(failures, ErrRangeVersioned)
	}
	return errors.Join(failures...)
}

/*
Quick version map:

- Go 1.18: type parameters, `any`, and the `comparable` constraint.
- Go 1.20: errors.Join and broader satisfaction of comparable constraints.
- Go 1.21: min/max/clear built-ins plus the slices, maps, and cmp packages.
- Go 1.22: each iteration of a freshly declared range variable gets distinct
  variables. Under Go 1.21 semantics, closures may observe the reused variable;
  explicitly copy it (`value := value`) before a closure when version portability
  matters.

Version discipline:

- Do not use a new helper merely because it exists. State the runtime version,
  prefer readable code, and know the short manual equivalent.
- Built-ins such as min and max require compatible ordered arguments. They do
  not accept arbitrary structs.
*/
