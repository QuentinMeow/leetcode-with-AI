// Go interview cheatsheet.
//
// Use this as a fast scan before an interview. It combines Go syntax,
// standard-library patterns, common data structures, and reusable
// algorithm templates in one executable file.
//
// The examples favor interview clarity over production abstraction.
// Adjacent lines often show equivalent operations and the caveats that
// commonly cause bugs.
//
// After editing this file, run:
//
//	gofmt -w templates/golang/cheatsheet.go
//	go run templates/golang/cheatsheet.go
//
// Target: Go 1.21+, matching the version documented for this repository.
package main

import (
	"container/heap"
	"context"
	"errors"
	"fmt"
	"io"
	"maps"
	"math"
	"math/bits"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"
)

// ===================================================================
// 0. Version / Syntax Quick Map
// ===================================================================

/*
Go 1.18+: generics, any alias, comparable constraint
Go 1.20+: errors.Join, comparable types satisfy comparable constraints
Go 1.21+: built-in min/max/clear, slices/maps standard packages

Core declarations:

	var x int                 // zero value: 0
	x := 3                    // infer type inside a function
	const mod = 1_000_000_007
	nums := []int{1, 2, 3}    // slice
	fixed := [3]int{1, 2, 3}  // array; length is part of the type
	count := map[string]int{} // map

Multiple assignment makes swaps easy:

	a, b = b, a

Only false is false. Integers, strings, slices, and pointers do not
implicitly convert to bool. Braces are required; semicolons are normally
inserted automatically.
*/

// ===================================================================
// 1. Common Imports
// ===================================================================

/*
container/heap: heap.Interface, heap.Push, heap.Pop, heap.Init
sort:           sort.Ints, sort.Strings, sort.Slice, sort.Search
slices:         slices.Sort, slices.Clone, slices.Equal, slices.Reverse
maps:           maps.Clone, maps.Equal
strings:        strings.Split, Join, Fields, Builder
strconv:        strconv.Atoi, Itoa, ParseInt
math:           math.Sqrt, math.Inf, math.Abs
math/bits:      bits.OnesCount, bits.Len
unicode/utf8:   utf8.RuneCountInString, utf8.ValidString
*/

// ===================================================================
// 2. Go Data Structures - CRUD Cheat Sheet
// ===================================================================

/*
Master memory table:

| Operation       | slice               | map                 | set          |
|-----------------|---------------------|---------------------|--------------|
| Empty           | []T{} / make([]T,0) | make(map[K]V)       | map[T]struct |
| Add one         | append              | m[k] = v            | s[x]=struct{}|
| Add many        | append(a, b...)     | loop assignments    | loop         |
| Read            | a[i]                | v, ok := m[k]       | _, ok := s[x]|
| Remove          | reslice/copy trick  | delete(m, k)        | delete(s, x) |
| Empty all       | clear / a = a[:0]   | clear(m)            | clear(s)     |

Key caveats:

- Arrays copy by value. Slices are small headers pointing to backing arrays.
- Map reads return the value type's zero value when a key is absent.
- Use the comma-ok form when zero and missing mean different things.
- append may reuse the backing array or allocate a new one.
- Map iteration order is deliberately unspecified.
- nil slices can be read, ranged over, and appended to.
- Writing to a nil map panics; reading from it is safe.
*/

