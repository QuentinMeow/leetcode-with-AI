// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"math/bits"
)

// ===================================================================
// 14. Bit Manipulation
// ===================================================================

// bitManipulationExamples demonstrates read, set, clear, and toggle operations.
// Requires: import "math/bits"
func bitManipulationExamples(mask uint) {
	mask |= 1 << 3
	hasBit := mask&(1<<3) != 0
	isOdd := mask&1 == 1
	mask &^= 1 << 3 // &^ means AND NOT; it clears selected bits.
	mask ^= 1 << 3
	clearedLowestSetBit := mask & (mask - 1)
	isolatedLowestSetBit := mask & -mask
	setBitCount := bits.OnesCount(mask)
	bitLength := bits.Len(mask)
	trailingZeroCount := bits.TrailingZeros(mask)
	leadingZeroCount := bits.LeadingZeros(mask)
	_ = []any{hasBit, isOdd, clearedLowestSetBit, isolatedLowestSetBit, setBitCount, bitLength, trailingZeroCount, leadingZeroCount}
}

// isPowerOfTwo is true when a positive value has exactly one set bit.
func isPowerOfTwo(value uint) bool {
	return value != 0 && value&(value-1) == 0
}

// allSubsetsUsingBitmask maps each bit to one input position. It is compact and useful
// when n is small enough for O(n*2^n) enumeration.
// Requires: import "math/bits"
func allSubsetsUsingBitmask(nums []int) [][]int {
	if len(nums) >= bits.UintSize {
		return nil
	}
	answer := make([][]int, 0, 1<<len(nums))
	for mask := 0; mask < 1<<len(nums); mask++ {
		subset := make([]int, 0, bits.OnesCount(uint(mask)))
		for index, value := range nums {
			if mask&(1<<index) != 0 {
				subset = append(subset, value)
			}
		}
		answer = append(answer, subset)
	}
	return answer
}
