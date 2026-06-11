# LeetCode 1 — Two Sum
# https://leetcode.com/problems/two-sum/

from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # create a dict to store the value and the index
        val_to_idx = {}
        for i, val in enumerate(nums):
            if target - val in val_to_idx:
                return [val_to_idx[target - val], i]
            val_to_idx[val] = i
        
        return []


def run_assertion_tests():
    solution = Solution()

    def check(nums, target):
        result = solution.twoSum(nums, target)
        assert len(result) == 2
        assert result[0] != result[1]
        assert nums[result[0]] + nums[result[1]] == target

    check([2, 7, 11, 15], 9)
    check([3, 2, 4], 6)
    check([3, 3], 6)


if __name__ == "__main__":
    solution = Solution()
    print(f"Test case 1 (expected: indices [0, 1]): {solution.twoSum([2, 7, 11, 15], 9)}")
    print(f"Test case 2 (expected: indices [1, 2]): {solution.twoSum([3, 2, 4], 6)}")
    print(f"Test case 3 (expected: indices [0, 1]): {solution.twoSum([3, 3], 6)}")

    run_assertion_tests()