func sliceCRUDPatterns(n, rows, cols int, nums []int) {
	x := 2

	// CREATE
	var nilSlice []int
	empty := []int{}
	withLength := make([]int, n)
	withCapacity := make([]int, 0, n)
	fromData := []int{1, 2, 3}
	_ = [3]int{1, 2, 3} // Array, not slice.

	// Matrix create: each row gets its own backing array.
	grid := make([][]int, rows)
	for r := range grid {
		grid[r] = make([]int, cols)
	}

	// READ
	first := nums[0] // Panics if empty.
	last := nums[len(nums)-1]
	middle := nums[1:3] // View sharing the backing array.
	size := len(nums)
	capacity := cap(nums)
	firstX := slices.Index(nums, x) // -1 if absent.
	containsX := slices.Contains(nums, x)

	// COPY
	copy1 := slices.Clone(nums)
	copy2 := append([]int(nil), nums...)
	copy3 := make([]int, len(nums))
	copy(copy3, nums)

	matrixCopy := make([][]int, len(grid))
	for r := range grid {
		matrixCopy[r] = slices.Clone(grid[r])
	}

	// UPDATE / ADD
	withCapacity = append(withCapacity, x)
	withCapacity = append(withCapacity, 4, 5)
	withCapacity = append(withCapacity, nums...)
	nums[0] = 99
	slices.Sort(nums)    // In-place ascending.
	slices.Reverse(nums) // In-place.

	// Insert x before index i.
	i := min(1, len(withCapacity))
	withCapacity = append(withCapacity, 0)
	copy(withCapacity[i+1:], withCapacity[i:])
	withCapacity[i] = x

	// DELETE while preserving order.
	if len(withCapacity) > 0 {
		i = 0
		copy(withCapacity[i:], withCapacity[i+1:])
		withCapacity = withCapacity[:len(withCapacity)-1]
	}

	// DELETE without preserving order, O(1).
	if len(fromData) > 0 {
		i = 0
		fromData[i] = fromData[len(fromData)-1]
		fromData = fromData[:len(fromData)-1]
	}

	clear(withLength)           // Zero elements; length is unchanged.
	withLength = withLength[:0] // Empty but retain capacity.
	nilSlice = nil              // Release reference to backing array.

	// Stack: append and trim the right end.
	stack := []int{1, 2}
	top := stack[len(stack)-1]
	popped := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	_ = []any{
		nilSlice, empty, first, last, middle, size, capacity,
		firstX, containsX, copy1, copy2, copy3, matrixCopy,
		top, popped, stack,
	}
}

func mapCRUDPatterns(words []string, nums []int) {
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
		key := sortString(word)
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

type Set[T comparable] map[T]struct{}

func newSet[T comparable](values ...T) Set[T] {
	set := make(Set[T], len(values))
	for _, value := range values {
		set[value] = struct{}{}
	}
	return set
}

func (s Set[T]) add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) contains(value T) bool {
	_, ok := s[value]
	return ok
}

func (s Set[T]) remove(value T) {
	delete(s, value)
}

func setCRUDPatterns(nums []int) {
	set := newSet(nums...)
	set.add(42)
	exists := set.contains(42)
	set.remove(42)
	size := len(set)
	clear(set)
	_ = []any{exists, size}
}

// Queue with a head index gives amortized O(1) dequeue. Periodically
// compact long-lived queues so consumed elements can be garbage-collected.
type Queue[T any] struct {
	data []T
	head int
}

func (q *Queue[T]) push(value T) {
	q.data = append(q.data, value)
}

func (q *Queue[T]) pop() (T, bool) {
	var zero T
	if q.head == len(q.data) {
		return zero, false
	}
	value := q.data[q.head]
	q.data[q.head] = zero
	q.head++

	if q.head > 1024 && q.head*2 >= len(q.data) {
		q.data = append([]T(nil), q.data[q.head:]...)
		q.head = 0
	}
	return value, true
}

func (q *Queue[T]) front() (T, bool) {
	var zero T
	if q.head == len(q.data) {
		return zero, false
	}
	return q.data[q.head], true
}

func (q *Queue[T]) len() int {
	return len(q.data) - q.head
}

// IntHeap is a min-heap. Reverse Less for a max-heap.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(value any) {
	*h = append(*h, value.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	value := old[n-1]
	*h = old[:n-1]
	return value
}

func heapCRUDPatterns(nums []int) {
	h := IntHeap(slices.Clone(nums))
	heap.Init(&h)     // O(n).
	heap.Push(&h, 42) // O(log n).
	smallest := h[0]  // Peek, O(1); requires non-empty heap.
	popped := heap.Pop(&h).(int)
	heap.Fix(&h, 0)    // Restore after changing h[0].
	heap.Remove(&h, 0) // Remove by index, O(log n).
	clear(h)
	h = h[:0]
	_ = []any{smallest, popped}
}

// ===================================================================
// 3. Copying / Pointers / Mutation
// ===================================================================

func copyAndAliasing(grid [][]int, nums []int) {
	alias := nums // Shared backing array.
	shallow1 := slices.Clone(nums)
	shallow2 := append([]int(nil), nums...)

	matrixShallow := slices.Clone(grid) // Rows still shared.
	matrixCopy := make([][]int, len(grid))
	for r := range grid {
		matrixCopy[r] = slices.Clone(grid[r])
	}

	path := []int{1, 2}
	result := make([][]int, 0)
	result = append(result, slices.Clone(path)) // Save snapshot.

	_ = []any{
		alias, shallow1, shallow2, matrixShallow, matrixCopy, result,
	}
}

