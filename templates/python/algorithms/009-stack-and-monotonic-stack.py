"""
009 - Stack and monotonic stack patterns

Use when the newest unresolved item matters most, or when you need the next
greater/smaller element to the left or right.
"""


# Variant 1: monotonic increasing stack for next smaller / previous smaller.
# Example problems: daily temperatures, largest rectangle, sum of subarray minimums.
# Time: O(n)
# Space: O(n)
def next_smaller_indices(nums: list[int]) -> list[int]:
    result = [-1] * len(nums)
    stack: list[int] = []

    for i, x in enumerate(nums):
        while stack and nums[stack[-1]] > x:
            result[stack.pop()] = i
        stack.append(i)

    return result


# Variant 2: monotonic decreasing stack for next greater element.
# Example problems: daily temperatures, next greater element, stock span.
# Time: O(n)
# Space: O(n)
def days_until_warmer(temperatures: list[int]) -> list[int]:
    answer = [0] * len(temperatures)
    stack: list[int] = []

    for i, temp in enumerate(temperatures):
        while stack and temperatures[stack[-1]] < temp:
            prev = stack.pop()
            answer[prev] = i - prev
        stack.append(i)

    return answer


# Variant 3: plain stack for matching delimiters / cancellation.
# Example problems: valid parentheses, simplify path, remove adjacent duplicates.
# Time: O(n)
# Space: O(n)
def is_valid_parentheses(s: str) -> bool:
    close_to_open = {")": "(", "]": "[", "}": "{"}
    stack: list[str] = []

    for ch in s:
        if ch in close_to_open.values():
            stack.append(ch)
        elif not stack or stack.pop() != close_to_open[ch]:
            return False

    return not stack


# Variant 4: min stack with paired state.
# Example problems: Min Stack, stack with extra aggregate.
# Time: O(1) per operation
# Space: O(n)
class MinStack:
    def __init__(self) -> None:
        self.stack: list[tuple[int, int]] = []

    def push(self, val: int) -> None:
        current_min = val if not self.stack else min(val, self.stack[-1][1])
        self.stack.append((val, current_min))

    def pop(self) -> int:
        return self.stack.pop()[0]

    def top(self) -> int:
        return self.stack[-1][0]

    def get_min(self) -> int:
        return self.stack[-1][1]
