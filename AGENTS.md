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
| `.cursor/skills/<name>/SKILL.md` | Cursor skills (coaching, interview simulation) |
| `.cursor/MEMORY.md` | Cross-session memory for progress and mistakes (gitignored) |
| `progress.md` | Problem-solving progress tracker |
| `README.md` | Human-facing repo overview |
| `AGENTS.md` | This file (agent-facing contract) |

## Handy Commands

```bash
# List all rules
ls .cursor/rules/

# List all skills
ls .cursor/skills/

# List solved problems by language
ls solutions/python/
ls solutions/golang/
```

## Guidance to AI Agent Tasks

### Read Order

1. Read this file first for repo-level orientation.
2. Read `.cursor/rules/repo-discovery/RULE.md` for layout and rule routing.
3. If the user asks about LeetCode MCP setup or tools are missing: explain that `.cursor/mcp.json` is gitignored; they must copy `.cursor/mcp-example.json` → `.cursor/mcp.json`, replace every placeholder path, optionally set `LEETCODE_SESSION`, then reload Cursor. Do not commit `.cursor/mcp.json`.
4. Read `.cursor/MEMORY.md` (if it exists) for cross-session context.
5. Read the target skill's `SKILL.md` before invoking skill-specific workflows.

### Solution File Convention

- Filename: `<4-digit-number>-<leetcode-slug>.<ext>` (e.g. `0001-two-sum.py`)
- One solution per file. If the user wants multiple approaches, use comments to separate them within the file.

### Banned Tools

Never use `claude-code`, `@anthropic-ai/claude-code`, or any CLI tool containing "claude code" in its name or invocation. Never run unsandboxed AI CLI agents. See `.cursor/rules/banned-tools/RULE.md`.

### Doc Ownership

- `README.md` is human-facing. Do not add agent instructions to it.
- `AGENTS.md` is agent-facing. Do not add human usage guides to it.
- Child docs own their own details; parent docs summarize and link.
