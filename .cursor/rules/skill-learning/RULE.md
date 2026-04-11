---
description: Skill learning framework — issue tracking and lesson extraction for all skills
globs: ".cursor/skills/**"
alwaysApply: false
---

## Purpose

Every skill should accumulate operational knowledge over time. When the agent encounters a coaching failure, misdiagnosis of a problem pattern, or gives a wrong hint, it must record what happened, why, and what generic principle prevents recurrence.

## Required Skill Feedback Files

Every skill under `.cursor/skills/<skill-name>/` should account for these feedback files:

| Path | Purpose |
|------|---------|
| `issues/` | Structured issue records (one file per issue) |
| `LESSONS.md` | Accumulated generic lessons extracted from resolved issues |

## Issue Recording

When the agent gives incorrect guidance, misidentifies a pattern, or the user corrects a mistake:

1. Create a new file at `issues/<NNN>-<slug>.md` (zero-padded 3-digit sequence).
2. Use this template:

```markdown
---
id: <NNN>
title: <concise title>
status: open | resolved
severity: critical | high | medium | low
root_cause: <one-sentence root cause>
resolved_at: <ISO date or empty>
lesson_extracted: true | false
---

## Symptoms

What happened and how the failure manifested.

## Root Cause

Why it happened — the underlying mechanism, not just the surface trigger.

## Fix

What was changed to resolve it.

## Lesson

The generic principle extracted (copied verbatim into LESSONS.md when resolved).
```

3. Mark `status: resolved` and `lesson_extracted: true` once verified and added to `LESSONS.md`.

## Lesson Extraction

After resolving an issue, extract a **generic lesson** and append to `LESSONS.md`.

### Quality Filter

| Gate | Good example | Bad example |
|------|-------------|-------------|
| **Generic** — applies beyond the specific bug | "Sliding window problems with variable-length windows need two pointers, not a fixed offset" | "Problem 3 needs a set" |
| **Actionable** — an agent can follow it | "When the user says 'subarray', check whether the problem wants contiguous elements before suggesting two pointers" | "Be careful with arrays" |
| **Non-obvious** — the agent wouldn't already know | "Monotonic stack problems often appear disguised as 'next greater element' or 'largest rectangle' variants" | "Read the problem carefully" |

Reject lessons that fail any gate.

### LESSONS.md Format

```markdown
# Lessons

## <Category>

- **<Principle>**: <Explanation of why this matters and when to apply it.>
```

Group by category (e.g., "Pattern Recognition", "Edge Cases", "Complexity Analysis"). Each lesson is one bullet: bold principle, then explanation.

## Continuous Improvement Loop

```
coaching mistake → record issue → fix approach → extract lesson → update LESSONS.md
```

Agents should follow this loop automatically when they give wrong guidance. Do not wait for human instruction.

## Anti-Patterns

- **Overfitting**: Writing a lesson that only applies to one problem. Ask: "Would this help on a different problem in the same category?"
- **Workarounds over understanding**: A lesson that says "use this template" is weaker than one that says "this pattern applies when you see X because Y."
- **Noisy logging**: Only record problems that caused incorrect guidance or wasted the user's time.

## Read Order

1. Check `LESSONS.md` before coaching on a topic — it contains hard-won operational knowledge.
2. Check `issues/` for open issues before modifying a skill.
