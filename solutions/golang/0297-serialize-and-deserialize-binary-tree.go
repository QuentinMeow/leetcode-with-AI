package main

// LeetCode 297 — Serialize and Deserialize Binary Tree
// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/


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
// ser := Constructor();
// deser := Constructor();
// data := ser.serialize(root);
// ans := deser.deserialize(data);


// Local compile hook (LeetCode runs your func without this).
func main() {}
