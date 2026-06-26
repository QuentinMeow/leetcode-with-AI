# Database Choices Cheatsheet

## Visual Decision Tree

Use this first under interview pressure. Every leaf names concrete choices; open the matching `individual-analysis/` file for details.

```text
Data/storage requirement
|-- Is this the durable source of truth?
|   |-- Needs transactions, joins, constraints, flexible queries?
|   |   |-- Single-region or moderate scale
|   |   |   +-- PostgreSQL, MySQL, MariaDB, SQL Server
|   |   +-- Horizontal or multi-region SQL scale
|   |       +-- Google Spanner, CockroachDB, YugabyteDB, TiDB
|   |-- Mostly point/range lookup by key at huge scale?
|   |   +-- DynamoDB, FoundationDB, RocksDB-backed KV service
|   |-- Aggregate-shaped flexible documents?
|   |   +-- MongoDB, Couchbase, Firestore
|   +-- Write-heavy partition-key-first events/messages/telemetry?
|       +-- Cassandra, Bigtable, ScyllaDB, HBase
|-- Is this a derived query/index surface?
|   |-- Full-text, fuzzy search, ranking, faceting, logs?
|   |   +-- Elasticsearch, OpenSearch, Solr
|   |-- Time-window metrics, telemetry, retention, downsampling?
|   |   +-- TimescaleDB, InfluxDB, Prometheus, VictoriaMetrics
|   |-- Deep relationship traversal or graph algorithms?
|   |   +-- Neo4j, Amazon Neptune, JanusGraph, TigerGraph
|   |-- Nearby, radius, geofence, map viewport?
|   |   +-- PostGIS, Elasticsearch geo, Redis GEO, S2/H3/geohash
|   |-- Large historical analytical scans and dashboards?
|   |   +-- ClickHouse, BigQuery, Snowflake, Redshift, Pinot, Druid
|   +-- Semantic similarity or RAG retrieval?
|       +-- pgvector, Pinecone, Milvus, Weaviate, FAISS
|-- Is this an append-only event backbone?
|   +-- Kafka, Kinesis, Pulsar, Redpanda, EventStoreDB
|-- Is this large immutable/file/blob data?
|   +-- S3, Google Cloud Storage, Azure Blob Storage, MinIO
+-- Is this low-latency hot derived data?
    +-- Use cache tree: Redis/Valkey, Memcached, Caffeine/Guava, CloudFront/Cloudflare/Fastly
```

## Rendered Decision Diagram

If your Markdown preview supports Mermaid, this renders as a flowchart.

```mermaid
flowchart TD
    start["Data/storage requirement"] --> truth{"Durable source of truth?"}
    truth -->|Yes| tx{"Transactions / joins / constraints?"}
    tx -->|Single-region or moderate scale| sql["PostgreSQL / MySQL / MariaDB / SQL Server"]
    tx -->|Horizontal or multi-region SQL| dsql["Google Spanner / CockroachDB / YugabyteDB / TiDB"]
    truth -->|Yes| key{"Mostly key-based access?"}
    key --> kv["DynamoDB / FoundationDB / RocksDB-backed KV"]
    truth -->|Yes| doc{"Aggregate-shaped flexible records?"}
    doc --> docs["MongoDB / Couchbase / Firestore"]
    truth -->|Yes| wide{"Write-heavy partition-key-first?"}
    wide --> wc["Cassandra / Bigtable / ScyllaDB / HBase"]

    truth -->|No| derived{"Derived query or index surface?"}
    derived --> search{"Text search / fuzzy / ranking?"}
    search --> es["Elasticsearch / OpenSearch / Solr"]
    derived --> ts{"Time-window metrics or telemetry?"}
    ts --> tsdb["TimescaleDB / InfluxDB / Prometheus / VictoriaMetrics"]
    derived --> graph{"Deep relationship traversal?"}
    graph --> gdb["Neo4j / Neptune / JanusGraph / TigerGraph"]
    derived --> geo{"Nearby / radius / geofence?"}
    geo --> geodb["PostGIS / Elasticsearch geo / Redis GEO / S2-H3"]
    derived --> olap{"Analytical scans and dashboards?"}
    olap --> warehouse["ClickHouse / BigQuery / Snowflake / Redshift"]
    derived --> vector{"Semantic similarity or RAG?"}
    vector --> vdb["pgvector / Pinecone / Milvus / Weaviate / FAISS"]

    truth -->|No| event{"Replay / CDC / audit / fanout?"}
    event --> log["Kafka / Kinesis / Pulsar / Redpanda / EventStoreDB"]
    truth -->|No| blob{"Large files or immutable blobs?"}
    blob --> obj["S3 / Google Cloud Storage / Azure Blob / MinIO"]
    truth -->|No| hot{"Hot derived low-latency reads?"}
    hot --> cache["Redis / Valkey / Memcached / CloudFront / Cloudflare"]
```

