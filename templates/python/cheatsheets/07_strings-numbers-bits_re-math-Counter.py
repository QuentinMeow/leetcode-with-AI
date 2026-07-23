"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

import collections
import math
import re

# ====================================================================
# 7. Strings / Numbers / Bits
# ====================================================================

# Shows common string creation, slicing, searching, joining, counting, and regular expressions.
# Requires: import collections
# Requires: import re
def string_patterns(
    s: str, chars: list[str], nums: list[int]
) -> None:
    # Split / join / trim.
    words = s.split()  # Split on runs of whitespace.
    csv_parts = [
        part.strip() for part in s.split(",") if part.strip()
    ]
    lines = s.splitlines()
    joined = " ".join(words)
    answer_line = ",".join(str(x) for x in nums)

    # Efficient build: list append + join.
    pieces: list[str] = []
    for ch in chars:
        pieces.append(ch)
    built = "".join(pieces)

    # Slices: end index is exclusive.
    first_three = s[:3]
    without_last = s[:-1]
    reversed_s = s[::-1]
    every_other = s[::2]

    # Search / replace.
    contains = "needle" in s
    first_idx = s.find("needle")  # -1 if absent.
    count_a = s.count("a")
    replaced_once = s.replace("old", "new", 1)
    replaced_all = s.replace("old", "new")

    # Case / whitespace.
    lower = s.lower()
    upper = s.upper()
    stripped = s.strip()
    left_stripped = s.lstrip()
    right_stripped = s.rstrip()
    starts = s.startswith("pre")
    ends = s.endswith("suf")

    # Character classes are common in palindrome / parsing problems.
    normalized = "".join(ch.lower() for ch in s if ch.isalnum())
    is_clean_palindrome = normalized == normalized[::-1]
    kind_flags = [
        (ch.isalpha(), ch.isdigit(), ch.isalnum(), ch.isspace())
        for ch in s[:3]
    ]

    # Ord / chr for compact fixed alphabet arrays.
    idx = ord("c") - ord("a")
    ch = chr(ord("a") + idx)
    digit = ord("7") - ord("0")
    freq26 = [0] * 26
    for ch2 in s:
        if "a" <= ch2 <= "z":
            freq26[ord(ch2) - ord("a")] += 1
    anagram_key_counts = tuple(freq26)
    anagram_key_sorted = "".join(sorted(s))
    char_counts = collections.Counter(s)
    top_chars = char_counts.most_common(3)

    # Parsing numbers from strings.
    ints_from_spaces = [int(part) for part in s.split()]
    ints_from_text = [int(m) for m in re.findall(r"-?\d+", s)]
    digits_from_string = [
        ord(ch) - ord("0") for ch in s if ch.isdigit()
    ]
    zero_padded = f"{len(s):04d}"  # Useful for labels/debug output.




# Shows integer conversion, exact square roots, number theory helpers, and bit operations.
# Requires: import math
def numeric_bit_patterns(a: int, b: int, n: int, mask: int) -> None:
    sign = (a > 0) - (a < 0)
    absolute = abs(a)
    clamped = max(0, min(n, 100))

    quotient, remainder = divmod(a, b)  # Same as (a // b, a % b).
    true_division = a / b
    floor_division = a // b  # Floors toward -inf.
    trunc_toward_zero = (abs(a) // abs(b)) * (1 if a * b >= 0 else -1)
    ceil_div_positive = -(-a // b)  # For positive b.
    remainder = a % b  # If b > 0, result is in [0, b).
    normalized_mod = (a % b + b) % b

    inf = float("inf")
    neg_inf = float("-inf")
    also_inf = math.inf

    gcd = math.gcd(a, b)
    lcm = math.lcm(a, b)
    root_floor = math.isqrt(n)
    is_square = root_floor * root_floor == n
    choose_two = math.comb(n, 2) if n >= 2 else 0
    permutations = math.perm(n, 2) if n >= 2 else 0

    mod = 10**9 + 7
    add_mod = (a + b) % mod
    mul_mod = (a * b) % mod
    pow_mod = pow(a, n, mod)
    # Modular inverse when gcd(a, mod) == 1.
    # inv_mod = pow(a, -1, mod)

    # Decimal digits.
    digits = [int(ch) for ch in str(abs(n))]
    digit_sum = sum(digits)
    x = abs(n)
    reversed_digits = 0
    while x:
        reversed_digits = reversed_digits * 10 + x % 10
        x //= 10

    one_bit = 1 << 3
    mask |= 1 << 3  # Set bit.
    has_bit = (mask & (1 << 3)) != 0
    mask &= ~(1 << 3)  # Clear bit.
    mask ^= 1 << 3  # Toggle bit.
    bits = mask.bit_count()
