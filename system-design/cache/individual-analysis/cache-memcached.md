# Cache - Memcached

Examples: Memcached, managed Memcached clusters.

## 30-Second Interview Answer

Use Memcached when you need a simple distributed cache for string/blob values and do not need Redis data structures, persistence, streams, sorted sets, or complex atomic operations. It is a strong interview choice for straightforward cache-aside object caching.

## Use When

- You need a simple key-value cache.
- Values are disposable and can be recomputed or reloaded from a database.
- Cache-aside is enough.
- You want predictable simple behavior and low operational surface.
- Data structures like sorted sets, hashes, or streams are unnecessary.

## Avoid When

- You need counters with richer atomic behavior, sorted sets, leaderboards, queues, or rate limit scripts.
- You need persistence or recovery semantics.
- You need complex invalidation logic inside the cache.
- You need multi-key operations with stronger semantics.

## Core Model

Memcached is an in-memory distributed key-value cache. It stores opaque byte values by key and evicts data under memory pressure.

Common values:

- Serialized user/profile object.
- Product detail JSON.
- Rendered HTML fragment.
- Permission or feature flag snapshot.

## Query And Indexing

Memcached only supports key lookup. There are no secondary indexes, scans, or rich queries.

Interview detail: the application must already know the key. If you need filtering/search, use the database or search index, then cache final results carefully.

## Consistency

Memcached is usually used with cache-aside:

1. Read cache.
2. On miss, read database.
3. Populate cache with TTL.
4. On write, update database and delete the cache key.

It is eventually consistent by design. The database remains authoritative.

## Scaling

Memcached scales by adding nodes and distributing keys through client-side hashing or a proxy layer.

Failure modes:

- Node loss causes many cache misses.
- Rehashing can invalidate a large part of the cache.
- Hot keys overload one node.
- Eviction under memory pressure lowers hit rate.
- Large values waste memory and network bandwidth.

## Data Modeling

Use clear, versionable keys:

- `user:v1:{id}`
- `product:v3:{id}`
- `page:v2:{route}:{locale}`

Keep values small. Include schema version in keys when deployments may change serialized shape.

## Interview Examples

- Cache product detail pages.
- Cache user profile snapshots.
- Cache rendered feed/page fragments.
- Cache expensive DB query results with short TTL.

## Senior-Level Tradeoffs

- Memcached is simpler than Redis for pure cache-aside workloads, but less flexible.
- Losing Memcached data should only degrade latency, not correctness.
- Client-side sharding is simple but can cause churn when nodes change.
- It is a cache, not a coordination system or durable store.

## Common Mistakes

- Choosing Redis by default when simple key-value caching is enough.
- Treating Memcached values as durable.
- Forgetting warmup behavior after node loss.
- Caching huge values or unbounded query results.
