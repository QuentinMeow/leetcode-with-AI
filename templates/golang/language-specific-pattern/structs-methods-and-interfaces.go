// Go structs, methods, and interfaces for coding interviews.
//
// Go favors small concrete structs and implicit interfaces. There are no
// classes or constructors in the language; a constructor is an ordinary
// function returning a useful initialized value.
package languagepatterns

// ListNodePattern and TreeNodePattern mirror common platform-provided types.
type ListNodePattern struct {
	Value int
	Next  *ListNodePattern
}

type TreeNodePattern struct {
	Value       int
	Left, Right *TreeNodePattern
}

// BuildListPattern uses a dummy node to avoid a special first-insertion branch.
func BuildListPattern(values []int) *ListNodePattern {
	dummy := &ListNodePattern{}
	current := dummy
	for _, value := range values {
		current.Next = &ListNodePattern{Value: value}
		current = current.Next
	}
	return dummy.Next
}

// PointPattern is comparable because both fields are comparable. It can be a
// map key without custom equality or hashing code.
type PointPattern struct {
	Row int
	Col int
}

func VisitedPointsPattern(points []PointPattern) map[PointPattern]struct{} {
	visited := make(map[PointPattern]struct{}, len(points))
	for _, point := range points {
		visited[point] = struct{}{}
	}
	return visited
}

// CounterPattern demonstrates value and pointer receivers.
type CounterPattern struct {
	total int
}

// Total uses a value receiver because it only reads a small value.
func (counter CounterPattern) Total() int {
	return counter.total
}

// Add uses a pointer receiver because it mutates the original CounterPattern.
func (counter *CounterPattern) Add(delta int) {
	counter.total += delta
}

// NewCounterPattern is a conventional constructor, not special syntax.
func NewCounterPattern(initial int) *CounterPattern {
	return &CounterPattern{total: initial}
}

// NamedPattern is a small behavior-focused interface. A type satisfies it
// automatically by having the method; there is no `implements` declaration.
type NamedPattern interface {
	Name() string
}

type TaskPattern struct {
	Label    string
	Priority int
}

func (task TaskPattern) Name() string {
	return task.Label
}

func CollectNamesPattern(values []NamedPattern) []string {
	names := make([]string, len(values))
	for index, value := range values {
		names[index] = value.Name()
	}
	return names
}

// ResettablePattern is satisfied only by *CounterPattern because Add has a
// pointer receiver. Its method set is different from CounterPattern's.
type ResettablePattern interface {
	Add(int)
	Total() int
}

func IncrementPattern(counter ResettablePattern) int {
	counter.Add(1)
	return counter.Total()
}

// TaggedPattern demonstrates embedding for composition and method promotion.
type TaggedPattern struct {
	TaskPattern
	Tags []string
}

func EmbeddedNamePattern(tagged TaggedPattern) string {
	return tagged.Name() // Promoted from the embedded TaskPattern field.
}

// DescribeDynamicPattern shows a type switch over an interface value.
func DescribeDynamicPattern(value any) string {
	switch typed := value.(type) {
	case nil:
		return "nil"
	case int:
		if typed == 0 {
			return "zero int"
		}
		return "int"
	case string:
		return "string"
	case []int:
		return "int slice"
	default:
		return "other"
	}
}

// -----------------------------------------------------------------------------
// Advanced: typed nil inside an interface
// -----------------------------------------------------------------------------

type OptionalNodePattern struct {
	Label string
}

func (node *OptionalNodePattern) Name() string {
	if node == nil {
		return "<nil node>"
	}
	return node.Label
}

// TypedNilPattern returns false for interfaceIsNil: an interface is nil only
// when both its dynamic type and dynamic value are absent. Here the dynamic type
// is *OptionalNodePattern even though the stored pointer is nil.
func TypedNilPattern() (pointerIsNil, interfaceIsNil bool, safeName string) {
	var node *OptionalNodePattern
	var named NamedPattern = node
	return node == nil, named == nil, named.Name()
}

/*
Struct and method rules:

- Struct literals may use field names (`PointPattern{Row: 1, Col: 2}`), which
  are safer than positional literals when field order can change.
- The zero value is often useful. Initialize maps/channels inside constructors
  when their nil form cannot support required operations.
- Choose a pointer receiver for mutation, large structs, or consistency when
  any method needs a pointer. A value receiver operates on a copy.
- Go may insert & or * for convenient method calls on addressable concrete
  values, but interface satisfaction follows exact method sets.
- Embedding promotes fields/methods; it is composition, not subclassing.
- Interfaces are best kept small and defined near the consumer.

Interface representation:

- An interface value carries a dynamic type and dynamic value.
- Type assertions use `value, ok := x.(T)`. Omit ok only when failure should
  panic.
- Interface values are comparable, but comparing two interfaces panics when
  their dynamic values have a non-comparable type such as a slice.

Typed-nil guidance:

- A typed nil pointer stored in an interface makes the interface non-nil.
- Prefer returning a literal nil interface on failure. If nil pointer receivers
  are valid, methods must check the receiver before dereferencing it.
- This edge case matters in Go API/debugging interviews, but ordinary algorithm
  code should avoid designs that depend on it.
*/
