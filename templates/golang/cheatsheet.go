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
	"cmp"
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
	// The indexed read examples below require len(nums) >= 3.
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
	popped := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	_ = []any{
		nilSlice, empty, first, last, middle, size, capacity,
		firstX, containsX, copy1, copy2, copy3, matrixCopy,
		popped, stack,
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

func setCRUDPatterns(nums []int) {
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

func queueCRUDPatterns(values []int) []int {
	// A slice plus a head index gives amortized O(1) dequeue.
	queue := append([]int(nil), values...)
	head := 0
	queue = append(queue, 42)
	if head < len(queue) {
		front := queue[head]
		queue[head] = 0 // Release references here for pointer element types.
		head++
		_ = front
	}
	return queue[head:]
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
	if len(h) > 0 {
		h[0]++          // Example priority change.
		heap.Fix(&h, 0) // Restore after changing h[0].
		heap.Remove(&h, 0)
	}
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
		return cmp.Compare(b, a) // Descending without subtraction overflow.
	})

	sortedWords := slices.Clone(words)
	slices.SortFunc(sortedWords, func(a, b string) int {
		if len(a) != len(b) {
			return cmp.Compare(len(a), len(b))
		}
		return cmp.Compare(a, b)
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

	quotient, remainder, normalizedMod := 0, 0, 0
	if b != 0 {
		quotient = a / b  // Integer division truncates toward zero.
		remainder = a % b // Remainder has the sign of a.
		normalizedMod = ((remainder % b) + b) % b
	}

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

// ===================================================================
// 9. Goroutines / Channels / Synchronization
// ===================================================================

func squareWorkerPool(ctx context.Context, nums []int, workers int) []int {
	if workers <= 0 {
		return nil
	}

	type job struct {
		index int
		value int
	}
	type result struct {
		index int
		value int
	}

	jobs := make(chan job, len(nums))
	results := make(chan result, len(nums))
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
				select {
				case <-ctx.Done():
					return
				case results <- result{
					index: current.index,
					value: current.value * current.value,
				}:
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

func mutexCounterPattern(workers int) int {
	if workers <= 0 {
		return 0
	}
	var mu sync.Mutex
	var wg sync.WaitGroup
	count := 0

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
	wg.Wait()
	return count
}

func bufferedChannelPattern(values []int) []int {
	channel := make(chan int, len(values))
	for _, value := range values {
		channel <- value
	}
	close(channel)

	result := make([]int, 0, len(values))
	for value := range channel {
		result = append(result, value)
	}
	return result
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

	queue := []state{{node: start}}
	head := 0
	seen := map[int]struct{}{start: {}}

	for head < len(queue) {
		current := queue[head]
		head++
		if current.node == target {
			return current.distance
		}
		for _, neighbor := range graph[current.node] {
			if _, ok := seen[neighbor]; ok {
				continue
			}
			seen[neighbor] = struct{}{} // Mark on enqueue.
			queue = append(queue, state{
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
	// Interview grids are normally rectangular.
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
	if k <= 0 || k > len(nums) {
		return 0
	}
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
	groups := make(map[string][]string)
	for _, word := range words {
		key := sortString(word)
		groups[key] = append(groups[key], word)
	}

	answer := make([][]string, 0, len(groups))
	for _, group := range groups {
		answer = append(answer, group)
	}
	return answer
}

func longestConsecutive(nums []int) int {
	values := make(map[int]struct{}, len(nums))
	for _, value := range nums {
		values[value] = struct{}{}
	}
	best := 0
	for value := range values {
		if _, exists := values[value-1]; exists {
			continue
		}
		end := value
		for {
			if _, exists := values[end]; !exists {
				break
			}
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
	queue := []*TreeNode{root}
	head := 0

	for head < len(queue) {
		levelSize := len(queue) - head
		level := make([]int, 0, levelSize)
		for levelIndex := 0; levelIndex < levelSize; levelIndex++ {
			node := queue[head]
			head++
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
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

	queue := make([]int, 0, n)
	for node, degree := range indegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	order := make([]int, 0, n)
	for head := 0; head < len(queue); head++ {
		node := queue[head]
		order = append(order, node)
		for _, neighbor := range graph[node] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
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

func (s *MinStack) pop() (int, bool) {
	if len(s.stack) == 0 {
		return 0, false
	}
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return top[0], true
}

func (s *MinStack) minimum() (int, bool) {
	if len(s.stack) == 0 {
		return 0, false
	}
	return s.stack[len(s.stack)-1][1], true
}

func kadaneMaxSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
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
// 12. General Algorithm Pattern Add-ons
// ===================================================================

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

type Trie struct {
	root *TrieNode
}

func newTrie() *Trie {
	return &Trie{
		root: &TrieNode{children: make(map[rune]*TrieNode)},
	}
}

func (t *Trie) insert(word string) {
	node := t.root
	for _, char := range word {
		if node.children[char] == nil {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		node = node.children[char]
	}
	node.isWord = true
}

func (t *Trie) walk(text string) *TrieNode {
	node := t.root
	for _, char := range text {
		node = node.children[char]
		if node == nil {
			return nil
		}
	}
	return node
}

func (t *Trie) search(word string) bool {
	node := t.walk(word)
	return node != nil && node.isWord
}

func (t *Trie) startsWith(prefix string) bool {
	return t.walk(prefix) != nil
}

func (t *Trie) wildcardSearch(pattern string) bool {
	patternRunes := []rune(pattern)
	var search func(int, *TrieNode) bool
	search = func(index int, node *TrieNode) bool {
		if index == len(patternRunes) {
			return node.isWord
		}
		char := patternRunes[index]
		if char != '.' {
			child := node.children[char]
			return child != nil && search(index+1, child)
		}
		for _, child := range node.children {
			if search(index+1, child) {
				return true
			}
		}
		return false
	}
	return search(0, t.root)
}

func topKFrequent(nums []int, k int) []int {
	if k <= 0 {
		return nil
	}

	frequency := make(map[int]int)
	for _, value := range nums {
		frequency[value]++
	}

	buckets := make([][]int, len(nums)+1)
	for value, count := range frequency {
		buckets[count] = append(buckets[count], value)
	}

	answer := make([]int, 0, min(k, len(frequency)))
	for count := len(buckets) - 1; count > 0 && len(answer) < k; count-- {
		for _, value := range buckets[count] {
			answer = append(answer, value)
			if len(answer) == k {
				break
			}
		}
	}
	return answer // Ties may appear in any order.
}

func coinChangeTopDown(coins []int, amount int) int {
	impossible := amount + 1
	memo := make(map[int]int)

	var dp func(int) int
	dp = func(remaining int) int {
		switch {
		case remaining == 0:
			return 0
		case remaining < 0:
			return impossible
		}
		if cached, ok := memo[remaining]; ok {
			return cached
		}

		best := impossible
		for _, coin := range coins {
			best = min(best, 1+dp(remaining-coin))
		}
		memo[remaining] = best
		return best
	}

	answer := dp(amount)
	if answer >= impossible {
		return -1
	}
	return answer
}

func medianTwoSortedPartition(a, b []int) float64 {
	if len(a) > len(b) {
		return medianTwoSortedPartition(b, a)
	}
	if len(a)+len(b) == 0 {
		panic("at least one input must be non-empty")
	}

	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	m, n := len(a), len(b)
	half := (m + n + 1) / 2
	left, right := 0, m

	for left <= right {
		aCut := left + (right-left)/2
		bCut := half - aCut

		aLeft, aRight := minInt, maxInt
		bLeft, bRight := minInt, maxInt
		if aCut > 0 {
			aLeft = a[aCut-1]
		}
		if aCut < m {
			aRight = a[aCut]
		}
		if bCut > 0 {
			bLeft = b[bCut-1]
		}
		if bCut < n {
			bRight = b[bCut]
		}

		if aLeft <= bRight && bLeft <= aRight {
			if (m+n)%2 == 1 {
				return float64(max(aLeft, bLeft))
			}
			leftMax := float64(max(aLeft, bLeft))
			rightMin := float64(min(aRight, bRight))
			return (leftMax + rightMin) / 2
		}
		if aLeft > bRight {
			right = aCut - 1
		} else {
			left = aCut + 1
		}
	}
	panic("inputs must be sorted")
}

func meetingRoomsTwoPointer(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	starts := make([]int, 0, len(intervals))
	ends := make([]int, 0, len(intervals))
	for _, interval := range intervals {
		if interval.start >= interval.end {
			continue
		}
		starts = append(starts, interval.start)
		ends = append(ends, interval.end)
	}
	if len(starts) == 0 {
		return 0
	}
	slices.Sort(starts)
	slices.Sort(ends)

	rooms, best := 0, 0
	startIndex, endIndex := 0, 0
	for startIndex < len(starts) {
		if starts[startIndex] < ends[endIndex] {
			rooms++
			best = max(best, rooms)
			startIndex++
		} else {
			rooms--
			endIndex++
		}
	}
	return best
}

func insertInterval(
	intervals []Interval,
	newInterval Interval,
) []Interval {
	answer := make([]Interval, 0, len(intervals)+1)
	index := 0
	for index < len(intervals) &&
		intervals[index].end < newInterval.start {
		answer = append(answer, intervals[index])
		index++
	}
	for index < len(intervals) &&
		intervals[index].start <= newInterval.end {
		newInterval.start = min(newInterval.start, intervals[index].start)
		newInterval.end = max(newInterval.end, intervals[index].end)
		index++
	}
	answer = append(answer, newInterval)
	return append(answer, intervals[index:]...)
}

func eraseOverlapIntervals(intervals []Interval) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].end < intervals[j].end
	})

	removed, previousEnd := 0, 0
	hasPrevious := false
	for _, interval := range intervals {
		if !hasPrevious || interval.start >= previousEnd {
			previousEnd = interval.end
			hasPrevious = true
		} else {
			removed++
		}
	}
	return removed
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 {
		return head
	}

	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for step := 0; step < n; step++ {
		fast = fast.Next
		if fast == nil {
			return head
		}
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

type LRUNode struct {
	key   int
	value int
	prev  *LRUNode
	next  *LRUNode
}

type LRUCache struct {
	capacity int
	nodes    map[int]*LRUNode
	head     *LRUNode
	tail     *LRUNode
}

func newLRUCache(capacity int) *LRUCache {
	head, tail := &LRUNode{}, &LRUNode{}
	head.next = tail
	tail.prev = head
	return &LRUCache{
		capacity: capacity,
		nodes:    make(map[int]*LRUNode),
		head:     head,
		tail:     tail,
	}
}

func (c *LRUCache) remove(node *LRUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) addMostRecent(node *LRUNode) {
	previous := c.tail.prev
	node.prev, node.next = previous, c.tail
	previous.next = node
	c.tail.prev = node
}

func (c *LRUCache) get(key int) int {
	node, ok := c.nodes[key]
	if !ok {
		return -1
	}
	c.remove(node)
	c.addMostRecent(node)
	return node.value
}

func (c *LRUCache) put(key, value int) {
	if node, ok := c.nodes[key]; ok {
		c.remove(node)
		delete(c.nodes, key)
	}

	node := &LRUNode{key: key, value: value}
	c.nodes[key] = node
	c.addMostRecent(node)

	if len(c.nodes) > c.capacity {
		victim := c.head.next
		c.remove(victim)
		delete(c.nodes, victim.key)
	}
}

func gridBFSDistance(grid [][]int, start Point) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return nil
	}

	rows, cols := len(grid), len(grid[0])
	distance := make([][]int, rows)
	for row := range distance {
		distance[row] = make([]int, cols)
		for col := range distance[row] {
			distance[row][col] = -1
		}
	}
	if start.Row < 0 || start.Row >= rows ||
		start.Col < 0 || start.Col >= cols ||
		grid[start.Row][start.Col] != 0 {
		return distance
	}

	queue := []Point{start}
	distance[start.Row][start.Col] = 0
	for head := 0; head < len(queue); head++ {
		current := queue[head]
		for _, direction := range directions4 {
			next := Point{
				Row: current.Row + direction[0],
				Col: current.Col + direction[1],
			}
			if next.Row < 0 || next.Row >= rows ||
				next.Col < 0 || next.Col >= cols ||
				grid[next.Row][next.Col] != 0 ||
				distance[next.Row][next.Col] != -1 {
				continue
			}
			distance[next.Row][next.Col] =
				distance[current.Row][current.Col] + 1
			queue = append(queue, next) // Marked before enqueue, so no duplicates.
		}
	}
	return distance
}

type ArrayCursor struct {
	value      int
	arrayIndex int
	valueIndex int
}

type ArrayCursorHeap []ArrayCursor

func (h ArrayCursorHeap) Len() int           { return len(h) }
func (h ArrayCursorHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h ArrayCursorHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ArrayCursorHeap) Push(value any) {
	*h = append(*h, value.(ArrayCursor))
}

func (h *ArrayCursorHeap) Pop() any {
	old := *h
	value := old[len(old)-1]
	*h = old[:len(old)-1]
	return value
}

func mergeKSortedArrays(arrays [][]int) []int {
	pq := &ArrayCursorHeap{}
	totalLength := 0
	for arrayIndex, values := range arrays {
		totalLength += len(values)
		if len(values) > 0 {
			heap.Push(pq, ArrayCursor{
				value:      values[0],
				arrayIndex: arrayIndex,
			})
		}
	}

	answer := make([]int, 0, totalLength)
	for pq.Len() > 0 {
		current := heap.Pop(pq).(ArrayCursor)
		answer = append(answer, current.value)
		nextIndex := current.valueIndex + 1
		values := arrays[current.arrayIndex]
		if nextIndex < len(values) {
			heap.Push(pq, ArrayCursor{
				value:      values[nextIndex],
				arrayIndex: current.arrayIndex,
				valueIndex: nextIndex,
			})
		}
	}
	return answer
}

func combinationSum(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	answer := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(int, int)
	backtrack = func(start, remaining int) {
		if remaining == 0 {
			answer = append(answer, slices.Clone(path))
			return
		}
		for index := start; index < len(candidates); index++ {
			value := candidates[index]
			if value > remaining {
				break
			}
			path = append(path, value)
			backtrack(index, remaining-value)
			path = path[:len(path)-1]
		}
	}

	backtrack(0, target)
	return answer
}

func applyRangeAdditions(length int, updates [][3]int) []int {
	if length <= 0 {
		return nil
	}
	difference := make([]int, length+1)
	for _, update := range updates {
		left, right, delta := update[0], update[1], update[2]
		if left < 0 || left >= length || right < left || right >= length {
			continue
		}
		difference[left] += delta
		difference[right+1] -= delta
	}

	result := make([]int, length)
	running := 0
	for index := range result {
		running += difference[index]
		result[index] = running
	}
	return result
}

func multiSourceGridDistance(grid [][]int, sources []Point) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return nil
	}
	rows, cols := len(grid), len(grid[0])
	distance := make([][]int, rows)
	for row := range distance {
		distance[row] = make([]int, cols)
		for col := range distance[row] {
			distance[row][col] = -1
		}
	}

	queue := make([]Point, 0, len(sources))
	for _, source := range sources {
		if source.Row < 0 || source.Row >= rows ||
			source.Col < 0 || source.Col >= cols ||
			grid[source.Row][source.Col] != 0 ||
			distance[source.Row][source.Col] != -1 {
			continue
		}
		distance[source.Row][source.Col] = 0
		queue = append(queue, source)
	}

	for head := 0; head < len(queue); head++ {
		current := queue[head]
		for _, direction := range directions4 {
			next := Point{
				Row: current.Row + direction[0],
				Col: current.Col + direction[1],
			}
			if next.Row < 0 || next.Row >= rows ||
				next.Col < 0 || next.Col >= cols ||
				grid[next.Row][next.Col] != 0 ||
				distance[next.Row][next.Col] != -1 {
				continue
			}
			distance[next.Row][next.Col] =
				distance[current.Row][current.Col] + 1
			queue = append(queue, next)
		}
	}
	return distance
}

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))
	for index, temperature := range temperatures {
		for len(stack) > 0 &&
			temperatures[stack[len(stack)-1]] < temperature {
			previous := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[previous] = index - previous
		}
		stack = append(stack, index)
	}
	return answer
}

func countComponents(n int, edges [][2]int) int {
	dsu := newDSU(n)
	for _, edge := range edges {
		dsu.union(edge[0], edge[1])
	}
	return dsu.components
}

func findRedundantEdge(edges [][2]int) [2]int {
	maxNode := 0
	for _, edge := range edges {
		maxNode = max(maxNode, edge[0], edge[1])
	}
	dsu := newDSU(maxNode + 1)
	for _, edge := range edges {
		if !dsu.union(edge[0], edge[1]) {
			return edge
		}
	}
	return [2]int{}
}

func maxOverlappingIntervals(intervals []Interval) int {
	type event struct {
		time  int
		delta int
	}
	events := make([]event, 0, 2*len(intervals))
	for _, interval := range intervals {
		if interval.start >= interval.end {
			continue
		}
		events = append(events,
			event{time: interval.start, delta: 1},
			event{time: interval.end, delta: -1},
		)
	}
	slices.SortFunc(events, func(a, b event) int {
		if a.time != b.time {
			return cmp.Compare(a.time, b.time)
		}
		return cmp.Compare(a.delta, b.delta) // End before start on ties.
	})

	active, best := 0, 0
	for _, current := range events {
		active += current.delta
		best = max(best, active)
	}
	return best
}

func buildBinaryTreeLevelOrder(values []*int) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}
	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	index := 1
	for head := 0; head < len(queue) && index < len(values); head++ {
		node := queue[head]
		if index < len(values) && values[index] != nil {
			node.Left = &TreeNode{Val: *values[index]}
			queue = append(queue, node.Left)
		}
		index++
		if index < len(values) && values[index] != nil {
			node.Right = &TreeNode{Val: *values[index]}
			queue = append(queue, node.Right)
		}
		index++
	}
	return root
}

func buildTreePreorderInorder(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	inorderIndex := make(map[int]int, len(inorder))
	for index, value := range inorder {
		inorderIndex[value] = index
	}
	preorderIndex := 0
	var build func(int, int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		value := preorder[preorderIndex]
		preorderIndex++
		node := &TreeNode{Val: value}
		split := inorderIndex[value]
		node.Left = build(left, split-1)
		node.Right = build(split+1, right)
		return node
	}
	return build(0, len(inorder)-1)
}

func buildRootedTree(n int, edges [][2]int, root int) [][]int {
	graph := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	children := make([][]int, n)
	type state struct {
		node   int
		parent int
	}
	stack := []state{{node: root, parent: -1}}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, neighbor := range graph[current.node] {
			if neighbor == current.parent {
				continue
			}
			children[current.node] = append(children[current.node], neighbor)
			stack = append(stack, state{node: neighbor, parent: current.node})
		}
	}
	return children
}

func memoizedGridPaths(rows, cols int) int {
	memo := make(map[[2]int]int)
	var count func(int, int) int
	count = func(row, col int) int {
		if row == rows-1 && col == cols-1 {
			return 1
		}
		if row >= rows || col >= cols {
			return 0
		}
		key := [2]int{row, col}
		if cached, ok := memo[key]; ok {
			return cached
		}
		memo[key] = count(row+1, col) + count(row, col+1)
		return memo[key]
	}
	if rows <= 0 || cols <= 0 {
		return 0
	}
	return count(0, 0)
}

func runLengthEncode(s string) string {
	if s == "" {
		return ""
	}
	var builder strings.Builder
	for start := 0; start < len(s); {
		end := start + 1
		for end < len(s) && s[end] == s[start] {
			end++
		}
		builder.WriteByte(s[start])
		builder.WriteString(strconv.Itoa(end - start))
		start = end
	}
	return builder.String()
}

func intFromBase10Digits(digits []int) (int, bool) {
	value := 0
	for _, digit := range digits {
		if digit < 0 || digit > 9 {
			return 0, false
		}
		value = value*10 + digit
	}
	return value, true
}

func primeFactorCounts(n int) map[int]int {
	factors := make(map[int]int)
	if n < 0 {
		n = -n
	}
	for factor := 2; factor <= n/factor; factor++ {
		for n%factor == 0 {
			factors[factor]++
			n /= factor
		}
	}
	if n > 1 {
		factors[n]++
	}
	return factors
}

// ===================================================================
// 13. Main / Local Script Pattern
// ===================================================================

func solve(nums []int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
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
	assert(fixedWindowMaxSum([]int{1, 2, 3, 4}, 2) == 7, "fixed window")
	assert(fixedWindowMaxSum([]int{1, 2}, 3) == 0, "invalid fixed window")
	assert(kadaneMaxSubarray(nil) == 0, "empty Kadane")
	assert(
		slices.Equal(
			applyRangeAdditions(
				3,
				[][3]int{{0, 1, 2}, {1, 2, 3}},
			),
			[]int{2, 5, 3},
		),
		"difference array",
	)
	assert(
		slices.Equal(
			dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}),
			[]int{1, 1, 4, 2, 1, 1, 0, 0},
		),
		"daily temperatures",
	)
	assert(
		countComponents(5, [][2]int{{0, 1}, {1, 2}, {3, 4}}) == 2,
		"DSU components",
	)
	assert(
		findRedundantEdge([][2]int{{1, 2}, {1, 3}, {2, 3}}) == [2]int{2, 3},
		"redundant edge",
	)
	assert(
		maxOverlappingIntervals(
			[]Interval{{start: 0, end: 30}, {start: 5, end: 10}},
		) == 2,
		"interval sweep",
	)
	assert(memoizedGridPaths(3, 3) == 6, "memoized recursion")
	assert(runLengthEncode("aaabb") == "a3b2", "run-length encoding")
	parsedDigits, digitsOK := intFromBase10Digits([]int{1, 2, 3})
	assert(parsedDigits == 123 && digitsOK, "digits to integer")
	assert(maps.Equal(primeFactorCounts(12), map[int]int{2: 2, 3: 1}), "factors")
	assert(
		slices.Equal(
			topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2),
			[]int{1, 2},
		),
		"top K frequent",
	)
	assert(coinChangeTopDown([]int{1, 2, 5}, 11) == 3, "coin change")
	assert(
		medianTwoSortedPartition([]int{1, 3}, []int{2}) == 2,
		"median partition",
	)
	assert(
		meetingRoomsTwoPointer(
			[]Interval{{start: 0, end: 30}, {start: 5, end: 10}},
		) == 2,
		"meeting rooms",
	)
	assert(
		slices.Equal(
			insertInterval(
				[]Interval{{start: 1, end: 3}, {start: 6, end: 9}},
				Interval{start: 2, end: 5},
			),
			[]Interval{{start: 1, end: 5}, {start: 6, end: 9}},
		),
		"insert interval",
	)
	assert(
		eraseOverlapIntervals(
			[]Interval{
				{start: 1, end: 2},
				{start: 2, end: 3},
				{start: 1, end: 3},
			},
		) == 1,
		"erase overlap intervals",
	)
	assert(
		slices.Equal(
			mergeKSortedArrays(
				[][]int{{1, 4}, {1, 3, 5}, {2, 6}},
			),
			[]int{1, 1, 2, 3, 4, 5, 6},
		),
		"merge K arrays",
	)
	assert(len(combinationSum([]int{2, 3, 6, 7}, 7)) == 2, "combination sum")

	trie := newTrie()
	trie.insert("apple")
	assert(trie.search("apple"), "trie search")
	assert(!trie.search("app") && trie.startsWith("app"), "trie prefix")
	assert(trie.wildcardSearch("a..le"), "trie wildcard")

	cache := newLRUCache(2)
	cache.put(1, 1)
	cache.put(2, 2)
	assert(cache.get(1) == 1, "LRU get")
	cache.put(3, 3)
	assert(cache.get(2) == -1, "LRU eviction")

	list := &ListNode{
		Val:  1,
		Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}},
	}
	list = removeNthFromEnd(list, 2)
	assert(list.Val == 1 && list.Next.Val == 3, "remove Nth from end")

	distance := gridBFSDistance(
		[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		Point{Row: 0, Col: 0},
	)
	assert(distance[2][2] == 4, "grid BFS distance")
	multiDistance := multiSourceGridDistance(
		[][]int{{0, 0, 0}, {0, 0, 0}},
		[]Point{{Row: 0, Col: 0}, {Row: 1, Col: 2}},
	)
	assert(multiDistance[0][2] == 1, "multi-source BFS")
	assert(mutexCounterPattern(4) == 4, "mutex counter")
	assert(
		slices.Equal(bufferedChannelPattern([]int{1, 2, 3}), []int{1, 2, 3}),
		"buffered channel",
	)

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
