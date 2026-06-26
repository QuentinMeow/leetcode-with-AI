# Analytics - OLAP Column Stores

Examples: ClickHouse, BigQuery, Snowflake, Redshift, Apache Pinot, Druid.

## 30-Second Interview Answer

Use an OLAP column store when the workload scans large historical datasets for aggregations, dashboards, reports, or exploratory analytics. Keep OLAP separate from OLTP so analytical scans do not hurt user-facing transactions.

## Use When

- Queries aggregate across millions or billions of rows.
- Workload is read-heavy analytics, not per-request transactions.
- Data is append-heavy and historical.
- Dashboards need group-by, filters, rollups, and joins at scale.
- Business intelligence, product analytics, or event analysis is a requirement.

## Avoid When

- The path is user-facing OLTP with small transactional writes.
- You need low-latency point updates and row-level transactions.
- Data volume is small enough for SQL read replicas or materialized views.
- The system cannot tolerate ingestion delay.

## Core Model

OLAP systems often store data column-wise, which improves compression and scan speed for analytical queries that read only selected columns.

Common table types:

- Event fact tables.
- Orders/payments fact tables.
- User/product dimension tables.
- Aggregated rollup tables.

## Query And Indexing

OLAP queries emphasize:

- Filtering by time and dimensions.
- Aggregating metrics.
- Joining facts to dimensions.
- Computing percentiles, funnels, cohorts, and trends.

Partitioning by date/time is common. Clustering/sorting by frequent filters improves pruning.

## Consistency

OLAP is usually eventually consistent from OLTP. Data arrives through:

- Event streams.
- Batch ETL.
- Change data capture.
- Periodic exports from object storage.

Interview detail: state acceptable freshness. "Dashboards can lag by 5 minutes" is usually fine; fraud blocking may not be.

## Scaling

OLAP scales compute and storage for scans. Cost depends heavily on data scanned and query patterns.

Failure modes:

- Expensive unbounded queries.
- Broken ETL causing incomplete reports.
- Late-arriving events.
- Duplicate events without dedupe keys.
- Schema drift from producers.
- Dashboards competing with ad hoc exploration.

## Data Modeling

Common modeling choices:

- Star schema for BI: facts plus dimensions.
- Wide event table for product analytics.
- Pre-aggregated rollups for common dashboards.
- Materialized views for expensive repeated queries.

## Interview Examples

- Product analytics dashboard.
- Sales reporting.
- Fraud/risk offline analysis.
- Ad campaign performance.
- Logs/events long-term analysis.

## Senior-Level Tradeoffs

- OLAP protects OLTP from heavy scans, but introduces data pipeline complexity.
- Freshness can be traded for cost and simplicity.
- Raw events preserve flexibility, but rollups make dashboards cheaper and faster.
- Exactly-once analytics is usually implemented through idempotent event IDs and deduplication, not magical delivery guarantees.

## Common Mistakes

- Running dashboard scans on the production OLTP primary.
- Ignoring event dedupe and late-arriving data.
- Claiming real-time analytics without defining freshness.
- Not separating operational metrics from business analytics.