func appendCaveat(nums []int) []int {
	// The returned slice must be assigned by the caller because append may
	// allocate a new backing array and change the slice header.
	return append(nums, 42)
}

func mutateElements(nums []int) {
	// Element writes are visible through every slice sharing this array.
	for i := range nums {
		nums[i] *= 2
	}
}

func rebindLocally(nums []int) {
	// This changes only the local slice header, not the caller's header.
	nums = append(nums, 42)
	_ = nums
}

// ===================================================================
// 4. Sorting / Comparison / Heap
// ===================================================================

type ScoreRecord struct {
	name  string
	score int
	age   int
}

type Interval struct {
	start int
	end   int
}

func sortingPatterns(nums []int, words []string, records []ScoreRecord) {
	sortedNums := slices.Clone(nums)
	slices.Sort(sortedNums)
	slices.SortFunc(sortedNums, func(a, b int) int {
		return b - a // Descending; avoid subtraction if overflow matters.
	})

	sortedWords := slices.Clone(words)
	slices.SortFunc(sortedWords, func(a, b string) int {
		if len(a) != len(b) {
			return len(a) - len(b)
		}
		return strings.Compare(a, b)
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

func binarySearchLibrary(sortedNums []int, target int) (int, bool) {
	index, found := slices.BinarySearch(sortedNums, target)
	return index, found
}

func lowerBoundLibrary(sortedNums []int, target int) int {
	// First index i where sortedNums[i] >= target.
	return sort.Search(len(sortedNums), func(i int) bool {
		return sortedNums[i] >= target
	})
}

func upperBoundLibrary(sortedNums []int, target int) int {
	// First index i where sortedNums[i] > target.
	return sort.Search(len(sortedNums), func(i int) bool {
		return sortedNums[i] > target
	})
}

// ===================================================================
// 5. Iteration / Ranges
// ===================================================================

func iterationPatterns(nums []int, matrix [][]int) {
	for index, value := range nums {
		_, _ = index, value
	}

	for index := range nums {
		_ = index
	}

	for index := len(nums) - 1; index >= 0; index-- {
		_ = nums[index]
	}

	for row := range matrix {
		for col := range matrix[row] {
			_ = matrix[row][col]
		}
	}

	// range copies each element into value. Mutate through nums[i].
	for i := range nums {
		nums[i] *= 2
	}

	// Before Go 1.22, closures could accidentally capture the reused range
	// variable. Rebinding remains clear and portable interview code.
	functions := make([]func() int, 0, len(nums))
	for _, value := range nums {
		value := value
		functions = append(functions, func() int { return value })
	}

	// Maps are unordered. Sort keys for deterministic output.
	counts := map[string]int{"b": 2, "a": 1}
	keys := make([]string, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	_ = []any{functions, keys}
}

// ===================================================================
// 6. Strings / Numbers / Bits
// ===================================================================

func stringPatterns(s string, nums []int) {
	// A string is immutable bytes, commonly UTF-8.
	byteLength := len(s)
	firstByte := s[0] // Panics if empty.
	validUTF8 := utf8.ValidString(s)
	runeCount := utf8.RuneCountInString(s)

	// range decodes UTF-8 and yields byte index plus rune.
	for byteIndex, r := range s {
		_, _ = byteIndex, r
	}

	bytesCopy := []byte(s)
	runes := []rune(s)
	firstRune := runes[0] // Panics if there are no runes.
	reconstructed := string(runes)

	words := strings.Fields(s)
	csvParts := strings.Split(s, ",")
	joined := strings.Join(words, " ")
	contains := strings.Contains(s, "needle")
	firstIndex := strings.Index(s, "needle") // Byte index, -1 if absent.
	count := strings.Count(s, "a")
	replacedOnce := strings.Replace(s, "old", "new", 1)
	replacedAll := strings.ReplaceAll(s, "old", "new")
	trimmed := strings.TrimSpace(s)
	lower := strings.ToLower(s)
	upper := strings.ToUpper(s)
	starts := strings.HasPrefix(s, "pre")
	ends := strings.HasSuffix(s, "suf")

	var builder strings.Builder
	for _, word := range words {
		if builder.Len() > 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(word)
	}
	built := builder.String()

	numberText := make([]string, len(nums))
	for i, value := range nums {
		numberText[i] = strconv.Itoa(value)
	}
	answerLine := strings.Join(numberText, ",")

	parsed, err := strconv.Atoi("42")
	parsed64, err64 := strconv.ParseInt("-42", 10, 64)
	formatted := strconv.Itoa(parsed)

	_ = []any{
		byteLength, firstByte, validUTF8, runeCount, bytesCopy,
		firstRune, reconstructed, csvParts, joined, contains,
		firstIndex, count, replacedOnce, replacedAll, trimmed,
		lower, upper, starts, ends, built, answerLine, parsed,
		err, parsed64, err64, formatted,
	}
}

func normalizeAlnumLower(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(unicode.ToLower(r))
		}
	}
	return builder.String()
}

func numericBitPatterns(a, b, n int, mask uint) {
	absolute := a
	if absolute < 0 {
		absolute = -absolute
	}
	clamped := max(0, min(n, 100))

	quotient := a / b  // Integer division truncates toward zero.
	remainder := a % b // Remainder has the sign of a.
	normalizedMod := ((a % b) + b) % b

	positiveInfinity := math.Inf(1)
	squareRoot := math.Sqrt(float64(n))
	gcdValue := gcd(a, b)

	mask |= 1 << 3
	hasBit := mask&(1<<3) != 0
	mask &^= 1 << 3 // Clear bit; &^ is AND NOT.
	mask ^= 1 << 3  // Toggle bit.
	ones := bits.OnesCount(mask)
	bitLength := bits.Len(mask)

	_ = []any{
		absolute, clamped, quotient, remainder, normalizedMod,
		positiveInfinity, squareRoot, gcdValue, hasBit, ones, bitLength,
	}
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func sortString(s string) string {
	runes := []rune(s)
	slices.Sort(runes)
	return string(runes)
}

// ===================================================================
// 7. Functions / Closures / Errors / Defer
// ===================================================================

func variadicSum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func namedAndFirstClassFunctions(nums []int) {
	add := func(a, b int) int { return a + b }
	operation := add
	result := operation(2, 3)
	total := variadicSum(nums...)
	_ = []any{result, total}
}

func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

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

- Deferred calls run LIFO when the surrounding function returns.
- Arguments are evaluated when defer is registered.
- Use defer for cleanup after successfully acquiring a resource.
- panic/recover is not ordinary error handling; return errors instead.
*/

// ===================================================================
// 8. Structs / Methods / Interfaces / Generics
// ===================================================================

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Point struct {
	Row int
	Col int
}

type Person struct {
	Name string
	Age  int
}

func (p Person) label() string {
	return fmt.Sprintf("%s:%d", p.Name, p.Age)
}

func (p *Person) birthday() {
	p.Age++
}

type Stringer interface {
	String() string
}

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) push(value T) {
	s.data = append(s.data, value)
}

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

type Number interface {
	~int | ~int64 | ~float64
}

func sumValues[T Number](values []T) T {
	var total T
	for _, value := range values {
		total += value
	}
	return total
}

func cloneMap[K comparable, V any](source map[K]V) map[K]V {
	result := make(map[K]V, len(source))
	for key, value := range source {
		result[key] = value
	}
	return result
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

// ===================================================================
// 9. Goroutines / Channels / Synchronization
// ===================================================================

func squareWorkerPool(ctx context.Context, nums []int, workers int) []int {
	type job struct {
		index int
		value int
	}
	type result struct {
		index int
		value int
	}

	jobs := make(chan job)
	results := make(chan result)
	var wg sync.WaitGroup

	worker := func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case current, ok := <-jobs:
				if !ok {
					return
				}
				results <- result{
					index: current.index,
					value: current.value * current.value,
				}
			}
		}
	}

	wg.Add(workers)
	for workerIndex := 0; workerIndex < workers; workerIndex++ {
		go worker()
	}

	go func() {
		defer close(jobs)
		for index, value := range nums {
			select {
			case <-ctx.Done():
				return
			case jobs <- job{index: index, value: value}:
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	squared := make([]int, len(nums))
	for result := range results {
		squared[result.index] = result.value
	}
	return squared
}

func timeoutPattern(parent context.Context) error {
	ctx, cancel := context.WithTimeout(parent, 100*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(10 * time.Millisecond):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

/*
Concurrency memory rules:

- Start a goroutine only when you know how it stops.
- The sender normally closes a channel; receivers should not.
- Sending to or closing a closed channel panics.
- Receiving from a closed channel returns buffered values, then zero,false.
- A nil channel blocks forever in send/receive cases.
- Run `go test -race ./...` for concurrent production code.
*/

// ===================================================================
// 10. Algorithm Skeletons
// ===================================================================

func twoSumHash(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for index, value := range nums {
		if otherIndex, ok := seen[target-value]; ok {
			return []int{otherIndex, index}
		}
		seen[value] = index
	}
	return nil
}

func twoPointersSorted(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		total := nums[left] + nums[right]
		switch {
		case total == target:
			return []int{left, right}
		case total < target:
			left++
		default:
			right--
		}
	}
	return nil
}

func slidingWindowAtMostKDistinct(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	count := make(map[int]int)
	left, total := 0, 0
	for right, value := range nums {
		count[value]++
		for len(count) > k {
			leftValue := nums[left]
			count[leftValue]--
			if count[leftValue] == 0 {
				delete(count, leftValue)
			}
			left++
		}
		total += right - left + 1
	}
	return total
}

func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		middle := left + (right-left)/2
		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return left
}

func firstFeasible(low, high int, can func(int) bool) int {
	for low < high {
		middle := low + (high-low)/2
		if can(middle) {
			high = middle
		} else {
			low = middle + 1
		}
	}
	return low
}

func prefixSumSubarrayCount(nums []int, target int) int {
	seen := map[int]int{0: 1}
	prefix, answer := 0, 0
	for _, value := range nums {
		prefix += value
		answer += seen[prefix-target]
		seen[prefix]++
	}
	return answer
}

func bfsGraph(graph map[int][]int, start, target int) int {
	type state struct {
		node     int
		distance int
	}

	queue := Queue[state]{}
	queue.push(state{node: start})
	seen := newSet(start)

	for queue.len() > 0 {
		current, _ := queue.pop()
		if current.node == target {
			return current.distance
		}
		for _, neighbor := range graph[current.node] {
			if seen.contains(neighbor) {
				continue
			}
			seen.add(neighbor) // Mark on enqueue.
			queue.push(state{
				node:     neighbor,
				distance: current.distance + 1,
			})
		}
	}
	return -1
}

var directions4 = [][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func dfsIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])

	var dfs func(int, int)
	dfs = func(row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols {
			return
		}
		if grid[row][col] != '1' {
			return
		}
		grid[row][col] = '0'
		for _, direction := range directions4 {
			dfs(row+direction[0], col+direction[1])
		}
	}

	count := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				count++
				dfs(row, col)
			}
		}
	}
	return count
}

