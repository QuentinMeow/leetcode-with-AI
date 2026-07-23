// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"slices"
	"unicode"
)

// ===================================================================
// 20. Two Pointers, Palindromes, and Linked Lists
// ===================================================================

// uniqueTripletsSummingToZero returns unique value triples whose sum is zero. Sorting enables a fixed
// first value plus inward-moving pointers; skipping equal values prevents duplicate
// triples. It mutates nums. Time O(n^2); output excluded from space O(1).
// Requires: import "slices"
func uniqueTripletsSummingToZero(nums []int) [][]int {
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

// moveZeroes moves zero values to the end in place while preserving non-zero order.
// Time O(n); space O(1).
func moveZeroes(nums []int) {
	write := 0
	for read, value := range nums {
		if value != 0 {
			nums[write], nums[read] = nums[read], nums[write]
			write++
		}
	}
}

// maximumWaterContainerAreaUsingTwoPointers solves “Container With Most Water.” Width
// shrinks every step, so only moving the shorter wall can possibly increase the
// limiting height. Time O(n); space O(1).
func maximumWaterContainerAreaUsingTwoPointers(height []int) int {
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

// countPalindromicSubstringsByExpandingCenters counts every palindromic substring. Each
// character and each gap is a center; matching characters expand outward until they
// differ. Time O(n^2); space O(1).
func countPalindromicSubstringsByExpandingCenters(s string) int {
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

// isPalindromeIgnoringNonLettersAndDigits compares Unicode letters and digits
// case-insensitively while skipping punctuation and spaces. Time O(number of runes);
// rune-copy space O(number of runes).
// Requires: import "unicode"
func isPalindromeIgnoringNonLettersAndDigits(s string) bool {
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

// mergeTwoLists reuses nodes from two ascending linked lists to produce one ascending
// list. Time O(a+b); space O(1).
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

// reverseList reverses a singly linked list in place by redirecting each Next pointer.
// Time O(n); space O(1).
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

// linkedListHasCycleUsingFloydAlgorithm uses Floyd's tortoise-and-hare method: one
// pointer advances one node and the other two. In a cycle the faster pointer must
// eventually meet the slower one; at a nil link no cycle exists. Time O(n); space O(1).
func linkedListHasCycleUsingFloydAlgorithm(head *ListNode) bool {
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

// removeNthLinkedListNodeFromEnd keeps fast n nodes ahead of slow. When fast reaches
// the tail, slow is immediately before the node to remove. A dummy node handles removal
// of the original head. Time O(n); space O(1).
func removeNthLinkedListNodeFromEnd(head *ListNode, n int) *ListNode {
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
