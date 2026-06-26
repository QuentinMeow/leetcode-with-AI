# Cache - CDN / Edge / Client

Examples: CloudFront, Cloudflare, Fastly, Akamai, browser HTTP cache, mobile client cache.

## 30-Second Interview Answer

Use CDN, edge, or client caching when content can be served close to users and is public, static, immutable, or safely stale. It is the best interview answer for images, videos, static assets, public pages, and globally read-heavy content.

## Use When

- Content is static or changes infrequently.
- Data is public or can be safely cached per user with strict headers.
- Latency and bandwidth matter globally.
- Large files or media are served repeatedly.
- Origin/database load should be reduced.

## Avoid When

- Content is private and cache keys/auth headers are not carefully designed.
- Every request needs fresh authoritative state.
- Response varies by many headers/cookies and creates cache fragmentation.
- Invalidating globally on every write is required.

## Core Model

CDNs cache responses or objects at edge locations near users. Browsers/mobile apps can also cache responses locally based on HTTP headers or app logic.

Common controls:

- `Cache-Control`
- `ETag`
- `Last-Modified`
- `max-age`
- `s-maxage`
- immutable asset names
- signed URLs/cookies

## Query And Indexing

CDN cache keys usually include URL plus selected headers/query parameters. Bad cache-key design either leaks data or destroys hit rate.

Interview detail: public static assets should use content-hashed URLs and long TTLs. Mutable content should use shorter TTLs, revalidation, or purge/versioning.

## Consistency

CDN and client caches are intentionally stale within a configured window.

Options:

- Content-hashed filenames for immutable assets.
- Versioned URLs for media variants.
- Short TTL plus revalidation for mutable public content.
- Explicit purge for urgent invalidation.
- Signed URLs for protected objects.

## Scaling

CDNs reduce origin bandwidth and latency by serving from edge locations.

Failure modes:

- Serving stale or deleted content.
- Private data cached publicly.
- Cache fragmentation from cookies/query strings.
- Purge delays during invalidation.
- Origin overload when cache expires globally.

## Data Modeling

Good candidates:

- Images, videos, thumbnails.
- JavaScript/CSS bundles.
- Public profile images.
- Public article pages.
- Downloadable exports with signed URLs.

Poor candidates:

- Account balance.
- Checkout state.
- Private inbox pages without strict per-user cache controls.
- Inventory that must be exact.

## Interview Examples

- Video streaming service: object storage plus CDN.
- Photo sharing app: thumbnails and originals at edge.
- News site: public article cache with purge/revalidation.
- Static frontend assets with content hashes.
- File download service with signed URLs.

## Senior-Level Tradeoffs

- Edge caching gives the largest global latency and bandwidth win, but invalidation can be slow.
- Long TTLs are safe with immutable versioned assets.
- Personalized content needs careful `Vary`, private cache headers, or no CDN caching.
- CDN protects origin, but global expiry can create origin traffic spikes.

## Common Mistakes

- Caching private responses publicly.
- Forgetting cache key variation by auth/cookie/language.
- Using short TTLs for immutable assets.
- Depending on instant global purge for correctness.
