// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"container/heap"
)

// ===================================================================
// 24. Breadth-First Search, Topological Sort, and Shortest Paths
// ===================================================================

// binaryTreeValuesByLevelBreadthFirst returns one value slice per tree depth. The queue
// segment present at the start of an outer iteration is exactly one level. Time
// O(nodes); space O(maximum tree width).
func binaryTreeValuesByLevelBreadthFirst(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	answer := make([][]int, 0)
	queue := []*TreeNode{root}
	head := 0

	for head < len(queue) {
		levelSize := len(queue) - head
		level := make([]int, 0, levelSize)
		for levelIndex := 0; levelIndex < levelSize; levelIndex++ {
			node := queue[head]
			head++
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		answer = append(answer, level)
	}
	return answer
}

// topologicalOrderUsingKahnAlgorithm orders a directed graph where each edge is
// [prerequisite,course]. Kahn's algorithm repeatedly removes zero-indegree
// nodes; if fewer than n are removed, a directed cycle exists and nil is
// returned. For V vertices and E edges, time O(V+E).
func topologicalOrderUsingKahnAlgorithm(n int, edges [][2]int) []int {
	graph := make([][]int, n)
	indegree := make([]int, n)
	for _, edge := range edges {
		prerequisite, course := edge[0], edge[1]
		graph[prerequisite] = append(graph[prerequisite], course)
		indegree[course]++
	}

	queue := make([]int, 0, n)
	for node, degree := range indegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	order := make([]int, 0, n)
	for head := 0; head < len(queue); head++ {
		node := queue[head]
		order = append(order, node)
		for _, neighbor := range graph[node] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	if len(order) != n {
		return nil
	}
	return order
}

type WeightedEdge struct {
	to     int
	weight int
}

type DistanceState struct {
	distance int
	node     int
}

// DistanceMinHeap orders graph states by their currently known distance.
type DistanceMinHeap []DistanceState

// Len reports the number of elements required by heap.Interface.
func (h DistanceMinHeap) Len() int { return len(h) }

// Less defines which of two elements has higher heap priority.
func (h DistanceMinHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

// Swap exchanges two heap positions while container/heap restores its invariant.
func (h DistanceMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push appends a concrete value received through heap.Interface's any parameter.
func (h *DistanceMinHeap) Push(value any) {
	*h = append(*h, value.(DistanceState))
}

// Pop removes the final slice item after container/heap has moved the root there.
func (h *DistanceMinHeap) Pop() any {
	old := *h
	value := old[len(old)-1]
	*h = old[:len(old)-1]
	return value
}

// shortestDistancesUsingDijkstraAlgorithm returns minimum distances from start in a
// graph with non-negative edge weights. The min-heap chooses the unsettled
// shortest candidate; stale entries are skipped instead of decreased in place.
// For V vertices and E edges, time O((V+E) log V).
// Requires: import "container/heap"
func shortestDistancesUsingDijkstraAlgorithm(graph map[int][]WeightedEdge, start int) map[int]int {
	distance := map[int]int{start: 0}
	priorityQueue := &DistanceMinHeap{{distance: 0, node: start}}
	heap.Init(priorityQueue)

	for priorityQueue.Len() > 0 {
		current := heap.Pop(priorityQueue).(DistanceState)
		if current.distance != distance[current.node] {
			continue // Stale heap entry.
		}
		for _, edge := range graph[current.node] {
			nextDistance := current.distance + edge.weight
			oldDistance, seen := distance[edge.to]
			if !seen || nextDistance < oldDistance {
				distance[edge.to] = nextDistance
				heap.Push(priorityQueue, DistanceState{
					distance: nextDistance,
					node:     edge.to,
				})
			}
		}
	}
	return distance
}

// gridShortestDistancesUsingBreadthFirstSearch treats 0 cells as open and non-zero
// cells as blocked. It returns shortest four-way distances from start, with -1 for
// unreachable cells. Time/space O(rows*cols).
func gridShortestDistancesUsingBreadthFirstSearch(grid [][]int, start Point) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return nil
	}

	rows, cols := len(grid), len(grid[0])
	distance := make([][]int, rows)
	for row := range distance {
		distance[row] = make([]int, cols)
		for col := range distance[row] {
			distance[row][col] = -1
		}
	}
	if start.Row < 0 || start.Row >= rows ||
		start.Col < 0 || start.Col >= cols ||
		grid[start.Row][start.Col] != 0 {
		return distance
	}

	queue := []Point{start}
	distance[start.Row][start.Col] = 0
	for head := 0; head < len(queue); head++ {
		current := queue[head]
		for _, direction := range fourWayGridDirections {
			next := Point{
				Row: current.Row + direction[0],
				Col: current.Col + direction[1],
			}
			if next.Row < 0 || next.Row >= rows ||
				next.Col < 0 || next.Col >= cols ||
				grid[next.Row][next.Col] != 0 ||
				distance[next.Row][next.Col] != -1 {
				continue
			}
			distance[next.Row][next.Col] =
				distance[current.Row][current.Col] + 1
			queue = append(queue, next) // Marked before enqueue, so no duplicates.
		}
	}
	return distance
}

// gridDistancesFromMultipleSources enqueues every valid source at distance zero before
// breadth-first search. The first visit to a cell is therefore its distance to the
// nearest source. Blocked cells are non-zero; unreachable cells stay -1.
func gridDistancesFromMultipleSources(grid [][]int, sources []Point) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return nil
	}
	rows, cols := len(grid), len(grid[0])
	distance := make([][]int, rows)
	for row := range distance {
		distance[row] = make([]int, cols)
		for col := range distance[row] {
			distance[row][col] = -1
		}
	}

	queue := make([]Point, 0, len(sources))
	for _, source := range sources {
		if source.Row < 0 || source.Row >= rows ||
			source.Col < 0 || source.Col >= cols ||
			grid[source.Row][source.Col] != 0 ||
			distance[source.Row][source.Col] != -1 {
			continue
		}
		distance[source.Row][source.Col] = 0
		queue = append(queue, source)
	}

	for head := 0; head < len(queue); head++ {
		current := queue[head]
		for _, direction := range fourWayGridDirections {
			next := Point{
				Row: current.Row + direction[0],
				Col: current.Col + direction[1],
			}
			if next.Row < 0 || next.Row >= rows ||
				next.Col < 0 || next.Col >= cols ||
				grid[next.Row][next.Col] != 0 ||
				distance[next.Row][next.Col] != -1 {
				continue
			}
			distance[next.Row][next.Col] =
				distance[current.Row][current.Col] + 1
			queue = append(queue, next)
		}
	}
	return distance
}

// orientUndirectedTreeFromRoot converts undirected edges into parent-to-children
// adjacency lists relative to root. Carrying the parent avoids walking immediately back
// along an edge. Input must be a valid tree. Time O(n).
func orientUndirectedTreeFromRoot(n int, edges [][2]int, root int) [][]int {
	graph := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	children := make([][]int, n)
	type state struct {
		node   int
		parent int
	}
	stack := []state{{node: root, parent: -1}}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, neighbor := range graph[current.node] {
			if neighbor == current.parent {
				continue
			}
			children[current.node] = append(children[current.node], neighbor)
			stack = append(stack, state{node: neighbor, parent: current.node})
		}
	}
	return children
}
