# Vector - Similarity Search

Examples: pgvector, Pinecone, Milvus, Weaviate, FAISS-backed services, Elasticsearch/OpenSearch vector search.

## 30-Second Interview Answer

Use a vector store when the core query is semantic similarity: find items with embeddings near a query embedding. It is usually a derived index for search, recommendations, deduplication, or retrieval-augmented generation.

## Use When

- Users search by meaning, not exact keywords.
- You need nearest-neighbor lookup over embeddings.
- Recommendations are based on similarity.
- The system needs RAG retrieval over documents.
- Duplicate or near-duplicate detection matters.

## Avoid When

- Exact filters, transactions, or joins are the core requirement.
- Keyword search alone solves the product need.
- Embeddings are not available or cannot be refreshed.
- Explainability and deterministic ranking matter more than semantic similarity.

## Core Model

An embedding model converts text, images, users, products, or documents into high-dimensional vectors. The vector store indexes those vectors and returns nearest neighbors by distance metric.

Common metrics:

- Cosine similarity.
- Dot product.
- Euclidean distance.

## Query And Indexing

Vector search often uses approximate nearest-neighbor indexes, trading exactness for speed.

Concepts:

- Embedding dimension.
- Top-K retrieval.
- Metadata filters.
- Approximate index recall.
- Re-ranking with a stronger model or business rules.

Interview detail: separate retrieval from ranking. Vector search finds candidates; a ranker/filter may enforce permissions, freshness, popularity, or business logic.

## Consistency

Vector indexes are usually eventually consistent. Updates require re-embedding if source content changes. Deletions must propagate to avoid returning removed/private data.

For permissions, do not rely only on the vector index unless metadata filters are reliable. Recheck authorization before returning results.

## Scaling

Scaling depends on:

- Number of vectors.
- Embedding dimension.
- Update rate.
- Top-K size.
- Metadata filter selectivity.
- Recall/latency target.

Failure modes:

- Stale embeddings after content changes.
- Returning private/deleted documents.
- Poor recall from overly aggressive approximation.
- Expensive re-embedding/backfills.
- Embedding model version drift.

## Data Modeling

Store:

- Vector ID.
- Embedding vector.
- Source entity/document ID.
- Chunk ID for long documents.
- Metadata filters: tenant, visibility, language, timestamp, category.
- Embedding model version.

Keep source text and authoritative metadata in the primary database or object storage.

## Interview Examples

- Semantic document search.
- RAG knowledge base.
- Similar product recommendations.
- Image similarity search.
- Near-duplicate detection.

## Senior-Level Tradeoffs

- Vector search improves semantic recall, but exact filters and permissions still need careful handling.
- Approximate indexes reduce latency, but can miss relevant results.
- Chunking improves retrieval for long documents, but creates more vectors and ranking complexity.
- Embeddings age as content and models change, so versioning and reindexing are part of the design.

## Common Mistakes

- Treating vector DB as the only database.
- Forgetting authorization checks after retrieval.
- Ignoring embedding refresh and model versioning.
- Using vector search when normal full-text search is enough.
