# Cache - Local / In-Process

Examples: Caffeine, Guava Cache, LRU maps, process memory caches, language runtime caches.

## 30-Second Interview Answer

Use a local in-process cache for tiny, very hot, low-change data that every service instance can keep independently. It is the fastest cache, but hardest to invalidate globally and disappears when the process restarts.

## Use When

- Data is read extremely often and changes rarely.
- Data is small enough to fit in every instance's memory.
- Staleness is acceptable for a short bounded window.
- You need to avoid a network hop to Redis/Memcached.
- The value is configuration, metadata, permissions snapshot, or static lookup data.

## Avoid When

- Cache values must be globally consistent immediately.
- The working set is large.
- Updates are frequent and invalidation must be instant.
- You need shared counters, sessions, presence, or leaderboards.
- Memory pressure could affect the application process.

## Core Model

Each application instance keeps its own cache in process memory. Common eviction policies include LRU, LFU, size-based eviction, and TTL expiration.

## Query And Indexing

Lookups are direct by key in local memory. There are no remote calls and no cross-instance coordination by default.

## Consistency

Local caches are usually eventually consistent. Options:

- Short TTL.
- Versioned config.
- Push invalidation events from a control plane.
- Periodic refresh.
- Bypass cache for correctness-sensitive operations.

Interview detail: always say how many seconds/minutes of staleness are acceptable.

## Scaling

Local caches scale naturally with service instances, but total memory usage is duplicated across instances.

Failure modes:

- Inconsistent values across instances.
- Thundering herd after deploy/restart.
- Memory leak or unbounded cache growth.
- Old values survive until TTL after a write.
- Different versions of services interpret cached data differently.

## Data Modeling

Good candidates:

- Country/region metadata.
- Feature flag snapshots.
- Public key/JWKS cache.
- Tenant configuration with TTL.
- Product category taxonomy.

Bad candidates:

- Inventory counts.
- Account balances.
- User sessions that must survive restarts.
- Large feed or search result caches.

## Interview Examples

- Cache authorization public keys for JWT verification.
- Cache feature flag/config snapshots.
- Cache small reference data used on every request.
- Cache compiled templates or expensive regex/routing metadata.

## Senior-Level Tradeoffs

- Fastest read path, but no global invalidation guarantee.
- Good for low-change data, risky for user-specific mutable data.
- Reduces load on distributed cache, but duplicates memory across instances.
- Deployment and restart behavior matter because every instance can cold start at once.

## Common Mistakes

- Using local cache for data that needs immediate global freshness.
- No max size or eviction policy.
- Forgetting cold-start storms after deployments.
- Letting cache memory compete with application heap.
