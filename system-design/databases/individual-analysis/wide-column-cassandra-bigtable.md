# Wide-Column - Cassandra / Bigtable

Examples: Apache Cassandra, Google Bigtable, HBase, ScyllaDB.

## 30-Second Interview Answer

Use wide-column stores for huge write throughput and predictable queries by partition key, often with time-ordered data. They are strong for logs, events, messages, IoT telemetry, and activity feeds when the access pattern is known.

## Use When

- The workload is write-heavy and horizontally scalable.
- Queries are known and can be modeled by partition key plus clustering/sort key.
- Data is append-heavy or time-ordered.
- Availability matters more than complex transactions.
- You can denormalize tables per query.

## Avoid When

- You need ad hoc queries, joins, or flexible filtering.
- Strong cross-row transactions are central.
- Query patterns are unknown.
- You cannot choose a good partition key.

## Core Model

Wide-column stores organize data into tables keyed by partition keys and clustering columns. Rows can be sparse and wide. The physical layout is designed for fast writes and efficient reads within a partition.

Common patterns:

- `messages_by_conversation(conversation_id, created_at)`.
- `metrics_by_device(device_id, timestamp)`.
- `events_by_user(user_id, timestamp)`.
- `feed_by_user(user_id, score_or_time)`.

## Query And Indexing

You query by partition key. Clustering columns support ordered reads inside that partition. Secondary indexes are limited and can be expensive at scale.

Interview detail: create one table per important query instead of expecting one normalized table to serve every query.

## Consistency

Cassandra-style systems often allow tunable consistency: choose how many replicas must acknowledge reads/writes. Bigtable-style systems emphasize strong consistency within a single cluster/tablet model but have different replication tradeoffs.

Interview phrase: "For this append-heavy feed, eventual consistency is acceptable; for financial inventory it is not."

## Scaling

Wide-column stores scale by partitioning data across many nodes. They handle large write volumes well if partitions are balanced.

Failure modes:

- Hot partitions from popular keys.
- Very large partitions causing slow reads and compaction pressure.
- Tombstone buildup from deletes.
- Cross-partition queries becoming scatter-gather.
- Repair/compaction consuming IO.

## Data Modeling

Design backwards from queries:

1. Identify partition key.
2. Bound partition size.
3. Choose clustering order.
4. Duplicate data into query-specific tables if needed.
5. Add TTL for expiring time-series/event data.

Hot-key mitigation:

- Bucket by time: `user_id + yyyy_mm`.
- Add shard suffix: `celebrity_id + bucket`.
- Split large tenants by sub-entity.

## Interview Examples

- Chat messages by conversation and timestamp.
- Activity events by user and time.
- IoT readings by device and hour/day.
- High-scale notification inbox.
- Write-heavy logs before moving older data to object storage/OLAP.

## Senior-Level Tradeoffs

- Excellent write scale, but query flexibility is intentionally constrained.
- Denormalized tables are normal; write amplification is the cost of predictable reads.
- Eventual consistency may be acceptable for feeds/logs, but not for core financial invariants.
- Partition design is the architecture, not an implementation detail.

## Common Mistakes

- Designing normalized relational tables in Cassandra/Bigtable.
- Querying without partition key.
- Allowing unbounded partitions.
- Using it for small relational systems where SQL is simpler.
