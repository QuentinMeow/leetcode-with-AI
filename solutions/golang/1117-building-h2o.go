package main

// LeetCode 1117 — Building H2O
// https://leetcode.com/problems/building-h2o/

type H2O struct {

}

func NewH2O() *H2O {
	h := &H2O{}
	return h
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	// releaseHydrogen() outputs "H". Do not change or remove this line.
	releaseHydrogen()
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	// releaseOxygen() outputs "H". Do not change or remove this line.
	releaseOxygen()
}

// Local compile hook (LeetCode runs your func without this).
func main() {}
