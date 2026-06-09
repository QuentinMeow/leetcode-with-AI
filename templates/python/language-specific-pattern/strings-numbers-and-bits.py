"""
Python string, number, and bit-operation patterns for interviews.

Use this when a problem is conceptually simple but Python syntax details matter:
parsing, character math, joining strings, integer division, modulo, or bit masks.
"""


# -----------------------------------------------------------------------------
# Strings are immutable
# -----------------------------------------------------------------------------


def build_string(chars: list[str]) -> str:
    # Strings are immutable. Repeated `s += ch` creates new string objects and
    # can become O(n^2). Accumulate pieces and join once.
    pieces: list[str] = []
    for ch in chars:
        pieces.append(ch)
    return "".join(pieces)


def reverse_words(sentence: str) -> str:
    words = sentence.split()
    return " ".join(reversed(words))


# -----------------------------------------------------------------------------
# Character operations
# -----------------------------------------------------------------------------


def char_index(ch: str) -> int:
    # ord returns the Unicode code point for a one-character string.
    return ord(ch) - ord("a")


def index_char(i: int) -> str:
    # chr converts a Unicode code point back to a one-character string.
    return chr(ord("a") + i)


def frequency_26(s: str) -> list[int]:
    counts = [0] * 26

    for ch in s:
        counts[ord(ch) - ord("a")] += 1

    return counts


# -----------------------------------------------------------------------------
# Parsing and formatting
# -----------------------------------------------------------------------------


def parse_ints(line: str) -> list[int]:
    return [int(part) for part in line.split()]


def parse_csv_ints(line: str) -> list[int]:
    return [int(part.strip()) for part in line.split(",") if part.strip()]


def format_answer(nums: list[int]) -> str:
    return ",".join(str(x) for x in nums)


# -----------------------------------------------------------------------------
# Division and modulo
# -----------------------------------------------------------------------------


def division_examples(a: int, b: int) -> tuple[float, int, int]:
    true_division = a / b  # Always returns float.
    floor_division = a // b  # Floors toward negative infinity.
    remainder = a % b  # Has the same sign as b.
    return true_division, floor_division, remainder


def truncate_toward_zero(a: int, b: int) -> int:
    # Python // floors toward negative infinity. int(a / b) truncates toward 0.
    return int(a / b)


def positive_mod_index(i: int, n: int) -> int:
    # In Python, -1 % 5 == 4. This is useful for circular arrays.
    return i % n


# -----------------------------------------------------------------------------
# Infinity and sentinels
# -----------------------------------------------------------------------------


def min_value(nums: list[int]) -> int | None:
    if not nums:
        return None

    best = float("inf")
    for x in nums:
        best = min(best, x)

    return int(best)


# -----------------------------------------------------------------------------
# Bit operations
# -----------------------------------------------------------------------------


def bit_mask_add_remove(values: list[int]) -> int:
    mask = 0

    for x in values:
        mask |= 1 << x  # Set bit x to 1.

    for x in values[::2]:
        mask &= ~(1 << x)  # Set bit x to 0.

    return mask


def has_bit(mask: int, bit: int) -> bool:
    return (mask & (1 << bit)) != 0


def count_bits(mask: int) -> int:
    return mask.bit_count()


def single_number(nums: list[int]) -> int:
    ans = 0
    for x in nums:
        ans ^= x
    return ans


"""
Interview notes:

- Strings are immutable; build with list append + `"".join(...)`.
- `split()` without arguments collapses repeated whitespace.
- `ord` and `chr` convert between characters and Unicode code points.
- `//` floors, which matters for negative numbers.
- `%` is always non-negative when the divisor is positive.
- Python integers are arbitrary precision; no overflow in normal interview code.
- `x.bit_count()` counts set bits in Python 3.8+.

Concept explanations:

- Python strings are immutable sequences of Unicode characters. Operations that
  appear to modify a string create a new string instead.
- `"".join(parts)` is efficient because Python can allocate the final string
  once after seeing all pieces.
- `split()` without arguments treats any run of whitespace as one separator and
  drops leading/trailing whitespace. `split(",")` uses exactly comma separators.
- `strip()` removes leading and trailing whitespace by default. It is common
  when parsing input that may include spaces after commas.
- `ord(ch)` returns a Unicode code point. For lowercase English letters,
  `ord(ch) - ord("a")` maps `"a"` to 0 through `"z"` to 25.
- `chr(i)` is the inverse operation for a Unicode code point. It is useful for
  reconstructing characters from numeric offsets.
- `/` is true division and returns a float. `//` is floor division, not
  truncation toward zero. This differs from languages where integer division
  truncates.
- Python modulo follows floor-division semantics. If `n` is positive,
  `i % n` is always in `0..n-1`, which is useful for circular arrays.
- `float("inf")` and `float("-inf")` are convenient sentinels for min/max
  scans, but return `None` or handle empty input explicitly when no answer
  exists.
- Python integers do not overflow, so bit manipulation is about logic rather
  than fixed-width integer limits unless the problem explicitly defines a width.
- `1 << bit` creates a mask with one bit set. `|` sets bits, `&` tests or keeps
  bits, `~` flips bits, and `^` toggles bits.
- `x.bit_count()` counts 1-bits directly and is clearer than manual loops when
  the Python version supports it.
"""
