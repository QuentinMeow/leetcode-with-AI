"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

# ====================================================================
# 21. Main Guard and Executable Checks
# ====================================================================

# solve is a minimal local-script placeholder; LeetCode normally calls a
# method on Solution instead.
def solve(nums: list[int]) -> int:
    return sum(nums)


# assert_file_width fails when generated lines become too wide to scan
# comfortably. The 100-column limit leaves room for explanatory names.
def assert_file_width(max_columns: int = 100) -> None:
    with open(__file__, encoding="utf-8") as source:
        for line_no, line in enumerate(source, 1):
            text = line.rstrip("\n")
            width = len(text)
            if width <= max_columns:
                continue
            message = (
                f"line {line_no} is {width} columns "
                f"(limit {max_columns}): {text!r}"
            )
            raise AssertionError(message)


# run_smoke_checks exercises representative language and algorithm sections.
# These are integration checks for generation and renaming, not exhaustive
# correctness tests.
def run_smoke_checks() -> None:
    assert two_sum_indices_using_map([2, 7, 11], 9) == (0, 1)
    assert first_index_at_least_target([1, 3, 3, 8], 3) == 1
    assert count_subarrays_with_target_sum_using_prefix_sums(
        [1, 1, 1], 2
    ) == 2
    assert (
        shortest_path_length_unweighted_breadth_first_search(
            {0: [1], 1: [2], 2: []}, 0, 2
        )
        == 2
    )
    assert all_subsets_using_backtracking([1, 2]) == [
        [],
        [1],
        [1, 2],
        [2],
    ]
    assert unique_triplets_summing_to_zero(
        [-1, 0, 1, 2, -1, -4]
    ) == [[-1, -1, 2], [-1, 0, 1]]
    assert search_rotated_sorted_array([4, 5, 6, 0, 1, 2], 0) == 3
    assert longest_consecutive_sequence_length([100, 4, 2, 1, 3]) == 4
    assert maximum_subarray_sum_using_kadane_algorithm(
        [-2, 1, -3, 4, -1, 2, 1]
    ) == 6
    assert reduce_fraction(8, 12) == (2, 3)
    union_find = DisjointSetUnion(3)
    assert union_find.union(0, 1)
    assert union_find.find(0) == union_find.find(1)


# main runs the sample, smoke checks, and readability guard when this
# generated aggregate is executed as a local script.
def main() -> None:
    sample = [1, 2, 3]
    print(solve(sample))
    run_smoke_checks()
    assert_file_width()


if __name__ == "__main__":
    # LeetCode normally calls `Solution` methods directly, so do
    # not include this
    # in submissions unless you are building a local script or
    # template.
    main()
