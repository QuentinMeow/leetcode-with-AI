"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

# ====================================================================
# 4. Copying / Binding / Mutability
# ====================================================================

# Contrasts shared bindings, shallow copies, nested aliasing, and safe matrix construction.
def copy_and_aliasing(grid: list[list[int]], nums: list[int]) -> None:
    alias = nums  # Same list object.
    shallow1 = nums[:]
    shallow2 = nums.copy()
    shallow3 = list(nums)

    matrix_shallow = grid[:]  # New outer list, shared rows.
    matrix_copy = [row[:] for row in grid]  # Common matrix copy.

    path: list[int] = []
    result: list[list[int]] = []
    result.append(path.copy())  # Save a snapshot in backtracking.
    # result.append(path)
    # AVOID: later path mutations affect saved row.

    # Mutable default args: avoid `def f(bucket=[]): ...`; use
    # None sentinel.
