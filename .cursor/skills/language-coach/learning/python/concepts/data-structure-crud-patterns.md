---
concept_id: py-data-structure-crud-patterns
title: Python Data Structure CRUD Patterns
lang: python
depth: standard
last_session: 2026-06-25
related_slugs:
  - python-specific-interview-patterns
  - common-libraries
---

# Python Data Structure CRUD Patterns

## At a glance

- Python containers use repeated CRUD naming patterns: add one, add many, read, remove-and-return, remove-by-value, and clear.
- `pop*` methods return the removed element; `remove`, `discard`, `del`, `clear`, `sort`, and `reverse` mutate without returning a useful value.
- `list`, `set`, `dict`, `heapq`, `deque`, and `queue.Queue` share concepts but use different verbs.
- The cheat sheet prefers module-qualified calls such as `collections.Counter(...)`, `heapq.heappush(...)`, and `bisect.bisect_left(...)` so examples teach where helpers come from.
- The compact reference now lives in `templates/python/cheatsheet.py` section 2.

## What you already know

- You want Python language notes organized as memory-first interview references, not long prose tutorials.
- You prefer side-by-side naming logic and section-by-section examples for common data structures.
- You prefer `import module` plus qualified calls such as `collections.deque(...)`, `heapq.heappush(...)`, and `functools.cache(...)` so the library source stays visible in examples.

## Gaps / not yet covered

- Practice choosing the right container from problem constraints: ordering, uniqueness, priority, FIFO/LIFO, and concurrency.
- Add advanced heap lazy-deletion examples only when a problem needs them.
- Compare asymptotic costs for middle operations, especially `list.pop(0)`, `deque` indexing, and heap rebuilds.

## Detailed notes

### 2026-06-25

Reorganized `templates/python/cheatsheet.py` section 2 into a "Python Data Structures - CRUD Cheat Sheet" modeled after a visual table style:

- Big naming logic table for add/read/update/delete verbs.
- Master memory table comparing `list`, `set`, `dict`, `heapq`, and `deque`.
- Focused CRUD functions for `list`, `set`, `dict`, `heapq`, `deque`, and thread-safe `queue.Queue` variants.
- Inline reminders for common pitfalls: `{}` is a dict, `discard` is soft while `remove` is strict, `dict.get` avoids missing-key errors, `heapq` is min-heap only, `deque.popleft()` is the algorithmic queue primitive, and `queue.Queue` is mainly for concurrency.

### 2026-06-25

Updated `templates/python/cheatsheet.py` to prefer module-qualified standard-library usage:

- Use `import module` instead of `from module import helper`.
- Write calls as `collections.Counter(...)`, `collections.deque(...)`, `heapq.heappush(...)`, `bisect.bisect_left(...)`, `functools.cache(...)`, `operator.itemgetter(...)`, `dataclasses.dataclass`, and `collections.abc.Iterable`.
- The purpose is pedagogical: repeated module prefixes make it easier to remember which helpers come from which library.

## Syntax and equivalents

- Empty set: `set()`, not `{}`.
- Soft dict read: `d.get(key, default)`; strict read: `d[key]`.
- Soft set delete: `s.discard(x)`; strict set delete: `s.remove(x)`.
- Queue for interviews: `collections.deque`; thread-safe queue for concurrency: `queue.Queue`.
- Import style: prefer `import module` and qualified calls like `heapq.heappush(...)`, `collections.Counter(...)`, and `bisect.bisect_left(...)` in the cheat sheet.
- Max-heap with `heapq`: push negated numeric priorities or use tuple priorities when appropriate.

## Interview angles

- Explain the difference between "remove and return" and "remove without returning".
- Choose `deque` over `list.pop(0)` for FIFO queues.
- Explain why arbitrary heap deletion is not efficient without rebuilding or lazy deletion.
- Use `dict`/`set` membership for average O(1) lookups, while remembering hashability constraints.

## Related concepts

- [python-specific-interview-patterns](python-specific-interview-patterns.md) - broader Python interview syntax and idioms.
- [common-libraries](common-libraries.md) - standard-library modules used by the CRUD patterns.

## Addenda

