---
concept_id: py-interview-algorithm-templates
title: Interview Algorithm Templates
lang: python
depth: standard
last_session: 2026-06-09
related_slugs: [python-specific-interview-patterns]
---

# Interview Algorithm Templates

## At a glance

- Created ranked Python template files under `templates/python/algorithms/`.
- Each file starts with the most common interview variant, then lists additional variants.
- Templates include Python-specific implementation habits such as `defaultdict`, `Counter`, `deque`, `heapq`, dummy linked-list nodes, nested DFS/backtracking helpers, and `@cache`.
- Each variant includes time and space complexity notes.

## What you already know

- You are comfortable with most algorithm ideas and want the templates to emphasize Python-specific syntax and patterns used in coding interviews.

## Gaps / not yet covered

- Add less common advanced patterns such as trie, topological sort, bit manipulation, segment tree, Fenwick tree, Dijkstra, and A* only when needed.
- Expand each algorithm file with problem-specific gotchas after solving representative LeetCode problems.
- Continue linking Python-specific idiom notes to algorithm templates as repeated patterns appear in solved problems.

## Detailed notes

### 2026-06-09

Seeded the algorithm-pattern folder with numbered files ordered roughly from most commonly used coding-interview patterns to less common but still important patterns:

1. Hash map and counting
2. Two pointers
3. Sliding window
4. Binary search
5. Prefix sum
6. Linked list
7. BFS
8. DFS and backtracking
9. Stack and monotonic stack
10. Heap / priority queue
11. Intervals
12. Dynamic programming
13. Union-find

The files are intentionally practical templates rather than tutorials. The goal is to support quick recall during interview preparation while still exposing the Python syntax and standard-library tools that make each pattern concise.

## Syntax and equivalents

- `defaultdict(int)` avoids explicit missing-key initialization for counts.
- `Counter` is a concise frequency-table tool, but a plain `dict` or `defaultdict(int)` is better when updates are tightly controlled inside a loop.
- `deque` gives O(1) queue pops from the left for BFS.
- `heapq` is a min-heap; max-heap behavior usually comes from negating numeric priorities.
- Nested helper functions are common for DFS/backtracking because they can close over `result`, `path`, grid dimensions, or memoized state.

## Interview angles

- Be ready to state why a pattern applies before coding: sorted input, contiguous window, monotonic predicate, repeated range query, shortest unweighted path, overlapping subproblems, or dynamic connectivity.
- Mention complexity for both the dominant loop/search and the auxiliary data structure.
- In Python, explain standard-library choices briefly when they affect asymptotic behavior, such as `deque.popleft()` versus `list.pop(0)`.

## Related concepts

- [python-specific-interview-patterns](python-specific-interview-patterns.md) — companion notes for Python syntax, containers, mutability, version differences, and standard-library habits used by the algorithm templates.

## Addenda
