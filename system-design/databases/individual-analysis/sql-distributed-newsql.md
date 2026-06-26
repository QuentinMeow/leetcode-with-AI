# SQL - Distributed SQL / NewSQL

Examples: Google Spanner, CockroachDB, YugabyteDB, TiDB.

## 30-Second Interview Answer

Use distributed SQL when you need relational transactions and SQL semantics, but a single-region primary database is not enough for scale, availability, or geographic requirements.

## Use When

- You need ACID transactions and SQL across horizontally distributed data.
- Multi-region availability is a product requirement.
- Data must stay close to users in different geographies.
- The system needs strong consistency for important records but cannot rely on one primary region.
- You are willing to pay operational/cost complexity for correctness at scale.

## Avoid When

- A normal PostgreSQL/MySQL primary plus replicas is enough.
- The workload is simple key-value access at extreme scale.
- Latency budget cannot tolerate distributed consensus for writes.
- The team cannot operate or reason about distributed transactions.

## Core Model

Distributed SQL splits relational data into ranges or shards, replicates them across nodes, and uses consensus protocols to keep replicas consistent. The system still exposes SQL tables, indexes, and transactions.

## Query And Indexing

SQL remains flexible, but physical locality matters more than in a single-node database.

Interview details:

- Primary key design affects range distribution and hot spots.
- Secondary indexes are distributed structures with write amplification.
- Cross-region queries can be slower if they touch far-away replicas.
- Co-locating related rows can reduce distributed transaction cost.

## Consistency And Transactions

Distributed SQL usually offers strong consistency through consensus. This is powerful, but not free.

Tradeoffs:

- Writes may require quorum across replicas.
- Multi-region writes add network latency.
- Large transactions can lock or coordinate across many ranges.
- Clock synchronization or timestamp ordering may be part of the design.

## Scaling

Distributed SQL scales by adding nodes and splitting ranges. It is strongest when data and traffic distribute evenly.

Watch for:

- Hot rows, hot tenants, or monotonic keys.
- Large cross-shard joins.
- Secondary index write amplification.
- Rebalancing pressure during traffic spikes.

## Replication And Failover

Replicas are placed across zones or regions. If a node fails, quorum can keep serving traffic. If a region fails, the system may continue if enough replicas remain.

Interview reasoning: explain the availability target. For example, "I would place replicas across three zones in one region first. I would move to multi-region only if the product requires regional disaster recovery or local reads."

## Data Modeling

Use relational modeling, but be more deliberate about locality.

Patterns:

- Use tenant or region in keys when it matches access patterns.
- Avoid sequential IDs if they create hot ranges.
- Keep transactions small.
- Use async workflows for operations that do not need one distributed transaction.

## Interview Examples

- Global inventory reservation with strong consistency.
- Multi-region financial ledger.
- SaaS control plane that must survive zone or region failure.
- Metadata database for a globally distributed platform.

## Senior-Level Tradeoffs

- You preserve SQL correctness, but pay in latency, cost, and operational complexity.
- Multi-region strong consistency protects invariants, but user-visible writes may be slower.
- It can remove manual sharding burden, but does not remove data modeling discipline.
- It is often overkill for early-stage systems.

## Common Mistakes

- Choosing distributed SQL before proving normal SQL is insufficient.
- Ignoring write latency across regions.
- Assuming horizontal scale fixes bad indexes or hot keys.
- Running huge analytical queries on the transactional cluster.
