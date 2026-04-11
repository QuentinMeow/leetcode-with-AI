---
description: Memory keeper — automatically read and update MEMORY.md for persistent cross-session LeetCode learning
alwaysApply: true
---

## Purpose

Maintain durable knowledge across conversations in `.cursor/MEMORY.md`. This file tracks algorithm strengths, weaknesses, mistakes, and preferences so the agent can personalize coaching over time.

## Session Start

At the beginning of every conversation, silently read `.cursor/MEMORY.md` if it exists. Do not mention reading it unless the user asks. Use its contents to inform coaching, problem selection, and hint depth.

## Session End

At the end of every **substantive** conversation (solved a problem, identified a weakness, learned a pattern), update `.cursor/MEMORY.md`:

1. **Add** new entries that pass all three quality gates.
2. **Remove** entries that are no longer true or have been superseded.
3. **Reorganize** sections if they have grown unwieldy.

Trivial conversations (single-line answers, typo fixes) do not require updates.

### Pre-Final Checklist

Before sending the final response for a substantive conversation:

1. Re-read `.cursor/MEMORY.md` so edits are based on the latest state.
2. Decide whether there are new durable entries and/or stale entries to remove.
3. Apply memory updates (or intentionally skip because no entries passed quality gates).

## Quality Gates

Every new entry must pass **all three** gates:

| Gate | Passes | Fails |
|------|--------|-------|
| **Durable** — still true next month | "Struggles with DP state transition definitions" | "Got stuck on problem 42 today" |
| **Non-obvious** — agent wouldn't know from code | "Prefers to write brute force first then optimize" | "Uses Python" |
| **Actionable** — changes how the agent should coach | "Needs extra practice on graph BFS/DFS problems" | "The repo has solutions" |

## File Creation

If `.cursor/MEMORY.md` does not exist and you have entries that pass quality gates, create it using the seed template at `.cursor/rules/memory-keeper/templates/MEMORY.md`.

## Section Format

```
# Memory

## Weak Areas
- Algorithm categories and patterns where the user struggles

## Mistakes
- Recurring errors and anti-patterns to watch for

## Strengths
- Patterns and categories the user handles well

## Preferences
- Preferred languages, coding style, explanation depth

## Progress
- Milestones, patterns mastered, problem counts
```

Each entry is a markdown bullet under the appropriate section heading. Keep entries concise — one to two sentences. Group related entries and deduplicate.

## Scope Rules

- Track algorithm learning progress, not transient session details
- Promote recurring observations; demote one-off events unless they reveal a pattern
- When a weakness becomes a strength, move the entry
