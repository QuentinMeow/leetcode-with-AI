// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"slices"
	"sort"
)

// ===================================================================
// 15. Core Algorithm Templates
// ===================================================================

// Learn the trigger for each family, then copy the smallest template that fits.
// Detailed variants live in Appendix D.
// Fixed-length sliding window: maximumSumFixedLengthWindow in D.3.
// Fast/slow pointers: linkedListHasCycleUsingFloydAlgorithm in D.2.
// Memoized dynamic programming: countGridPathsUsingMemoization in D.11.

// -------------------------------------------------------------------
// 15.1 Hash Map and Counting (companion: algorithms/001-hash-map-and-counting.go)
// -------------------------------------------------------------------

// twoSumIndicesUsingMap returns indices of two values whose sum is target, or nil when
// no pair exists. The map stores each previously seen value's index, so target-current
// can be checked before current is inserted. Time O(n); extra space O(n).
func twoSumIndicesUsingMap(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for index, value := range nums {
		if otherIndex, ok := seen[target-value]; ok {
			return []int{otherIndex, index}
		}
		seen[value] = index
	}
	return nil
}

// -------------------------------------------------------------------
// 15.2 Two Pointers (companion: algorithms/002-two-pointers.go)
// -------------------------------------------------------------------

// sortedTwoSumIndicesUsingTwoPointers requires ascending input. If the current sum is
// too small only the left pointer can increase it; if too large only the right pointer
// can decrease it. It returns two indices or nil. Time O(n); space O(1).
func sortedTwoSumIndicesUsingTwoPointers(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		total := nums[left] + nums[right]
		switch {
		case total == target:
			return []int{left, right}
		case total < target:
			left++
		default:
			right--
		}
	}
	return nil
}

// -------------------------------------------------------------------
// 15.3 Sliding Window (companion: algorithms/003-sliding-window.go)
// -------------------------------------------------------------------

// countSubarraysWithAtMostKDistinctValues counts contiguous subarrays containing no
// more than k distinct numbers. The window [left,right] is shrunk until valid; every
// suffix ending at right and starting from left through right is then valid,
// contributing right-left+1. Time O(n); extra space O(k).
func countSubarraysWithAtMostKDistinctValues(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	count := make(map[int]int)
	left, total := 0, 0
	for right, value := range nums {
		count[value]++
		for len(count) > k {
			leftValue := nums[left]
			count[leftValue]--
			if count[leftValue] == 0 {
				delete(count, leftValue)
			}
			left++
		}
		total += right - left + 1
	}
	return total
}

// -------------------------------------------------------------------
// 15.4 Binary Search on Answer (companion: algorithms/004-binary-search.go)
// -------------------------------------------------------------------

// smallestFeasibleValueUsingBinarySearch requires a monotonic predicate: once
// can(value) becomes true, it must stay true for all larger values. It returns the
// first true value in the inclusive range [low,high]. Time O(log(high-low)) predicate
// calls; space O(1).
func smallestFeasibleValueUsingBinarySearch(low, high int, can func(int) bool) int {
	for low < high {
		middle := low + (high-low)/2
		if can(middle) {
			high = middle
		} else {
			low = middle + 1
		}
	}
	return low
}

// -------------------------------------------------------------------
// 15.5 Prefix Sum (companion: algorithms/005-prefix-sum.go)
// -------------------------------------------------------------------

// countSubarraysWithTargetSumUsingPrefixSums counts contiguous subarrays whose sum
// equals target, including arrays with negative values. If
// currentPrefix-previousPrefix=target, a prior prefix of currentPrefix-target starts a
// valid subarray. seen[0]=1 represents the empty prefix. Time O(n); space O(n).
func countSubarraysWithTargetSumUsingPrefixSums(nums []int, target int) int {
	seen := map[int]int{0: 1}
	prefix, answer := 0, 0
	for _, value := range nums {
		prefix += value
		answer += seen[prefix-target]
		seen[prefix]++
	}
	return answer
}

// -------------------------------------------------------------------
// 15.6 Breadth-First Search (companion: algorithms/007-bfs.go)
// -------------------------------------------------------------------

// shortestPathLengthUnweightedBreadthFirstSearch returns the fewest edges from start to
// target, or -1 if unreachable. Breadth-first search explores vertices by
// increasing distance, and marking on enqueue prevents duplicate work. For V
// vertices and E edges, time O(V+E); space O(V).
func shortestPathLengthUnweightedBreadthFirstSearch(graph map[int][]int, start, target int) int {
	type state struct {
		node     int
		distance int
	}

	queue := []state{{node: start}}
	head := 0
	seen := map[int]struct{}{start: {}}

	for head < len(queue) {
		current := queue[head]
		head++
		if current.node == target {
			return current.distance
		}
		for _, neighbor := range graph[current.node] {
			if _, ok := seen[neighbor]; ok {
				continue
			}
			seen[neighbor] = struct{}{} // Mark on enqueue.
			queue = append(queue, state{
				node:     neighbor,
				distance: current.distance + 1,
			})
		}
	}
	return -1
}

