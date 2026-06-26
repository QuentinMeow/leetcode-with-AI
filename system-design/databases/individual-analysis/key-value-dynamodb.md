# Key-Value - DynamoDB / Key-Based Stores

Examples: Amazon DynamoDB, FoundationDB key-value layer, RocksDB-based services, persistent KV systems.

## 30-Second Interview Answer

Use a key-value store when the dominant access pattern is predictable point lookup or range lookup by key at very high scale. You trade query flexibility for low latency, high availability, and horizontal scale.

## Use When

- Reads and writes are mostly by primary key.
- You can design partition keys from known access patterns.
- You need very high throughput with predictable latency.
- Records are independent or have limited transactional coupling.
- The workload is carts, device state, feature flags, user settings, feed entries, or idempotency keys.

## Avoid When

- You need joins or ad hoc filtering.
- You need many secondary query patterns.
- Cross-item transactions are central to correctness.
- Hot keys are unavoidable and cannot be split.

## Core Model

Data is stored as key -> value. Some systems add sort keys, secondary indexes, conditional writes, TTL, streams, and limited transactions.

Common key shapes:

- `USER#123` -> profile/settings.
- `CART#user_id` -> cart state.
- `CONVERSATION#id`, sort key `MESSAGE#timestamp`.
- `IDEMPOTENCY#request_id` -> retry protection.

## Query And Indexing

The primary key is the design. Secondary indexes help, but they are not free and often have consistency/throughput limits.

Interview detail: state every access pattern before schema design. In DynamoDB-style modeling, you often denormalize into multiple items to support multiple reads.

## Consistency

Many key-value stores offer eventual reads by default, with optional strong reads or conditional writes. Use conditional writes for compare-and-set patterns like "create if not exists" or "update if version matches".

## Scaling

Key-value stores partition by hash/range of key. They scale well when keys are high-cardinality and traffic is evenly distributed.

Failure modes:

- Hot partition from celebrity users or popular items.
- Monotonic keys causing write concentration.
- Secondary index hot spots.
- Large values increasing latency and cost.

## Data Modeling

Design from access patterns:

1. List the queries.
2. Choose partition key and sort key.
3. Decide whether denormalized duplicate items are acceptable.
4. Add TTL for expiring data.
5. Use streams/CDC to update derived views.

## Interview Examples

- Shopping cart by user ID.
- Device latest state by device ID.
- Idempotency keys for retries.
- Feed inbox entries by user and timestamp.
- User settings or feature flags.

## Senior-Level Tradeoffs

- Excellent scale and latency for known queries, weak flexibility for unknown queries.
- Denormalization speeds reads, but creates write amplification and repair needs.
- Conditional writes can enforce simple invariants, but complex relational constraints are harder.
- Hot-key mitigation must be explicit: bucketing, salting, per-entity sharding, or special handling for celebrities.

## Common Mistakes

- Treating a key-value store like a relational database.
- Adding many global secondary indexes and expecting unlimited write scale.
- Forgetting that partition-key choice determines scalability.
- Storing large blobs instead of using object storage.
