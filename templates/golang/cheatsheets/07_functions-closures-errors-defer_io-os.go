// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

// ===================================================================
// 7. Functions, Closures, Errors, and Defer
// ===================================================================

// variadicSum adds zero or more integer arguments. A slice is expanded into variadic
// arguments with values....
func variadicSum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

// namedAndFirstClassFunctions demonstrates assigning an anonymous function to a
// variable, calling it through another variable, and expanding a slice into variadic
// arguments.
func namedAndFirstClassFunctions(nums []int) {
	add := func(a, b int) int { return a + b }
	operation := add
	result := operation(2, 3)
	total := variadicSum(nums...)
	_ = []any{result, total}
}

// makeCounter returns a closure that retains and increments its own count across calls.
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// divide performs integer division and returns an error instead of panicking when the
// divisor is zero.
// Requires: import "errors"
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// parsePositive converts decimal text to an integer, wraps conversion errors with
// context, and rejects zero or negative results.
// Requires: import "fmt"
// Requires: import "strconv"
func parsePositive(text string) (int, error) {
	value, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("parse %q: %w", text, err)
	}
	if value <= 0 {
		return 0, fmt.Errorf("value must be positive: %d", value)
	}
	return value, nil
}

// readAndClose reads a file and joins any close failure with an earlier read failure.
// The named error result lets the deferred cleanup update the returned error.
// Requires: import "errors"
// Requires: import "io"
// Requires: import "os"
func readAndClose(path string) (_ []byte, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = errors.Join(err, file.Close())
	}()
	return io.ReadAll(file)
}

/*
defer notes:

- Deferred calls run last-in-first-out (LIFO) when the function returns.
- Arguments are evaluated when defer is registered.
- Use defer for cleanup after successfully acquiring a resource.
- panic/recover is not ordinary error handling; return errors instead.
*/

// minAndMaxValues demonstrates multiple return values. The bool is false for
// empty input because no minimum or maximum exists.
func minAndMaxValues(values []int) (minimum, maximum int, ok bool) {
	if len(values) == 0 {
		return 0, 0, false
	}
	minimum, maximum = values[0], values[0]
	for _, value := range values[1:] {
		minimum = min(minimum, value)
		maximum = max(maximum, value)
	}
	return minimum, maximum, true
}

// fibonacciRecursiveTeaching is intentionally exponential: it demonstrates
// recursive syntax, not the implementation to choose for large n.
func fibonacciRecursiveTeaching(n int) int {
	if n <= 1 {
		return max(n, 0)
	}
	return fibonacciRecursiveTeaching(n-1) + fibonacciRecursiveTeaching(n-2)
}
