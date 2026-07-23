// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

// ===================================================================
// 19. Hash Maps and Counting
// ===================================================================

// groupAnagramsBySortedRunes groups words containing the same runes with the same
// multiplicities. Sorting each word's runes creates a canonical map key: anagrams
// produce identical keys. For maximum word length m, time O(totalWords*m log m).
// Requires via helper: import "slices"
func groupAnagramsBySortedRunes(words []string) [][]string {
	groups := make(map[string][]string)
	for _, word := range words {
		key := sortedRunesKey(word)
		groups[key] = append(groups[key], word)
	}

	answer := make([][]string, 0, len(groups))
	for _, group := range groups {
		answer = append(answer, group)
	}
	return answer
}

// longestConsecutiveSequenceLength returns the longest run of consecutive integers in
// any input order. It starts counting only at values whose predecessor is absent, so
// each stored value participates in one run. Expected time O(n); space O(n).
func longestConsecutiveSequenceLength(nums []int) int {
	values := make(map[int]struct{}, len(nums))
	for _, value := range nums {
		values[value] = struct{}{}
	}
	best := 0
	for value := range values {
		if _, exists := values[value-1]; exists {
			continue
		}
		end := value
		for {
			if _, exists := values[end]; !exists {
				break
			}
			end++
		}
		best = max(best, end-value)
	}
	return best
}

// topKFrequentValuesUsingBuckets returns k values with highest frequencies. A value
// occurring f times is placed in bucket f; scanning buckets from high to low avoids
// sorting all distinct values. Tie order is unspecified. Time O(n); space O(n).
func topKFrequentValuesUsingBuckets(nums []int, k int) []int {
	if k <= 0 {
		return nil
	}

	frequency := make(map[int]int)
	for _, value := range nums {
		frequency[value]++
	}

	buckets := make([][]int, len(nums)+1)
	for value, count := range frequency {
		buckets[count] = append(buckets[count], value)
	}

	answer := make([]int, 0, min(k, len(frequency)))
	for count := len(buckets) - 1; count > 0 && len(answer) < k; count-- {
		for _, value := range buckets[count] {
			answer = append(answer, value)
			if len(answer) == k {
				break
			}
		}
	}
	return answer // Ties may appear in any order.
}
