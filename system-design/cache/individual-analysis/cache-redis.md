# Cache - Redis

Examples: Redis, Memcached, managed Redis-compatible stores.

## 30-Second Interview Answer

Use Redis as a fast, usually derived data layer for hot reads, counters, sessions, rate limits, queues/light streams, leaderboards, and ephemeral coordination. Do not make it the primary source of truth unless you explicitly design persistence, recovery, and data loss behavior.

## Use When

- You need sub-millisecond or low-millisecond access.
- Data is hot, repeated, expensive to compute, or safe to recreate.
- You need TTL-based expiration.
- You need atomic counters, sets, sorted sets, or rate limits.
- You need presence/session state where slight loss is acceptable.

## Avoid When

- Data must be durable and Redis persistence is not part of the design.
- Dataset is much larger than memory budget.
- Complex querying, joins, or analytical scans are needed.
- Strong distributed locking correctness is required and not carefully designed.

## Core Model

Redis is an in-memory data structure server. Useful structures:

- String: cache value, session blob, feature flag.
- Hash: small object fields.
- Set: membership, dedupe.
- Sorted set: leaderboard, time-ranked items.
- List/stream: lightweight queue or event stream.
- HyperLogLog/bitmap: approximate counts and flags.

## Cache Patterns

### Cache-Aside

Application reads cache first. On miss, read database, then populate cache. Common and simple.

### Write-Through

Writes go through cache and database together. More consistent, but more complex.

### Write-Behind

Writes hit cache first and flush later. Fast but risky for data loss and ordering.

### Read-Through

Cache layer knows how to load from database. Less common in hand-built system design answers.

## Consistency

Redis usually serves stale or derived data. Explain acceptable staleness and invalidation:

- TTL for eventual refresh.
- Explicit delete/update on write.
- Versioned keys to avoid old overwrites.
- Event-driven invalidation from database changes.

## Scaling

Redis scales with replication, clustering, sharding, and memory sizing.

Failure modes:

- Cache stampede: many requests miss at once.
- Cache penetration: repeated misses for nonexistent keys.
- Hot keys: one key receives too much traffic.
- Eviction surprises: memory pressure removes needed data.
- Replica failover causing lost recent writes.

## Data Modeling

Use Redis for small, frequently accessed values. Avoid giant keys and huge values.

Patterns:

- `user_profile:{id}` with TTL.
- `rate:{user_id}:{minute}` counters.
- `leaderboard:{game_id}` sorted set.
- `presence:{user_id}` ephemeral status.
- `feed:{user_id}` recent item IDs.

## Interview Examples

- Rate limiting API requests.
- Caching product details or user profiles.
- Session storage.
- Top-N leaderboard.
- Recent feed cache.
- Distributed lock with fencing tokens if correctness matters.

## Senior-Level Tradeoffs

- Cache improves latency and protects databases, but adds invalidation complexity.
- TTL is simple, but creates staleness.
- Explicit invalidation is fresher, but failure-prone.
- Redis locks are easy to misuse; for hard correctness prefer database constraints or lease/fencing design.

## Common Mistakes

- Making cache authoritative without durability discussion.
- Forgetting cache stampede mitigation.
- Caching data that changes too frequently to be useful.
- Ignoring memory limits, eviction policy, and hot keys.
