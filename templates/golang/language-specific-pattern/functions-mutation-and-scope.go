// Go function, mutation, scope, error, and defer patterns.
//
// Every Go argument is passed by value. The copied value may itself contain a
// pointer or descriptor that reaches shared data, as slices and maps do.
package languagepatterns

import (
	"errors"
	"fmt"
)

// SwapFunction uses pointers when the caller's integer variables must change.
func SwapFunction(left, right *int) {
	*left, *right = *right, *left
}

// MutateSliceFunction can change existing elements because the copied slice
// header points at the same backing array.
func MutateSliceFunction(nums []int) {
	for index := range nums {
		nums[index]++
	}
}

// AppendSliceFunction returns the updated header because append may change the
// slice's length, capacity, and backing-array pointer.
func AppendSliceFunction(nums []int, values ...int) []int {
	return append(nums, values...)
}

// MutateMapFunction can update the caller-visible map through a copied map
// header. A non-nil map must be initialized before assignment.
func MutateMapFunction(counts map[string]int, key string) {
	counts[key]++
}

// ArrayCopyFunction receives a complete array copy.
func ArrayCopyFunction(values [3]int) [3]int {
	values[0] = -1
	return values
}

// MultipleResultsFunction demonstrates multiple return values and an explicit
// validity flag, a common alternative to sentinel values.
func MultipleResultsFunction(nums []int) (minimum int, ok bool) {
	if len(nums) == 0 {
		return 0, false
	}
	minimum = nums[0]
	for _, value := range nums[1:] {
		minimum = min(minimum, value)
	}
	return minimum, true
}

// RunningTotalFunction returns a closure. The captured total escapes its stack
// frame safely when needed; Go manages that allocation.
func RunningTotalFunction() func(int) int {
	total := 0
	return func(delta int) int {
		total += delta
		return total
	}
}

// DeferOrderFunction demonstrates last-in, first-out execution. Deferred calls
// run after a return statement sets named results but before the caller receives
// them.
func DeferOrderFunction() (order []int) {
	defer func() { order = append(order, 3) }()
	defer func() { order = append(order, 2) }()
	order = append(order, 1)
	return order
}

// DeferArgumentFunction shows that ordinary deferred-call arguments are
// evaluated when defer executes, not when the surrounding function returns.
func DeferArgumentFunction() (captured int) {
	value := 1
	defer func(snapshot int) { captured = snapshot }(value)
	value = 2
	_ = value
	return 0
}

// PatternCloserFunction is the minimal shape needed by the cleanup example.
type PatternCloserFunction interface {
	Close() error
}

// WorkAndCloseFunction preserves a work error; if work succeeds, a cleanup
// error becomes the result. Use defer immediately after acquiring a resource.
func WorkAndCloseFunction(closer PatternCloserFunction, work func() error) (err error) {
	defer func() {
		closeErr := closer.Close()
		if err == nil {
			err = closeErr
		}
	}()
	return work()
}

var ErrNegativeFunction = errors.New("negative input")

// ValidateFunction returns an error as a normal value.
func ValidateFunction(value int) error {
	if value < 0 {
		return fmt.Errorf("validate %d: %w", value, ErrNegativeFunction)
	}
	return nil
}

// IsNegativeFunction checks an error chain. Use errors.Is rather than direct
// equality when an error may have been wrapped with %w.
func IsNegativeFunction(err error) bool {
	return errors.Is(err, ErrNegativeFunction)
}

// BoundsErrorFunction carries structured error details.
type BoundsErrorFunction struct {
	Index int
	Size  int
}

func (err *BoundsErrorFunction) Error() string {
	return fmt.Sprintf("index %d outside length %d", err.Index, err.Size)
}

// AtFunction returns a structured error rather than panicking on caller input.
func AtFunction(nums []int, index int) (int, error) {
	if index < 0 || index >= len(nums) {
		return 0, &BoundsErrorFunction{Index: index, Size: len(nums)}
	}
	return nums[index], nil
}

// BoundsDetailsFunction uses errors.As so wrapped BoundsErrorFunction values are
// still recognized.
func BoundsDetailsFunction(err error) (index, size int, ok bool) {
	var boundsErr *BoundsErrorFunction
	if !errors.As(err, &boundsErr) {
		return 0, 0, false
	}
	return boundsErr.Index, boundsErr.Size, true
}

/*
Function and scope rules:

- Parameters, receivers, and return values are copied.
- A slice copy shares elements but not the slice header. A map copy refers to
  the same runtime map. A pointer copy still points to the same value.
- Variadic `values ...int` is a []int inside the function. Pass an existing
  slice with `f(nums...)`.
- Names declared in a smaller block shadow outer names. Accidental `:=` is a
  common reason an outer err/result is not updated.
- Closures capture variables, not frozen snapshots. Escape analysis decides
  whether captured state must move to the heap.

Error and defer rules:

- `error` is an interface. nil means success by convention.
- Add context with fmt.Errorf("operation: %w", err), then inspect chains with
  errors.Is/errors.As.
- defer is ideal for unlock/close/restore operations. Deferred calls execute
  LIFO even during panic unwinding.
- A deferred closure can modify named results. This is useful for cleanup errors
  but can make code hard to follow; keep it small and explicit.
- panic/recover is for unrecoverable invariants or framework boundaries, not
  ordinary invalid interview input. Prefer returned errors or ok flags.
*/
