package main

// LeetCode 1114 — Print in Order
// https://leetcode.com/problems/print-in-order/

type Foo struct {

}

func NewFoo() *Foo {
	return &Foo{
	}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
}

func (f *Foo) Second(printSecond func()) {
	/// Do not change this line
	printSecond()
}

func (f *Foo) Third(printThird func()) {
	// Do not change this line
	printThird()
}

// Local compile hook (LeetCode runs your func without this).
func main() {}
