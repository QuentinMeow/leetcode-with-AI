package main

// LeetCode 1367 — Linked List in Binary Tree
// https://leetcode.com/problems/linked-list-in-binary-tree/


// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}


// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	panic("TODO")
}

// Local compile hook (LeetCode runs your func without this).
func main() {}
