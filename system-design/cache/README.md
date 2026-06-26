# Cache

Human-facing navigation for system design cache strategy.

## Start Here

- `cheatsheet.md` - quick interview reference for caching decisions.
- `individual-analysis/` - detailed notes for cache systems and patterns.
- `AGENTS.md` - agent-facing navigation and edit rules.

## Scope

This topic covers:

- Cache placement and cache-aside/write-through/write-behind patterns.
- Redis, Memcached, local/in-process cache, and CDN/edge/client cache.
- TTLs, invalidation, consistency, stampedes, hot keys, and eviction.
- Cache use cases such as sessions, rate limiting, counters, leaderboards, presence, and hot read models.

Durable database selection belongs in sibling topic `../databases/`.

## Individual Analysis Files

- `individual-analysis/cache-redis.md`
- `individual-analysis/cache-memcached.md`
- `individual-analysis/cache-local-in-process.md`
- `individual-analysis/cache-cdn-edge-client.md`
