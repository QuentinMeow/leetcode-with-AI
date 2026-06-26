# Geospatial - PostGIS / S2 / H3 / Geo Indexes

Examples: PostGIS, Elasticsearch geo queries, Redis GEO, MongoDB geospatial indexes, S2, H3, Geohash.

## 30-Second Interview Answer

Use geospatial indexing when the core query is location-based: nearby drivers, restaurants, places, geofences, delivery zones, or map search. This is usually an indexing strategy layered on SQL, search, key-value, or cache, not always a standalone database.

## Use When

- Users query "near me" or within a radius/bounding box.
- You need nearest-neighbor lookup by latitude/longitude.
- Geofencing or region containment matters.
- Location updates are frequent and reads need low latency.
- Map/search ranking combines distance with filters.

## Avoid When

- Location is only metadata and not queried spatially.
- A simple city/region filter is enough.
- Exact transactional correctness is more important than nearest lookup.
- The system cannot tolerate approximate candidate generation.

## Core Model

Geospatial systems index latitude/longitude points, shapes, or cells. Common approaches:

- Spatial indexes in SQL: PostGIS GiST/SP-GiST/R-tree-like indexes.
- Search geo index: combine distance with text/filter queries.
- Grid/cell index: map coordinates to S2/H3/geohash cells.
- Cache geo set: keep hot live locations in Redis GEO or similar.

## Query And Indexing

Common queries:

- Find points within radius.
- Find points inside polygon.
- Find nearest K objects.
- Find objects in bounding box.
- Match moving users/drivers to nearby requests.

Interview detail: for high-scale nearby queries, first generate candidates by cell/bounding box, then compute exact distance and rank.

## Consistency

Geospatial freshness depends on the product:

- Restaurant/place search can be eventually consistent.
- Driver location matching needs fresh but not perfectly durable updates.
- Trip/payment source of truth still belongs in a durable database.

Use TTLs for moving locations so stale drivers/devices disappear.

## Scaling

Scaling concerns:

- Hot dense areas like airports or downtown.
- Frequent moving-object updates.
- Uneven population density by cell.
- Large-radius queries that scan many cells.
- Combining geo filters with text/category/ranking filters.

Failure modes:

- Too-large cells return too many candidates.
- Too-small cells require many lookups.
- Edge cases near cell boundaries miss candidates.
- Stale moving locations cause bad matches.
- Hot cells overload one partition.

## Data Modeling

Static places:

- Store canonical place/business in SQL/document store.
- Add geospatial index for coordinates and polygons.
- Use search index if text/category relevance matters.

Moving objects:

- Store durable entity state elsewhere.
- Store latest location in low-latency geo/cache store with TTL.
- Append history to time-series/wide-column/OLAP if needed.

## Interview Examples

- Find nearby drivers for ride sharing.
- Find restaurants within 5 km with filters.
- Track delivery courier location.
- Geofence alerts when a device enters a region.
- Search hotels by map viewport.

## Senior-Level Tradeoffs

- Geo is often an index/access pattern, not the source of truth.
- Approximate cell lookup is fast, but needs exact distance filtering after candidate generation.
- Fresh moving locations belong in a hot low-latency store; durable trip/order data belongs elsewhere.
- Dense areas require hot-cell mitigation through smaller cells, sharding, or load-aware matching.

## Common Mistakes

- Scanning all rows and computing distance without an index.
- Forgetting boundary cells and missing nearby candidates.
- Treating live location cache as durable trip history.
- Ignoring stale locations and TTLs.
