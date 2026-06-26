# Time-Series - TimescaleDB / InfluxDB

Examples: TimescaleDB, InfluxDB, Prometheus TSDB, VictoriaMetrics, OpenTSDB.

## 30-Second Interview Answer

Use a time-series database when the dominant workload is appending timestamped measurements and querying by time windows, tags, and rollups. It is strongest for metrics, telemetry, monitoring, IoT, and event trends.

## Use When

- Every record has a timestamp and time is the main query dimension.
- Writes are append-heavy.
- Queries ask for recent windows, aggregations, downsampling, or retention tiers.
- Data volume is high but old data can be compressed, downsampled, or expired.
- Tags/labels define filtering dimensions.

## Avoid When

- Data is mutable relational business entities.
- Queries need many joins across normalized tables.
- Timestamp is incidental, not central.
- Cardinality is uncontrolled and can explode.

## Core Model

Time-series records usually contain:

- Timestamp.
- Metric/measurement name.
- Tags/labels for dimensions.
- Fields/values.

Example: CPU metric by host, region, service, and timestamp.

## Query And Indexing

Time-series systems optimize time-window queries:

- Last 5 minutes.
- Average latency by service over 1 hour.
- p95 CPU by host for 7 days.
- Error count grouped by region.

Indexes usually emphasize time plus tags. High-cardinality labels increase memory/storage pressure.

## Consistency

Most monitoring/telemetry systems can tolerate eventual consistency or delayed ingestion. If metrics arrive late, the system may update recent windows or mark late data separately.

Interview detail: dashboards can be slightly stale; alerting needs bounded ingestion lag.

## Scaling

Scaling concerns:

- Write ingestion rate.
- Label cardinality.
- Retention duration.
- Downsampling jobs.
- Query fanout across long time ranges.

Failure modes:

- Cardinality explosion from labels like request ID or user ID.
- Hot series from one metric with huge traffic.
- Expensive long-range queries.
- Late/out-of-order data.
- Retention jobs competing with ingestion.

## Data Modeling

Good tags:

- service
- region
- host
- endpoint
- status_class

Bad tags:

- request_id
- raw_url_with_params
- user_id for high-scale systems
- unbounded error strings

## Interview Examples

- Metrics platform.
- IoT sensor ingestion.
- API latency dashboards.
- Ride location history by time.
- Fraud/risk event trend monitoring.

## Senior-Level Tradeoffs

- Time-series stores compress and expire data well, but are not general OLTP stores.
- More labels improve filtering, but high cardinality can break the system.
- Downsampling saves cost, but loses raw detail.
- Alerting needs predictable ingestion lag and clear missing-data behavior.

## Common Mistakes

- Putting every request ID into labels.
- Using a time-series database for mutable user/order records.
- Keeping raw high-frequency data forever without retention tiers.
- Ignoring late events and clock skew.
