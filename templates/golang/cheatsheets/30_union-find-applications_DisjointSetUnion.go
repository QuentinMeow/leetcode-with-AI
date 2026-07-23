// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

// ===================================================================
// 30. Union-Find Applications
// ===================================================================

// countConnectedComponentsUsingDisjointSetUnion starts with n separate vertices and
// unions every undirected edge. The structure's remaining group count is the number of
// connected components. Nearly O(n+edges) amortized.
func countConnectedComponentsUsingDisjointSetUnion(n int, edges [][2]int) int {
	dsu := newDisjointSetUnion(n)
	for _, edge := range edges {
		dsu.unionSets(edge[0], edge[1])
	}
	return dsu.components
}

// findRedundantEdge returns the first undirected edge whose endpoints are already
// connected. Adding that edge would create a cycle. It returns a zero pair when no
// redundant edge exists.
func findRedundantEdge(edges [][2]int) [2]int {
	maxNode := 0
	for _, edge := range edges {
		maxNode = max(maxNode, edge[0], edge[1])
	}
	dsu := newDisjointSetUnion(maxNode + 1)
	for _, edge := range edges {
		if !dsu.unionSets(edge[0], edge[1]) {
			return edge
		}
	}
	return [2]int{}
}
