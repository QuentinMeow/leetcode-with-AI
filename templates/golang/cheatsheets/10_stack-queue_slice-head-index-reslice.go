// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

// ===================================================================
// 10. Stack and Queue
// ===================================================================

// A slice handles stacks directly. For queues, a head index is a strong default.
// container/heap is verbose because the type owns ordering through Less.

// queueWithHeadIndexExample demonstrates amortized O(1) dequeue with a monotonically
// increasing head index. It returns the unconsumed suffix.
func queueWithHeadIndexExample(values []int) []int {
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

// stackOperationsExample demonstrates last-in-first-out stack operations with
// a slice. Push with append; the final index peeks; shortening by one pops.
func stackOperationsExample(values []int) (top int, remaining []int, ok bool) {
	stack := append([]int(nil), values...)
	if len(stack) == 0 {
		return 0, stack, false
	}
	top = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return top, stack, true
}
