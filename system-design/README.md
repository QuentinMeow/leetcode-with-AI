# System Design Study Guide

Use this folder for senior+ system design interview preparation. The goal is not to memorize product names; the goal is to make clear tradeoffs under time pressure.

## Folder Pattern

Every system design topic should follow the same shape:

- `README.md` - human-facing overview and navigation.
- `AGENTS.md` - AI-agent-facing navigation, conventions, and edit guidance.
- `cheatsheet.md` - fast interview reference.
- `individual-analysis/` - deep dives that are too detailed for the cheatsheet.

Do not add extra root-level reference files inside a topic unless the structure changes intentionally. Put detailed learning notes under `individual-analysis/`.

## Topics

- `databases/` - source-of-truth stores, derived data stores, geospatial indexes, event logs, and durable object storage.
- `cache/` - Redis, Memcached, local/in-process cache, CDN/edge/client cache, invalidation, and hot-path acceleration.
- `cheatsheet.md` - general system design strategy for interviews.

## Interview Loop

1. Clarify the product goal, users, core flows, read/write patterns, latency, availability, durability, and scale.
2. State the data model and access patterns before naming technology.
3. Pick the simplest data store that preserves correctness for the critical invariants.
4. Add scale only where the bottleneck is explicit: partitioning, replication, caching, indexes, queues, or async processing.
5. Discuss failure modes, consistency, backfills, observability, and operational ownership.
6. Close with how the design evolves when scale or requirements change.

## Senior-Level Signals

- You separate hard requirements from convenient implementation details.
- You defend choices with access patterns, invariants, and failure modes.
- You know when not to use a specialized database.
- You mention migration paths instead of pretending the first design is final.
- You balance user experience, correctness, cost, and operability.
