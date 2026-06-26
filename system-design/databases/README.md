# Databases

Human-facing navigation for system design database choices.

## Start Here

- `cheatsheet.md` - quick decision tree and comparison for interviews.
- `individual-analysis/` - detailed learning notes for each database family.
- `AGENTS.md` - agent-facing navigation and edit rules.

## Scope

This topic covers durable source-of-truth stores and interview-useful data stores:

- SQL OLTP and distributed SQL.
- Key-value, document, and wide-column stores.
- Search, time-series, graph, geospatial, OLAP, vector, event-log, and object storage when they are used as data stores or derived indexes.

Caching is intentionally separated into sibling topic `../cache/`.

## Recommended Flow

1. Use `cheatsheet.md` to pick the database family during an interview.
2. Open the matching file in `individual-analysis/` when you need details.
3. If the design needs hot derived reads, invalidation, or Redis-style structures, switch to `../cache/cheatsheet.md`.

## Individual Analysis Files

- `individual-analysis/sql-relational-oltp.md`
- `individual-analysis/sql-distributed-newsql.md`
- `individual-analysis/key-value-dynamodb.md`
- `individual-analysis/document-mongodb.md`
- `individual-analysis/wide-column-cassandra-bigtable.md`
- `individual-analysis/search-elasticsearch-opensearch.md`
- `individual-analysis/time-series-timescaledb-influxdb.md`
- `individual-analysis/graph-neo4j.md`
- `individual-analysis/geospatial-postgis-s2-h3.md`
- `individual-analysis/analytics-olap-clickhouse-bigquery.md`
- `individual-analysis/vector-pinecone-faiss-pgvector.md`
- `individual-analysis/event-log-kafka-kinesis-pulsar.md`
- `individual-analysis/object-storage-s3-gcs.md`
