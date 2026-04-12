# Solutions

LeetCode solutions organized by language.

## Folder structure

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

## Naming convention

- **Filename**: `<problem-id>-<leetcode-slug>.<extension>`
- **Slug** comes from the URL path (e.g. `leetcode.com/problems/two-sum/` → `two-sum`).
- **Problem id** is LeetCode’s `questionFrontendId`. Pad with leading zeros to **at least four** digits (e.g. `1` → `0001`). If the id is already four or more digits (e.g. `3701`, `10000`), use it as-is in the filename.

## Problem lists and skeleton files

The JSON under `data/leetcode/` is only a **catalog index** (titles, slugs, difficulties). It is not a substitute for solution files.

### 1. Refresh the catalog (optional)

From the repo root (requires network to [leetcode.com](https://leetcode.com)):

```bash
npm run leetcode:fetch-lists
```

This writes `data/leetcode/python/first-100.json`, `last-100.json`, and the same payloads under `data/leetcode/golang/` (same problems; tracks are for your workflow, not different APIs).

### 2. Generate starter code under `solutions/`

```bash
npm run leetcode:generate-skeletons
```

This reads those JSON files, calls LeetCode’s API for each slug, and writes **Python 3** and **Go** files using the **official** starter signatures (plus minimal fixes so files parse locally):

- **Python**: inserts `pass` in empty method bodies and adds common `typing` imports when hints use them.
- **Go**: unwraps doc-block helper types (e.g. `ListNode`), adds `panic("TODO")` in empty functions, and appends `func main() {}` so **`go build path/to/one-file.go`** succeeds. Each file is its own `package main`; do not `go build *.go` in the whole directory (duplicate `main`). On LeetCode, submit only the solution function(s) if their editor wraps your code; you can delete `main` locally if you prefer.

**Overwrite existing files** (including ones you already edited):

```bash
node scripts/generate-leetcode-skeletons.mjs --force
```

**Only one language** (still fetches each problem once from the API):

```bash
node scripts/generate-leetcode-skeletons.mjs --python-only
node scripts/generate-leetcode-skeletons.mjs --golang-only
```

## Adding a new language

1. Create a folder under `solutions/` named after the language (lowercase): `solutions/java/`, `solutions/cpp/`, `solutions/rust/`, etc.
2. Add language-specific tooling files if needed (`go.mod`, `requirements.txt`, `Cargo.toml`, etc.) inside the language folder.
3. Follow the same `<problem-id>-<slug>.<ext>` naming convention.

## Language-specific notes

### Python

- Target: Python 3.11+ (matches LeetCode runtime).
- No external dependencies needed for solutions.

### Go

- Target: Go 1.21+ (matches LeetCode runtime).
- Generated stubs use `package main` with a no-op `main` for local tooling; adjust to match how you run tests or paste into LeetCode.
