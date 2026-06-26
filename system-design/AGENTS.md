# System Design - Agent Guide

## Summary

This folder stores senior+ system design interview prep. Keep the structure repeatable across topics so both humans and agents can navigate quickly under time pressure.

## Topic Folder Contract

Each topic under `system-design/` should use this layout:

| Path | Purpose |
|---|---|
| `README.md` | Human-facing overview and navigation |
| `AGENTS.md` | Agent-facing navigation, conventions, and edit guidance |
| `cheatsheet.md` | Fast interview reference and decision framework |
| `individual-analysis/` | Detailed learning notes and deep dives |

Avoid extra root-level topic files. If a note is detailed, put it in `individual-analysis/`. If it is a quick reference, fold it into `cheatsheet.md`. If it is navigation, fold it into `README.md` or `AGENTS.md`.

## Current Topics

- `databases/` - durable source-of-truth stores, derived stores, geospatial indexes, event logs, and object storage.
- `cache/` - Redis, Memcached, local/in-process cache, CDN/edge/client cache, and invalidation strategy.

## Read Order

1. Read `system-design/README.md` for human-facing orientation.
2. Read the topic `AGENTS.md` before editing or expanding that topic.
3. Read the topic `cheatsheet.md` before answering interview-style questions.
4. Read `individual-analysis/*.md` only when deeper explanation is needed.

## Editing Rules

- Preserve the topic folder contract unless the user explicitly changes it.
- Keep `cheatsheet.md` short, skimmable, and interview-actionable.
- If a cheatsheet compares multiple architectures, components, or products, place a `Visual Decision Tree` section immediately after the title and before detailed explanation.
- In visual decision trees, leaf nodes must name concrete choices: product names or close product families such as `PostgreSQL/MySQL`, `DynamoDB`, `Redis/Valkey`, `Kafka/Kinesis/Pulsar`, or `S3/GCS/Azure Blob`.
- Do not make a decision-tree leaf only a concept like "SQL", "NoSQL", "cache", or "queue"; include actual choices at the leaf.
- Keep deep dives consistent: 30-second answer, use when, avoid when, model, scaling/failure modes, examples, tradeoffs, mistakes.
- Prefer cross-links to sibling topic cheatsheets rather than duplicating full content.
- Use lowercase `cheatsheet.md`.
