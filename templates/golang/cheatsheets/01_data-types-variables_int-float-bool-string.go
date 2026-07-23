// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

// ===================================================================
// 1. Data Types, Variables, and Go Version Map
// ===================================================================

/*
Start here: know the zero values, use := inside functions, and convert types
explicitly. In interviews, prefer int unless constraints require int64.

Go 1.18+: generics, any alias, comparable constraint
Go 1.20+: errors.Join, comparable types satisfy comparable constraints
Go 1.21+: built-in min/max/clear, slices/maps standard packages

Core declarations:

	var x int                 // zero value: 0
	x := 3                    // infer type inside a function
	const mod = 1_000_000_007 // Common prime modulus; use only when required.
	nums := []int{1, 2, 3}    // slice
	fixed := [3]int{1, 2, 3}  // array; length is part of the type
	count := map[string]int{} // map
	var wide int64 = int64(x) // conversions are explicit
	var small int32 = 32          // fixed-width signed integer
	var flags uint = 0b101        // unsigned integer, useful for bit masks
	maxInt := math.MaxInt         // finite sentinel for unreachable/best values
	letter := string(rune(65))    // "A": integer -> rune -> string
	_, _, _, _ = small, flags, maxInt, letter

Multiple assignment makes swaps easy: a, b = b, a.
Only bool values are conditions. Braces are required; semicolons are normally
inserted automatically.

The math.MaxInt line requires: import "math".
*/
