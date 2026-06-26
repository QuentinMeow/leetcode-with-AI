# System Design Interview Cheatsheet

## Visual Decision Tree

Use this to pick the major architecture component first. Then open the topic-specific cheatsheet for the detailed tree.

```text
System design prompt
|-- Need core source of truth?
|   |-- Strong transactions / joins / constraints
|   |   +-- PostgreSQL, MySQL, MariaDB, SQL Server
|   |-- Strong SQL plus horizontal or multi-region scale
|   |   +-- Google Spanner, CockroachDB, YugabyteDB, TiDB
|   |-- Massive key-based access
|   |   +-- DynamoDB, FoundationDB, RocksDB-backed KV service
|   +-- Flexible aggregate-shaped records
|       +-- MongoDB, Couchbase, Firestore
|-- Need hot-path acceleration?
|   |-- Rich counters / sessions / leaderboards / rate limits
|   |   +-- Redis, Valkey
|   |-- Simple distributed object cache
|   |   +-- Memcached
|   |-- Tiny per-instance low-change data
|   |   +-- Caffeine, Guava Cache, in-process LRU cache
|   +-- Static / media / public global content
|       +-- CloudFront, Cloudflare, Fastly, Akamai
|-- Need async fanout, replay, or CDC?
|   +-- Kafka, Kinesis, Pulsar, Redpanda
|-- Need specialized read/query surface?
|   |-- Full-text / fuzzy / relevance search
|   |   +-- Elasticsearch, OpenSearch, Solr
|   |-- Metrics or time-window telemetry
|   |   +-- TimescaleDB, InfluxDB, Prometheus, VictoriaMetrics
|   |-- Deep relationship traversal
|   |   +-- Neo4j, Neptune, JanusGraph
|   |-- Nearby / radius / geofence queries
|   |   +-- PostGIS, Elasticsearch geo, Redis GEO, S2/H3
|   |-- Analytical scans / dashboards
|   |   +-- ClickHouse, BigQuery, Snowflake, Redshift
|   +-- Semantic similarity / RAG retrieval
|       +-- pgvector, Pinecone, Milvus, Weaviate, FAISS
+-- Need large binary object storage?
    +-- S3, Google Cloud Storage, Azure Blob Storage, MinIO
```

## 60-Second Frame

Start with requirements, not boxes: "The core invariant is X, the hot path is Y, and the expected scale is Z, so I will choose components that protect X while making Y fast enough at Z."

## First 5 Minutes

### Functional Requirements

Name concrete user actions: create post, read feed, upload video, search messages, send payment, track location.

### Non-Functional Requirements

- Latency: p50/p95/p99 target, sync vs async response.
- Scale: DAU, reads/sec, writes/sec, fanout, data size, retention.
- Availability: acceptable downtime, regional failover, degraded mode.
- Consistency: fresh reads, stale reads, transactions, reconciliation.
- Durability: can data be recreated, or must it never be lost?
- Security: authz, encryption, deletion, audit, tenant isolation.

### Invariants

Senior interviewers listen for invariants:

- Payments must not double charge.
- Inventory cannot sell below zero.
- Message ordering matters per conversation.
- Users must not see private content.
- Analytics can lag but should be complete.

## Design In Layers

- API: operations, idempotency, pagination, retries.
- Data model: entities, cardinality, ownership, read paths.
- Storage: source of truth, derived indexes, blobs, analytics.
- Async: queues/streams for fanout, indexing, email, metrics, retries.
- Reliability: replication, failover, backpressure, rate limits, observability.

## Storage Shortcuts

- Relational SQL: transactional consistency, joins, constraints.
- Key-value: huge scale point lookups.
- Document: aggregate-shaped flexible records.
- Wide-column: write-heavy partition-key-first workloads.
- Search index: relevance, fuzzy matching, text/log search.
- Time-series: metric/event series by time window.
- Graph: deep relationship traversal.
- Geospatial: nearby/radius/geofence/location queries.
- OLAP: analytical scans and dashboards.
- Event log: replayable ordered events, CDC, fanout, audit trail.
- Cache: derived, disposable, low-latency hot data.
- Object storage: files, media, backups, large immutable objects.

See `databases/cheatsheet.md` and `cache/cheatsheet.md`.

## Tradeoffs To Say Out Loud

### Correctness Vs Latency

If the invariant is strong, pay for transactions, serializable isolation, compare-and-set, leases, or single-writer ownership. If weak, use eventual consistency, caches, async replication, and reconciliation.

### Normalization Vs Denormalization

Normalize when correctness and many query shapes matter. Denormalize when the read path is known, latency matters, and write amplification plus repair jobs are acceptable.

### Consistency Vs Availability

During partitions, a strongly consistent system may reject writes. An eventually consistent system may accept writes and reconcile later. Say which user experience is acceptable.

### Simplicity Vs Specialization

Start simple. Add a specialized store only when the primary database cannot meet a clear latency, scale, query, or product requirement.

## Capacity Shortcuts

- Request rate: users * sessions/user * actions/session.
- Storage: record size * records/day * retention * replication * indexes.
- Bandwidth: payload size * request rate.
- Cache: working set, not total dataset size.

## Failure Modes To Mention

- Hot partitions or celebrity users.
- Cache stampede, penetration, stale data.
- Search/indexing lag after primary writes.
- Replica lag and read-after-write confusion.
- Rebalancing storms during shard movement.
- Backfills competing with production traffic.
- Exactly-once assumptions without idempotency.
- Schema migration and dual-write inconsistency.
- Regional outage and split-brain risk.

## If You Are Unsure

Use a conservative baseline: relational SQL for core records, object storage for blobs, cache for hot derived reads, queue/event log for async work, search index only if search is core, OLAP only if analytics scans hurt OLTP.

## Closing Structure

- Phase 1: simple architecture that meets correctness.
- Phase 2: add replicas/cache/indexes for measured bottlenecks.
- Phase 3: partition or specialize when volume/access patterns require it.
- Phase 4: multi-region, disaster recovery, observability, and playbooks.
