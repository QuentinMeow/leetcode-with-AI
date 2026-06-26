# Databases - Agent Guide

## Summary

This topic helps with senior+ system design database selection. Keep the root of this folder limited to:

- `README.md` - human-facing navigation.
- `AGENTS.md` - agent-facing navigation and edit rules.
- `cheatsheet.md` - quick interview reference.
- `individual-analysis/` - detailed database deep dives.

Do not add separate root-level reference files. Fold reference guidance into the three root files or place detailed analysis under `individual-analysis/`.

## Read Order

1. Read `README.md` for human-facing scope and file list.
2. Read `cheatsheet.md` before answering database choice questions.
3. Read the matching `individual-analysis/*.md` only when deeper details are needed.
4. If the user asks about caches, Redis, invalidation, cache stampede, rate limiting, sessions, or leaderboards, use `../cache/` instead.

## File Naming

Use database type prefixes so related files group together:

- `sql-`
- `key-value-`
- `document-`
- `wide-column-`
- `search-`
- `time-series-`
- `graph-`
- `geospatial-`
- `analytics-`
- `vector-`
- `event-log-`
- `object-storage-`

## Deep Dive Header Contract

Each `individual-analysis/*.md` file should remain easy to skim:

1. `30-Second Interview Answer`
2. `Use When`
3. `Avoid When`
4. `Core Model`
5. Query/indexing and consistency sections as relevant
6. `Scaling`
7. `Data Modeling`
8. `Interview Examples`
9. `Senior-Level Tradeoffs`
10. `Common Mistakes`

## Decision Guidance

Prefer conservative interview reasoning:

- Start with relational SQL for core transactional source of truth unless requirements push otherwise.
- Add specialized databases only for clear access patterns or scale constraints.
- Treat search/vector/analytics as derived stores unless explicitly chosen as the source of truth.
- Treat event logs as append-only replay/fanout stores, not arbitrary query databases.
- Treat geospatial as an access pattern/index that may sit on SQL, search, cache, or key-value systems.
- Mention consistency, partition keys, hot keys, indexing cost, and rebuild/reconciliation strategy.
