"""
Python function, mutability, and scope patterns.

These are common sources of bugs for programmers who already know the algorithm
but are newer to Python's object model and name binding rules.
"""

from __future__ import annotations

from copy import deepcopy


# -----------------------------------------------------------------------------
# Names bind to objects; assignment does not copy containers
# -----------------------------------------------------------------------------


def aliasing_example(nums: list[int]) -> tuple[list[int], list[int]]:
    # Both alias and nums point at the same list object. copied points at a new
    # outer list with the same elements.
    alias = nums
    copied = nums.copy()

    alias.append(99)
    copied.append(100)

    return nums, copied


def copy_matrix(grid: list[list[int]]) -> tuple[list[list[int]], list[list[int]]]:
    shallow = grid.copy()
    deep = deepcopy(grid)

    if shallow:
        shallow[0][0] = -1  # Also mutates grid[0][0].
    if deep:
        deep[0][0] = -2  # Does not mutate grid.

    return shallow, deep


# -----------------------------------------------------------------------------
# Mutable default arguments
# -----------------------------------------------------------------------------


def bad_append(value: int, bucket: list[int] = []) -> list[int]:
    # Default arguments are evaluated once when the function is defined, not
    # each time the function is called. The same list is reused across calls.
    bucket.append(value)
    return bucket


def good_append(value: int, bucket: list[int] | None = None) -> list[int]:
    if bucket is None:
        bucket = []
    bucket.append(value)
    return bucket


# -----------------------------------------------------------------------------
# Return values vs in-place mutation
# -----------------------------------------------------------------------------


def sorted_copy(nums: list[int]) -> list[int]:
    return sorted(nums)


def sort_in_place(nums: list[int]) -> None:
    nums.sort()


# Interview habit: say whether your helper mutates its input or returns a new
# value. This avoids accidental state sharing in backtracking and graph problems.


# -----------------------------------------------------------------------------
# Nested helper functions and nonlocal
# -----------------------------------------------------------------------------


def count_tree_nodes(root: TreeNode | None) -> int:
    count = 0

    def dfs(node: TreeNode | None) -> None:
        # Without nonlocal, `count += 1` would create a new local variable named
        # count inside dfs and fail because it reads before assignment.
        nonlocal count
        if node is None:
            return

        count += 1
        dfs(node.left)
        dfs(node.right)

    dfs(root)
    return count


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


# `nonlocal` is needed when assigning to a variable from an enclosing function.
# Mutating an enclosed list/dict does not require nonlocal because the binding
# itself is not being reassigned.


def collect_paths(root: TreeNode | None) -> list[list[int]]:
    result: list[list[int]] = []
    path: list[int] = []

    def dfs(node: TreeNode | None) -> None:
        if node is None:
            return

        path.append(node.val)

        if node.left is None and node.right is None:
            result.append(path.copy())
        else:
            dfs(node.left)
            dfs(node.right)

        path.pop()

    dfs(root)
    return result


# -----------------------------------------------------------------------------
# Late binding in closures
# -----------------------------------------------------------------------------


def bad_multipliers() -> list:
    funcs = []
    for i in range(3):
        funcs.append(lambda x: x * i)
    return funcs


def good_multipliers() -> list:
    funcs = []
    for i in range(3):
        funcs.append(lambda x, factor=i: x * factor)
    return funcs


# In the bad version, every lambda reads the same final `i`.


"""
Interview notes:

- Assignment binds a name; it does not clone the object.
- `list.copy()` is shallow. For nested containers, use a comprehension or
  `copy.deepcopy` when a real deep copy is needed.
- Avoid mutable default arguments.
- `sorted(nums)` returns a new list; `nums.sort()` mutates in place and returns
  None.
- `nonlocal` is for rebinding an outer function variable from an inner helper.
- In backtracking, append/pop around recursion and copy the path when saving.

Concept explanations:

- Python variables are names bound to objects. `alias = nums` does not copy the
  list; it creates another name for the same list object.
- `is` checks whether two names refer to the same object. `==` checks equality
  by value according to the object's equality rules.
- `list.copy()` and `nums[:]` make a shallow copy: the outer list is new, but
  nested lists or objects inside it are shared.
- A deep copy recursively copies nested objects. It is sometimes useful for
  examples, but in interviews a targeted copy such as `[row[:] for row in grid]`
  is often clearer and faster for matrices.
- Default argument expressions are evaluated once at function definition time.
  This is why `bucket=[]` is dangerous. Use `None` as the default sentinel and
  create a fresh list inside the function.
- Python functions can mutate objects passed by the caller. They cannot rebind
  the caller's variable name itself, but they can mutate the object that name
  refers to.
- Nested helper functions are common in DFS/backtracking because they can read
  `result`, `path`, dimensions, or other context from the outer function.
- `nonlocal` lets an inner function reassign a name from the nearest enclosing
  function scope. It is different from `global`, which refers to module scope.
- Mutating an outer list or dict does not need `nonlocal`; rebinding the name
  does. `result.append(path)` mutates, but `result = []` rebinds.
- Closures capture variables, not snapshots of values. Lambdas created in a
  loop all see the final loop variable unless you freeze the value with a
  default argument such as `factor=i`.
"""
