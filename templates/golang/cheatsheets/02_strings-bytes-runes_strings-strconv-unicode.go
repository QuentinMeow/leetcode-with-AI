// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"slices"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// ===================================================================
// 2. Strings, Bytes, and Runes
// ===================================================================

// Use byte indexing for guaranteed ASCII input. Use range or []rune when a
// logical character may occupy multiple UTF-8 bytes. Strings are immutable.

// stringPatterns is a non-returning catalog of byte, rune, strings-package, and strconv
// operations; pass a non-empty string because the indexing examples intentionally show
// direct access.
// Requires: import "strconv"
// Requires: import "strings"
// Requires: import "unicode"
// Requires: import "unicode/utf8"
func stringPatterns(s string, nums []int) {
	// A string is immutable bytes, commonly UTF-8.
	byteLength := len(s)
	firstByte := s[0] // Panics if empty.
	prefix := s[:min(5, len(s))]
	suffix := s[min(5, len(s)):]
	middle := s[min(1, len(s)):min(5, len(s))]
	validUTF8 := utf8.ValidString(s)
	runeCount := utf8.RuneCountInString(s)

	// range decodes UTF-8 and yields byte index plus rune.
	for byteIndex, r := range s {
		_, _ = byteIndex, r
	}
	firstRuneIsUpper, firstRuneIsLower := false, false
	uppercasedFirstRune := rune(0)
	if len([]rune(s)) > 0 {
		first := []rune(s)[0]
		firstRuneIsUpper = unicode.IsUpper(first)
		firstRuneIsLower = unicode.IsLower(first)
		uppercasedFirstRune = unicode.ToUpper(first)
	}

	bytesCopy := []byte(s)
	runes := []rune(s)
	firstRune := runes[0] // Panics if there are no runes.
	reconstructed := string(runes)

	words := strings.Fields(s)
	csvParts := strings.Split(s, ",")
	joined := strings.Join(words, " ")
	contains := strings.Contains(s, "needle")
	firstIndex := strings.Index(s, "needle") // Byte index, -1 if absent.
	lastIndex := strings.LastIndex(s, "needle")
	repeated := strings.Repeat("go", 3)
	trimmedCutset := strings.Trim(s, "-_")
	count := strings.Count(s, "a")
	replacedOnce := strings.Replace(s, "old", "new", 1)
	replacedAll := strings.ReplaceAll(s, "old", "new")
	trimmed := strings.TrimSpace(s)
	lower := strings.ToLower(s)
	upper := strings.ToUpper(s)
	starts := strings.HasPrefix(s, "pre")
	ends := strings.HasSuffix(s, "suf")

	var builder strings.Builder
	for _, word := range words {
		if builder.Len() > 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(word)
	}
	built := builder.String()

	numberText := make([]string, len(nums))
	for i, value := range nums {
		numberText[i] = strconv.Itoa(value)
	}
	answerLine := strings.Join(numberText, ",")

	parsed, err := strconv.Atoi("42")
	parsed64, err64 := strconv.ParseInt("-42", 10, 64)
	parsedHex, hexErr := strconv.ParseInt("ff", 16, 64)
	formattedHex := strconv.FormatInt(255, 16)
	formatted := strconv.Itoa(parsed)

	_ = []any{
		byteLength, firstByte, prefix, suffix, middle, validUTF8, runeCount, bytesCopy,
		firstRune, reconstructed, csvParts, joined, contains,
		firstIndex, lastIndex, repeated, trimmedCutset, count, replacedOnce, replacedAll, trimmed,
		lower, upper, starts, ends, built, answerLine, parsed,
		err, parsed64, err64, parsedHex, hexErr, formattedHex, formatted,
		firstRuneIsUpper, firstRuneIsLower, uppercasedFirstRune,
	}
}

// normalizeLettersAndDigitsLowercase keeps Unicode letters and digits, lowercases
// letters, and removes every other rune.
// Requires: import "strings"
// Requires: import "unicode"
func normalizeLettersAndDigitsLowercase(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(unicode.ToLower(r))
		}
	}
	return builder.String()
}

// sortedRunesKey returns the input runes in sorted order. It is useful as a canonical
// key for grouping anagrams, not for locale-aware display ordering.
// Requires: import "slices"
func sortedRunesKey(s string) string {
	runes := []rune(s)
	slices.Sort(runes)
	return string(runes)
}
