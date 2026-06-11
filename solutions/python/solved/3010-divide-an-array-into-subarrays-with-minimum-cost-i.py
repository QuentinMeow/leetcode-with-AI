# LeetCode 3010 — Divide an Array Into Subarrays With Minimum Cost I
# https://leetcode.com/problems/divide-an-array-into-subarrays-with-minimum-cost-i/

from typing import List

class Solution:
    def minimumCost(self, nums: List[int]) -> int:
        if len(nums) < 3:
            return -1

        fixed_cost = nums[0]
        first_min = min(nums[1], nums[2])
        second_min = max(nums[1], nums[2])

        for i in range(3, len(nums)):
            if nums[i] < second_min:
                second_min = max(nums[i], first_min)
                first_min = min(nums[i], first_min)

        return fixed_cost + first_min + second_min


def run_assertion_tests():
    solution = Solution()

    assert solution.minimumCost([1, 2, 3, 12]) == 6
    assert solution.minimumCost([5, 4, 3]) == 12
    assert solution.minimumCost([10, 3, 1, 1]) == 12


if __name__ == "__main__":
    solution = Solution()
    print(f"Test case 1 (expected: 6): {solution.minimumCost([1, 2, 3, 12])}")
    print(f"Test case 2 (expected: 12): {solution.minimumCost([5, 4, 3])}")
    print(f"Test case 3 (expected: 12): {solution.minimumCost([10, 3, 1, 1])}")

    run_assertion_tests()
