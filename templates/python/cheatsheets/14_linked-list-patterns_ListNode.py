"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

# ====================================================================
# 14. Linked-List Patterns
# ====================================================================

# Reuses nodes from two ascending linked lists to form one ascending list. Time O(a + b).
def merge_two_lists(
    a: ListNode | None, b: ListNode | None
) -> ListNode | None:
    dummy = tail = ListNode()
    while a and b:
        if a.val <= b.val:
            tail.next, a = a, a.next
        else:
            tail.next, b = b, b.next
        tail = tail.next
    tail.next = a or b
    return dummy.next


# Reverses a singly linked list in place. Time O(n), space O(1).
def reverse_list(head: ListNode | None) -> ListNode | None:
    prev = None
    cur = head
    while cur:
        next_node = cur.next
        cur.next = prev
        prev = cur
        cur = next_node
    return prev


# Uses Floyd slow/fast pointers; meeting proves a cycle. Time O(n), space O(1).
def linked_list_has_cycle_using_floyd_algorithm(head: ListNode | None) -> bool:
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow is fast:
            return True
    return False


# Keeps two pointers n nodes apart and removes the requested node in one pass.
def remove_nth_linked_list_node_from_end(
    head: ListNode | None, n: int
) -> ListNode | None:
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
