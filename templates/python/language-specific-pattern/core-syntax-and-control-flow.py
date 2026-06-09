"""
Core Python syntax and control-flow semantics.

These are the Python-language rules that often surprise programmers coming from
Java, Go, C++, or JavaScript. They are separate from algorithm patterns and
separate from standard-library helpers.
"""

from __future__ import annotations


# -----------------------------------------------------------------------------
# Blocks are indentation-based
# -----------------------------------------------------------------------------


def indentation_example(nums: list[int]) -> int:
    total = 0

    for x in nums:
        if x > 0:
            total += x

    return total


# Python uses indentation to define blocks. Braces are not optional syntax; they
# are not part of block structure at all.


# -----------------------------------------------------------------------------
# Truthiness and None
# -----------------------------------------------------------------------------


def truthiness_examples(value: object, nums: list[int] | None) -> tuple[bool, bool]:
    is_truthy = bool(value)

    # Use `is None`, not `== None`, because None is a singleton object.
    missing = nums is None

    return is_truthy, missing


def empty_container_check(nums: list[int]) -> bool:
    # Empty containers are falsy; non-empty containers are truthy.
    return not nums


# Falsy built-ins include None, False, numeric zero, empty strings, and empty
# containers such as [], {}, set(), and ().


# -----------------------------------------------------------------------------
# Boolean operators return operands
# -----------------------------------------------------------------------------


def default_name(name: str | None) -> str:
    # `or` returns the first truthy operand, not necessarily a bool.
    return name or "anonymous"


def guarded_first(nums: list[int] | None) -> int | None:
    # `and` returns the first falsy operand or the final operand.
    return nums and nums[0]


# Use this idiom carefully: `0`, "", and [] are falsy even when they may be valid
# answers.


# -----------------------------------------------------------------------------
# Comparison chaining and membership
# -----------------------------------------------------------------------------


def in_bounds(r: int, c: int, rows: int, cols: int) -> bool:
    return 0 <= r < rows and 0 <= c < cols


def membership(x: int, values: set[int]) -> bool:
    return x in values


# `a < b < c` is equivalent to `a < b and b < c`, but `b` is evaluated once.
# `in` calls the container's membership logic, which is O(1) average for sets
# and dict keys, but O(n) for lists.


# -----------------------------------------------------------------------------
# Conditional expression
# -----------------------------------------------------------------------------


def sign_label(x: int) -> str:
    return "positive" if x > 0 else "non-positive"


# Python's ternary form is `A if condition else B`, not `condition ? A : B`.


# -----------------------------------------------------------------------------
# pass, Ellipsis, and placeholders
# -----------------------------------------------------------------------------


def not_implemented_yet() -> None:
    pass


def also_placeholder() -> None:
    ...


# `pass` is a statement that does nothing. `...` is the Ellipsis object; it is
# often used as a placeholder in stubs, but `pass` is clearer inside real code.


# -----------------------------------------------------------------------------
# Loop controls
# -----------------------------------------------------------------------------


def first_even(nums: list[int]) -> int | None:
    for x in nums:
        if x % 2 != 0:
            continue
        return x

    return None


def search_with_break(nums: list[int], target: int) -> bool:
    found = False

    for x in nums:
        if x == target:
            found = True
            break

    return found


# `break` exits the nearest loop. `continue` jumps to the next iteration of the
# nearest loop.


# -----------------------------------------------------------------------------
# Exceptions
# -----------------------------------------------------------------------------


def parse_non_negative_int(text: str) -> int | None:
    try:
        value = int(text)
    except ValueError:
        return None

    if value < 0:
        raise ValueError("expected a non-negative integer")

    return value


def cleanup_example(nums: list[int]) -> int:
    try:
        return nums[0]
    except IndexError:
        return -1
    finally:
        # finally runs whether the try block returned, raised, or completed.
        nums.append(0)


# In interviews, exceptions are useful for invalid API inputs, but most LeetCode
# problems expect explicit condition checks instead of exception-driven flow.


# -----------------------------------------------------------------------------
# Context managers and `with`
# -----------------------------------------------------------------------------


class TemporaryFlag:
    def __init__(self):
        self.enabled = False

    def __enter__(self) -> TemporaryFlag:
        self.enabled = True
        return self

    def __exit__(self, exc_type, exc_value, traceback) -> bool:
        self.enabled = False
        return False  # False means do not suppress exceptions.


def context_manager_example() -> bool:
    with TemporaryFlag() as flag:
        return flag.enabled


"""
Concept explanations:

- Python syntax uses significant indentation. A wrong indent changes program
  meaning or raises `IndentationError`.
- `None` is Python's null-like singleton. Use `x is None` and `x is not None`
  because identity is the intended check.
- Truthiness is broader than booleans. Empty containers are falsy, which is
  convenient, but dangerous when an empty value is different from missing data.
- `and` and `or` short-circuit and return one of their operands. They do not
  coerce the result to `True` or `False` unless you wrap with `bool(...)`.
- Chained comparisons are first-class syntax: `0 <= i < n` is preferred for
  bounds checks.
- `in` means membership, not substring-only search. It works with strings,
  lists, sets, dicts, tuples, and custom containers.
- The conditional expression is expression-level branching:
  `value_if_true if condition else value_if_false`.
- `raise` throws an exception. `try` / `except` handles selected exception
  types. `finally` runs cleanup code.
- A `with` statement uses the context-manager protocol: call `__enter__` before
  the block and `__exit__` after the block. Files, locks, and temporary state
  often use this shape.
- `assert condition` is for debugging assumptions, not user-facing validation.
  Python can remove asserts when run with optimization flags.

Interview guidance:

- Use Python's syntax to make invariants visible: `0 <= r < rows`,
  `if node is None`, `if not queue`, and `x in seen`.
- Avoid clever truthiness when `0`, empty string, or empty list is a valid
  answer distinct from "missing".
- Prefer explicit checks in algorithm problems; use exceptions for genuinely
  exceptional invalid states or API design questions.
"""