func backtrackingSubsets(nums []int) [][]int {
	answer := make([][]int, 0, 1<<len(nums))
	path := make([]int, 0, len(nums))

	var backtrack func(int)
	backtrack = func(start int) {
		answer = append(answer, slices.Clone(path))
		for index := start; index < len(nums); index++ {
			path = append(path, nums[index])
			backtrack(index + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return answer
}

func monotonicNextGreater(nums []int) []int {
	answer := make([]int, len(nums))
	for index := range answer {
		answer[index] = -1
	}

	stack := make([]int, 0) // Indices; values decrease.
	for index, value := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < value {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[top] = index
		}
		stack = append(stack, index)
	}
	return answer
}

func mergeIntervals(intervals []Interval) []Interval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	merged := make([]Interval, 0, len(intervals))
	for _, current := range intervals {
		if len(merged) == 0 ||
			current.start > merged[len(merged)-1].end {
			merged = append(merged, current)
			continue
		}
		last := &merged[len(merged)-1]
		last.end = max(last.end, current.end)
	}
	return merged
}

type DSU struct {
	parent     []int
	size       []int
	components int
}

func newDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for index := 0; index < n; index++ {
		parent[index] = index
		size[index] = 1
	}
	return &DSU{parent: parent, size: size, components: n}
}

