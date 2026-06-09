"""
Python modules, imports, and the `__main__` guard.

This matters less for LeetCode's single-function submissions, but it matters for
interviews that ask you to write a small script, test helpers locally, or explain
why code should not run at import time.
"""


# -----------------------------------------------------------------------------
# Module execution model
# -----------------------------------------------------------------------------


def module_identity_explanation() -> str:
    return (
        "Every .py file is a module. When Python runs a file directly, that "
        "module's __name__ is '__main__'. When another file imports it, "
        "__name__ is the module's import name."
    )


def solve(nums: list[int]) -> int:
    return sum(nums)


def main() -> None:
    sample = [1, 2, 3]
    print(solve(sample))


if __name__ == "__main__":
    # This block runs only when this file is executed directly:
    #
    #     python modules-main-and-imports.py
    #
    # It does not run when another module imports this file.
    main()


# -----------------------------------------------------------------------------
# Import styles
# -----------------------------------------------------------------------------


"""
Common import forms:

- `import math`
  Keeps the module namespace explicit: `math.sqrt(9)`.

- `from collections import deque`
  Imports one name directly: `deque()`.

- `import heapq as hq`
  Uses an alias. Prefer common aliases only when they improve readability.

- `from module import *`
  Avoid in interview and production code. It hides where names came from and can
  overwrite existing names.

Import-time rule:

- Top-level statements run at import time. Function and class definitions are
  executed to create objects, but their bodies do not run until called.
- Keep expensive work, input parsing, and print/debug code inside `main()` or
  inside the `if __name__ == "__main__":` block.
"""


# -----------------------------------------------------------------------------
# Package basics and relative imports
# -----------------------------------------------------------------------------


"""
Concept explanations:

- A module is one Python file loaded as an object. Its global variables,
  functions, and classes become attributes of that module object.
- `__name__` is a special module global. It is `"__main__"` for the entry-point
  file and usually the dotted import path for imported modules.
- The `if __name__ == "__main__":` guard separates reusable code from script
  behavior. It lets tests import your functions without executing sample runs.
- `__file__` is usually the path of the current module file. It is useful in
  scripts but not needed in normal LeetCode solutions.
- Python caches imported modules in `sys.modules`. Importing the same module
  again usually reuses the existing module object instead of rerunning the file.
- An `__init__.py` file marks a directory as a traditional package. Modern
  Python also supports namespace packages, but interview code rarely needs that.
- Relative imports like `from .utils import helper` are for code inside a
  package. Single-file interview solutions normally use absolute imports.

Interview guidance:

- In LeetCode, do not add a `main()` guard unless you are making a local
  template or helper file. The platform calls your `Solution` methods directly.
- In take-home scripts or live coding outside LeetCode, put demo code and input
  parsing behind the `__main__` guard.
- If an interviewer asks "what is `__main__`?", answer: it is the name Python
  gives to the module used as the program entry point.
"""
