# SQL - Relational OLTP

Examples: PostgreSQL, MySQL, MariaDB, SQL Server, Oracle.

## 30-Second Interview Answer

Use relational SQL as the default source of truth when correctness, transactions, constraints, joins, and query flexibility matter. It is the safest baseline for users, orders, payments, inventory, permissions, and metadata.

## Use When

- You need ACID transactions across multiple rows/tables.
- You need uniqueness, foreign keys, constraints, or strong data integrity.
- Query patterns are not fully known up front.
- You need joins, secondary indexes, filtering, sorting, and pagination.
- Data volume fits a single primary plus replicas, or can be partitioned carefully.

## Avoid When

- The main workload is massive simple key-value reads/writes.
- The schema is only append-only metrics/events at very high write volume.
- The system primarily needs fuzzy text search, vector similarity, graph traversal, or analytical scans.
- You cannot tolerate a single-writer primary and are not ready for sharding or distributed SQL.

## Core Model

Relational databases store normalized tables with schemas, constraints, indexes, and SQL queries. They are optimized for OLTP: many small reads/writes with correctness.

Common entities:

- Users, organizations, permissions.
- Orders, payments, invoices.
- Inventory and reservations.
- Posts, comments, likes, metadata.
- Application configuration and audit records.

## Query And Indexing

Primary keys make point lookups fast. Secondary indexes speed filters and sorts but increase write cost and storage. Composite indexes should match query predicates in order.

Interview detail: say which queries need indexes. For example, `orders(user_id, created_at)` supports "show a user's recent orders"; `inventory(product_id)` supports stock checks.

## Transactions And Isolation

Use transactions when multiple writes must succeed or fail together. Isolation levels matter:

- Read committed: common default; avoids dirty reads.
- Repeatable read: stable reads within a transaction.
- Serializable: strongest, safest for hard invariants, but can reduce throughput.

Senior signal: enforce invariants in the database when possible. Use unique constraints, check constraints, foreign keys, and transactional updates instead of relying only on service code.

## Scaling

### Read Scaling

- Add read replicas.
- Cache hot reads.
- Add covering indexes or materialized views.
- Split analytical scans into OLAP.

### Write Scaling

- Reduce unnecessary indexes.
- Batch writes.
- Partition/shard by tenant, user, region, or entity ID.
- Move append-only events to a log or specialized store.

### Sharding

Sharding adds routing, cross-shard queries, rebalancing, hot-shard handling, and operational complexity. In interviews, use it only after simpler scaling options are exhausted or the prompt requires it.

## Replication And Failover

Common setup: one primary accepts writes; replicas serve reads. Replication can be synchronous or asynchronous.

Failure modes:

- Replica lag causes stale reads.
- Failover can briefly reject writes.
- Split brain can corrupt data if two primaries accept writes.
- Backups must be tested with restore drills.

## Schema Design

Normalize for correctness first. Denormalize when a read path is hot and stable. If denormalizing, explain how you repair derived data after failed updates.

Patterns:

- Use generated IDs or UUIDs depending on sharding and ordering needs.
- Use status columns for workflows, but avoid ambiguous state machines.
- Use soft deletes only if audit/recovery needs them.
- Use outbox pattern for reliable async side effects after transactions.

## Interview Examples

- Banking ledger: SQL for accounts, transactions, idempotency keys, and constraints.
- E-commerce orders: SQL for inventory reservation, payment state, and order rows.
- Permission system: SQL for users, groups, roles, and auditability.
- SaaS app: SQL by tenant; shard later if tenants grow.

## Senior-Level Tradeoffs

- SQL gives correctness and flexibility, but global write scaling is harder than key-value stores.
- Strong constraints simplify application logic, but migrations require care.
- Indexes speed reads, but slow writes and complicate backfills.
- Read replicas improve read capacity, but stale reads can violate user expectations.

## Common Mistakes

- Saying SQL "does not scale" without discussing replicas, partitioning, and sharding.
- Using NoSQL for transactional data just because traffic is high.
- Forgetting idempotency keys for retries in payments/orders.
- Letting analytics dashboards run expensive scans on the OLTP primary.