func (d *DSU) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) union(a, b int) bool {
	rootA, rootB := d.find(a), d.find(b)
	if rootA == rootB {
		return false
	}
	if d.size[rootA] < d.size[rootB] {
		rootA, rootB = rootB, rootA
	}
	d.parent[rootB] = rootA
	d.size[rootA] += d.size[rootB]
	d.components--
	return true
}

// ===================================================================
// 11. High-Frequency Add-ons
// ===================================================================

func setBasedUniqueWindow(s string) int {
	// Byte-oriented version, appropriate when the problem guarantees ASCII.
	seen := make(map[byte]struct{})
	left, best := 0, 0
	for right := 0; right < len(s); right++ {
		for {
			if _, ok := seen[s[right]]; !ok {
				break
			}
			delete(seen, s[left])
			left++
		}
		seen[s[right]] = struct{}{}
		best = max(best, right-left+1)
	}
	return best
}

func minSubarrayLenAtLeast(nums []int, target int) int {
	left, total := 0, 0
	best := len(nums) + 1
	for right, value := range nums {
		total += value
		for total >= target {
			best = min(best, right-left+1)
			total -= nums[left]
			left++
		}
	}
	if best == len(nums)+1 {
		return 0
	}
	return best
}

func fixedWindowMaxSum(nums []int, k int) int {
	window := 0
	for _, value := range nums[:k] {
		window += value
	}
	best := window
	for right := k; right < len(nums); right++ {
		window += nums[right] - nums[right-k]
		best = max(best, window)
	}
	return best
}

