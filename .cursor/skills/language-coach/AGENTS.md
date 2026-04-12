# Language coach — agent guide

## Summary

Explains language syntax, semantics, sugar, and interview-oriented details. **State lives in `learning/<language>/`**: `TRACKER.md` (global shallow index) plus `concepts/<slug>.md` (one file per concept, append-only detail).

## Folder structure

| Path | Purpose |
|------|---------|
| `SKILL.md` | Harness workflow, depth rules, disclosure layers |
| `README.md` | Human-facing overview |
| `AGENTS.md` | This file |
| `learning/README.md` | Corpus layout |
| `learning/python/TRACKER.md` | Python concept index — **read first** for Python |
| `learning/python/concepts/*.md` | Per-concept notes (`_TEMPLATE.md` is copy source) |
| `references/harness-and-progressive-disclosure.md` | Why tracker + layers |
| `references/interview-language-topics.md` | Optional coverage checklist |
| `LESSONS.md` | Coaching mistakes (this skill), not user vocabulary |
| `issues/` | Per skill-learning RULE |

## Read order

1. `SKILL.md` for workflow.
2. `learning/<lang>/TRACKER.md` before teaching.
3. `learning/<lang>/concepts/<slug>.md` when the question needs more than the tracker row.
4. `references/harness-and-progressive-disclosure.md` when depth/disclosure is ambiguous.
5. `LESSONS.md` before repeating a known coaching mistake.

## Edit rules

- Never delete tracker rows for concepts that were studied; update depth and summaries.
- Prefer **append** in **Detailed notes** with dates.
- New languages: duplicate `learning/python/` layout as `learning/<lang>/`.
- New lessons and issues follow `.cursor/rules/skill-learning/RULE.md`.
