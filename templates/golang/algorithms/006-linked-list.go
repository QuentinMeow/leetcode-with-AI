package algorithms

/*
006 - Linked list patterns

Use when pointer rewiring is central. In Go interviews, keep a small ListNode
definition nearby and rely on dummy nodes to simplify edge cases.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// Variant 1: dummy head for insert/delete/merge.
// Example problems: merge two sorted lists, remove elements, partition list.
// Time: O(n + m)
// Space: O(1)
func MergeTwoLists(a, b *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next = a
			a = a.Next
		} else {
			tail.Next = b
			b = b.Next
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

// Variant 2: reverse a linked list.
// Example problems: reverse list, reverse between, palindrome linked list helper.
// Time: O(n)
// Space: O(1)
func ReverseList(head *ListNode) *ListNode {
	var previous *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}

// Variant 3: fast and slow pointers.
// Example problems: middle of linked list, cycle detection, split list.
// Time: O(n)
// Space: O(1)
func HasCycle(head *ListNode) bool {
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

// Variant 4: remove nth node from end.
// Example problems: remove nth from end, keep fixed gap between pointers.
// Time: O(n)
// Space: O(1)
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return dummy.Next
}