func exactlyKDistinct(nums []int, k int) int {
	return slidingWindowAtMostKDistinct(nums, k) -
		slidingWindowAtMostKDistinct(nums, k-1)
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	answer := make([][]int, 0)

	for index, value := range nums {
		if index > 0 && value == nums[index-1] {
			continue
		}
		left, right := index+1, len(nums)-1
		for left < right {
			total := value + nums[left] + nums[right]
			switch {
			case total == 0:
				answer = append(
					answer,
					[]int{value, nums[left], nums[right]},
				)
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			case total < 0:
				left++
			default:
				right--
			}
		}
	}
	return answer
}

func moveZeroes(nums []int) {
	write := 0
	for read, value := range nums {
		if value != 0 {
			nums[write], nums[read] = nums[read], nums[write]
			write++
		}
	}
}

func maxAreaContainer(height []int) int {
	left, right, best := 0, len(height)-1, 0
	for left < right {
		best = max(
			best,
			(right-left)*min(height[left], height[right]),
		)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return best
}

func countPalindromesExpand(s string) int {
	expand := func(left, right int) int {
		count := 0
		for left >= 0 && right < len(s) && s[left] == s[right] {
			count++
			left--
			right++
		}
		return count
	}

	count := 0
	for center := 0; center < len(s); center++ {
		count += expand(center, center)
		count += expand(center, center+1)
	}
	return count
}

func isAlnumPalindrome(s string) bool {
	runes := []rune(s)
	left, right := 0, len(runes)-1
	for left < right {
		for left < right && !unicode.IsLetter(runes[left]) &&
			!unicode.IsDigit(runes[left]) {
			left++
		}
		for left < right && !unicode.IsLetter(runes[right]) &&
			!unicode.IsDigit(runes[right]) {
			right--
		}
		if unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func addDecimalStrings(a, b string) string {
	i, j, carry := len(a)-1, len(b)-1, 0
	answer := make([]byte, 0, max(len(a), len(b))+1)
	for i >= 0 || j >= 0 || carry > 0 {
		x, y := 0, 0
		if i >= 0 {
			x = int(a[i] - '0')
			i--
		}
		if j >= 0 {
			y = int(b[j] - '0')
			j--
		}
		total := x + y + carry
		answer = append(answer, byte(total%10)+'0')
		carry = total / 10
	}
	slices.Reverse(answer)
	return string(answer)
}

func compareVersionNumbers(first, second string) int {
	a := strings.Split(first, ".")
	b := strings.Split(second, ".")
	for index := 0; index < max(len(a), len(b)); index++ {
		x, y := 0, 0
		if index < len(a) {
			x, _ = strconv.Atoi(a[index])
		}
		if index < len(b) {
			y, _ = strconv.Atoi(b[index])
		}
		if x < y {
			return -1
		}
		if x > y {
			return 1
		}
	}
	return 0
}

func searchRotated(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		}
		if nums[left] <= nums[middle] {
			if nums[left] <= target && target < nums[middle] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else if nums[middle] < target && target <= nums[right] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}

func groupAnagramsCountKey(words []string) [][]string {
	groups := make(map[[26]int][]string)
	for _, word := range words {
		var count [26]int
		for _, char := range word {
			count[char-'a']++
		}
		groups[count] = append(groups[count], word)
	}

	answer := make([][]string, 0, len(groups))
	for _, group := range groups {
		answer = append(answer, group)
	}
	return answer
}

func longestConsecutive(nums []int) int {
	values := newSet(nums...)
	best := 0
	for value := range values {
		if values.contains(value - 1) {
			continue
		}
		end := value
		for values.contains(end) {
			end++
		}
		best = max(best, end-value)
	}
	return best
}

type PrefixSum1D struct {
	prefix []int
}

func newPrefixSum1D(nums []int) PrefixSum1D {
	prefix := make([]int, len(nums)+1)
	for index, value := range nums {
		prefix[index+1] = prefix[index] + value
	}
	return PrefixSum1D{prefix: prefix}
}

func (p PrefixSum1D) sumRange(left, right int) int {
	return p.prefix[right+1] - p.prefix[left]
}

type PrefixSum2D struct {
	prefix [][]int
}

func newPrefixSum2D(matrix [][]int) PrefixSum2D {
	rows := len(matrix)
	cols := 0
	if rows > 0 {
		cols = len(matrix[0])
	}

	prefix := make([][]int, rows+1)
	for row := range prefix {
		prefix[row] = make([]int, cols+1)
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			prefix[row+1][col+1] = matrix[row][col] +
				prefix[row][col+1] +
				prefix[row+1][col] -
				prefix[row][col]
		}
	}
	return PrefixSum2D{prefix: prefix}
}

func (p PrefixSum2D) sumRegion(r1, c1, r2, c2 int) int {
	return p.prefix[r2+1][c2+1] -
		p.prefix[r1][c2+1] -
		p.prefix[r2+1][c1] +
		p.prefix[r1][c1]
}

func mergeTwoLists(a, b *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next, a = a, a.Next
		} else {
			tail.Next, b = b, b.Next
		}
		tail = tail.Next
	}
	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}
	return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
	var previous *ListNode
	for head != nil {
		next := head.Next
		head.Next = previous
		previous = head
		head = next
	}
	return previous
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	answer := make([][]int, 0)
	queue := Queue[*TreeNode]{}
	queue.push(root)

	for queue.len() > 0 {
		levelSize := queue.len()
		level := make([]int, 0, levelSize)
		for levelIndex := 0; levelIndex < levelSize; levelIndex++ {
			node, _ := queue.pop()
			level = append(level, node.Val)
			if node.Left != nil {
				queue.push(node.Left)
			}
			if node.Right != nil {
				queue.push(node.Right)
			}
		}
		answer = append(answer, level)
	}
	return answer
}

