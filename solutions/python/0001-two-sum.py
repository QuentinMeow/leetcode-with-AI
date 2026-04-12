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
        
