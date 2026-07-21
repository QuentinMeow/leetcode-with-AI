package algorithms

/*
013 - Union-find / disjoint set patterns

Use when the problem asks about dynamic connectivity, grouping components,
cycle detection in undirected graphs, or merging accounts/equivalence classes.
*/

// Variant 1: path compression + union by size.
// Example problems: number of connected components, redundant connection.
// Amortized time: almost O(1) per find/union, commonly written as O(alpha(n)).
// Space: O(n)
type DSU struct {
	parent     []int
	size       []int
	components int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{parent: parent, size: size, components: n}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(a, b int) bool {
	rootA, rootB := d.Find(a), d.Find(b)
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

func (d *DSU) Components() int {
	return d.components
}

// Variant 2: count connected components from edges.
// Example problems: number of connected components in an undirected graph.
// Time: O((n + e) * alpha(n))
// Space: O(n)
func CountComponents(n int, edges [][2]int) int {
	dsu := NewDSU(n)
	for _, edge := range edges {
		dsu.Union(edge[0], edge[1])
	}
	return dsu.Components()
}

// Variant 3: detect cycle / redundant edge.
// Example problems: redundant connection, validate tree.
// Time: O(e * alpha(n))
// Space: O(n)
func FindRedundantEdge(edges [][2]int) ([2]int, bool) {
	largest := -1
	for _, edge := range edges {
		largest = max(largest, max(edge[0], edge[1]))
	}
	dsu := NewDSU(largest + 1)
	for _, edge := range edges {
		if !dsu.Union(edge[0], edge[1]) {
			return edge, true
		}
	}
	return [2]int{}, false
}

// Variant 4: union by arbitrary keys using ID compression.
// Example problems: accounts merge, sentence similarity, equation satisfiability.
// Time: O((n + e) * alpha(n))
// Space: O(n)
func GroupEquivalentItems(pairs [][2]string) map[int][]string {
	ids := make(map[string]int)
	items := make([]string, 0)
	getID := func(item string) int {
		if id, ok := ids[item]; ok {
			return id
		}
		id := len(ids)
		ids[item] = id
		items = append(items, item)
		return id
	}

	for _, pair := range pairs {
		getID(pair[0])
		getID(pair[1])
	}
	dsu := NewDSU(len(ids))
	for _, pair := range pairs {
		dsu.Union(ids[pair[0]], ids[pair[1]])
	}

	groups := make(map[int][]string)
	for _, item := range items {
		root := dsu.Find(ids[item])
		groups[root] = append(groups[root], item)
	}
	return groups
}
