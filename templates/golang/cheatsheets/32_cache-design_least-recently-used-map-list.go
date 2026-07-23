// Source section for the generated Go interview cheatsheet.
// Edit this file, then run: npm run cheatsheets:generate
package cheatsheets

// ===================================================================
// 32. Cache Design
// ===================================================================

// LeastRecentlyUsedCacheNode participates in a doubly linked recency list.
type LeastRecentlyUsedCacheNode struct {
	key   int
	value int
	prev  *LeastRecentlyUsedCacheNode
	next  *LeastRecentlyUsedCacheNode
}

// LeastRecentlyUsedCache combines O(1) map lookup with O(1) recency updates.
// The head and tail are sentinel nodes, so insertion/removal needs no nil cases.
type LeastRecentlyUsedCache struct {
	capacity int
	nodes    map[int]*LeastRecentlyUsedCacheNode
	head     *LeastRecentlyUsedCacheNode
	tail     *LeastRecentlyUsedCacheNode
}

// newLeastRecentlyUsedCache creates a least-recently-used cache. A map gives O(1) key
// lookup; a doubly linked list orders nodes from least recent near head to most recent
// near tail.
func newLeastRecentlyUsedCache(capacity int) *LeastRecentlyUsedCache {
	head, tail := &LeastRecentlyUsedCacheNode{}, &LeastRecentlyUsedCacheNode{}
	head.next = tail
	tail.prev = head
	return &LeastRecentlyUsedCache{
		capacity: capacity,
		nodes:    make(map[int]*LeastRecentlyUsedCacheNode),
		head:     head,
		tail:     tail,
	}
}

// remove detaches one cache node from the doubly linked recency list.
func (c *LeastRecentlyUsedCache) remove(node *LeastRecentlyUsedCacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// addMostRecent inserts one cache node immediately before the tail sentinel.
func (c *LeastRecentlyUsedCache) addMostRecent(node *LeastRecentlyUsedCacheNode) {
	previous := c.tail.prev
	node.prev, node.next = previous, c.tail
	previous.next = node
	c.tail.prev = node
}

// get returns a cached value or -1 when absent. A hit moves its node to the most-recent
// end in O(1).
func (c *LeastRecentlyUsedCache) get(key int) int {
	node, ok := c.nodes[key]
	if !ok {
		return -1
	}
	c.remove(node)
	c.addMostRecent(node)
	return node.value
}

// put inserts or replaces a value, marks it most recent, and evicts the least-recent
// node when capacity is exceeded. All operations are O(1).
func (c *LeastRecentlyUsedCache) put(key, value int) {
	if node, ok := c.nodes[key]; ok {
		c.remove(node)
		delete(c.nodes, key)
	}

	node := &LeastRecentlyUsedCacheNode{key: key, value: value}
	c.nodes[key] = node
	c.addMostRecent(node)

	if len(c.nodes) > c.capacity {
		victim := c.head.next
		c.remove(victim)
		delete(c.nodes, victim.key)
	}
}
