# Interview Simulator — Agent Guide

## Summary

This skill simulates coding interview dynamics. The agent plays the role of an interviewer, asking follow-up questions and evaluating the user's responses across multiple dimensions.

## Folder Structure

| Path | Purpose |
|------|---------|
| `SKILL.md` | Agent-facing entry point with interview workflow and rubric |
| `README.md` | Human-facing overview |
| `AGENTS.md` | This file |
| `references/follow-up-bank.md` | Categorized follow-up questions by problem type |
| `references/evaluation-rubric.md` | Detailed scoring criteria |
| `LESSONS.md` | Accumulated lessons about effective interview coaching |
| `issues/` | Structured records of simulation failures |

## Read Order

1. Read `SKILL.md` for the interview workflow, styles, and rubric.
2. Read `LESSONS.md` (if it exists) before simulating — it contains calibration notes.
3. Read `references/follow-up-bank.md` for specific follow-up questions by problem category.
4. Read `.cursor/MEMORY.md` for the user's interview performance history.

## Edit Rules

- New lessons must pass the quality filter in `.cursor/rules/skill-learning/RULE.md`.
- Evaluation feedback should be specific and actionable, not generic praise.
- Follow-up questions should match the problem's difficulty and category.
