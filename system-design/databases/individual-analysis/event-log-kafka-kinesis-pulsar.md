# Event Log - Kafka / Kinesis / Pulsar

Examples: Apache Kafka, Amazon Kinesis, Apache Pulsar, Redpanda, EventStoreDB.

## 30-Second Interview Answer

Use a durable event log when the system needs ordered append-only events, replay, fanout to multiple consumers, change data capture, audit history, or event sourcing. It is not a replacement for an OLTP database, but it is often the backbone for derived stores and async workflows.

## Use When

- Multiple consumers need the same event stream.
- Events must be replayable for backfills or rebuilding indexes.
- You need CDC from OLTP into search, cache, OLAP, or ML pipelines.
- The system is append-heavy and order matters within a key.
- Audit trail or event sourcing is a core requirement.

## Avoid When

- You need random point lookups by ID.
- You need relational transactions and constraints.
- A simple queue with delete-on-consume semantics is enough.
- Consumers cannot handle duplicates or replays.
- Retention/replay is not needed.

## Core Model

An event log stores records in append-only topics/streams. Records are partitioned by key. Consumers track offsets and can replay from earlier offsets while data remains within retention.

Key concepts:

- Topic/stream.
- Partition/shard.
- Offset/sequence number.
- Consumer group.
- Retention policy.
- Compaction for latest value by key.

## Query And Indexing

Event logs are not query databases. The main access pattern is sequential read by partition and offset.

Interview detail: use the log to feed queryable stores. For example, write events to Kafka, then build Elasticsearch for search, Redis/cache for hot reads, and OLAP for analytics.

## Consistency And Ordering

Ordering is usually guaranteed only within a partition key. Choose the partition key based on ordering needs:

- `conversation_id` for chat message order.
- `order_id` for order lifecycle events.
- `user_id` for user activity sequence.

Exactly-once behavior usually depends on idempotent producers, transactional writes where supported, dedupe keys, and idempotent consumers.

## Scaling

Scale through partitions/shards and consumer groups.

Failure modes:

- Hot partition from a celebrity/user/tenant.
- Consumer lag after traffic spikes.
- Poison messages blocking consumers.
- Reprocessing duplicates after retry.
- Retention too short for backfills.
- Schema evolution breaking old consumers.

## Data Modeling

Good event design:

- Stable event name and version.
- Entity ID and partition key.
- Event ID for dedupe.
- Occurred-at timestamp.
- Producer metadata.
- Backward-compatible schema.

Prefer immutable facts: `OrderPlaced`, `PaymentCaptured`, `MessageSent`, not vague commands like `ProcessThing`.

## Interview Examples

- CDC from SQL to search index.
- Rebuild a feed, cache, or OLAP table from events.
- Payment/order audit trail.
- Notification fanout.
- Activity stream ingestion.
- Event-sourced account ledger with snapshots.

## Senior-Level Tradeoffs

- Event logs decouple producers and consumers, but introduce lag, schema governance, and replay complexity.
- Retention enables recovery and backfills, but costs storage.
- Partitioning gives scale, but limits global ordering.
- A log can preserve facts, but user-facing reads usually need materialized views.

## Common Mistakes

- Calling Kafka a database for arbitrary queries.
- Assuming global ordering across partitions.
- Ignoring duplicate delivery and idempotent consumers.
- Forgetting schema evolution and retention for replay.
- Letting consumer lag silently break freshness assumptions.
