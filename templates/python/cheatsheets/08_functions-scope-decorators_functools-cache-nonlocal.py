"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

import functools

# ====================================================================
# 8. Functions / Scope / Decorators
# ====================================================================

# Shows positional, keyword, default, variadic positional, and variadic keyword arguments.
def function_call_syntax(
    required: int, *args: int, **kwargs: int
) -> None:
    collected_args = args  # tuple
    collected_kwargs = kwargs  # dict

    values = [1, 2, 3]
    options = {"required": 1}
    # function_call_syntax(*values)
    # function_call_syntax(**options)


# ScopeTreeNode is a minimal binary node for the nested-scope example below.
class ScopeTreeNode:
    # Local copy so this section does not depend on the Classes
    # section.
    # Initializes a new instance and establishes its invariants.
    def __init__(
        self,
        val: int = 0,
        left: "ScopeTreeNode | None" = None,
        right: "ScopeTreeNode | None" = None,
    ) -> None:
        self.val = val
        self.left = left
        self.right = right


# Shows nonlocal mutation: the nested traversal updates a binding in its enclosing function.
def nested_helper_and_nonlocal(root: ScopeTreeNode | None) -> int:
    count = 0

    # visit_connected_land marks every land cell connected to this cell.
    def visit_subtree(node: ScopeTreeNode | None) -> None:
        # nonlocal means "when assigning count, use the
        # count in the outer
        # function scope." Without it, count += 1 would
        # create/read a local name.
        nonlocal count
        if node is None:
            return
        count += 1
        visit_subtree(node.left)
        visit_subtree(node.right)

    visit_subtree(root)
    return count




# Demonstrates functools.cache memoization: repeated arguments reuse the stored return value.
# Requires: import functools
@functools.cache
def cached_dp(i: int, remaining: int) -> int:
    # @functools.cache means: cached_dp = functools.cache(cached_dp).
    # It memoizes by the argument tuple (i, remaining), so args
    # must be hashable.
    if remaining == 0:
        return 1
    if i == 0 or remaining < 0:
        return 0
    return cached_dp(i - 1, remaining) + cached_dp(
        i - 1, remaining - 1
    )




# Demonstrates lru_cache(maxsize=None), the pre-Python-3.9 spelling of unbounded memoization.
# Requires: import functools
@functools.lru_cache(maxsize=None)
def cached_dp_old_spelling(i: int) -> int:
    return (
        i
        if i <= 1
        else cached_dp_old_spelling(i - 1)
        + cached_dp_old_spelling(i - 2)
    )
