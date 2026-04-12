---
description:
alwaysApply: true
---

# Agent Contract

## Summary

This repo is a LeetCode practice environment with AI coaching. Solutions live under `solutions/<language>/`. Cursor rules enforce workflow discipline. Cursor skills provide coaching and interview simulation. A **local** LeetCode MCP server (see below) handles problem fetch/search and related tools; submit/run on LeetCode depends on MCP package version and `LEETCODE_SESSION` configuration.

## Folder Structure

| Path | Purpose |
|------|---------|
| `solutions/<language>/` | Solution files, one per problem, named `<NNNN>-<slug>.<ext>` |
| `.cursor/mcp-example.json` | **Committed** template for LeetCode MCP — copy to `.cursor/mcp.json` and edit (see README). |
| `.cursor/mcp.json` | **Gitignored** — your real MCP config (paths, optional `LEETCODE_SESSION`). Cursor reads this file name only after you create it. |
| `.cursor/rules/<name>/RULE.md` | Cursor project rules |
| `.cursor/skills/<name>/SKILL.md` | Cursor skills (coaching, interview simulation, language study) |
| `.cursor/skills/language-coach/learning/<lang>/` | Per-language tracker + `concepts/*.md` for language-coach progress |
| `.cursor/MEMORY.md` | Cross-session memory for progress and mistakes (gitignored) |
| `progress.md` | Problem-solving progress tracker |
| `README.md` | Human-facing repo overview |
| `AGENTS.md` | This file (agent-facing contract) |
| `data/leetcode/` | Fetched problem-list JSON (from `scripts/fetch-leetcode-problem-lists.mjs`) used by `scripts/generate-leetcode-skeletons.mjs` |
| `scripts/fetch-leetcode-problem-lists.mjs` | Downloads first/last N problems per track into `data/leetcode/{python,golang}/` |
| `scripts/generate-leetcode-skeletons.mjs` | Writes `solutions/{python,golang}/NNNN-slug.*` stubs from LeetCode starter code + those JSON lists |

## Handy Commands

```bash
# List all rules
ls .cursor/rules/

# List all skills
ls .cursor/skills/

# List solved problems by language
ls solutions/python/
ls solutions/golang/

# Fetch catalog JSON → data/leetcode/… then generate solution stubs
npm run leetcode:fetch-lists
npm run leetcode:generate-skeletons
```

## Guidance to AI Agent Tasks

### Read Order

1. Read this file first for repo-level orientation.
2. Read `.cursor/rules/repo-discovery/RULE.md` for layout and rule routing.
3. If the user asks about LeetCode MCP setup or tools are missing: explain that `.cursor/mcp.json` is gitignored; they must copy `.cursor/mcp-example.json` → `.cursor/mcp.json`, replace every placeholder path, optionally set `LEETCODE_SESSION`, then reload Cursor. Do not commit `.cursor/mcp.json`.
4. Read `.cursor/MEMORY.md` (if it exists) for cross-session context.
5. Read the target skill's `SKILL.md` before invoking skill-specific workflows.
6. For **language-coach**, read `learning/<lang>/TRACKER.md` before long explanations; load `concepts/<slug>.md` only when needed.

### Solution File Convention

- Filename: `<4-digit-number>-<leetcode-slug>.<ext>` (e.g. `0001-two-sum.py`). Use at least four digits; if `questionFrontendId` exceeds 9999, the unpadded id is used (e.g. `10000-problem-slug.py`).
- One solution per file. If the user wants multiple approaches, use comments to separate them within the file.
- To (re)generate empty stubs for problems listed under `data/leetcode/`, run `npm run leetcode:generate-skeletons` (see `solutions/README.md`). Use `--force` to overwrite existing files.

### Banned Tools

Never use `claude-code`, `@anthropic-ai/claude-code`, or any CLI tool containing "claude code" in its name or invocation. Never run unsandboxed AI CLI agents. See `.cursor/rules/banned-tools/RULE.md`.

### Doc Ownership

- `README.md` is human-facing. Do not add agent instructions to it.
- `AGENTS.md` is agent-facing. Do not add human usage guides to it.
- Child docs own their own details; parent docs summarize and link.
