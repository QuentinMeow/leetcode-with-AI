// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

// ===================================================================
// 23. Prefix Sums and Difference Arrays
// ===================================================================

// OneDimensionalPrefixSum stores cumulative sums for constant-time range queries.
type OneDimensionalPrefixSum struct {
	prefix []int
}

// newOneDimensionalPrefixSum stores prefix[i] as the sum of values before index i. The
// leading zero makes an inclusive range [left,right] equal
// prefix[right+1]-prefix[left]. Build time/space O(n).
func newOneDimensionalPrefixSum(nums []int) OneDimensionalPrefixSum {
	prefix := make([]int, len(nums)+1)
	for index, value := range nums {
		prefix[index+1] = prefix[index] + value
	}
	return OneDimensionalPrefixSum{prefix: prefix}
}

// sumRange returns the inclusive one-dimensional range sum [left,right] in O(1).
// Indices must be valid for the original input.
func (p OneDimensionalPrefixSum) sumRange(left, right int) int {
	return p.prefix[right+1] - p.prefix[left]
}

// TwoDimensionalPrefixSum stores cumulative rectangle sums.
type TwoDimensionalPrefixSum struct {
	prefix [][]int
}

// newTwoDimensionalPrefixSum builds a table with one zero border. prefix[r+1][c+1] is
// the sum of the rectangle from (0,0) through (r,c). Build time/space O(rows*cols).
func newTwoDimensionalPrefixSum(matrix [][]int) TwoDimensionalPrefixSum {
	rows := len(matrix)
	cols := 0
	if rows > 0 {
		cols = len(matrix[0])
	}

	prefix := make([][]int, rows+1)
	for row := range prefix {
		prefix[row] = make([]int, cols+1)
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			prefix[row+1][col+1] = matrix[row][col] +
				prefix[row][col+1] +
				prefix[row+1][col] -
				prefix[row][col]
		}
	}
	return TwoDimensionalPrefixSum{prefix: prefix}
}

// sumRegion returns an inclusive rectangle sum using inclusion-exclusion: whole prefix
// minus the two outside strips plus their twice-subtracted overlap. Query time O(1).
func (p TwoDimensionalPrefixSum) sumRegion(r1, c1, r2, c2 int) int {
	return p.prefix[r2+1][c2+1] -
		p.prefix[r1][c2+1] -
		p.prefix[r2+1][c1] +
		p.prefix[r1][c1]
}

// applyRangeAdditionsUsingDifferenceArray applies inclusive updates [left,right,delta]
// without touching every covered element. +delta marks where an effect starts and
// -delta after right marks where it stops; one prefix scan materializes values. Time
// O(length+updates).
func applyRangeAdditionsUsingDifferenceArray(length int, updates [][3]int) []int {
	if length <= 0 {
		return nil
	}
	difference := make([]int, length+1)
	for _, update := range updates {
		left, right, delta := update[0], update[1], update[2]
		if left < 0 || left >= length || right < left || right >= length {
			continue
		}
		difference[left] += delta
		difference[right+1] -= delta
	}

	result := make([]int, length)
	running := 0
	for index := range result {
		running += difference[index]
		result[index] = running
	}
	return result
}
