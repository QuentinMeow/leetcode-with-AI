"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

import sys

# ====================================================================
# 20. Matrix Simulation and Recursion Limits
# ====================================================================

# Compacts non-zero matrix values downward within each column.
def apply_gravity(board: list[list[int]]) -> None:
    rows, cols = len(board), len(board[0])
    for c in range(cols):
        write = rows - 1
        for r in range(rows - 1, -1, -1):
            if board[r][c] != 0:
                board[write][c] = board[r][c]
                write -= 1
        for r in range(write, -1, -1):
            board[r][c] = 0


# Returns the first coordinate starting three equal non-zero values in a row.
def find_three_in_a_row(
    board: list[list[int]],
) -> set[tuple[int, int]]:
    rows, cols = len(board), len(board[0])
    crush: set[tuple[int, int]] = set()
    for r in range(1, rows - 1):
        for c in range(cols):
            if (
                board[r][c]
                and board[r - 1][c] == board[r][c] == board[r + 1][c]
            ):
                crush.update({(r - 1, c), (r, c), (r + 1, c)})
    for r in range(rows):
        for c in range(1, cols - 1):
            if (
                board[r][c]
                and board[r][c - 1] == board[r][c] == board[r][c + 1]
            ):
                crush.update({(r, c - 1), (r, c), (r, c + 1)})
    return crush




# Raises Python recursion depth for unusually deep recursive traversal; use carefully because
# stack memory still grows.
# Requires: import sys
def raise_recursion_limit_for_deep_depth_first_search() -> None:
    sys.setrecursionlimit(10**6)
