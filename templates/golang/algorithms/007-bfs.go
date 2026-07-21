package algorithms

/*
007 - Breadth-first search patterns

Use for shortest paths in unweighted graphs, level-order traversal, grid
distance, and problems where "minimum number of steps" is the signal.
*/

// Variant 1: queue BFS on a graph.
// Example problems: shortest path in unweighted graph, word ladder state graph.
// Time: O(V + E)
// Space: O(V)
func ShortestPathLength(graph map[int][]int, start, target int) int {
	type state struct {
		node int
		dist int
	}
	queue := []state{{node: start}}
	seen := map[int]bool{start: true}
	for head := 0; head < len(queue); head++ {
		current := queue[head]
		if current.node == target {
			return current.dist
		}
		for _, neighbor := range graph[current.node] {
			if !seen[neighbor] {
				seen[neighbor] = true
				queue = append(queue, state{neighbor, current.dist + 1})
			}
		}
	}
	return -1
}

type BFSTreeNode struct {
	Val         int
	Left, Right *BFSTreeNode
}

// Variant 2: level-order traversal.
// Example problems: binary tree level order, right side view, zigzag traversal.
// Time: O(n)
// Space: O(w), where w is max tree width.
func LevelOrder(root *BFSTreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)
	queue := []*BFSTreeNode{root}
	for head := 0; head < len(queue); {
		levelEnd := len(queue)
		level := make([]int, 0, levelEnd-head)
		for head < levelEnd {
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
		result = append(result, level)
	}
	return result
}

// Variant 3: grid BFS with four directions.
// Example problems: rotting oranges, walls and gates, shortest path in binary matrix.
// Time: O(rows * cols)
// Space: O(rows * cols)
func ShortestGridPath(grid [][]int, start, target [2]int) int {
	rows, cols := len(grid), len(grid[0])
	type state struct {
		row  int
		col  int
		dist int
	}
	queue := []state{{start[0], start[1], 0}}
	seen := make([][]bool, rows)
	for row := range seen {
		seen[row] = make([]bool, cols)
	}
	seen[start[0]][start[1]] = true
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for head := 0; head < len(queue); head++ {
		current := queue[head]
		if current.row == target[0] && current.col == target[1] {
			return current.dist
		}
		for _, direction := range directions {
			nextRow := current.row + direction[0]
			nextCol := current.col + direction[1]
			if nextRow >= 0 && nextRow < rows &&
				nextCol >= 0 && nextCol < cols &&
				grid[nextRow][nextCol] == 0 &&
				!seen[nextRow][nextCol] {
				seen[nextRow][nextCol] = true
				queue = append(queue, state{nextRow, nextCol, current.dist + 1})
			}
		}
	}
	return -1
}

// Variant 4: multi-source BFS.
// Example problems: rotting oranges, nearest zero, map of highest peak.
// Time: O(rows * cols)
// Space: O(rows * cols)
func DistanceFromSources(grid [][]int, sources [][2]int) [][]int {
	rows, cols := len(grid), len(grid[0])
	dist := make([][]int, rows)
	for row := range dist {
		dist[row] = make([]int, cols)
		for col := range dist[row] {
			dist[row][col] = -1
		}
	}

	queue := make([][2]int, 0, len(sources))
	for _, source := range sources {
		if dist[source[0]][source[1]] == -1 {
			dist[source[0]][source[1]] = 0
			queue = append(queue, source)
		}
	}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for head := 0; head < len(queue); head++ {
		row, col := queue[head][0], queue[head][1]
		for _, direction := range directions {
			nextRow := row + direction[0]
			nextCol := col + direction[1]
			if nextRow >= 0 && nextRow < rows &&
				nextCol >= 0 && nextCol < cols &&
				dist[nextRow][nextCol] == -1 {
				dist[nextRow][nextCol] = dist[row][col] + 1
				queue = append(queue, [2]int{nextRow, nextCol})
			}
		}
	}
	return dist
}