## 30-Second Rule

Choose by access pattern and correctness invariant, not by brand name.

Say this in interviews: "The source of truth needs X consistency and Y query pattern. I will use Z because it optimizes for that, and I will avoid overloading it with search, analytics, or blob storage."

## Decision Tree

### 1. Is This The Durable Source Of Truth?

If yes, ask whether the core invariant needs transactions.

- Multi-row constraints, joins, unique constraints, foreign keys, or flexible ad hoc queries: choose `individual-analysis/sql-relational-oltp.md`.
- SQL transactions plus global or multi-region scale: choose `individual-analysis/sql-distributed-newsql.md`.
- Simple primary-key access at massive scale: choose `individual-analysis/key-value-dynamodb.md`.
- Flexible aggregate-shaped records and queries mostly within one aggregate: choose `individual-analysis/document-mongodb.md`.
- Enormous write throughput with partition-key-first queries: choose `individual-analysis/wide-column-cassandra-bigtable.md`.

If no, ask what derived access pattern you are optimizing:

- Text search, fuzzy matching, relevance ranking, log search: choose `individual-analysis/search-elasticsearch-opensearch.md`.
- Time-window metrics, events, or telemetry: choose `individual-analysis/time-series-timescaledb-influxdb.md`.
- Deep relationship traversal: choose `individual-analysis/graph-neo4j.md`.
- Nearby/location/radius/geofence queries: choose `individual-analysis/geospatial-postgis-s2-h3.md`.
- Analytical scans, dashboards, aggregations across large history: choose `individual-analysis/analytics-olap-clickhouse-bigquery.md`.
- Semantic similarity search or RAG retrieval: choose `individual-analysis/vector-pinecone-faiss-pgvector.md`.
- Replayable ordered events, CDC, audit trail, fanout to consumers: choose `individual-analysis/event-log-kafka-kinesis-pulsar.md`.
- Large files, media, backups: choose `individual-analysis/object-storage-s3-gcs.md`.
- Low-latency hot derived reads, counters, sessions, rate limits: use `../cache/cheatsheet.md`.

### 2. What Is The Hottest Query?

- Point lookup by ID: key-value or relational primary key.
- Lookup by many filters/sorts: relational first; search if relevance/fuzzy/full-text dominates.
- Append events and query recent windows: time-series or wide-column.
- Scan billions of rows for aggregates: OLAP column store.
- Traverse friends-of-friends or fraud rings: graph.
- Find nearby objects or geofence matches: geospatial index.
- Find similar embeddings: vector store.
- Replay events or feed many consumers: event log.

### 3. What Is The Hardest Correctness Requirement?

- Money, inventory, permissions, uniqueness: relational SQL or distributed SQL.
- Idempotent event ingestion: append log plus dedupe key; wide-column/time-series can work.
- User-visible read-after-write: primary reads, session consistency, or strongly consistent reads.
- Derived search/feed/analytics: eventual consistency is usually acceptable if you explain lag and rebuilds.

### 4. What Is The Main Scale Pressure?

- Read heavy: read replicas, cache, CDN, materialized views.
- Write heavy: partition by stable key, batch writes, reduce secondary indexes.
- Hot users/keys: shard hot entities, isolate celebrities, use per-key limits.
- Large blobs: object storage, signed URLs, CDN.
- Analytics pressure: move scans off OLTP into OLAP through CDC/events.

## Comparison Matrix

