# LeetCode 2 — Add Two Numbers
# https://leetcode.com/problems/add-two-numbers/

from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        carry = 0
        start = ListNode()
        current = start
        while l1 or l2 or carry:
            val1 = l1.val if l1 else 0
            val2 = l2.val if l2 else 0
            total = val1 + val2 + carry
            carry = total // 10
            current.next = ListNode(total % 10)
            current = current.next
            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None
        return start.next


def build_list(values):
    dummy = ListNode()
    curr = dummy
    for value in values:
        curr.next = ListNode(value)
        curr = curr.next
    return dummy.next


def to_list(node):
    values = []
    while node:
        values.append(node.val)
        node = node.next
    return values


def run_assertion_tests():
    solution = Solution()

    result = solution.addTwoNumbers(build_list([2, 4, 3]), build_list([5, 6, 4]))
    assert to_list(result) == [7, 0, 8]

    result = solution.addTwoNumbers(build_list([0]), build_list([0]))
    assert to_list(result) == [0]

    result = solution.addTwoNumbers(
        build_list([9, 9, 9, 9, 9, 9, 9]),
        build_list([9, 9, 9, 9]),
    )
    assert to_list(result) == [8, 9, 9, 9, 0, 0, 0, 1]


if __name__ == "__main__":
    solution = Solution()

    result = solution.addTwoNumbers(build_list([2, 4, 3]), build_list([5, 6, 4]))
    print(f"Test case 1 (expected: [7, 0, 8]): {to_list(result)}")

    result = solution.addTwoNumbers(build_list([0]), build_list([0]))
    print(f"Test case 2 (expected: [0]): {to_list(result)}")

    result = solution.addTwoNumbers(
        build_list([9, 9, 9, 9, 9, 9, 9]),
        build_list([9, 9, 9, 9]),
    )
    print(f"Test case 3 (expected: [8, 9, 9, 9, 0, 0, 0, 1]): {to_list(result)}")

    run_assertion_tests()
