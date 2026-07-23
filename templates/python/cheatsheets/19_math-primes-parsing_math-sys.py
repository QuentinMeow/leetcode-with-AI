"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

import collections.abc
import math

# ====================================================================
# 19. Math, Primes, and Parsing
# ====================================================================

# Returns decimal digits from most significant to least significant.
def base10_digits(n: int) -> list[int]:
    n = abs(n)
    if n == 0:
        return [0]
    digits: list[int] = []
    while n:
        n, digit = divmod(n, 10)
        digits.append(digit)
    return digits[::-1]




# Builds an integer from decimal digits and rejects values outside 0 through 9.
# Requires: import collections.abc
def int_from_base10_digits(
    digits: collections.abc.Iterable[int],
) -> int:
    result = 0
    for digit in digits:
        result = result * 10 + digit
    return result




# Divides numerator and denominator by their greatest common divisor.
# Requires: import math
def reduce_fraction(
    numerator: int, denominator: int
) -> tuple[int, int]:
    g = math.gcd(numerator, denominator)
    numerator //= g
    denominator //= g
    if denominator < 0:
        numerator, denominator = -numerator, -denominator
    return numerator, denominator


# Tests possible divisors through the square root. Time O(sqrt(n)).
def is_prime_trial(n: int) -> bool:
    if n < 2:
        return False
    if n % 2 == 0:
        return n == 2
    factor = 3
    while factor * factor <= n:
        if n % factor == 0:
            return False
        factor += 2
    return True


# Sieve of Eratosthenes marks composite multiples and returns primality through n.
def sieve_is_prime(n: int) -> list[bool]:
    is_prime = [True] * (n + 1)
    if n >= 0:
        is_prime[0] = False
    if n >= 1:
        is_prime[1] = False
    p = 2
    while p * p <= n:
        if is_prime[p]:
            for multiple in range(p * p, n + 1, p):
                is_prime[multiple] = False
        p += 1
    return is_prime


# Returns each prime factor mapped to its exponent using trial division.
def prime_factor_counts(n: int) -> dict[int, int]:
    n = abs(n)
    factors: dict[int, int] = {}
    divisor = 2
    while divisor * divisor <= n:
        while n % divisor == 0:
            factors[divisor] = factors.get(divisor, 0) + 1
            n //= divisor
        divisor += 1 if divisor == 2 else 2
    if n > 1:
        factors[n] = factors.get(n, 0) + 1
    return factors


# Assumes all values occur twice except one; exclusive OR cancels equal pairs.
def single_number_xor(nums: list[int]) -> int:
    result = 0
    for x in nums:
        result ^= x
    return result


# Reverses decimal digits and returns 0 when the signed 32-bit range would overflow.
def reverse_integer_32(x: int) -> int:
    sign = -1 if x < 0 else 1
    x = abs(x)
    limit = 2**31 - 1 if sign > 0 else 2**31
    result = 0
    while x:
        digit = x % 10
        x //= 10
        if result > limit // 10 or (
            result == limit // 10 and digit > limit % 10
        ):
            return 0
        result = result * 10 + digit
    return sign * result


# Parses optional sign and decimal digits, clamping to the signed 32-bit range.
def parse_signed_int32_clamped(s: str) -> int:
    i = 0
    while i < len(s) and s[i] == " ":
        i += 1
    sign = 1
    if i < len(s) and s[i] in "+-":
        sign = -1 if s[i] == "-" else 1
        i += 1
    result = 0
    while i < len(s) and s[i].isdigit():
        result = result * 10 + int(s[i])
        i += 1
    return max(-(2**31), min(2**31 - 1, sign * result))


# Checks whether digit runs skip exactly the stated number of word characters.
def valid_word_abbreviation(word: str, abbr: str) -> bool:
    i = j = 0
    while i < len(abbr):
        if j >= len(word):
            return False
        if abbr[i] == word[j]:
            i += 1
            j += 1
        elif not abbr[i].isdigit() or abbr[i] == "0":
            return False
        else:
            skip = 0
            while i < len(abbr) and abbr[i].isdigit():
                skip = skip * 10 + int(abbr[i])
                i += 1
            j += skip
    return j == len(word)


# Parses a comma-separated transaction into typed fields after validating field count.
def parse_transaction_record(raw: str) -> tuple[str, int, int, str]:
    name, time, amount, city = raw.split(",")
    return name, int(time), int(amount), city
