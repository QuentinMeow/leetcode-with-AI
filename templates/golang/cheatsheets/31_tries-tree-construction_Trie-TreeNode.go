// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

// ===================================================================
// 31. Tries and Tree Construction
// ===================================================================

// TrieNode is one node in a trie (prefix tree); each child edge consumes a rune.
type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

// Trie stores words by shared prefixes.
type Trie struct {
	root *TrieNode
}

// newTrie creates a trie, also called a prefix tree. Each edge stores one rune, so
// shared word prefixes share nodes.
func newTrie() *Trie {
	return &Trie{
		root: &TrieNode{children: make(map[rune]*TrieNode)},
	}
}

// insert adds one word to the trie, creating missing rune edges.
func (t *Trie) insert(word string) {
	node := t.root
	for _, char := range word {
		if node.children[char] == nil {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		node = node.children[char]
	}
	node.isWord = true
}

// followPrefixPath follows a complete rune path and returns its terminal node, or nil when the path
// is absent.
func (t *Trie) followPrefixPath(text string) *TrieNode {
	node := t.root
	for _, char := range text {
		node = node.children[char]
		if node == nil {
			return nil
		}
	}
	return node
}

// search returns true only when the complete word was inserted, not merely when it is a
// prefix.
func (t *Trie) search(word string) bool {
	node := t.followPrefixPath(word)
	return node != nil && node.isWord
}

// startsWith returns true when any stored word begins with the supplied prefix.
func (t *Trie) startsWith(prefix string) bool {
	return t.followPrefixPath(prefix) != nil
}

// wildcardSearch returns whether the pattern matches a stored word; '.' matches exactly
// one arbitrary rune. A wildcard branches to every child, so worst-case time can be
// exponential in wildcard count.
func (t *Trie) wildcardSearch(pattern string) bool {
	patternRunes := []rune(pattern)
	var search func(int, *TrieNode) bool
	search = func(index int, node *TrieNode) bool {
		if index == len(patternRunes) {
			return node.isWord
		}
		char := patternRunes[index]
		if char != '.' {
			child := node.children[char]
			return child != nil && search(index+1, child)
		}
		for _, child := range node.children {
			if search(index+1, child) {
				return true
			}
		}
		return false
	}
	return search(0, t.root)
}

// buildBinaryTreeLevelOrder interprets pointer values in breadth-first order; nil
// entries represent missing children.
func buildBinaryTreeLevelOrder(values []*int) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}
	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	index := 1
	for head := 0; head < len(queue) && index < len(values); head++ {
		node := queue[head]
		if index < len(values) && values[index] != nil {
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

// buildTreePreorderInorder reconstructs a binary tree with unique values. Preorder
// supplies each subtree root; its position in inorder separates left and right subtree
// values. The index map gives O(n) time and O(n) space.
func buildTreePreorderInorder(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	inorderIndex := make(map[int]int, len(inorder))
	for index, value := range inorder {
		inorderIndex[value] = index
	}
	preorderIndex := 0
	var build func(int, int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		value := preorder[preorderIndex]
		preorderIndex++
		node := &TreeNode{Val: value}
		split := inorderIndex[value]
		node.Left = build(left, split-1)
		node.Right = build(split+1, right)
		return node
	}
	return build(0, len(inorder)-1)
}
