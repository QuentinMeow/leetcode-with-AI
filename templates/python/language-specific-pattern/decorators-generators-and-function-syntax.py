"""
Decorators, generators, and Python function-call syntax.

These constructs are Python-specific enough that a non-Python programmer may
understand the algorithm but still be slowed down by the syntax.
"""

from __future__ import annotations

from collections.abc import Callable, Generator
from functools import wraps


# -----------------------------------------------------------------------------
# Function objects and decorators
# -----------------------------------------------------------------------------


def trace(func: Callable[..., int]) -> Callable[..., int]:
    @wraps(func)
    def wrapper(*args, **kwargs) -> int:
        print(f"calling {func.__name__}")
        return func(*args, **kwargs)

    return wrapper


@trace
def add(a: int, b: int) -> int:
    return a + b


# Decorator syntax:
#
#     @trace
#     def add(...):
#         ...
#
# is roughly equivalent to:
#
#     def add(...):
#         ...
#     add = trace(add)
#
# A decorator runs when the function is defined, not every time the function is
# called. The returned wrapper runs when the decorated function is called.


# -----------------------------------------------------------------------------
# *args and **kwargs
# -----------------------------------------------------------------------------


def collect_arguments(required: int, *args: int, **kwargs: int) -> tuple[int, tuple[int, ...], dict[str, int]]:
    return required, args, kwargs


def call_with_unpacking(nums: list[int], options: dict[str, int]) -> tuple[int, tuple[int, ...], dict[str, int]]:
    # `*nums` expands a sequence into positional arguments.
    # `**options` expands a dict into keyword arguments.
    return collect_arguments(*nums, **options)


# In a function definition, `*args` collects extra positional arguments into a
# tuple and `**kwargs` collects extra keyword arguments into a dict.
# In a function call, `*` and `**` unpack values instead.


# -----------------------------------------------------------------------------
# Positional-only and keyword-only parameters
# -----------------------------------------------------------------------------


def parameter_kinds(pos_only: int, /, normal: int, *, keyword_only: int) -> int:
    return pos_only + normal + keyword_only


def parameter_kinds_example() -> int:
    return parameter_kinds(1, 2, keyword_only=3)


# `pos_only` must be passed positionally because it appears before `/`.
# `keyword_only` must be passed by name because it appears after bare `*`.
# You rarely need this in LeetCode, but you may see it in Python APIs.


# -----------------------------------------------------------------------------
# Generators and yield
# -----------------------------------------------------------------------------


def countdown(n: int) -> Generator[int, None, None]:
    while n > 0:
        yield n
        n -= 1


def flatten(nested: list[list[int]]) -> Generator[int, None, None]:
    for row in nested:
        yield from row


def generator_example() -> list[int]:
    values = []

    for x in countdown(3):
        values.append(x)

    values.extend(flatten([[4, 5], [6]]))
    return values


# A function containing `yield` is a generator function. Calling it returns a
# generator object immediately; the body runs only when you iterate over it.


# -----------------------------------------------------------------------------
# Properties, static methods, and class methods
# -----------------------------------------------------------------------------


class Score:
    scale = 100

    def __init__(self, raw: int):
        self.raw = raw

    @property
    def percent(self) -> float:
        return self.raw / self.scale

    @staticmethod
    def clamp(value: int) -> int:
        return max(0, min(Score.scale, value))

    @classmethod
    def from_fraction(cls, numerator: int, denominator: int) -> Score:
        return cls(round(cls.scale * numerator / denominator))


def score_example() -> tuple[float, int, Score]:
    score = Score(80)
    return score.percent, Score.clamp(120), Score.from_fraction(1, 2)


"""
Concept explanations:

- Functions are objects in Python. You can pass them to other functions, store
  them in variables, and return them.
- A decorator is a callable that receives a function or class and returns a
  replacement. `@decorator` is syntax sugar for reassignment after definition.
- `functools.wraps` copies metadata such as `__name__` and docstring from the
  original function to the wrapper. Without it, debugging and test output may
  show the wrapper's name instead.
- `*args` and `**kwargs` have two meanings depending on location: collect in a
  function definition, unpack in a function call.
- `/` marks positional-only parameters. Bare `*` marks the start of keyword-only
  parameters. These are common in Python's own APIs but uncommon in LeetCode.
- `yield` pauses a function and produces a value. The function resumes after the
  yield when the next value is requested.
- Generators are lazy. They can save memory, but they are single-pass: once
  consumed, you need to create a new generator to iterate again.
- `yield from iterable` delegates yielding to another iterable. It is a concise
  way to flatten or compose generators.
- `@property` lets method logic be accessed like an attribute. Use it when a
  value is computed from object state.
- `@staticmethod` defines a function namespaced inside a class that does not
  receive `self` or `cls`.
- `@classmethod` receives the class as `cls`, which makes alternate constructors
  work correctly with subclasses.

Interview guidance:

- For algorithm interviews, decorators and custom generators are usually not
  required, but you should recognize them in starter code and explain them.
- Avoid clever decorators in live coding unless the problem is explicitly about
  design or reusable instrumentation.
- Use generator expressions for simple lazy pipelines; use `yield` when the
  logic needs multiple statements or recursive/layered generation.
"""
