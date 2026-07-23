"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

import itertools

# ====================================================================
# 6. Iteration / Comprehensions / Ranges
# ====================================================================

# Shows ranges, enumerate, zip, pairwise iteration, comprehensions, generators, and unpacking.
# Requires: import itertools
def iteration_patterns(
    nums: list[int], a: list[int], b: list[int]
) -> None:
    for i, x in enumerate(nums):
        pass

    for i in range(len(nums)):
        pass

    for i in range(len(nums) - 1, -1, -1):
        pass

    pairs1 = list(itertools.pairwise(nums))  # Python 3.10+
    pairs2 = list(
        zip(nums, nums[1:])
    )  # Older equivalent; slice copies.

    zipped = list(zip(a, b))  # Stops at shortest input.
    dot = sum(x * y for x, y in zip(a, b))

    # [<output_expr> for <element> in <iterable> if
    # <filter_condition>]
    # x is each number from nums; x * x is what gets stored.
    squares = [x * x for x in nums if x >= 0]

    # Braces with one expression make a set comprehension.
    unique_mods = {x % 10 for x in nums}

    # Braces with key: value make a dict comprehension.
    # enumerate(nums) yields (index, value), unpacked as i, x.
    index_by_value = {x: i for i, x in enumerate(nums)}

    # Generator expression: no list is built; sum/any/all pull
    # values lazily.
    total_pos = sum(x for x in nums if x > 0)
    has_even = any(x % 2 == 0 for x in nums)
    all_positive = all(x > 0 for x in nums)

    # Star unpacking: middle becomes a list of the leftover values.
    first, *middle, last = nums if len(nums) >= 2 else [0, 0]
    a0, b0 = 1, 2
    a0, b0 = b0, a0
