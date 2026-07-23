// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"slices"
	"strconv"
	"strings"
)

// ===================================================================
// 33. String, Number, and Bit Extras
// ===================================================================

// addDecimalStrings adds non-negative base-10 integers represented as strings,
// processing digits from right to left with carry.
// Requires: import "slices"
func addDecimalStrings(a, b string) string {
	i, j, carry := len(a)-1, len(b)-1, 0
	answer := make([]byte, 0, max(len(a), len(b))+1)
	for i >= 0 || j >= 0 || carry > 0 {
		x, y := 0, 0
		if i >= 0 {
			x = int(a[i] - '0')
			i--
		}
		if j >= 0 {
			y = int(b[j] - '0')
			j--
		}
		total := x + y + carry
		answer = append(answer, byte(total%10)+'0')
		carry = total / 10
	}
	slices.Reverse(answer)
	return string(answer)
}

// compareVersionNumbers compares dot-separated non-negative integer components and
// treats missing trailing components as zero.
// Requires: import "strconv"
// Requires: import "strings"
func compareVersionNumbers(first, second string) int {
	a := strings.Split(first, ".")
	b := strings.Split(second, ".")
	for index := 0; index < max(len(a), len(b)); index++ {
		x, y := 0, 0
		if index < len(a) {
			x, _ = strconv.Atoi(a[index])
		}
		if index < len(b) {
			y, _ = strconv.Atoi(b[index])
		}
		if x < y {
			return -1
		}
		if x > y {
			return 1
		}
	}
	return 0
}

// base10Digits returns the decimal digits of an integer in most-significant-first
// order.
// Requires: import "slices"
func base10Digits(n int) []int {
	if n < 0 {
		n = -n
	}
	if n == 0 {
		return []int{0}
	}
	digits := make([]int, 0)
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	slices.Reverse(digits)
	return digits
}

// isPrimeUsingTrialDivision tests divisors only through sqrt(n), written factor <=
// n/factor to avoid factor*factor overflow. After handling 2, only odd candidates are
// needed. Time O(sqrt(n)); space O(1).
func isPrimeUsingTrialDivision(n int) bool {
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return n == 2
	}
	for factor := 3; factor <= n/factor; factor += 2 {
		if n%factor == 0 {
			return false
		}
	}
	return true
}

// primeTableUsingSieveOfEratosthenes returns a bool table from 0 through n. Each
// discovered prime marks multiples starting at prime^2 because smaller multiples were
// handled earlier. Time O(n log log n); space O(n).
func primeTableUsingSieveOfEratosthenes(n int) []bool {
	isPrime := make([]bool, n+1)
	for value := 2; value <= n; value++ {
		isPrime[value] = true
	}
	for prime := 2; prime <= n/prime; prime++ {
		if !isPrime[prime] {
			continue
		}
		for multiple := prime * prime; multiple <= n; multiple += prime {
			isPrime[multiple] = false
		}
	}
	return isPrime
}

// singleUnpairedNumberUsingExclusiveOr assumes every value appears twice except one.
// Exclusive OR cancels equal pairs (x XOR x = 0) and zero changes nothing, leaving the
// unpaired value. Time O(n); space O(1).
func singleUnpairedNumberUsingExclusiveOr(nums []int) int {
	answer := 0
	for _, value := range nums {
		answer ^= value
	}
	return answer
}

// runLengthEncode converts each consecutive byte run to the byte followed by its
// decimal count, so "aaabb" becomes "a3b2". It is byte-oriented rather than
// Unicode-rune-oriented. Time O(n).
// Requires: import "strconv"
// Requires: import "strings"
func runLengthEncode(s string) string {
	if s == "" {
		return ""
	}
	var builder strings.Builder
	for start := 0; start < len(s); {
		end := start + 1
		for end < len(s) && s[end] == s[start] {
			end++
		}
		builder.WriteByte(s[start])
		builder.WriteString(strconv.Itoa(end - start))
		start = end
	}
	return builder.String()
}

// intFromBase10Digits combines decimal digits into an integer; false reports a value
// outside 0 through 9.
func intFromBase10Digits(digits []int) (int, bool) {
	value := 0
	for _, digit := range digits {
		if digit < 0 || digit > 9 {
			return 0, false
		}
		value = value*10 + digit
	}
	return value, true
}

// primeFactorCounts returns prime factors mapped to their exponents. Repeated trial
// division removes each factor before moving on; any remainder greater than one is
// prime. Time O(sqrt(n)) worst case.
func primeFactorCounts(n int) map[int]int {
	factors := make(map[int]int)
	if n < 0 {
		n = -n
	}
	for factor := 2; factor <= n/factor; factor++ {
		for n%factor == 0 {
			factors[factor]++
			n /= factor
		}
	}
	if n > 1 {
		factors[n]++
	}
	return factors
}
