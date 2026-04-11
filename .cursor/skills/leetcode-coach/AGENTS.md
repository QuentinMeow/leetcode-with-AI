# LeetCode Coach — Agent Guide

## Summary

This skill provides algorithm coaching for LeetCode practice. The agent guides users through problems with progressive hints rather than direct answers, and tracks learning over time.

## Folder Structure

| Path | Purpose |
|------|---------|
| `SKILL.md` | Agent-facing entry point with coaching workflow |
| `README.md` | Human-facing overview |
| `AGENTS.md` | This file |
| `references/algorithm-patterns.md` | Catalog of algorithm patterns with trigger signals |
| `references/complexity-guide.md` | Time/space complexity quick reference |
| `LESSONS.md` | Accumulated coaching lessons (auto-populated) |
| `issues/` | Structured records of coaching mistakes |

## Read Order

1. Read `SKILL.md` for the coaching workflow and hint levels.
2. Read `LESSONS.md` (if it exists) before coaching — it contains corrections from past mistakes.
3. Read `references/algorithm-patterns.md` when identifying which pattern a problem belongs to.
4. Check `issues/` for open issues before modifying this skill.

## Edit Rules

- Do not modify `references/algorithm-patterns.md` without reading it first.
- New lessons must pass the quality filter defined in `.cursor/rules/skill-learning/RULE.md`.
- Issue files follow the template in `.cursor/rules/skill-learning/RULE.md`.
