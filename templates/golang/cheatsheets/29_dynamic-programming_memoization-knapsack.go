// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

// ===================================================================
// 29. Dynamic Programming
// ===================================================================

// maximumSubarraySumUsingKadaneAlgorithm returns the largest sum of a non-empty
// contiguous subarray, or 0 for empty input. At each position the best ending there
// either starts fresh or extends the previous best ending. Time O(n); space O(1).
func maximumSubarraySumUsingKadaneAlgorithm(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	best, current := nums[0], nums[0]
	for _, value := range nums[1:] {
		current = max(value, current+value)
		best = max(best, current)
	}
	return best
}

// countWaysToClimbStairsUsingRollingState counts sequences of one- and two-step moves.
// Ways(step)=ways(step-1)+ways(step-2); keeping only those two states reduces
// dynamic-programming space to O(1).
func countWaysToClimbStairsUsingRollingState(n int) int {
	if n <= 2 {
		return n
	}
	previous2, previous1 := 1, 2
	for step := 3; step <= n; step++ {
		previous2, previous1 = previous1, previous1+previous2
	}
	return previous1
}

// longestCommonSubsequence returns the maximum length obtainable by deleting characters
// without reordering either string. Matching characters advance both strings; otherwise
// choose the better one-string advance. Time/space O(len(a)*len(b)).
func longestCommonSubsequence(a, b string) int {
	dp := make([][]int, len(a)+1)
	for index := range dp {
		dp[index] = make([]int, len(b)+1)
	}
	for i := len(a) - 1; i >= 0; i-- {
		for j := len(b) - 1; j >= 0; j-- {
			if a[i] == b[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[0][0]
}

// canPartitionIntoEqualSumSubsets asks whether values can split into two equal-sum
// groups. It is a 0/1 knapsack: possible[s] means some processed values total s, and
// descending updates prevent reusing one value. Time O(n*target).
func canPartitionIntoEqualSumSubsets(nums []int) bool {
	total := 0
	for _, value := range nums {
		total += value
	}
	if total%2 != 0 {
		return false
	}

	target := total / 2
	possible := make([]bool, target+1)
	possible[0] = true
	for _, value := range nums {
		for current := target; current >= value; current-- {
			possible[current] = possible[current] ||
				possible[current-value]
		}
	}
	return possible[target]
}

// minimumCoinsTopDownDynamicProgramming returns the fewest coins totaling amount, or
// -1. Memoization stores each remaining amount; amount+1 is an impossible sentinel
// larger than any valid answer using positive coin values.
func minimumCoinsTopDownDynamicProgramming(coins []int, amount int) int {
	impossible := amount + 1
	memo := make(map[int]int)

	var dp func(int) int
	dp = func(remaining int) int {
		switch {
		case remaining == 0:
			return 0
		case remaining < 0:
			return impossible
		}
		if cached, ok := memo[remaining]; ok {
			return cached
		}

		best := impossible
		for _, coin := range coins {
			best = min(best, 1+dp(remaining-coin))
		}
		memo[remaining] = best
		return best
	}

	answer := dp(amount)
	if answer >= impossible {
		return -1
	}
	return answer
}

// countGridPathsUsingMemoization counts paths from the top-left to bottom-right using
// only down and right moves. The coordinate pair is a comparable map key, and
// memoization evaluates each cell once. Time/space O(rows*cols).
func countGridPathsUsingMemoization(rows, cols int) int {
	memo := make(map[[2]int]int)
	var count func(int, int) int
	count = func(row, col int) int {
		if row == rows-1 && col == cols-1 {
			return 1
		}
		if row >= rows || col >= cols {
			return 0
		}
		key := [2]int{row, col}
		if cached, ok := memo[key]; ok {
			return cached
		}
		memo[key] = count(row+1, col) + count(row, col+1)
		return memo[key]
	}
	if rows <= 0 || cols <= 0 {
		return 0
	}
	return count(0, 0)
}
