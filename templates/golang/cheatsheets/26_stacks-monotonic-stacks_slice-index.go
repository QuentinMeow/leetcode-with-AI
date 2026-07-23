// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

// ===================================================================
// 26. Stacks and Monotonic Stacks
// ===================================================================

// hasValidBracketNesting returns whether (), [], and {} are properly nested; unrelated
// bytes are ignored. Time O(n); space O(n).
func hasValidBracketNesting(s string) bool {
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

// StackWithMinimum stores each value beside the minimum at that depth, allowing
// push, pop, and minimum queries in O(1).
type StackWithMinimum struct {
	stack [][2]int // value, minimum through this position
}

// push appends one value to the receiver's stack storage.
func (s *StackWithMinimum) push(value int) {
	currentMin := value
	if len(s.stack) > 0 {
		currentMin = min(currentMin, s.stack[len(s.stack)-1][1])
	}
	s.stack = append(s.stack, [2]int{value, currentMin})
}

// pop removes and returns the top stack value; false means the stack was empty.
func (s *StackWithMinimum) pop() (int, bool) {
	if len(s.stack) == 0 {
		return 0, false
	}
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return top[0], true
}

// minimum returns the current stack minimum in O(1); false means the stack is empty.
func (s *StackWithMinimum) minimum() (int, bool) {
	if len(s.stack) == 0 {
		return 0, false
	}
	return s.stack[len(s.stack)-1][1], true
}

// daysUntilWarmerTemperature returns how many later positions must pass before a
// strictly warmer value, or 0 if none. The stack stores unresolved indices with
// non-increasing temperatures. Time O(n); space O(n).
func daysUntilWarmerTemperature(temperatures []int) []int {
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
