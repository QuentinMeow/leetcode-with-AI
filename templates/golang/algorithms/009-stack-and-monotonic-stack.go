package algorithms

/*
009 - Stack and monotonic stack patterns

Use when the newest unresolved item matters most, or when you need the next
greater/smaller element to the left or right.
*/

// Variant 1: monotonic increasing stack for next smaller / previous smaller.
// Example problems: daily temperatures, largest rectangle, sum of subarray minimums.
// Time: O(n)
// Space: O(n)
func NextSmallerIndices(nums []int) []int {
	result := make([]int, len(nums))
	for i := range result {
		result[i] = -1
	}
	stack := make([]int, 0, len(nums))
	for i, x := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > x {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top] = i
		}
		stack = append(stack, i)
	}
	return result
}

// Variant 2: monotonic decreasing stack for next greater element.
// Example problems: daily temperatures, next greater element, stock span.
// Time: O(n)
// Space: O(n)
func DaysUntilWarmer(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))
	for i, temperature := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperature {
			previous := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[previous] = i - previous
		}
		stack = append(stack, i)
	}
	return answer
}

// Variant 3: plain stack for matching delimiters / cancellation.
// Example problems: valid parentheses, simplify path, remove adjacent duplicates.
// Time: O(n)
// Space: O(n)
func IsValidParentheses(s string) bool {
	closeToOpen := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		default:
			open, isClose := closeToOpen[ch]
			if !isClose || len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

type minStackEntry struct {
	value   int
	minimum int
}

// Variant 4: min stack with paired state.
// Example problems: Min Stack, stack with extra aggregate.
// Time: O(1) per operation
// Space: O(n)
type MinStack struct {
	stack []minStackEntry
}

func (s *MinStack) Push(value int) {
	currentMinimum := value
	if len(s.stack) > 0 {
		currentMinimum = min(value, s.stack[len(s.stack)-1].minimum)
	}
	s.stack = append(s.stack, minStackEntry{value, currentMinimum})
}

func (s *MinStack) Pop() int {
	last := len(s.stack) - 1
	value := s.stack[last].value
	s.stack = s.stack[:last]
	return value
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1].value
}

func (s *MinStack) GetMin() int {
	return s.stack[len(s.stack)-1].minimum
}