func topologicalSortKahn(n int, edges [][2]int) []int {
	graph := make([][]int, n)
	indegree := make([]int, n)
	for _, edge := range edges {
		prerequisite, course := edge[0], edge[1]
		graph[prerequisite] = append(graph[prerequisite], course)
		indegree[course]++
	}

	queue := Queue[int]{}
	for node, degree := range indegree {
		if degree == 0 {
			queue.push(node)
		}
	}

	order := make([]int, 0, n)
	for queue.len() > 0 {
		node, _ := queue.pop()
		order = append(order, node)
		for _, neighbor := range graph[node] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue.push(neighbor)
			}
		}
	}
	if len(order) != n {
		return nil
	}
	return order
}

type WeightedEdge struct {
	to     int
	weight int
}

type DistanceState struct {
	distance int
	node     int
}

type DistanceHeap []DistanceState

func (h DistanceHeap) Len() int { return len(h) }

func (h DistanceHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h DistanceHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *DistanceHeap) Push(value any) {
	*h = append(*h, value.(DistanceState))
}

func (h *DistanceHeap) Pop() any {
	old := *h
	value := old[len(old)-1]
	*h = old[:len(old)-1]
	return value
}

func dijkstra(graph map[int][]WeightedEdge, start int) map[int]int {
	distance := map[int]int{start: 0}
	pq := &DistanceHeap{{distance: 0, node: start}}
	heap.Init(pq)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(DistanceState)
		if current.distance != distance[current.node] {
			continue // Stale heap entry.
		}
		for _, edge := range graph[current.node] {
			nextDistance := current.distance + edge.weight
			oldDistance, seen := distance[edge.to]
			if !seen || nextDistance < oldDistance {
				distance[edge.to] = nextDistance
				heap.Push(pq, DistanceState{
					distance: nextDistance,
					node:     edge.to,
				})
			}
		}
	}
	return distance
}

