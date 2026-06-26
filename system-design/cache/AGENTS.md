# Cache - Agent Guide

## Summary

This topic covers caching strategy for system design interviews. Keep this folder aligned with the system design topic contract:

- `README.md` - human-facing navigation.
- `AGENTS.md` - agent-facing navigation and edit rules.
- `cheatsheet.md` - quick interview reference.
- `individual-analysis/` - detailed cache deep dives.

## Read Order

1. Read `README.md` for scope and file list.
2. Read `cheatsheet.md` before answering cache strategy questions.
3. Read the matching `individual-analysis/cache-*.md` file for deep details.
4. Use `../databases/` when the user is choosing durable source-of-truth storage.

## Edit Rules

- Keep cache guidance separate from durable database guidance.
- Emphasize that caches are usually derived and rebuildable.
- Always discuss invalidation, TTLs, stale reads, stampede protection, hot keys, memory limits, and source of truth.
- If adding a new cache deep dive, use prefix `cache-`.

## Interview Cache Choices

- `cache-redis.md` - rich distributed cache and in-memory data structures.
- `cache-memcached.md` - simple distributed key-value cache.
- `cache-local-in-process.md` - fastest per-instance cache for small low-change data.
- `cache-cdn-edge-client.md` - edge/browser/mobile caching for static, media, and public read-heavy content.

## Interview Bias

Default stance: the database remains authoritative, cache accelerates hot or expensive reads, and correctness-sensitive workflows should use database transactions/constraints rather than cache-only logic.
