package algorithms

import (
	"sort"
	"strings"
)

/*
001 - Hash map and counting patterns

Use when the problem asks for pairs, frequencies, first/last seen positions,
membership checks, grouping, or "have we seen this before?" logic.
*/

// Variant 1: complement lookup, the most common pair pattern.
// Example problems: Two Sum, pair with target difference, pair existence.
// Time: O(n)
// Space: O(n)
func FindPairIndices(nums []int, target int) ([2]int, bool) {
	seen := make(map[int]int)
	for i, x := range nums {
		if j, ok := seen[target-x]; ok {
			return [2]int{j, i}, true
		}
		seen[x] = i
	}
	return [2]int{}, false
}

// Variant 2: frequency counting.
// Example problems: anagrams, majority checks, least/most frequent values.
// Time: O(n)
// Space: O(k), where k is the number of distinct values.
func FrequencyTable(items []int) map[int]int {
	counts := make(map[int]int)
	for _, item := range items {
		counts[item]++
	}
	return counts
}

// Variant 3: group by derived key.
// Example problems: group anagrams, bucket by normalized form, classify strings.
// Time: O(n * m log m) for sorted-string keys, where m is average word length.
// Space: O(n * m)
func GroupAnagrams(words []string) [][]string {
	groups := make(map[string][]string)
	order := make([]string, 0)
	for _, word := range words {
		letters := strings.Split(word, "")
		sort.Strings(letters)
		key := strings.Join(letters, "")
		if _, ok := groups[key]; !ok {
			order = append(order, key)
		}
		groups[key] = append(groups[key], word)
	}

	result := make([][]string, 0, len(groups))
	for _, key := range order {
		result = append(result, groups[key])
	}
	return result
}

// Variant 4: first seen index for longest distance / subarray transforms.
// Example problems: contiguous array, equal prefix states, first repeated state.
// Time: O(n)
// Space: O(k), where k is the number of distinct states.
func LongestSpanWithSameState(states []int) int {
	firstSeen := make(map[int]int)
	best := 0
	for i, state := range states {
		if first, ok := firstSeen[state]; ok {
			best = max(best, i-first)
		} else {
			firstSeen[state] = i
		}
	}
	return best
}