func validParentheses(s string) bool {
	closeToOpen := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0, len(s))
	for index := 0; index < len(s); index++ {
		char := s[index]
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
			continue
		}
		if expected, ok := closeToOpen[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != expected {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

type MinStack struct {
	stack [][2]int // value, minimum through this position
}

func (s *MinStack) push(value int) {
	currentMin := value
	if len(s.stack) > 0 {
		currentMin = min(currentMin, s.stack[len(s.stack)-1][1])
	}
	s.stack = append(s.stack, [2]int{value, currentMin})
}

func (s *MinStack) pop() int {
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return top[0]
}

func (s *MinStack) minimum() int {
	return s.stack[len(s.stack)-1][1]
}

func kadaneMaxSubarray(nums []int) int {
	best, current := nums[0], nums[0]
	for _, value := range nums[1:] {
		current = max(value, current+value)
		best = max(best, current)
	}
	return best
}

func climbStairsRolling(n int) int {
	if n <= 2 {
		return n
	}
	previous2, previous1 := 1, 2
	for step := 3; step <= n; step++ {
		previous2, previous1 = previous1, previous1+previous2
	}
	return previous1
}

func longestCommonSubsequence(a, b string) int {
	dp := make([][]int, len(a)+1)
	for index := range dp {
		dp[index] = make([]int, len(b)+1)
	}
	for i := len(a) - 1; i >= 0; i-- {
		for j := len(b) - 1; j >= 0; j-- {
			if a[i] == b[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[0][0]
}

func canPartition01Knapsack(nums []int) bool {
	total := 0
	for _, value := range nums {
		total += value
	}
	if total%2 != 0 {
		return false
	}

	target := total / 2
	possible := make([]bool, target+1)
	possible[0] = true
	for _, value := range nums {
		for current := target; current >= value; current-- {
			possible[current] = possible[current] ||
				possible[current-value]
		}
	}
	return possible[target]
}

func permutationsUsed(nums []int) [][]int {
	answer := make([][]int, 0)
	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			answer = append(answer, slices.Clone(path))
			return
		}
		for index, value := range nums {
			if used[index] {
				continue
			}
			used[index] = true
			path = append(path, value)
			backtrack()
			path = path[:len(path)-1]
			used[index] = false
		}
	}

	backtrack()
	return answer
}

func base10Digits(n int) []int {
	if n < 0 {
		n = -n
	}
	if n == 0 {
		return []int{0}
	}
	digits := make([]int, 0)
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	slices.Reverse(digits)
	return digits
}

func isPrimeTrial(n int) bool {
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return n == 2
	}
	for factor := 3; factor <= n/factor; factor += 2 {
		if n%factor == 0 {
			return false
		}
	}
	return true
}

func sieveIsPrime(n int) []bool {
	isPrime := make([]bool, n+1)
	for value := 2; value <= n; value++ {
		isPrime[value] = true
	}
	for prime := 2; prime <= n/prime; prime++ {
		if !isPrime[prime] {
			continue
		}
		for multiple := prime * prime; multiple <= n; multiple += prime {
			isPrime[multiple] = false
		}
	}
	return isPrime
}

func singleNumberXOR(nums []int) int {
	answer := 0
	for _, value := range nums {
		answer ^= value
	}
	return answer
}

// ===================================================================
// 12. Main / Local Script Pattern
// ===================================================================

func solve(nums []int) int {
	return sumValues(nums)
}

func assert(condition bool, message string) {
	if !condition {
		panic("check failed: " + message)
	}
}

func runSmokeChecks() {
	assert(solve([]int{1, 2, 3}) == 6, "solve")
	assert(
		slices.Equal(twoSumHash([]int{2, 7, 11, 15}, 9), []int{0, 1}),
		"two sum",
	)
	assert(lowerBound([]int{1, 2, 2, 4}, 2) == 1, "lower bound")
	assert(setBasedUniqueWindow("abcabcbb") == 3, "unique window")
	assert(addDecimalStrings("999", "1") == "1000", "decimal strings")
	assert(validParentheses("([]{})"), "valid parentheses")
	assert(kadaneMaxSubarray([]int{-2, 1, -3, 4, -1, 2, 1}) == 6, "Kadane")
	assert(canPartition01Knapsack([]int{1, 5, 11, 5}), "0/1 knapsack")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	assert(
		slices.Equal(
			squareWorkerPool(ctx, []int{1, 2, 3}, 2),
			[]int{1, 4, 9},
		),
		"worker pool",
	)
}

func main() {
	// LeetCode calls solution functions directly. Keep main only for local
	// execution and remove it when a submission editor requires that.
	fmt.Println(solve([]int{1, 2, 3}))
	runSmokeChecks()
}
