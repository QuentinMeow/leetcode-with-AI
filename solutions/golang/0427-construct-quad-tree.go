package main

// LeetCode 427 — Construct Quad Tree
// https://leetcode.com/problems/construct-quad-tree/


// Definition for a QuadTree node.
type Node struct {
    Val bool
    IsLeaf bool
    TopLeft *Node
    TopRight *Node
    BottomLeft *Node
    BottomRight *Node
}


func construct(grid [][]int) *Node {
	panic("TODO")
}

// Local compile hook (LeetCode runs your func without this).
func main() {}
