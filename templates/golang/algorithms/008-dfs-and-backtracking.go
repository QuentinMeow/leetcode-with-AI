package algorithms

import "sort"

/*
008 - DFS and backtracking patterns

Use when exploring all choices, connected components, recursion over trees,
or any search space where you choose, recurse, then undo.
*/

// Variant 1: recursive DFS over a graph/grid component.
// Example problems: number of islands, max area of island, connected components.
// Time: O(V + E), or O(rows * cols) for grids
// Space: O(V) recursion depth/visited set in the worst case.
func CountIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] != '1' {
			return
		}
		grid[row][col] = '0'
		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row, col-1)
	}

	islands := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				islands++
				dfs(row, col)
			}
		}
	}
	return islands
}

// Variant 2: choose -> recurse -> undo.
// Example problems: subsets, combinations, permutations, combination sum.
// Time: O(n * 2^n)
// Space: O(n) recursion path, excluding output.
func Subsets(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0, len(nums))
	var backtrack func(int)
	backtrack = func(start int) {
		result = append(result, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return result
}

// Variant 3: permutations with a used array.
// Example problems: permutations, generate arrangements, assignment search.
// Time: O(n * n!)
// Space: O(n) recursion path, excluding output.
func Permutations(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums))
	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			result = append(result, append([]int{}, path...))
			return
		}
		for i, x := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, x)
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack()
	return result
}

// Variant 4: backtracking with pruning.
// Example problems: combination sum, word search, N-Queens.
// Time: problem-specific exponential, reduced by pruning
// Space: O(depth)
func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)
	path := make([]int, 0)
	var backtrack func(int, int)
	backtrack = func(start, remaining int) {
		if remaining == 0 {
			result = append(result, append([]int{}, path...))
			return
		}
		for i := start; i < len(candidates); i++ {
			x := candidates[i]
			if x > remaining {
				break
			}
			path = append(path, x)
			backtrack(i, remaining-x)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, target)
	return result
}
