"""
Python iteration, unpacking, and comprehension patterns.

These are not algorithms by themselves, but they make interview code shorter
and less error-prone once you know what each construct does.
"""

from itertools import pairwise


# -----------------------------------------------------------------------------
# enumerate: index + value
# -----------------------------------------------------------------------------


def first_index(nums: list[int], target: int) -> int:
    # enumerate(iterable) yields (index, value) pairs lazily.
    for i, x in enumerate(nums):
        if x == target:
            return i
    return -1


# -----------------------------------------------------------------------------
# range: half-open intervals and reverse loops
# -----------------------------------------------------------------------------


def suffix_sums(nums: list[int]) -> list[int]:
    suffix = [0] * (len(nums) + 1)

    for i in range(len(nums) - 1, -1, -1):
        suffix[i] = nums[i] + suffix[i + 1]

    return suffix


# range(stop) -> 0, 1, ..., stop - 1
# range(start, stop) -> start, ..., stop - 1
# range(start, stop, step) excludes stop even when step is negative.


# -----------------------------------------------------------------------------
# zip: walk multiple sequences together
# -----------------------------------------------------------------------------


def dot_product(a: list[int], b: list[int]) -> int:
    # zip pairs values by position and stops when the shortest input ends.
    return sum(x * y for x, y in zip(a, b))


def adjacent_differences(nums: list[int]) -> list[int]:
    # pairwise is Python 3.10+. Equivalent: zip(nums, nums[1:])
    return [b - a for a, b in pairwise(nums)]


def adjacent_differences_older(nums: list[int]) -> list[int]:
    return [b - a for a, b in zip(nums, nums[1:])]


# -----------------------------------------------------------------------------
# Unpacking and swapping
# -----------------------------------------------------------------------------


def unpacking_examples(pair: tuple[int, int], nums: list[int]) -> tuple[int, int, list[int]]:
    a, b = pair

    if len(nums) >= 2:
        nums[0], nums[-1] = nums[-1], nums[0]

    first, *middle, last = nums if len(nums) >= 2 else [0, 0]
    return a + first, b + last, middle


# -----------------------------------------------------------------------------
# Comprehensions: list, set, dict
# -----------------------------------------------------------------------------


def comprehension_examples(nums: list[int]) -> tuple[list[int], set[int], dict[int, int]]:
    # The expression comes first, then the loop, then optional filters.
    squares = [x * x for x in nums if x >= 0]
    parity_set = {x % 2 for x in nums}
    index_by_value = {x: i for i, x in enumerate(nums)}

    return squares, parity_set, index_by_value


def matrix_comprehension(rows: int, cols: int) -> list[list[int]]:
    return [[0 for _ in range(cols)] for _ in range(rows)]


# -----------------------------------------------------------------------------
# Generator expressions: lazy values, often used with sum/min/max/any/all
# -----------------------------------------------------------------------------


def generator_examples(nums: list[int]) -> tuple[int, bool, bool]:
    # Generator expressions produce values one at a time instead of building a
    # whole list. any() and all() short-circuit.
    total_positive = sum(x for x in nums if x > 0)
    has_even = any(x % 2 == 0 for x in nums)
    all_non_negative = all(x >= 0 for x in nums)

    return total_positive, has_even, all_non_negative


# -----------------------------------------------------------------------------
# for/else: else runs only if the loop did not break
# -----------------------------------------------------------------------------


def contains_prime_candidate(nums: list[int]) -> bool:
    for x in nums:
        if x > 1 and all(x % d != 0 for d in range(2, int(x**0.5) + 1)):
            break
    else:
        return False

    return True


# for/else is useful but can surprise readers. In interviews, an explicit flag is
# often clearer unless the interviewer already knows the construct.


# -----------------------------------------------------------------------------
# Mutating while iterating
# -----------------------------------------------------------------------------


def remove_zeroes_copy(nums: list[int]) -> list[int]:
    # Prefer building a new list over removing from the list while iterating.
    return [x for x in nums if x != 0]


def remove_zeroes_in_place(nums: list[int]) -> None:
    write = 0
    for read, value in enumerate(nums):
        if value != 0:
            nums[write] = value
            write += 1

    del nums[write:]


"""
Interview notes:

- `enumerate` avoids manual index bookkeeping.
- `zip` stops at the shortest input.
- Comprehensions create concrete containers; generator expressions are lazy.
- `sum(generator)`, `any(generator)`, and `all(generator)` are idiomatic.
- Be careful with slices inside loops: `nums[1:]` copies O(n).
- Do not remove from a list while iterating over it unless you fully control the
  index movement.

Concept explanations:

- Iteration in Python is based on the iterator protocol. A `for` loop asks an
  object for values one at a time; it does not expose integer indexes unless you
  ask for them with `enumerate` or `range`.
- `enumerate(nums)` is preferred over `for i in range(len(nums))` when you need
  both index and value. It reduces off-by-one and indexing mistakes.
- `range` is lazy and half-open. `range(n)` does not allocate a list of length
  n, and the stop value is excluded.
- `zip(a, b)` is lazy and stops at the shortest iterable. This is helpful for
  pairwise comparison but can hide length mismatches if equal length matters.
- Unpacking assigns several names from one iterable shape. `a, b = pair` expects
  exactly two values; `first, *middle, last = nums` captures the leftover values
  into a list named `middle`.
- Swapping with `a, b = b, a` is safe because the right side is evaluated before
  assignments happen.
- A list comprehension like `[f(x) for x in nums if ok(x)]` creates a new list.
  Set and dict comprehensions create `set` and `dict` objects.
- A generator expression like `(f(x) for x in nums)` is lazy. Use it with
  `sum`, `min`, `max`, `any`, or `all` when you do not need to store every
  intermediate value.
- `any(...)` returns as soon as it finds a truthy value. `all(...)` returns as
  soon as it finds a falsy value. This short-circuiting can matter for runtime.
- `for` / `else` means "run `else` only if the loop did not break." This is
  Python-specific and often surprising, so use it sparingly in interviews.
- Mutating a list while looping over it can skip elements because indexes shift.
  Prefer building a new list or using a controlled read/write pointer.
"""
