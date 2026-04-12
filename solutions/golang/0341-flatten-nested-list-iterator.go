package main

// LeetCode 341 — Flatten Nested List Iterator
// https://leetcode.com/problems/flatten-nested-list-iterator/


// // This is the interface that allows for creating nested lists.
// // You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// // Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	panic("TODO")
}

// // Return the single integer that this NestedInteger holds, if it holds a single integer
// // The result is undefined if this NestedInteger holds a nested list
// // So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	panic("TODO")
}

// // Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	panic("TODO")
}

// // Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {
	panic("TODO")
}

// // Return the nested list that this NestedInteger holds, if it holds a nested list
// // The list length is zero if this NestedInteger holds a single integer
// // You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	panic("TODO")
}


type NestedIterator struct {
    
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	panic("TODO")
}

func (this *NestedIterator) Next() int {
	panic("TODO")
}

func (this *NestedIterator) HasNext() bool {
	panic("TODO")
}

// Local compile hook (LeetCode runs your func without this).
func main() {}
