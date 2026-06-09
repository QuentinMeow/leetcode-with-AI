"""
Python class and OOP patterns that show up in coding interviews.

LeetCode often provides classes such as ListNode, TreeNode, and Node. The main
interview skill is knowing how Python objects, attributes, identity, and
comparison behave while you implement algorithms around those classes.
"""

from __future__ import annotations

from collections.abc import Iterator
from dataclasses import dataclass, field


# -----------------------------------------------------------------------------
# LeetCode-style hand-written node classes
# -----------------------------------------------------------------------------


class ListNode:
    def __init__(self, val: int = 0, next: ListNode | None = None):
        self.val = val
        self.next = next


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


def linked_list_with_dummy(values: list[int]) -> ListNode | None:
    dummy = ListNode()
    cur = dummy

    for value in values:
        cur.next = ListNode(value)
        cur = cur.next

    return dummy.next


# -----------------------------------------------------------------------------
# dataclass: concise local data containers
# -----------------------------------------------------------------------------


@dataclass
class Interval:
    # @dataclass reads these annotations and generates __init__, __repr__, and
    # value-based __eq__ by default, so Interval(1, 3) can be created directly.
    start: int
    end: int


@dataclass
class TrieNode:
    # Use default_factory for mutable defaults. Do not write children: dict = {}.
    children: dict[str, TrieNode] = field(default_factory=dict)
    is_word: bool = False


class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        node = self.root
        for ch in word:
            node = node.children.setdefault(ch, TrieNode())
        node.is_word = True

    def search(self, word: str) -> bool:
        node = self.root
        for ch in word:
            if ch not in node.children:
                return False
            node = node.children[ch]
        return node.is_word


# -----------------------------------------------------------------------------
# Inheritance, overriding, and super()
# -----------------------------------------------------------------------------


class BaseCounter:
    def __init__(self):
        self.total = 0

    def add(self, value: str) -> None:
        raise NotImplementedError("subclasses must implement add")

    def add_many(self, values: list[str]) -> None:
        for value in values:
            self.add(value)


class FrequencyCounter(BaseCounter):
    def __init__(self):
        super().__init__()
        self.counts: dict[str, int] = {}

    def add(self, value: str) -> None:
        self.total += 1
        self.counts[value] = self.counts.get(value, 0) + 1


class LimitedFrequencyCounter(FrequencyCounter):
    def __init__(self, limit: int):
        super().__init__()
        self.limit = limit

    def add(self, value: str) -> None:
        if self.total >= self.limit:
            return
        super().add(value)


def inheritance_example(words: list[str]) -> dict[str, int]:
    counter = LimitedFrequencyCounter(limit=3)
    counter.add_many(words)
    return counter.counts


# -----------------------------------------------------------------------------
# Special methods / dunder methods
# -----------------------------------------------------------------------------


class Bag:
    def __init__(self, values: list[int] | None = None):
        self.values = values or []

    def __len__(self) -> int:
        return len(self.values)

    def __contains__(self, value: int) -> bool:
        return value in self.values

    def __iter__(self) -> Iterator[int]:
        return iter(self.values)

    def __repr__(self) -> str:
        return f"Bag({self.values!r})"


def special_method_example(bag: Bag) -> tuple[int, bool, list[int], str]:
    length = len(bag)  # Calls bag.__len__().
    has_one = 1 in bag  # Calls bag.__contains__(1).
    copied = [x for x in bag]  # Calls bag.__iter__().
    debug = repr(bag)  # Calls bag.__repr__().
    return length, has_one, copied, debug


# -----------------------------------------------------------------------------
# Instance variables vs class variables
# -----------------------------------------------------------------------------


class BadSharedState:
    values: list[int] = []  # Class variable shared by every instance.

    def add(self, value: int) -> None:
        self.values.append(value)


class GoodInstanceState:
    def __init__(self):
        self.values: list[int] = []

    def add(self, value: int) -> None:
        self.values.append(value)


