"""
Python version-specific syntax for interview code.

Most modern LeetCode Python submissions run on Python 3.10+ today, but older
examples and some interview environments still use Python 3.8 or 3.9 syntax.
Know the equivalents so syntax does not distract from the algorithm.
"""

from __future__ import annotations

from functools import cache, lru_cache
from typing import Optional


# -----------------------------------------------------------------------------
# Type hint syntax: Python 3.9 / 3.10 vs older style
# -----------------------------------------------------------------------------


def modern_builtin_generics(nums: list[int]) -> dict[int, list[int]]:
    groups: dict[int, list[int]] = {}

    for x in nums:
        groups.setdefault(x % 2, []).append(x)

    return groups


def modern_union(node: TreeNode | None) -> int:
    if node is None:
        return 0
    return 1 + max(modern_union(node.left), modern_union(node.right))


def older_optional_style(node: Optional["TreeNode"]) -> int:
    if node is None:
        return 0
    return 1 + max(older_optional_style(node.left), older_optional_style(node.right))


class TreeNode:
    def __init__(
        self,
        val: int = 0,
        left: TreeNode | None = None,
        right: TreeNode | None = None,
    ):
        self.val = val
        self.left = left
        self.right = right


# Older equivalents you may see:
#
# from typing import Dict, List, Optional
#
# def f(nums: List[int]) -> Dict[int, List[int]]: ...
# def g(node: Optional[TreeNode]) -> int: ...
#
# Modern equivalents:
#
# def f(nums: list[int]) -> dict[int, list[int]]: ...
# def g(node: TreeNode | None) -> int: ...


# -----------------------------------------------------------------------------
# Assignment expression :=, Python 3.8+
# -----------------------------------------------------------------------------


def walrus_example(nums: list[int], target: int) -> int:
    seen: dict[int, int] = {}

    for i, x in enumerate(nums):
        # := assigns and returns the assigned value inside an expression.
        if (j := seen.get(target - x)) is not None:
            return j
        seen[x] = i

    return -1


def no_walrus_equivalent(nums: list[int], target: int) -> int:
    seen: dict[int, int] = {}

    for i, x in enumerate(nums):
        j = seen.get(target - x)
        if j is not None:
            return j
        seen[x] = i

    return -1


# -----------------------------------------------------------------------------
# Structural pattern matching, Python 3.10+
# -----------------------------------------------------------------------------


def classify_token(token: str) -> str:
    # `match` compares the subject against cases from top to bottom. `case _`
    # is the wildcard case, like default in a switch statement.
    match token:
        case "+" | "-" | "*" | "/":
            return "operator"
        case "(" | ")":
            return "paren"
        case _ if token.isdigit():
            return "number"
        case _:
            return "other"


def classify_token_if_else(token: str) -> str:
    if token in {"+", "-", "*", "/"}:
        return "operator"
    if token in {"(", ")"}:
        return "paren"
    if token.isdigit():
        return "number"
    return "other"


# Pattern matching is powerful, but if/elif is usually clearer in LeetCode unless
# the problem naturally branches on structured shapes.


# -----------------------------------------------------------------------------
# functools.cache, Python 3.9+
# -----------------------------------------------------------------------------


@cache
def fib_cache(n: int) -> int:
    # @cache memoizes by argument tuple. Arguments must be hashable.
    if n <= 1:
        return n
    return fib_cache(n - 1) + fib_cache(n - 2)


@lru_cache(maxsize=None)
def fib_lru_cache(n: int) -> int:
    if n <= 1:
        return n
    return fib_lru_cache(n - 1) + fib_lru_cache(n - 2)


# @cache is equivalent to @lru_cache(maxsize=None) and is shorter.


# -----------------------------------------------------------------------------
# Dictionary merge operators, Python 3.9+
# -----------------------------------------------------------------------------


def dict_merge(a: dict[str, int], b: dict[str, int]) -> dict[str, int]:
    merged = a | b  # Values from b win on duplicate keys.
    copied_then_updated = a.copy()
    copied_then_updated.update(b)

    assert merged == copied_then_updated
    return merged


"""
Quick version map:

- Python 3.7: dict preserves insertion order as a language guarantee.
- Python 3.8: assignment expression `:=`.
- Python 3.9: built-in collection generics `list[int]`, `dict[str, int]`;
  `dict | dict`; `functools.cache`.
- Python 3.10: union types `int | None`; structural pattern matching `match`.
- Python 3.11: faster CPython runtime and better error messages; no major
  LeetCode syntax change you must rely on.

Interview habit:

- Prefer simple syntax the interviewer can read.
- Use modern hints if the platform supports them.
- Be ready to translate modern hints to older `typing.List` / `Optional` if
  an interviewer uses an older Python template.

Concept explanations:

- `from __future__ import annotations` postpones evaluation of type annotations.
  This lets a class mention itself in annotations before the class object fully
  exists, such as `left: TreeNode | None`. It affects type hints, not runtime
  algorithm behavior.
- Type hints are mostly optional metadata in normal Python execution. They help
  readers, editors, and static type checkers, but Python will still run if the
  actual values do not match the hints unless explicit runtime checks exist.
- `list[int]` and `dict[str, int]` are Python 3.9+ generic aliases for built-in
  containers. Older code uses `typing.List[int]` and `typing.Dict[str, int]`.
- `int | None` is Python 3.10+ union syntax. It means "an int or None" and is
  equivalent to older `Optional[int]`.
- The walrus operator `:=` is an assignment expression. It is useful when a
  computed value is needed both for a condition and inside that branch. Avoid it
  when it makes the control flow harder to read.
- `match` / `case` is structural pattern matching, not just a C-style switch.
  It can destructure lists, tuples, classes, and mappings. In many LeetCode
  solutions, ordinary `if` / `elif` remains clearer.
- `@cache` is a decorator that stores return values by function arguments. It is
  excellent for top-down DP when the state is made of hashable values such as
  integers, strings, or tuples. It will fail for unhashable arguments like lists.
- `@lru_cache(maxsize=None)` is the older spelling for an unbounded cache.
  `@lru_cache(maxsize=...)` can cap memory, but most interview DP uses
  `maxsize=None`.
- `dict_a | dict_b` creates a merged dict where right-side values win on key
  conflicts. It is concise, but `.copy()` plus `.update()` is understood in
  older Python versions.
"""
