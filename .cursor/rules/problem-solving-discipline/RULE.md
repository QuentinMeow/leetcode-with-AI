---
description: Structured problem-solving approach — enforces interview-style discipline when working on LeetCode problems
alwaysApply: true
---

## Purpose

Force a structured approach to every LeetCode problem, mirroring what a strong candidate does in a real coding interview. Prevents jumping straight to code and builds the habit of thinking before typing.

## The Problem-Solving Steps

When the user starts working on a new problem, guide them through these steps in order. Do not skip steps unless the user explicitly asks to.

### Step 1: Understand

- Read the problem statement carefully.
- Clarify constraints: input size, value ranges, edge cases.
- Walk through the provided examples to confirm understanding.
- Ask: "What are the edge cases?" (empty input, single element, duplicates, negative numbers, overflow)

### Step 2: Pattern Recognition

- Identify which algorithm pattern(s) apply. Check `.cursor/skills/leetcode-coach/references/algorithm-patterns.md` for the catalog.
- Consider multiple patterns before committing to one.
- If the user has struggled with this pattern before (check MEMORY.md), provide extra scaffolding.

### Step 3: Plan

- Outline the approach in plain language or pseudocode before writing real code.
- State the expected time and space complexity of the planned approach.
- If the plan has a suboptimal complexity, mention it and ask whether the user wants to try optimizing first.

### Step 4: Implement

- Write clean, readable code in the target language.
- Use meaningful variable names (not single letters unless idiomatic for the pattern).
- Handle edge cases identified in Step 1.

### Step 5: Verify

- Trace through the code with at least one example manually.
- Test edge cases mentally or by running code via the MCP `run_code` tool.
- Fix bugs before submitting.

### Step 6: Analyze

- State the actual time and space complexity of the implementation.
- Compare with the planned complexity from Step 3.
- If suboptimal, discuss what the optimal approach would be.

### Step 7: Record

- If the user made mistakes or learned something new, update `.cursor/MEMORY.md` (via memory-keeper rule).
- If the mistake reveals a generic lesson, record it in the relevant skill's `LESSONS.md`.
- Update `progress.md` with the problem.

## When Reviewing (not solving fresh)

If the user is reviewing or re-solving a problem they've seen before, abbreviate: skip Step 1 unless they ask, focus on whether they can recall the pattern (Step 2) and implement cleanly (Step 4).

## When the User is Stuck

Follow progressive hint levels instead of showing the answer:

1. **Category hint**: "This is a [pattern] problem."
2. **Approach hint**: "Think about using [data structure/technique]."
3. **Algorithm detail**: "The key insight is [specific observation]."
4. **Pseudocode**: Show the algorithm structure without language-specific code.

Only show a full solution if the user explicitly asks for it after receiving hints.
