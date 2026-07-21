// Go standard-library patterns that commonly support interview solutions.
//
// This shortlist focuses on bufio, strings, strconv, and container/heap rather
// than broad application cookbooks such as JSON, CSV, or date handling.
package languagepatterns

import (
	"bufio"
	"container/heap"
	"io"
	"strconv"
	"strings"
)

// ScanIntsLibrary reads whitespace-separated ints. Scanner is convenient for
// token input; its default token limit is about 64 KiB, so increase it when a
// single token may be larger.
func ScanIntsLibrary(reader io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 1024), 1_000_000)

	values := make([]int, 0)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return values, nil
}

// ScanLinesLibrary uses Scanner's default ScanLines split function.
func ScanLinesLibrary(reader io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(reader)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// WriteIntsLibrary buffers many small writes, then checks Flush because errors
// may surface only when buffered data is sent to the underlying writer.
func WriteIntsLibrary(writer io.Writer, values []int) error {
	buffered := bufio.NewWriter(writer)
	for index, value := range values {
		if index > 0 {
			if err := buffered.WriteByte(' '); err != nil {
				return err
			}
		}
		if _, err := buffered.WriteString(strconv.Itoa(value)); err != nil {
			return err
		}
	}
	if err := buffered.WriteByte('\n'); err != nil {
		return err
	}
	return buffered.Flush()
}

// StringHelpersLibrary groups high-frequency strings operations. Split uses an
// exact separator; Fields collapses runs of whitespace.
func StringHelpersLibrary(text string) (words []string, normalized string, hasPrefix bool) {
	words = strings.Fields(text)
	normalized = strings.Join(words, " ")
	hasPrefix = strings.HasPrefix(normalized, "go")
	return words, normalized, hasPrefix
}

// ReplaceTokensLibrary is suitable for a small fixed set of literal
// replacements. It does not use regular expressions.
func ReplaceTokensLibrary(text string) string {
	replacer := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	return replacer.Replace(text)
}

// ParseWidthsLibrary demonstrates width-aware strconv functions. ParseInt
// returns int64; FormatInt converts it back without depending on int width.
func ParseWidthsLibrary(text string) (int64, string, error) {
	value, err := strconv.ParseInt(strings.TrimSpace(text), 10, 64)
	if err != nil {
		return 0, "", err
	}
	return value, strconv.FormatInt(value, 10), nil
}

// PriorityItemLibrary stores an explicit tie-break order so equal-priority
// entries behave deterministically.
type PriorityItemLibrary struct {
	Priority int
	Order    int
	Value    string
}

// PriorityQueueLibrary implements heap.Interface as a min-heap.
type PriorityQueueLibrary []PriorityItemLibrary

func (queue PriorityQueueLibrary) Len() int {
	return len(queue)
}

func (queue PriorityQueueLibrary) Less(left, right int) bool {
	if queue[left].Priority != queue[right].Priority {
		return queue[left].Priority < queue[right].Priority
	}
	return queue[left].Order < queue[right].Order
}

func (queue PriorityQueueLibrary) Swap(left, right int) {
	queue[left], queue[right] = queue[right], queue[left]
}

func (queue *PriorityQueueLibrary) Push(value any) {
	*queue = append(*queue, value.(PriorityItemLibrary))
}

func (queue *PriorityQueueLibrary) Pop() any {
	old := *queue
	last := len(old) - 1
	value := old[last]
	old[last] = PriorityItemLibrary{}
	*queue = old[:last]
	return value
}

// PriorityOrderLibrary initializes in O(n), then removes entries in priority
// order at O(log n) each.
func PriorityOrderLibrary(items []PriorityItemLibrary) []string {
	queue := PriorityQueueLibrary(append([]PriorityItemLibrary(nil), items...))
	heap.Init(&queue)

	result := make([]string, 0, len(queue))
	for queue.Len() > 0 {
		item := heap.Pop(&queue).(PriorityItemLibrary)
		result = append(result, item.Value)
	}
	return result
}

// AddPriorityLibrary demonstrates heap.Push after initialization.
func AddPriorityLibrary(queue *PriorityQueueLibrary, item PriorityItemLibrary) {
	heap.Push(queue, item)
}

/*
Library shortlist:

- bufio.Scanner: simple token/line input. Always check Scanner.Err and increase
  the token buffer when constraints can exceed the default.
- bufio.Reader: finer control for very large input or delimiter-based reads.
- bufio.Writer: batches output; always Flush and handle its error.
- strings: Fields, Split, Join, Contains, HasPrefix/HasSuffix, TrimSpace,
  Builder, and NewReplacer.
- strconv: Atoi/Itoa for int; ParseInt/ParseUint and FormatInt/FormatUint when
  base or width matters.
- container/heap: priority queues over a user-defined heap.Interface.

Heap caveats:

- heap.Init establishes heap order in O(n). The root is minimal according to
  Less, but the backing slice is not globally sorted.
- heap.Push and heap.Pop call the interface methods; do not call the queue's
  Push/Pop methods directly when maintaining an existing heap.
- Reverse Less for a max-heap. Include a deterministic tie breaker when output
  order among equal priorities matters.

Complexity still matters: standard-library helpers do not change the cost of
the underlying operation. For example, splitting and parsing are linear in the
input size, and each heap update is logarithmic.
*/
