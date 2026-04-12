package main

// LeetCode 1116 — Print Zero Even Odd
// https://leetcode.com/problems/print-zero-even-odd/

type ZeroEvenOdd struct {
	n        int
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	zeo := &ZeroEvenOdd{
		n:        n
	}
	return zeo
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	panic("TODO")
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	panic("TODO")
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	panic("TODO")
}

// Local compile hook (LeetCode runs your func without this).
func main() {}
