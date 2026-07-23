"""Source section for the generated Python interview cheatsheet.

Edit this file, then run: npm run cheatsheets:generate
"""

from __future__ import annotations

# ====================================================================
# 1. Version / Syntax Quick Map
# ====================================================================

"""
Python 3.8+:  walrus `:=`, `int.bit_count()`
Python 3.9+:  `list[int]`, `dict[str, int]`, `dict_a | dict_b`,
`@cache`
Python 3.10+: `int | None`, `match/case`, `itertools.pairwise`

Older equivalents:
    list[int]        -> typing.List[int]
    dict[str, int]   -> typing.Dict[str, int]
    int | None       -> typing.Optional[int]
    @cache           -> @lru_cache(maxsize=None)

`from __future__ import annotations` must appear before ordinary
imports. Every numbered file and the generated aggregate place it
there so cross-topic type hints are not evaluated immediately.
Quoted forward references also remain safe when copied without the
future import.
"""
