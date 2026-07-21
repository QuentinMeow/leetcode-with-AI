package algorithms

/*
012 - Dynamic programming patterns

Use when a problem has overlapping subproblems and an optimal answer can be
built from smaller states. Start by naming the state and transition.
*/

// Variant 1: 1D rolling DP.
// Example problems: climbing stairs, house robber, min cost climbing stairs.
// Time: O(n)
// Space: O(1)
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	previousTwo, previousOne := 1, 2
	for step := 3; step <= n; step++ {
		previousTwo, previousOne = previousOne, previousOne+previousTwo
	}
	return previousOne
}

// Variant 2: top-down memoization.
// Example problems: coin change, word break, decode ways, recursive choices.
// Time: O(amount * len(coins))
// Space: O(amount)
func CoinChange(coins []int, amount int) int {
	impossible := amount + 1
	memo := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		memo[i] = -1
	}

	var dp func(int) int
	dp = func(remaining int) int {
		if remaining == 0 {
			return 0
		}
		if remaining < 0 {
			return impossible
		}
		if memo[remaining] != -1 {
			return memo[remaining]
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

// Variant 3: 2D table for two sequences.
// Example problems: longest common subsequence, edit distance, interleaving string.
// Time: O(m * n)
// Space: O(m * n)
func LongestCommonSubsequence(a, b string) int {
	rows, cols := len(a), len(b)
	dp := make([][]int, rows+1)
	for row := range dp {
		dp[row] = make([]int, cols+1)
	}
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if a[i] == b[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[0][0]
}

// Variant 4: knapsack-style capacity DP.
// Example problems: partition equal subset sum, 0/1 knapsack, target sum variants.
// Time: O(n * target)
// Space: O(target)
func CanPartition(nums []int) bool {
	total := 0
	for _, x := range nums {
		total += x
	}
	if total%2 != 0 {
		return false
	}
	target := total / 2
	possible := make([]bool, target+1)
	possible[0] = true
	for _, x := range nums {
		for current := target; current >= x; current-- {
			possible[current] = possible[current] || possible[current-x]
		}
	}
	return possible[target]
}
