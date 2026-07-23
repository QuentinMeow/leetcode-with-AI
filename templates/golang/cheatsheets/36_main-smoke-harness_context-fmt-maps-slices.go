// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

import (
	"context"
	"fmt"
	"maps"
	"slices"
)

// ===================================================================
// 36. Main and Executable Smoke Checks
// ===================================================================

// solve is a placeholder showing where a local script or contest solution entry point
// can call problem logic.
func solve(nums []int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}

// assert panics with a labeled message when a smoke-check condition is false.
func assert(condition bool, message string) {
	if !condition {
		panic("check failed: " + message)
	}
}

// runSmokeChecks executes representative examples for language helpers and algorithm
// templates.
// Requires: import "context"
// Requires: import "maps"
// Requires: import "slices"
func runSmokeChecks() {
	assert(solve([]int{1, 2, 3}) == 6, "solve")
	digit, isDigit := asciiDigitValue('7')
	assert(digit == 7 && isDigit, "ASCII digit conversion")
	boolSet, containedTwo := boolSetExamples([]int{1, 2, 3}, 2)
	assert(containedTwo && !boolSet[2], "bool set membership and delete")
	assert(
		maps.Equal(sliceToSet([]int{1, 2, 2}), map[int]struct{}{1: {}, 2: {}}),
		"slice to struct set",
	)
	legacySortedValues := []int{3, 1, 2}
	legacySortedWords := []string{"go", "array"}
	assert(
		legacySortPackageExamples(legacySortedValues, legacySortedWords),
		"legacy sort package",
	)
	assert(
		exactIndexBinarySearchClosedInterval([]int{1, 3, 5}, 3) == 1,
		"closed-interval binary search",
	)
	legacySearchIndex, legacySearchFound := legacySearchIntsExample(
		[]int{1, 3, 5},
		3,
	)
	assert(
		legacySearchIndex == 1 && legacySearchFound,
		"legacy SearchInts",
	)
	largestValues := kLargestValues([]int{3, 1, 5, 2, 4}, 3)
	slices.Sort(largestValues)
	assert(slices.Equal(largestValues, []int{3, 4, 5}), "k largest values")
	assert(greatestCommonDivisor(54, 24) == 6, "greatest common divisor")
	assert(integerSquareRoot(15) == 3, "integer square root")
	assert(isPowerOfTwo(8) && !isPowerOfTwo(6), "power of two")
	stackTop, stackRemainder, stackOK := stackOperationsExample([]int{1, 2})
	assert(
		stackTop == 2 && stackOK && slices.Equal(stackRemainder, []int{1}),
		"stack peek and pop",
	)
	row, col, found := findMatrixValue([][]int{{1, 2}, {3, 4}}, 3)
	assert(row == 1 && col == 0 && found, "labeled matrix search")
	leftSet := map[int]struct{}{1: {}, 2: {}}
	rightSet := map[int]struct{}{2: {}, 3: {}}
	assert(maps.Equal(setUnion(leftSet, rightSet), map[int]struct{}{1: {}, 2: {}, 3: {}}), "set union")
	assert(maps.Equal(setIntersection(leftSet, rightSet), map[int]struct{}{2: {}}), "set intersection")
	assert(maps.Equal(setDifference(leftSet, rightSet), map[int]struct{}{1: {}}), "set difference")
	assert(slices.Equal(drainQueueByReslicing([]int{1, 2, 3}), []int{1, 2, 3}), "reslice queue")
	assert(countOccurrences([]int{1, 2, 2, 2, 4}, 2) == 3, "count occurrences")
	assert(leastCommonMultiple(12, 18) == 36, "least common multiple")
	assert(modularPowerUsingBinaryExponentiation(2, 10, 1_000) == 24, "modular power")
	assert(len(allSubsetsUsingBitmask([]int{1, 2, 3})) == 8, "bitmask subsets")
	assert(slices.Equal(iterativeDepthFirstVisitOrder(map[int][]int{0: {1, 2}, 1: {2}}, 0), []int{0, 1, 2}), "iterative depth-first search")
	assert(
		slices.Equal(twoSumIndicesUsingMap([]int{2, 7, 11, 15}, 9), []int{0, 1}),
		"two sum",
	)
	assert(firstIndexAtLeastTarget([]int{1, 2, 2, 4}, 2) == 1, "lower bound")
	assert(longestSubstringWithoutRepeatedASCIIBytes("abcabcbb") == 3, "unique window")
	assert(addDecimalStrings("999", "1") == "1000", "decimal strings")
	assert(hasValidBracketNesting("([]{})"), "valid parentheses")
	assert(maximumSubarraySumUsingKadaneAlgorithm([]int{-2, 1, -3, 4, -1, 2, 1}) == 6, "maximum subarray using Kadane algorithm")
	assert(canPartitionIntoEqualSumSubsets([]int{1, 5, 11, 5}), "equal-sum partition using zero-or-one knapsack")
	assert(maximumSumFixedLengthWindow([]int{1, 2, 3, 4}, 2) == 7, "fixed window")
	assert(maximumSumFixedLengthWindow([]int{1, 2}, 3) == 0, "invalid fixed window")
	assert(maximumSubarraySumUsingKadaneAlgorithm(nil) == 0, "empty maximum subarray")
	assert(
		slices.Equal(
			applyRangeAdditionsUsingDifferenceArray(
				3,
				[][3]int{{0, 1, 2}, {1, 2, 3}},
			),
			[]int{2, 5, 3},
		),
		"difference array",
	)
	assert(
		slices.Equal(
			daysUntilWarmerTemperature([]int{73, 74, 75, 71, 69, 72, 76, 73}),
			[]int{1, 1, 4, 2, 1, 1, 0, 0},
		),
		"daily temperatures",
	)
	assert(
		countConnectedComponentsUsingDisjointSetUnion(5, [][2]int{{0, 1}, {1, 2}, {3, 4}}) == 2,
		"DisjointSetUnion components",
	)
	assert(
		findRedundantEdge([][2]int{{1, 2}, {1, 3}, {2, 3}}) == [2]int{2, 3},
		"redundant edge",
	)
	assert(
		maxOverlappingIntervals(
			[]Interval{{start: 0, end: 30}, {start: 5, end: 10}},
		) == 2,
		"interval sweep",
	)
	assert(countGridPathsUsingMemoization(3, 3) == 6, "memoized recursion")
	assert(runLengthEncode("aaabb") == "a3b2", "run-length encoding")
	parsedDigits, digitsOK := intFromBase10Digits([]int{1, 2, 3})
	assert(parsedDigits == 123 && digitsOK, "digits to integer")
	assert(maps.Equal(primeFactorCounts(12), map[int]int{2: 2, 3: 1}), "factors")
	assert(
		slices.Equal(
			topKFrequentValuesUsingBuckets([]int{1, 1, 1, 2, 2, 3}, 2),
			[]int{1, 2},
		),
		"top K frequent",
	)
	assert(minimumCoinsTopDownDynamicProgramming([]int{1, 2, 5}, 11) == 3, "coin change")
	assert(
		medianOfTwoSortedArraysUsingPartition([]int{1, 3}, []int{2}) == 2,
		"median partition",
	)
	assert(
		maximumConcurrentMeetingsUsingTwoPointers(
			[]Interval{{start: 0, end: 30}, {start: 5, end: 10}},
		) == 2,
		"meeting rooms",
	)
	assert(
		slices.Equal(
			insertInterval(
				[]Interval{{start: 1, end: 3}, {start: 6, end: 9}},
				Interval{start: 2, end: 5},
			),
			[]Interval{{start: 1, end: 5}, {start: 6, end: 9}},
		),
		"insert interval",
	)
	assert(
		eraseOverlapIntervals(
			[]Interval{
				{start: 1, end: 2},
				{start: 2, end: 3},
				{start: 1, end: 3},
			},
		) == 1,
		"erase overlap intervals",
	)
	assert(
		slices.Equal(
			mergeSortedArraysUsingMinHeap(
				[][]int{{1, 4}, {1, 3, 5}, {2, 6}},
			),
			[]int{1, 1, 2, 3, 4, 5, 6},
		),
		"merge K arrays",
	)
	assert(len(combinationsSummingToTargetWithReuse([]int{2, 3, 6, 7}, 7)) == 2, "combination sum")

	trie := newTrie()
	trie.insert("apple")
	assert(trie.search("apple"), "trie search")
	assert(!trie.search("app") && trie.startsWith("app"), "trie prefix")
	assert(trie.wildcardSearch("a..le"), "trie wildcard")

	cache := newLeastRecentlyUsedCache(2)
	cache.put(1, 1)
	cache.put(2, 2)
	assert(cache.get(1) == 1, "least-recently-used cache get")
	cache.put(3, 3)
	assert(cache.get(2) == -1, "least-recently-used cache eviction")

	list := &ListNode{
		Val:  1,
		Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}},
	}
	list = removeNthLinkedListNodeFromEnd(list, 2)
	assert(list.Val == 1 && list.Next.Val == 3, "remove Nth from end")

	distance := gridShortestDistancesUsingBreadthFirstSearch(
		[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		Point{Row: 0, Col: 0},
	)
	assert(distance[2][2] == 4, "grid BFS distance")
	multiDistance := gridDistancesFromMultipleSources(
		[][]int{{0, 0, 0}, {0, 0, 0}},
		[]Point{{Row: 0, Col: 0}, {Row: 1, Col: 2}},
	)
	assert(multiDistance[0][2] == 1, "multi-source BFS")
	assert(mutexCounterPattern(4) == 4, "mutex counter")
	assert(
		slices.Equal(bufferedChannelPattern([]int{1, 2, 3}), []int{1, 2, 3}),
		"buffered channel",
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	assert(
		slices.Equal(
			squareWorkerPool(ctx, []int{1, 2, 3}, 2),
			[]int{1, 4, 9},
		),
		"worker pool",
	)
}

// main demonstrates local execution. LeetCode usually calls the submitted solution
// method instead.
// Requires: import "fmt"
func main() {
	// LeetCode calls solution functions directly. Keep main only for local
	// execution and remove it when a submission editor requires that.
	fmt.Println(solve([]int{1, 2, 3}))
	runSmokeChecks()
}
