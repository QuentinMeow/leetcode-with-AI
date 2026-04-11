---
description: Repo discovery and routing — canonical entrypoint for repo layout and related rules
alwaysApply: true
---

## Repo Layout

| Path | What lives here |
|------|-----------------|
| `solutions/<language>/` | LeetCode solutions organized by language |
| `.cursor/mcp-example.json` | Template MCP config (committed); copy to `mcp.json` locally |
| `.cursor/mcp.json` | Real MCP config (gitignored — local paths and optional session cookie) |
| `.cursor/rules/<name>/RULE.md` | Cursor project rules (always-on or glob-scoped) |
| `.cursor/skills/<name>/SKILL.md` | Cursor skills with optional `references/` |
| `.cursor/MEMORY.md` | Cross-session memory for progress and mistakes (gitignored) |
| `progress.md` | Problem-solving progress tracker |
| `README.md` | Human-facing docs |
| `AGENTS.md` | Agent-facing contract |

## Navigation

- Start with `AGENTS.md` at repo root for the agent contract.
- Start with `README.md` at repo root for human orientation.
- Read a skill's `SKILL.md` before invoking skill-specific workflows.
- Read `.cursor/MEMORY.md` at session start for cross-session context.

## Related Rules

Use these supporting rules as the task demands:

1. `.cursor/rules/memory-keeper/RULE.md` for reading and updating persistent cross-session knowledge.
2. `.cursor/rules/problem-solving-discipline/RULE.md` for the structured problem-solving approach.
3. `.cursor/rules/skill-learning/RULE.md` when working inside `.cursor/skills/**`.
4. `.cursor/rules/banned-tools/RULE.md` for hard restrictions on banned CLI tools.

## Solution File Convention

- Filename: `<4-digit-number>-<leetcode-slug>.<ext>` (e.g. `0001-two-sum.py`)
- Language folder: lowercase language name under `solutions/`
- One solution per file
