// Go string, number, and bit patterns for coding interviews.
//
// The central string decision is whether a problem is defined over raw bytes
// (common for ASCII constraints) or Unicode code points (runes).
package languagepatterns

import (
	"math"
	"math/bits"
	"strconv"
	"strings"
	"unicode/utf8"
)

// BuildStringPattern avoids repeated immutable-string concatenation by using a
// strings.Builder. Builder.String returns the accumulated string.
func BuildStringPattern(parts []string) string {
	var builder strings.Builder
	for _, part := range parts {
		builder.WriteString(part)
	}
	return builder.String()
}

// ReverseWordsPattern demonstrates Fields and Join. Fields collapses runs of
// Unicode whitespace and discards leading/trailing whitespace.
func ReverseWordsPattern(sentence string) string {
	words := strings.Fields(sentence)
	for left, right := 0, len(words)-1; left < right; left, right = left+1, right-1 {
		words[left], words[right] = words[right], words[left]
	}
	return strings.Join(words, " ")
}

// ReverseRunesPattern reverses Unicode code points. Reversing bytes instead
// would corrupt multi-byte UTF-8 encodings.
func ReverseRunesPattern(text string) string {
	runes := []rune(text)
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}
	return string(runes)
}

// StringSizePattern distinguishes byte length from rune count.
func StringSizePattern(text string) (bytesCount, runesCount int, validUTF8 bool) {
	return len(text), utf8.RuneCountInString(text), utf8.ValidString(text)
}

// RuneOffsetsPattern records byte offsets produced by ranging over a string.
func RuneOffsetsPattern(text string) map[int]rune {
	atByteOffset := make(map[int]rune)
	for offset, value := range text {
		atByteOffset[offset] = value
	}
	return atByteOffset
}

// LowercaseFrequencyPattern is appropriate only when the problem guarantees
// ASCII 'a' through 'z'. A byte array is faster and clearer than Unicode logic.
func LowercaseFrequencyPattern(text string) [26]int {
	var counts [26]int
	for index := 0; index < len(text); index++ {
		value := text[index]
		if value >= 'a' && value <= 'z' {
			counts[value-'a']++
		}
	}
	return counts
}

// ParseIntsPattern parses whitespace-separated base-10 ints.
func ParseIntsPattern(line string) ([]int, error) {
	fields := strings.Fields(line)
	values := make([]int, 0, len(fields))
	for _, field := range fields {
		value, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}

// FormatIntsPattern converts numbers once and joins with a delimiter.
func FormatIntsPattern(values []int, separator string) string {
	parts := make([]string, len(values))
	for index, value := range values {
		parts[index] = strconv.Itoa(value)
	}
	return strings.Join(parts, separator)
}

// DivisionPattern shows Go integer division and remainder. Both truncate toward
// zero, so the remainder has the same sign as the dividend.
func DivisionPattern(a, b int) (quotient, remainder int, ok bool) {
	if b == 0 {
		return 0, 0, false
	}
	return a / b, a % b, true
}

// PositiveModuloPattern normalizes a remainder into [0, modulus). This is
// needed because -1%5 is -1 in Go, unlike languages with floor modulo.
func PositiveModuloPattern(value, modulus int) (int, bool) {
	if modulus <= 0 {
		return 0, false
	}
	remainder := value % modulus
	if remainder < 0 {
		remainder += modulus
	}
	return remainder, true
}

// CeilingDivisionPattern is valid for non-negative numerator and positive
// denominator. The rearranged form avoids numerator+denominator-1 overflow.
func CeilingDivisionPattern(numerator, denominator int) (int, bool) {
	if numerator < 0 || denominator <= 0 {
		return 0, false
	}
	if numerator == 0 {
		return 0, true
	}
	return 1 + (numerator-1)/denominator, true
}

// FloatMathPattern shows common math helpers. Converting a very large int to
// float64 can lose precision, so use integer algorithms when exactness matters.
func FloatMathPattern(value int) (float64, float64) {
	asFloat := float64(value)
	return math.Sqrt(asFloat), math.Abs(asFloat)
}

// BitMaskPattern sets, tests, toggles, and clears one bit.
func BitMaskPattern(mask uint, position uint) (set, toggled, cleared uint, present bool) {
	bit := uint(1) << position
	set = mask | bit
	present = mask&bit != 0
	toggled = mask ^ bit
	cleared = mask &^ bit // &^ is Go's bit-clear operator.
	return set, toggled, cleared, present
}

// BitCountsPattern uses math/bits. Len returns the minimum bits needed to
// represent value; both results are zero for value == 0.
func BitCountsPattern(value uint) (ones, width int) {
	return bits.OnesCount(value), bits.Len(value)
}

// SingleNumberPattern uses x^x == 0 and x^0 == x.
func SingleNumberPattern(nums []int) int {
	answer := 0
	for _, value := range nums {
		answer ^= value
	}
	return answer
}

/*
String rules:

- A string is an immutable sequence of bytes, conventionally UTF-8.
- text[i] is a byte. `for _, r := range text` decodes runes.
- Converting string to []byte or []rune allocates a mutable copy; converting
  back to string normally allocates again.
- Use bytes/ASCII when constraints guarantee it. Use runes when "character"
  means a Unicode code point. A user-perceived grapheme can contain multiple
  runes, but interview problems rarely require grapheme segmentation.

Number rules:

- Integer overflow follows fixed-width two's-complement arithmetic for signed
  integers. Choose int64 when bounds can exceed int on a 32-bit environment.
- strconv.ParseInt(text, base, bitSize) is preferable when width/base is part of
  the problem. Atoi parses into platform-sized int.
- Shifts require an unsigned or non-negative integer count at runtime. A shift
  that discards high bits follows the destination integer width.
*/
