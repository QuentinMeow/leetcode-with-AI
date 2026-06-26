# Search - Elasticsearch / OpenSearch

Examples: Elasticsearch, OpenSearch, Solr, Lucene-based services.

## 30-Second Interview Answer

Use a search index when users need full-text search, fuzzy matching, relevance ranking, faceting, filtering, autocomplete, or log search. Treat it as a derived index, not the source of truth.

## Use When

- Search is a product feature.
- Queries need relevance scoring or fuzzy matching.
- Users filter by many fields and sort by rank/time.
- Logs/events need fast text search.
- Autocomplete or prefix matching matters.

## Avoid When

- You only need exact lookup by primary key.
- Strong consistency is required immediately after every write.
- Relational filters and joins are enough.
- The team cannot operate indexing, mappings, and shard sizing.

## Core Model

Search engines build inverted indexes from documents. Text is analyzed into tokens, normalized, and indexed. Queries retrieve matching documents and score relevance.

Common document examples:

- Product search document.
- User/profile search document.
- Post/article search document.
- Log event document.

## Query And Indexing

Important concepts:

- Analyzer: tokenizes and normalizes text.
- Inverted index: maps term -> document list.
- Mapping: field types and analyzers.
- Facets/aggregations: category counts and filters.
- Relevance scoring: ranking based on match quality.

Interview detail: only index fields that support user-visible search/filtering. Every indexed field costs write and storage.

## Consistency

Search indexes are usually eventually consistent. Writes go to the source-of-truth database, then async indexing updates search.

Common pipeline:

1. Write primary database.
2. Publish event or use change data capture.
3. Index/update/delete search document.
4. Monitor lag and dead-letter failures.

## Scaling

Search clusters scale through shards and replicas.

Failure modes:

- Shard imbalance.
- Expensive wildcard queries.
- Mapping explosion from unbounded dynamic fields.
- Index lag after write spikes.
- Reindexing large datasets under load.
- Cluster memory pressure from aggregations.

## Data Modeling

Search documents are often denormalized. Include all fields needed to render search results, so search does not fan out to the primary DB for every hit.

Patterns:

- Product document: title, description, category, price, seller, availability.
- User document: display name, username, bio, public flags.
- Log document: timestamp, service, severity, message, trace ID.

## Interview Examples

- E-commerce product search.
- Search posts or messages.
- Autocomplete usernames.
- Log search/observability.
- Geospatial restaurant search with filters.

## Senior-Level Tradeoffs

- Search gives relevance and flexible filters, but sacrifices immediate consistency.
- Denormalized documents reduce query fanout, but require reindex/backfill strategy.
- More shards are not always better; too many shards increase overhead.
- Search can complement SQL, not replace transactional storage.

## Common Mistakes

- Using search as the only source of truth.
- Forgetting indexing lag and delete propagation.
- Not planning reindexing when mappings change.
- Running unbounded aggregations on user-facing paths.
