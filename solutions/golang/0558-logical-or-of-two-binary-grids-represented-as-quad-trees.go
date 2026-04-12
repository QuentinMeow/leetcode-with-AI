package main

// LeetCode 558 — Logical OR of Two Binary Grids Represented as Quad-Trees
// https://leetcode.com/problems/logical-or-of-two-binary-grids-represented-as-quad-trees/


// Definition for a QuadTree node.
type Node struct {
    Val bool
    IsLeaf bool
    TopLeft *Node
    TopRight *Node
    BottomLeft *Node
    BottomRight *Node
}


func intersect(quadTree1 *Node, quadTree2 *Node) *Node {
	panic("TODO")
}

// Local compile hook (LeetCode runs your func without this).
func main() {}
