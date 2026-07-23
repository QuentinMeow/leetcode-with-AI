// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

// ===================================================================
// 5. Structs and Common Interview Node Types
// ===================================================================

// Struct values copy on assignment. Use pointers for linked structures and
// mutation; use comparable value structs such as Point as map keys.

// Interval stores two endpoints. Each interval algorithm below states whether
// it treats end as inclusive or exclusive; do not mix those conventions.
type Interval struct {
	start int
	end   int
}

// ListNode matches the singly linked-list shape used by LeetCode prompts.
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode matches the binary-tree shape used by LeetCode prompts.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Point struct {
	Row int
	Col int
}

// newListNode is a small constructor for the pointer shape used by LeetCode.
func newListNode(value int) *ListNode {
	return &ListNode{Val: value}
}

// anonymousPairExample shows a one-off record type. Prefer a named type when
// methods, reuse, or package-level documentation would make the meaning clearer.
func anonymousPairExample(value, priority int) struct {
	value    int
	priority int
} {
	return struct {
		value    int
		priority int
	}{value: value, priority: priority}
}
