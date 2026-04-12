package main

// LeetCode 1226 — The Dining Philosophers
// https://leetcode.com/problems/the-dining-philosophers/

type DiningPhilosophers struct {
}

func (this *DiningPhilosophers) wantsToEat(
    philosopher int,
    pickLeftFork func(),
    pickRightFork func(),
    eat func(),
    putLeftFork func(),
    putRightFork func(),
) {
    // TODO: implement your solution here
}

// Local compile hook (LeetCode runs your func without this).
func main() {}
