# Progress

Track which problems have been solved and in which languages.

## Legend

- AC = Accepted (implementation present in repo; not necessarily submitted on LeetCode recently)
- WA = attempted but not yet accepted
- (blank) = not attempted

Python rows below list problems whose `solutions/python/*.py` files contain at least one non-`pass` method body (AST check). Difficulty and titles follow `data/leetcode/python/first-300.json` when the problem appears there. Go column is AC when a matching `solutions/golang/<id>-*.go` stub or solution file exists.

## Problems

| # | Problem | Difficulty | Python | Go | Pattern | Notes |
|---|---------|------------|--------|----|---------|-------|
| 1 | Two Sum | Easy | AC | AC | array, hash-table | |
| 2 | Add Two Numbers | Medium | AC | AC | linked-list, math, recursion | |
| 3 | Longest Substring Without Repeating Characters | Medium | AC | AC | hash-table, string, sliding-window | |
| 4 | Median of Two Sorted Arrays | Hard | AC | AC | array, binary-search, divide-and-conquer | |
| 5 | Longest Palindromic Substring | Medium | AC | AC | two-pointers, string, dynamic-programming | |
| 6 | Zigzag Conversion | Medium | AC | AC | string | |
| 7 | Reverse Integer | Medium | AC | AC | math | |
| 8 | String to Integer (atoi) | Medium | AC | AC | string | |
| 9 | Palindrome Number | Easy | AC | AC | math | |
| 31 | Next Permutation | Medium | AC | AC | array, two-pointers | |
| 37 | Sudoku Solver | Hard | AC | AC | array, hash-table, backtracking, matrix | |
| 48 | Rotate Image | Medium | AC | AC | array, math, matrix | |
| 73 | Set Matrix Zeroes | Medium | AC | AC | array, hash-table, matrix | |
| 75 | Sort Colors | Medium | AC | AC | array, two-pointers, sorting | |
| 88 | Merge Sorted Array | Easy | AC | AC | array, two-pointers, sorting | |
| 99 | Recover Binary Search Tree | Medium | AC | AC | tree, depth-first-search, binary-search-tree, binary-tree | |
| 114 | Flatten Binary Tree to Linked List | Medium | AC | AC | linked-list, stack, tree, depth-first-search, binary-tree | |
| 130 | Surrounded Regions | Medium | AC | AC | array, depth-first-search, breadth-first-search, union-find, matrix | |
| 143 | Reorder List | Medium | AC | AC | linked-list, two-pointers, stack, recursion | |
| 146 | LRU Cache | Medium | AC | AC | hash-table, linked-list, design, doubly-linked-list | |
| 189 | Rotate Array | Medium | AC | AC | array, math, two-pointers | |
| 237 | Delete Node in a Linked List | Medium | AC | AC | linked-list | |
| 283 | Move Zeroes | Easy | AC | AC | array, two-pointers | |
| 284 | Peeking Iterator | Medium | AC | AC | array, design, iterator | |
| 289 | Game of Life | Medium | AC | AC | array, matrix, simulation | |
| 297 | Serialize and Deserialize Binary Tree | Hard | AC | AC | string, tree, depth-first-search, breadth-first-search, design, binary-tree | |

*Regenerate the Python column with an AST pass over `solutions/python/` when you want to refresh; see agent notes in git history.*
