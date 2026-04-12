# Language coach

Helps you understand **how a language works**: syntax vs semantics, what is sugar, what is optional for humans but required by the grammar, and interview-style follow-ups.

## What makes it different

- **Progressive disclosure**: a **tracker** file lists everything you have touched at a glance; each **concept** has its own markdown file with layered sections (summary → gaps → deep notes).
- **Depth levels** (`glance` → `standard` → `deep`): once a topic is **deep**, the agent should not re-teach the whole thing from scratch unless you ask for review or new material — it builds on what is already written.
- **Per language**: start with `learning/python/`; add e.g. `learning/golang/` the same way when needed.

## How to use

- *“Explain this Python line by line — what is sugar?”*
- *“What have I already studied in Python?”* → reads the tracker only.
- *“Go deeper on iterators — we only did a glance before.”*
- *“Quiz me on closures, but don’t repeat the long explanation.”* (works best after the concept file is populated)

## Paths

| Path | Role |
|------|------|
| `learning/python/TRACKER.md` | Your Python concept index |
| `learning/python/concepts/<topic>.md` | Deep notes for one concept |
| `learning/python/concepts/_TEMPLATE.md` | Copy when starting a new concept file |

The agent follows the workflow in `SKILL.md` and the design notes in `references/harness-and-progressive-disclosure.md`.
