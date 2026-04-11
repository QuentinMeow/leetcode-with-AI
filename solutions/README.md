# Solutions

LeetCode solutions organized by language.

## Folder Structure

```
solutions/
  python/
    0001-two-sum.py
    0015-3sum.py
  golang/
    0001-two-sum.go
  <language>/
    <NNNN>-<slug>.<ext>
```

## Naming Convention

- **Filename**: `<4-digit-problem-number>-<leetcode-slug>.<extension>`
- The problem number and slug come directly from the LeetCode URL (e.g. `leetcode.com/problems/two-sum/` with problem number 1 becomes `0001-two-sum`)

## Adding a New Language

1. Create a folder under `solutions/` named after the language (lowercase): `solutions/java/`, `solutions/cpp/`, `solutions/rust/`, etc.
2. Add language-specific tooling files if needed (`go.mod`, `requirements.txt`, `Cargo.toml`, etc.) inside the language folder.
3. Follow the same `<NNNN>-<slug>.<ext>` naming convention.

## Language-Specific Notes

### Python

- Target: Python 3.11+ (matches LeetCode runtime)
- No external dependencies needed for solutions

### Go

- Target: Go 1.21+ (matches LeetCode runtime)
- Each file should use `package main` or a test-friendly package name
