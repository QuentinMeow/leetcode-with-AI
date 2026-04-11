# Algorithm Patterns

A catalog of common LeetCode patterns with trigger signals — use this to identify which pattern applies to a problem.

## Two Pointers

**Trigger signals**: sorted array, find pair with target sum, remove duplicates, partition
**When to use**: Problems involving pairs or comparisons in a sorted (or sortable) sequence
**Key idea**: Maintain two indices moving toward each other (or in the same direction) to avoid nested loops

**Common problems**: Two Sum II, 3Sum, Container With Most Water, Trapping Rain Water, Remove Duplicates

## Sliding Window

**Trigger signals**: contiguous subarray/substring, maximum/minimum length, "at most K distinct"
**When to use**: Finding an optimal contiguous segment of an array or string
**Key idea**: Maintain a window with two boundaries, expand right and contract left

**Variants**:
- Fixed-size window (simple)
- Variable-size window with constraint (track constraint violation to shrink)

**Common problems**: Longest Substring Without Repeating Characters, Minimum Window Substring, Maximum Sum Subarray of Size K

## Binary Search

**Trigger signals**: sorted array, "minimum/maximum that satisfies condition", O(log n) requirement
**When to use**: Searching in sorted data, or when the answer space is monotonic
**Key idea**: Eliminate half the search space each step

**Variants**:
- Classic (find exact value)
- Search on answer (binary search on the result, check feasibility)
- Bisect left/right (find insertion point)

**Common problems**: Search in Rotated Sorted Array, Find Minimum in Rotated Sorted Array, Koko Eating Bananas

## Hash Map / Set

**Trigger signals**: find duplicates, count frequency, find complement, O(1) lookup needed
**When to use**: Problems requiring fast lookup or counting
**Key idea**: Trade space for time — store seen elements for instant access

**Common problems**: Two Sum, Group Anagrams, Longest Consecutive Sequence, Subarray Sum Equals K

## Stack

**Trigger signals**: matching parentheses, next greater/smaller element, monotonic property, nested structure
**When to use**: Problems with LIFO ordering or where you need to track "what came before"
**Key idea**: Process elements and maintain state that can be unwound

**Variants**:
- Monotonic stack (increasing or decreasing)
- Expression evaluation
- Nested structure parsing

**Common problems**: Valid Parentheses, Daily Temperatures, Largest Rectangle in Histogram

## Linked List

**Trigger signals**: reverse, cycle detection, merge, find middle, reorder
**When to use**: In-place operations on linked list nodes
**Key techniques**: Fast/slow pointers, dummy head node, reverse in groups

**Common problems**: Reverse Linked List, Merge Two Sorted Lists, Linked List Cycle, LRU Cache

## Tree (BFS / DFS)

**Trigger signals**: binary tree, level order, path sum, serialize, lowest common ancestor
**When to use**: Traversal, search, or aggregation over tree structures

**Variants**:
- DFS (preorder, inorder, postorder) — path problems, validation
- BFS (level order) — level-based queries, shortest path in unweighted tree

**Common problems**: Maximum Depth, Level Order Traversal, Validate BST, Lowest Common Ancestor

## Graph

**Trigger signals**: nodes and edges, connected components, shortest path, topological order, "course schedule"
**When to use**: Problems modeled as graphs (explicit or implicit)

**Variants**:
- BFS — shortest path (unweighted), level exploration
- DFS — cycle detection, connected components, path finding
- Topological sort — dependency ordering (Kahn's or DFS-based)
- Union-Find — dynamic connectivity
- Dijkstra / Bellman-Ford — weighted shortest path

**Common problems**: Number of Islands, Course Schedule, Clone Graph, Network Delay Time

## Dynamic Programming

**Trigger signals**: "number of ways", "minimum cost", "maximum profit", overlapping subproblems, optimal substructure
**When to use**: Problems with overlapping subproblems where brute force would re-compute the same state

**Approach**:
1. Define the state: what information do you need to make a decision?
2. Define the transition: how do you get from one state to the next?
3. Define the base case: what's the starting point?
4. Determine the order: bottom-up (tabulation) or top-down (memoization)

**Variants**:
- 1D DP (Fibonacci-style, house robber)
- 2D DP (grid paths, string matching)
- Knapsack (0/1, unbounded, bounded)
- Interval DP (burst balloons, matrix chain)
- Bitmask DP (TSP, assignment)
- DP on trees

**Common problems**: Climbing Stairs, Coin Change, Longest Common Subsequence, Edit Distance, House Robber

## Backtracking

**Trigger signals**: "generate all", "find all combinations/permutations/subsets", constraint satisfaction
**When to use**: Exploring all possible configurations with pruning
**Key idea**: Build solution incrementally, abandon ("backtrack") when a constraint is violated

**Common problems**: Subsets, Permutations, Combination Sum, N-Queens, Word Search

## Greedy

**Trigger signals**: interval scheduling, "minimum number of", activity selection, local-to-global optimal
**When to use**: When making the locally optimal choice at each step leads to the global optimum
**Key warning**: Always prove the greedy choice property holds — many problems look greedy but need DP

**Common problems**: Jump Game, Merge Intervals, Task Scheduler, Gas Station

## Heap / Priority Queue

**Trigger signals**: "K-th largest/smallest", "top K", merge K sorted, stream median
**When to use**: Maintaining a dynamic set with efficient min/max extraction
**Key idea**: Use a min-heap or max-heap to efficiently track the extreme elements

**Common problems**: Kth Largest Element, Merge K Sorted Lists, Find Median from Data Stream, Top K Frequent Elements

## Trie

**Trigger signals**: prefix matching, autocomplete, word search with prefix constraints, "starts with"
**When to use**: Problems involving prefix-based operations on strings
**Key idea**: Tree where each node represents a character, shared prefixes share nodes

**Common problems**: Implement Trie, Word Search II, Design Add and Search Words

## Bit Manipulation

**Trigger signals**: XOR, single number, power of two, count bits, "without extra space"
**When to use**: Problems where bitwise operations can replace arithmetic or avoid extra storage
**Key techniques**: XOR (a ^ a = 0), bit masking, Brian Kernighan's algorithm

**Common problems**: Single Number, Number of 1 Bits, Counting Bits, Subsets (bitmask)

## Intervals

**Trigger signals**: merge intervals, insert interval, meeting rooms, overlapping ranges
**When to use**: Problems involving ranges on a number line
**Key technique**: Sort by start (or end), then sweep

**Common problems**: Merge Intervals, Insert Interval, Non-overlapping Intervals, Meeting Rooms II
