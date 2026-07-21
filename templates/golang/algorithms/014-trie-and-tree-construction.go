package algorithms

/*
014 - Trie and structure-construction patterns

Use when a problem needs prefix lookup, word dictionaries, or a setup step that
turns raw input into the tree/graph shape the real algorithm expects.
*/

type TrieNode struct {
	children map[byte]*TrieNode
	isWord   bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{children: make(map[byte]*TrieNode)}
}

// Variant 1: trie with map-backed children.
// Example problems: implement trie, search suggestions, replace words.
// Time: O(L) per insert/search, where L is the word/prefix length.
// Space: O(total characters inserted)
type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: newTrieNode()}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if node.children[ch] == nil {
			node.children[ch] = newTrieNode()
		}
		node = node.children[ch]
	}
	node.isWord = true
}

func (t *Trie) walk(prefix string) *TrieNode {
	node := t.root
	for i := 0; i < len(prefix); i++ {
		node = node.children[prefix[i]]
		if node == nil {
			return nil
		}
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.walk(word)
	return node != nil && node.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.walk(prefix) != nil
}

// Variant 2: trie with wildcard search.
// Example problems: design add and search words data structure.
// Time: O(L) without wildcards; worst-case O(number of trie nodes) with many "."
// Space: O(total characters inserted)
type WordDictionary struct {
	root *TrieNode
}

func NewWordDictionary() *WordDictionary {
	return &WordDictionary{root: newTrieNode()}
}

func (d *WordDictionary) AddWord(word string) {
	node := d.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if node.children[ch] == nil {
			node.children[ch] = newTrieNode()
		}
		node = node.children[ch]
	}
	node.isWord = true
}

func (d *WordDictionary) Search(pattern string) bool {
	var dfs func(int, *TrieNode) bool
	dfs = func(index int, node *TrieNode) bool {
		if index == len(pattern) {
			return node.isWord
		}
		ch := pattern[index]
		if ch == '.' {
			for _, child := range node.children {
				if dfs(index+1, child) {
					return true
				}
			}
			return false
		}
		child := node.children[ch]
		return child != nil && dfs(index+1, child)
	}
	return dfs(0, d.root)
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Variant 3: build a binary tree from LeetCode-style level-order values.
// A nil pointer represents a null entry.
// Example problems: local testing helpers, deserialize binary tree inputs.
// Time: O(n)
// Space: O(n)
func BuildBinaryTreeFromLevelOrder(values []*int) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}
	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	index := 1
	for head := 0; head < len(queue) && index < len(values); head++ {
		node := queue[head]
		if values[index] != nil {
			node.Left = &TreeNode{Val: *values[index]}
			queue = append(queue, node.Left)
		}
		index++

		if index < len(values) && values[index] != nil {
			node.Right = &TreeNode{Val: *values[index]}
			queue = append(queue, node.Right)
		}
		index++
	}
	return root
}

// Variant 4: rebuild binary tree from preorder and inorder traversals.
// Example problems: construct binary tree from preorder/inorder.
// Time: O(n)
// Space: O(n)
func BuildTreeFromPreorderInorder(preorder, inorder []int) *TreeNode {
	inorderIndex := make(map[int]int, len(inorder))
	for i, value := range inorder {
		inorderIndex[value] = i
	}
	preorderIndex := 0
	var build func(int, int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		rootValue := preorder[preorderIndex]
		preorderIndex++
		root := &TreeNode{Val: rootValue}
		split := inorderIndex[rootValue]
		root.Left = build(left, split-1)
		root.Right = build(split+1, right)
		return root
	}
	return build(0, len(inorder)-1)
}

// Variant 5: orient undirected edges into a rooted tree.
// Example problems: tree DP, subtree sizes, collect coins in a tree.
// Time: O(n + e)
// Space: O(n + e)
func BuildRootedTree(n int, edges [][2]int, root int) [][]int {
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
	stack := []state{{root, -1}}
	for len(stack) > 0 {
		last := len(stack) - 1
		current := stack[last]
		stack = stack[:last]
		for _, neighbor := range graph[current.node] {
			if neighbor == current.parent {
				continue
			}
			children[current.node] = append(children[current.node], neighbor)
			stack = append(stack, state{neighbor, current.node})
		}
	}
	return children
}
