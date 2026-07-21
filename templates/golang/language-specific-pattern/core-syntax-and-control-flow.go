// Go core syntax and control-flow patterns for coding interviews.
//
// The examples favor explicit control flow. Go has no truthiness, ternary
// operator, or while keyword: conditions must be bool, and `for` handles every
// loop shape.
package languagepatterns

// CoreSyntaxDeclarations demonstrates zero values, short declarations, and
// multiple assignment. `:=` is available only inside functions and must
// introduce at least one new name in the current scope.
func CoreSyntaxDeclarations(a, b int) (int, int, bool, string) {
	var count int    // 0
	var ready bool   // false
	var label string // ""

	count = a + b
	a, b = b, a
	return count, a - b, ready, label
}

// SumPositiveCore shows the three common for-loop forms.
func SumPositiveCore(nums []int) int {
	total := 0

	// Range loop: use `_` when an index or value is intentionally ignored.
	for _, value := range nums {
		if value <= 0 {
			continue
		}
		total += value
	}

	// Classic index loop.
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			break
		}
	}

	// Condition-only loop: Go's equivalent of `while total < 0`.
	for total < 0 {
		total++
	}

	return total
}

// ClassifyCore demonstrates if initialization and an expression switch.
func ClassifyCore(value int) string {
	if doubled := value * 2; doubled > 100 {
		return "large"
	}

	switch {
	case value < 0:
		return "negative"
	case value == 0:
		return "zero"
	default:
		return "positive"
	}
}

// TokenKindCore demonstrates a value switch. Cases do not fall through by
// default, and several values can share one case.
func TokenKindCore(token byte) string {
	switch token {
	case '+', '-', '*', '/':
		return "operator"
	case '(', ')':
		return "parenthesis"
	default:
		return "other"
	}
}

// LookupCore uses the comma-ok form because a missing key and a stored zero
// otherwise produce the same map read.
func LookupCore(counts map[string]int, key string) (int, bool) {
	value, ok := counts[key]
	return value, ok
}

// FirstMatchCore uses a label to leave nested loops. Unlabeled break and
// continue affect only the innermost loop.
func FirstMatchCore(grid [][]int, target int) (int, int, bool) {
	row, col := -1, -1

search:
	for r := range grid {
		for c, value := range grid[r] {
			if value == target {
				row, col = r, c
				break search
			}
		}
	}

	return row, col, row != -1
}

// ShadowingCore makes block scope visible. The inner `value` is a different
// variable; use `=` when the intent is to update the outer name.
func ShadowingCore(value int) (outer, inner int) {
	outer = value
	if value := value + 1; value > 0 {
		inner = value
	}
	return outer, inner
}

/*
Core rules to remember:

- Braces are required. gofmt defines the conventional formatting.
- Only bool values can be conditions; 0, "", nil, and empty slices are not
  implicitly false.
- `if` and `switch` may begin with a short statement whose names are scoped to
  that construct.
- `switch` stops after the first matching case. `fallthrough` exists but is
  rarely appropriate in interview code.
- `for {}` is an infinite loop. `break`, `continue`, and `return` leave it.
- Exported package names start with an uppercase letter; lowercase names are
  package-private.
- `const` values are compile-time values. `var` values receive their type's
  zero value when no initializer is supplied.
*/
