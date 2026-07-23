// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"fmt"
)

// ===================================================================
// 17. Methods, Interfaces, and Generics
// ===================================================================

type Person struct {
	Name string
	Age  int
}

// label formats a Person value without mutation, demonstrating a value receiver.
// Requires: import "fmt"
func (p Person) label() string {
	return fmt.Sprintf("%s:%d", p.Name, p.Age)
}

// birthday increments a Person age through a pointer receiver so the caller observes
// the change.
func (p *Person) birthday() {
	p.Age++
}

type Stringer interface {
	String() string
}

type Stack[T any] struct {
	data []T
}

// push appends one value to the receiver's stack storage.
func (s *Stack[T]) push(value T) {
	s.data = append(s.data, value)
}

// pop removes and returns the top stack value; false means the stack was empty.
func (s *Stack[T]) pop() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	index := len(s.data) - 1
	value := s.data[index]
	s.data[index] = zero
	s.data = s.data[:index]
	return value, true
}

type SupportedNumber interface {
	~int | ~int64 | ~float64
}

// sumValues adds values whose type satisfies the SupportedNumber generic constraint.
func sumValues[T SupportedNumber](values []T) T {
	var total T
	for _, value := range values {
		total += value
	}
	return total
}

// cloneMap returns a new map with copied keys and values; referenced values are still
// shallow copies.
func cloneMap[K comparable, V any](source map[K]V) map[K]V {
	result := make(map[K]V, len(source))
	for key, value := range source {
		result[key] = value
	}
	return result
}

// typedNilInterfacePattern demonstrates that an interface containing a typed nil
// pointer is not itself nil.
func typedNilInterfacePattern() bool {
	var node *ListNode
	var value any = node
	return value == nil // false: the interface has dynamic type *ListNode.
}

/*
Struct and interface notes:

- Struct assignment copies fields. Pointer fields still point to shared data.
- A value of type T has methods with receiver T; *T has methods for T and *T.
- Use pointer receivers to mutate or avoid copying a large struct.
- Exported names begin with an uppercase letter.
- A map key must be comparable; slices, maps, and functions are not.
- An interface is nil only when both dynamic type and dynamic value are nil.
  Storing a typed nil pointer in an interface produces a non-nil interface.
- Go has no generic methods with their own type parameters; use a generic
  function or a generic receiver type.
*/
