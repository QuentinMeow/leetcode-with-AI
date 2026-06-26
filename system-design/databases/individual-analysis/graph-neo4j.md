# Graph - Neo4j-Style Databases

Examples: Neo4j, Amazon Neptune, JanusGraph, TigerGraph.

## 30-Second Interview Answer

Use a graph database when relationship traversal is the core product feature and queries naturally walk multiple hops. It is strongest for social graphs, fraud rings, recommendations, dependency graphs, and knowledge graphs.

## Use When

- Multi-hop relationships are central.
- Query shape is "find connected things" rather than "filter rows".
- Edges have meaningful types/properties.
- Traversal depth is dynamic or hard to precompute.
- Graph algorithms matter: shortest path, community detection, centrality.

## Avoid When

- Relationships are one-hop and simple.
- SQL joins or a denormalized adjacency list are enough.
- The main workload is high-volume append-only events.
- You need simple key-value scale more than traversal flexibility.

## Core Model

Graph databases store nodes and edges. Edges can have labels and properties. Queries traverse relationships directly.

Example:

- Node: user, merchant, card, device.
- Edge: follows, purchased_from, used_device, transferred_to.

## Query And Indexing

Graph queries start from indexed nodes and traverse edges. Indexes find starting points; adjacency storage makes traversal fast.

Interview detail: graph databases do not make every graph query free. You still need bounded traversal depth, good starting filters, and careful high-degree node handling.

## Consistency

Many graph systems support transactions for graph updates, but distributed graph scaling can be hard. If the graph is derived from events, eventual consistency may be acceptable.

## Scaling

Graph scaling is challenging because relationships cut across partitions. Traversals may cross many machines.

Failure modes:

- Supernodes with millions of edges.
- Unbounded traversals.
- Expensive global graph algorithms on user-facing paths.
- Partitioning that splits highly connected subgraphs.
- Derived graph lag from source events.

## Data Modeling

Use graph when edges are first-class data.

Patterns:

- Social: user -> follows -> user.
- Fraud: account -> uses_device -> device -> used_by -> account.
- Recommendation: user -> bought -> product -> also_bought -> product.
- Dependency: service -> depends_on -> service.

## Interview Examples

- Detect fraud rings through shared devices/cards/addresses.
- Recommend friends-of-friends.
- Find shortest path in a professional network.
- Model service dependency blast radius.
- Knowledge graph search for entities and relationships.

## Senior-Level Tradeoffs

- Graphs make relationship queries expressive, but operating and scaling them is harder than SQL/key-value.
- For simple social feed reads, precomputed adjacency lists in key-value/wide-column may be better.
- For fraud analysis, graph can be an offline/nearline derived store rather than the write path.
- Supernodes need special handling: caps, sampling, precomputation, or separate storage.

## Common Mistakes

- Choosing graph for any data with relationships.
- Ignoring supernodes like celebrities or popular merchants.
- Running deep unbounded traversals synchronously.
- Using graph as source of truth when relational entities need stronger constraints.