| Choice | Best For | Avoid When | Interview Reasoning |
|---|---|---|---|
| Relational SQL | Core OLTP, transactions, joins, constraints, moderate scale | Massive write scale with simple access, schema-free event firehose | "Correctness and query flexibility matter more than extreme horizontal write scale." |
| Distributed SQL | SQL correctness with horizontal or multi-region scale | Simple app that one regional SQL database can handle | "I need relational semantics but cannot fit the availability/scale target into one primary." |
| Key-Value | Massive point reads/writes by key, carts, sessions, device state | Ad hoc queries, joins, multi-entity transactions | "The access pattern is key-based, so I trade query flexibility for predictable scale and latency." |
| Document | Aggregate records, flexible schema, catalog/profile documents | Cross-document transactions and relational reporting dominate | "Most reads/writes touch one aggregate, so documents reduce joins and allow schema evolution." |
| Wide-Column | High write throughput, time-ordered partition queries, IoT/events | Unknown query patterns or small relational data | "I know the partition key and query windows, so I can model for write scale and predictable reads." |
| Search Index | Full-text, fuzzy search, ranking, faceting, logs | Strongly consistent source-of-truth updates | "Search is a derived index; I will accept indexing lag and rebuild it from primary data." |
| Time-Series | Metrics, telemetry, time-window rollups, retention policies | Mutable relational entities and broad joins | "The dominant dimension is time, so retention, downsampling, compression, and time indexes matter." |
| Graph | Relationship traversal, recommendations, fraud rings, dependencies | Simple one-hop relationships or basic filters | "Traversal depth is the product feature; graph storage makes the query direct instead of join-heavy." |
| Geospatial Index | Nearby search, maps, drivers, geofences, location filters | Location is only display metadata | "Generate candidates with a spatial index, then exact-filter and rank by distance/business rules." |
| OLAP Column Store | Dashboards, BI, aggregates over large history | User-facing per-request transactions | "Analytical scans should not compete with OLTP; columnar storage optimizes scans and compression." |
| Vector Store | Semantic retrieval, similarity search, RAG | Exact filters/transactions without semantic matching | "The key operation is nearest-neighbor search over embeddings, usually as a derived index." |
| Event Log | CDC, replay, fanout, audit trail, event sourcing | Random point lookups or relational constraints | "The log preserves ordered facts for consumers and rebuilds; queryable views live elsewhere." |
| Object Storage | Images, videos, logs, backups, large immutable objects | Small transactional rows that need indexes and updates | "Blobs do not belong in OLTP; store metadata in DB and bytes in object storage plus CDN." |

## Common Combinations

### Social Feed

- Source of truth: relational SQL or key-value/document for posts and follows.
- Fanout: queue/stream.
- Feed cache: use `../cache/cheatsheet.md`.
- Search: search index if searching posts/users.
- Media: object storage plus CDN.

### Chat

- Source of truth: relational SQL for users/conversations; wide-column or key-value for messages by conversation/time at scale.
- Ordering: partition by conversation, sequence numbers.
- Search: derived search index.
- Recent messages/presence: use `../cache/cheatsheet.md`.

### E-Commerce

- Source of truth: relational SQL for orders, inventory, payments, users.
- Catalog: relational or document, depending on attribute flexibility.
- Search: Elasticsearch/OpenSearch for discovery.
- Product detail hot path: cache with careful staleness.
- Analytics: OLAP for sales dashboards.

### Metrics Platform

- Ingestion: queue/stream.
- Storage: time-series or wide-column.
- Rollups: background jobs and downsampling.
- Dashboards: time-series queries or OLAP for long-range analytics.

### Ride Sharing / Nearby Search

- Source of truth: relational SQL for users, trips, payments.
- Live locations: geospatial cache/index with TTL.
- Nearby matching: geospatial cell/index candidate generation plus exact distance ranking.
- History analytics: time-series or OLAP.

### Event-Driven System

- Source of truth: SQL or another durable primary store.
- Change stream: event log for CDC/fanout/replay.
- Derived views: search, cache, OLAP, or materialized read models.
- Correctness: idempotent consumers, dedupe keys, and offset tracking.

### Recommendation / AI Search

- Source of truth: relational/document for entities.
- Feature/events: OLAP or streaming pipeline.
- Similarity: vector store.
- Exact filters: keep metadata filters in vector index or join against primary/search carefully.

## Interview Phrases That Land Well

- "The search index is eventually consistent and can be rebuilt from the source of truth."
- "The partition key must match the hottest query, or this store will not scale."
- "This requirement needs a transaction, so I will keep it in SQL even if other parts are denormalized."
- "I would start with one primary data store and add specialized derived stores only after a measured bottleneck or a clear access pattern."
- "The tradeoff is write amplification: every extra index/materialized view makes reads faster but writes and backfills harder."

## Red Flags

- Choosing NoSQL only because "it scales".
- Storing large images/videos directly in relational rows.
- Using search as the source of truth.
- Adding graph/vector/time-series databases when SQL plus an index would be enough.
- Ignoring partition keys, hot keys, and secondary index costs.
- Claiming exactly-once distributed writes without explaining idempotency and dedupe.