# -----------------------------------------------------------------------------
# Identity, equality, and hashability
# -----------------------------------------------------------------------------


def identity_vs_value(a: ListNode, b: ListNode) -> tuple[bool, bool]:
    same_object = a is b
    same_value = a.val == b.val
    return same_object, same_value


@dataclass(frozen=True)
class Point:
    # frozen=True prevents attribute reassignment and makes this value object
    # hashable when all fields are hashable, so it can be stored in a set.
    row: int
    col: int


def point_set(points: list[Point]) -> set[Point]:
    # frozen=True makes Point immutable and hashable by value.
    return set(points)


# -----------------------------------------------------------------------------
# __slots__: optional memory optimization for many small objects
# -----------------------------------------------------------------------------


class CompactNode:
    __slots__ = ("val", "next")

    def __init__(self, val: int = 0, next: CompactNode | None = None):
        self.val = val
        self.next = next


"""
Interview notes:

- Use the provided LeetCode classes unless the problem asks you to design one.
- `self.x` creates or updates an instance attribute.
- A mutable class attribute is shared across instances; use `__init__` or
  `dataclasses.field(default_factory=...)` for per-instance containers.
- `is` checks object identity; `==` checks equality.
- User-defined objects are hashable by identity by default, but dataclasses with
  generated equality are not hashable unless frozen or configured.
- `from __future__ import annotations` lets type hints refer to the class being
  defined without quoting names in many Python versions.

Concept explanations:

- A decorator is syntax for passing a function or class through another callable
  at definition time. `@dataclass` above a class is roughly:

      Interval = dataclass(Interval)

  The class still exists, but `dataclass` modifies it by adding generated
  methods based on its annotated fields.
- `@dataclass` is useful for small data containers where writing `__init__` by
  hand would add noise. In interviews, it is convenient for intervals, points,
  trie nodes, graph states, heap entries, or records used in tests.
- Dataclass fields come from type annotations such as `start: int`. An
  unannotated class variable is not treated as a dataclass field.
- By default, `@dataclass` generates `__init__`, `__repr__`, and value-based
  `__eq__`. It does not generate ordering unless you pass `order=True`.
- `@dataclass(frozen=True)` makes instances behave like immutable value objects.
  This is the easiest way to make a dataclass hashable for use as a dict key or
  set element, assuming every field is also hashable.
- Mutable defaults are constrained: do not write `children: dict = {}` or
  `items: list = []` in a dataclass. Use `field(default_factory=dict)` or
  `field(default_factory=list)` so every instance gets a fresh container.
- A normal LeetCode `ListNode` / `TreeNode` class is often better than a
  dataclass because the platform already provides it and because identity of
  nodes can matter more than value equality.
- Class variables live on the class and are shared by instances. Instance
  variables live on `self` and are normally created in `__init__`.
- `is` answers "same object?" and is used for `None` checks and sometimes node
  identity. `==` answers "equal value?" and may call custom `__eq__` logic.
- `__slots__` tells Python not to create a per-instance `__dict__`. It can save
  memory for many tiny objects, but it is optional and can make classes less
  flexible because arbitrary new attributes are not allowed.
- Inheritance means a subclass reuses and extends behavior from a parent class.
  `class Child(Parent):` creates that relationship.
- Method overriding means a subclass defines a method with the same name as a
  parent method. Calls on the subclass use the subclass version first.
- `super()` returns a proxy to the next class in Python's method resolution order
  (MRO). In single inheritance, it usually means "call the parent method."
- Multiple inheritance exists in Python, and MRO decides lookup order. For
  coding interviews, prefer simple single inheritance unless a design question
  specifically asks for mixins or framework-style composition.
- `NotImplementedError` is a conventional way for a base class method to say
  "subclasses must implement this."
- Dunder methods such as `__len__`, `__contains__`, and `__iter__` let your
  objects participate in Python syntax like `len(obj)`, `x in obj`, and
  `for x in obj`.
- `__repr__` should return a useful debugging representation. It is what you see
  in many logs, REPL output, and failed assertions.
"""
