# Document - MongoDB-Style Databases

Examples: MongoDB, Couchbase, Firestore-like document stores.

## 30-Second Interview Answer

Use a document database when the data naturally fits aggregate documents, the schema evolves often, and most reads/writes touch one document or a small set of documents. It is strongest for flexible product records, profiles, content metadata, and nested aggregates.

## Use When

- Records are aggregate-shaped JSON/BSON documents.
- Reads commonly fetch the whole aggregate.
- Schema varies by record type or evolves frequently.
- You want to denormalize nested data for read performance.
- Cross-document relationships are limited.

## Avoid When

- The system needs many joins or relational reporting.
- Strong multi-document invariants dominate the design.
- Documents can grow without bound.
- Query patterns require complex ad hoc joins across many collections.

## Core Model

A document store stores records as structured documents. A document can contain nested objects and arrays, making it natural for profile, catalog, article, or configuration data.

Example aggregates:

- Product with flexible attributes.
- User profile with preferences.
- CMS article with embedded blocks.
- Game state or device config.

## Query And Indexing

Document stores index fields inside documents. Indexes support common filters and sorts, but each index increases write cost and storage.

Interview detail: model around aggregate boundaries. Embed data that is read together and changes together. Reference data that grows independently or is shared by many documents.

## Transactions And Consistency

Single-document operations are usually atomic. Many modern document databases support multi-document transactions, but if the design relies heavily on them, relational SQL may be simpler.

## Scaling

Document databases can shard by a shard key. The shard key should have high cardinality and match common access patterns.

Failure modes:

- Poor shard key creates hot shards.
- Large documents exceed limits or become expensive to update.
- Unbounded arrays make writes and reads slow.
- Secondary indexes multiply write cost.

## Data Modeling

### Embed When

- Child data is owned by the parent.
- Child data is read with the parent.
- Child cardinality is bounded.
- Updates happen together.

### Reference When

- Child data is shared.
- Child cardinality is unbounded.
- Child data changes independently.
- You need separate access control or lifecycle.

## Interview Examples

- Product catalog with varied attributes by category.
- User profile/preferences.
- CMS content documents.
- Feature configuration by tenant.
- Metadata for uploaded media, with bytes in object storage.

## Senior-Level Tradeoffs

- Documents reduce join complexity for aggregate reads, but cross-aggregate queries become harder.
- Flexible schema speeds product iteration, but validation and migrations still matter.
- Denormalization improves reads, but stale duplicated fields need repair strategy.
- Sharding can scale, but shard-key mistakes are expensive.

## Common Mistakes

- Saying "schema-less" as if there is no schema. The schema moved into application expectations.
- Embedding unbounded event lists in one document.
- Using document storage for highly relational financial/order workflows.
- Ignoring index design until queries are slow.
