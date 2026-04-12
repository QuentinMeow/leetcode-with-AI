# Language learning corpus (per language)

This directory holds **your** language-learning state: one **tracker** per language (shallow index) and **one file per concept** (detail on demand).

- **Committed** so the agent and you share the same ground truth across machines (if you push the branch).
- **Personal**: only concept files you create or the agent adds during sessions appear here.

## Layout

```
learning/
  README.md                 # This file
  python/
    TRACKER.md              # Global index for Python — read first every time
    concepts/
      <kebab-slug>.md       # One concept per file
```

Add more languages as sibling folders (e.g. `golang/`) with the same `TRACKER.md` + `concepts/` pattern.

## Naming concepts

Use **kebab-case** filenames that match the tracker `slug` column, e.g. `list-comprehensions.md`, `garbage-collection.md`.
