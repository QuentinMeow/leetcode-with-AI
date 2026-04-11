# Complexity Analysis Guide

Quick reference for analyzing time and space complexity of common patterns.

## Common Time Complexities (fastest to slowest)

| Complexity | Name | Typical Pattern | Example |
|------------|------|-----------------|---------|
| O(1) | Constant | Hash lookup, math formula | Array access, hash map get |
| O(log n) | Logarithmic | Binary search, balanced BST | Binary search, heap operations |
| O(n) | Linear | Single pass, two pointers | Linear scan, sliding window |
| O(n log n) | Linearithmic | Efficient sorting, divide & conquer | Merge sort, heap sort |
| O(n^2) | Quadratic | Nested loops | Brute-force pairs, 2D DP |
| O(n^3) | Cubic | Triple nested loops | Floyd-Warshall, interval DP |
| O(2^n) | Exponential | Subsets, backtracking without pruning | Generate all subsets |
| O(n!) | Factorial | Permutations | Generate all permutations |

## How to Determine Complexity

### Time

1. Count the loops: nested loops multiply, sequential loops add.
2. Recursive calls: draw the recursion tree, count total nodes. Use the Master Theorem for divide-and-conquer.
3. Built-in operations matter: sorting is O(n log n), `in` on a list is O(n) but O(1) on a set/dict.

### Space

1. Count auxiliary data structures (hash maps, arrays, stacks).
2. Recursion depth = stack space. DFS on a tree uses O(h) where h = height.
3. "In-place" means O(1) extra space (excluding input and output).

## Pattern-to-Complexity Mapping

| Pattern | Typical Time | Typical Space |
|---------|-------------|---------------|
| Two pointers (sorted) | O(n) | O(1) |
| Sliding window | O(n) | O(k) where k = window content |
| Binary search | O(log n) | O(1) |
| Hash map single pass | O(n) | O(n) |
| BFS/DFS on graph | O(V + E) | O(V) |
| Topological sort | O(V + E) | O(V) |
| Dynamic programming | O(states * transition cost) | O(states) |
| Backtracking | O(branches^depth) | O(depth) |
| Sorting + greedy | O(n log n) | O(1) to O(n) |
| Heap operations | O(n log k) | O(k) |
| Trie operations | O(L) per word, L = length | O(total characters) |

## Recognizing Complexity from Constraints

LeetCode constraints hint at the expected complexity:

| n range | Expected complexity | Reasoning |
|---------|-------------------|-----------|
| n <= 10 | O(n!) or O(2^n) | Brute force / backtracking feasible |
| n <= 20 | O(2^n) | Bitmask DP |
| n <= 100 | O(n^3) | Triple loop or interval DP |
| n <= 1,000 | O(n^2) | Quadratic DP or nested loops |
| n <= 10,000 | O(n^2) borderline | Optimize if possible |
| n <= 100,000 | O(n log n) | Sort + binary search, or divide & conquer |
| n <= 1,000,000 | O(n) | Linear scan, two pointers, sliding window |
| n <= 10^9 | O(log n) or O(1) | Binary search or math |

## Common Pitfalls

- **String concatenation in a loop**: In Python, `s += char` in a loop is O(n^2). Use `''.join(list)` instead.
- **List `in` operator**: `x in list` is O(n). Use a set for O(1) lookup.
- **Sorting stability**: Python's `sort()` is stable (TimSort). Go's `sort.Slice` is not guaranteed stable — use `sort.SliceStable` when needed.
- **Recursion depth**: Python default recursion limit is 1000. For deep recursion, use `sys.setrecursionlimit()` or convert to iterative.
- **Integer overflow**: Not an issue in Python (arbitrary precision), but matters in Go/Java/C++.