var fourWayGridDirections = [][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

// -------------------------------------------------------------------
// 15.7 Depth-First Search (companion: algorithms/008-dfs-and-backtracking.go)
// -------------------------------------------------------------------

// iterativeDepthFirstVisitOrder returns node IDs in depth-first visit order from
// start; it does not compute shortest paths or disconnected components. Marking
// when pushed prevents duplicates. Reversed insertion preserves neighbor order.
func iterativeDepthFirstVisitOrder(graph map[int][]int, start int) []int {
	visited := map[int]bool{start: true}
	stack := []int{start}
	order := make([]int, 0, len(graph))
	for len(stack) > 0 {
		last := len(stack) - 1
		node := stack[last]
		stack = stack[:last]
		order = append(order, node)
		neighbors := graph[node]
		for index := len(neighbors) - 1; index >= 0; index-- {
			neighbor := neighbors[index]
			if !visited[neighbor] {
				visited[neighbor] = true
				stack = append(stack, neighbor)
			}
		}
	}
	return order
}

// countIslandsUsingDepthFirstSearch counts four-directionally connected groups of ASCII
// byte '1'. Each depth-first search changes visited land to '0', so this function
// mutates grid. Time O(rows*cols); recursion space can reach O(rows*cols).
func countIslandsUsingDepthFirstSearch(grid [][]byte) int {
	// Interview grids are normally rectangular.
	if len(grid) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])

	var visitConnectedLand func(int, int)
	visitConnectedLand = func(row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols {
			return
		}
		if grid[row][col] != '1' {
			return
		}
		grid[row][col] = '0'
		for _, direction := range fourWayGridDirections {
			visitConnectedLand(row+direction[0], col+direction[1])
		}
	}

	count := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				count++
				visitConnectedLand(row, col)
			}
		}
	}
	return count
}

// -------------------------------------------------------------------
// 15.8 Backtracking (companion: algorithms/008-dfs-and-backtracking.go)
// -------------------------------------------------------------------

// allSubsetsUsingBacktracking returns the power set, including the empty subset. The
// path holds the current choices; each recursive call chooses a later index, and
// slices.Clone saves a snapshot before path is changed again. Duplicate input values
// produce duplicate subsets. Time/output O(n*2^n).
// Requires: import "slices"
func allSubsetsUsingBacktracking(nums []int) [][]int {
	answer := make([][]int, 0, 1<<len(nums))
	path := make([]int, 0, len(nums))

	var backtrack func(int)
	backtrack = func(start int) {
		answer = append(answer, slices.Clone(path))
		for index := start; index < len(nums); index++ {
			path = append(path, nums[index])
			backtrack(index + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return answer
}

// -------------------------------------------------------------------
// 15.9 Monotonic Stack (companion: algorithms/009-stack-and-monotonic-stack.go)
// -------------------------------------------------------------------

// nextGreaterElementIndexUsingMonotonicStack returns, for every position, the index of
// the first strictly larger value to its right, or -1. The stack stores unresolved
// indices whose values are non-increasing. Each index is pushed and popped once: time
// O(n), space O(n).
func nextGreaterElementIndexUsingMonotonicStack(nums []int) []int {
	answer := make([]int, len(nums))
	for index := range answer {
		answer[index] = -1
	}

	stack := make([]int, 0) // Indices; values decrease.
	for index, value := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < value {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[top] = index
		}
		stack = append(stack, index)
	}
	return answer
}

// -------------------------------------------------------------------
// 15.10 Intervals (companion: algorithms/011-intervals.go)
// -------------------------------------------------------------------

// mergeIntervals sorts intervals by start, then combines each overlap with the last
// output interval. This version treats touching closed intervals as overlapping because
// next.start <= previous.end. Time O(n log n); output space O(n).
// Requires: import "sort"
func mergeIntervals(intervals []Interval) []Interval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	merged := make([]Interval, 0, len(intervals))
	for _, current := range intervals {
		if len(merged) == 0 ||
			current.start > merged[len(merged)-1].end {
			merged = append(merged, current)
			continue
		}
		last := &merged[len(merged)-1]
		last.end = max(last.end, current.end)
	}
	return merged
}

// -------------------------------------------------------------------
// 15.11 Union-Find (companion: algorithms/013-union-find.go)
// -------------------------------------------------------------------

// DisjointSetUnion, often abbreviated DSU and also called union-find, maintains
// a partition of items into non-overlapping groups. parent identifies each
// group's representative, size guides balanced merging, and components tracks
// how many groups remain.
type DisjointSetUnion struct {
	parent     []int
	size       []int
	components int
}

// newDisjointSetUnion creates n separate groups. A disjoint-set union structure tracks
// which items belong to the same connected component without storing every member
// together.
func newDisjointSetUnion(n int) *DisjointSetUnion {
	parent := make([]int, n)
	size := make([]int, n)
	for index := 0; index < n; index++ {
		parent[index] = index
		size[index] = 1
	}
	return &DisjointSetUnion{parent: parent, size: size, components: n}
}

// findRepresentative returns the root identifying x's group. Path compression rewires
// visited nodes directly to the root, making later operations nearly constant time.
func (d *DisjointSetUnion) findRepresentative(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.findRepresentative(d.parent[x])
	}
	return d.parent[x]
}

// unionSets merges the groups containing a and b and returns true. It returns false
// when they were already connected. Attaching the smaller tree below the larger tree
// limits height; amortized operations are nearly O(1).
func (d *DisjointSetUnion) unionSets(a, b int) bool {
	rootA, rootB := d.findRepresentative(a), d.findRepresentative(b)
	if rootA == rootB {
		return false
	}
	if d.size[rootA] < d.size[rootB] {
		rootA, rootB = rootB, rootA
	}
	d.parent[rootB] = rootA
	d.size[rootA] += d.size[rootB]
	d.components--
	return true
}
