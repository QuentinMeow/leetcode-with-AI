// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"math"
)

// ===================================================================
// 13. Math and Numbers
// ===================================================================

// numberAndMathExamples shows common numeric helpers. The math package uses
// float64; Go 1.21 built-ins min and max work directly with integers.
// Requires: import "math"
func numberAndMathExamples(a, b, n int) {
	absolute := a
	if absolute < 0 {
		absolute = -absolute
	}
	clamped := max(0, min(n, 100))
	quotient, remainder, nonNegativeRemainder := 0, 0, 0
	if b != 0 {
		quotient = a / b // Integer division truncates toward zero.
		remainder = a % b
		nonNegativeRemainder = ((remainder % b) + b) % b
	}
	squareRoot := math.Sqrt(float64(max(n, 0)))
	floor, ceiling := math.Floor(squareRoot), math.Ceil(squareRoot)
	absoluteFloat := math.Abs(float64(a))
	power := math.Pow(2, 10)
	baseTwoLogarithm := math.Log2(max(power, 1))
	_ = []any{absolute, clamped, quotient, remainder, nonNegativeRemainder, squareRoot, floor, ceiling, absoluteFloat, power, baseTwoLogarithm, math.MaxInt}
}

// integerSquareRoot returns floor(sqrt(n)). It avoids floating-point rounding.
func integerSquareRoot(n int) int {
	if n < 0 {
		return -1
	}
	left, right := 0, n
	for left <= right {
		middle := left + (right-left)/2
		if middle != 0 && middle > n/middle {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return right
}

// greatestCommonDivisor returns the largest non-negative integer that divides both
// inputs without a remainder. Euclid's algorithm repeatedly replaces (a, b) with (b,
// a%b), which preserves their common divisors. Time O(log min(|a|, |b|)); space O(1).
func greatestCommonDivisor(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// leastCommonMultiple returns the smallest non-negative integer divisible by
// both inputs. The identity “least common multiple * greatest common divisor =
// |a*b|” avoids searching through multiples. Time is dominated by
// greatestCommonDivisor.
func leastCommonMultiple(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	value := (a / greatestCommonDivisor(a, b)) * b // Divide first to reduce overflow risk.
	if value < 0 {
		value = -value
	}
	return value
}

// modularPowerUsingBinaryExponentiation computes base^exponent mod modulus in O(log exponent).
func modularPowerUsingBinaryExponentiation(base, exponent, modulus int64) int64 {
	if exponent < 0 || modulus <= 0 {
		return 0
	}
	base = ((base % modulus) + modulus) % modulus
	result := int64(1 % modulus)
	for exponent > 0 {
		if exponent&1 == 1 {
			result = result * base % modulus
		}
		base = base * base % modulus
		exponent >>= 1
	}
	return result
}
