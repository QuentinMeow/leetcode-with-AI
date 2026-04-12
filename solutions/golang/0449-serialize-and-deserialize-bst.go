package main

// LeetCode 449 — Serialize and Deserialize BST
// https://leetcode.com/problems/serialize-and-deserialize-bst/


// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


type Codec struct {
    
}

func Constructor() Codec {
	panic("TODO")
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	panic("TODO")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	panic("TODO")
}



// Your Codec object will be instantiated and called as such:
// ser := Constructor()
// deser := Constructor()
// tree := ser.serialize(root)
// ans := deser.deserialize(tree)
// return ans


// Local compile hook (LeetCode runs your func without this).
func main() {}
