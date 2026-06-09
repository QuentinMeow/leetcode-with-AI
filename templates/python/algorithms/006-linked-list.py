"""
006 - Linked list patterns

Use when pointer rewiring is central. In Python interviews, keep a small
ListNode definition nearby and rely on dummy nodes to simplify edge cases.
"""

from __future__ import annotations


class ListNode:
    def __init__(self, val: int = 0, next: ListNode | None = None) -> None:
        self.val = val
        self.next = next


# Variant 1: dummy head for insert/delete/merge.
# Example problems: merge two sorted lists, remove elements, partition list.
# Time: O(n + m)
# Space: O(1)
def merge_two_lists(a: ListNode | None, b: ListNode | None) -> ListNode | None:
    dummy = ListNode()
    tail = dummy

    while a and b:
        if a.val <= b.val:
            tail.next = a
            a = a.next
        else:
            tail.next = b
            b = b.next
        tail = tail.next

    tail.next = a or b
    return dummy.next


# Variant 2: reverse a linked list.
# Example problems: reverse list, reverse between, palindrome linked list helper.
# Time: O(n)
# Space: O(1)
def reverse_list(head: ListNode | None) -> ListNode | None:
    prev = None
    curr = head

    while curr:
        nxt = curr.next
        curr.next = prev
        prev = curr
        curr = nxt

    return prev


# Variant 3: fast and slow pointers.
# Example problems: middle of linked list, cycle detection, split list.
# Time: O(n)
# Space: O(1)
def has_cycle(head: ListNode | None) -> bool:
    slow = fast = head

    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow is fast:
            return True

    return False


# Variant 4: remove nth node from end.
# Example problems: remove nth from end, keep fixed gap between pointers.
# Time: O(n)
# Space: O(1)
def remove_nth_from_end(head: ListNode | None, n: int) -> ListNode | None:
    dummy = ListNode(0, head)
    fast = slow = dummy

    for _ in range(n):
        fast = fast.next

    while fast and fast.next:
        fast = fast.next
        slow = slow.next

    if slow.next:
        slow.next = slow.next.next

    return dummy.next
