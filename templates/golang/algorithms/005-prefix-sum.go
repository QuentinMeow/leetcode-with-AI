package algorithms

/*
005 - Prefix sum patterns

Use when repeated range sums are needed, or when subarray problems can be
converted into differences between prefix states.
*/

// Variant 1: prefix sum with a hash map of seen prefix counts.
// Example problems: subarray sum equals k, path sum with prefix counts.
// Time: O(n)
// Space: O(n)
func CountSubarraysSumK(nums []int, k int) int {
	seen := map[int]int{0: 1}
	prefix, total := 0, 0
	for _, x := range nums {
		prefix += x
		total += seen[prefix-k]
		seen[prefix]++
	}
	return total
}

// Variant 2: prefix array for immutable range sum queries.
// Example problems: range sum query, sum between i and j many times.
// Build time: O(n)
// Query time: O(1)
// Space: O(n)
type PrefixSum struct {
	prefix []int
}

func NewPrefixSum(nums []int) *PrefixSum {
	prefix := make([]int, len(nums)+1)
	for i, x := range nums {
		prefix[i+1] = prefix[i] + x
	}
	return &PrefixSum{prefix: prefix}
}

// RangeSum returns the sum of the inclusive range [left, right].
func (p *PrefixSum) RangeSum(left, right int) int {
	return p.prefix[right+1] - p.prefix[left]
}

// Variant 3: 2D prefix sum.
// Example problems: range sum query 2D, matrix block sum.
// Build time: O(m * n)
// Query time: O(1)
// Space: O(m * n)
type PrefixSum2D struct {
	prefix [][]int
}

func NewPrefixSum2D(matrix [][]int) *PrefixSum2D {
	rows := len(matrix)
	cols := 0
	if rows > 0 {
		cols = len(matrix[0])
	}
	prefix := make([][]int, rows+1)
	for r := range prefix {
		prefix[r] = make([]int, cols+1)
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			prefix[r+1][c+1] = matrix[r][c] +
				prefix[r][c+1] +
				prefix[r+1][c] -
				prefix[r][c]
		}
	}
	return &PrefixSum2D{prefix: prefix}
}

// RegionSum returns the inclusive rectangle from (r1, c1) to (r2, c2).
func (p *PrefixSum2D) RegionSum(r1, c1, r2, c2 int) int {
	return p.prefix[r2+1][c2+1] -
		p.prefix[r1][c2+1] -
		p.prefix[r2+1][c1] +
		p.prefix[r1][c1]
}

// Variant 4: difference array for many range updates.
// Example problems: range addition, car pooling, meeting capacity.
// Time: O(n + q)
// Space: O(n)
// Each update is [left, right, delta] with an inclusive range.
func ApplyRangeAdditions(length int, updates [][3]int) []int {
	diff := make([]int, length+1)
	for _, update := range updates {
		left, right, delta := update[0], update[1], update[2]
		diff[left] += delta
		if right+1 < len(diff) {
			diff[right+1] -= delta
		}
	}

	result := make([]int, length)
	running := 0
	for i := range result {
		running += diff[i]
		result[i] = running
	}
	return result
}
