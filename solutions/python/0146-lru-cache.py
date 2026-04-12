# LeetCode 146 — LRU Cache
# https://leetcode.com/problems/lru-cache/

class ListNode:
    def __init__(self, key, val):
        self.key = key
        self.val = val
        self.prev = None
        self.next = None

class LRUCache:

    def __init__(self, capacity: int):
        self.capacity = capacity
        self.dict = {}
        # dummy head and tail (not inserted into dict; keys unused)
        self.head = ListNode(0, 0)
        self.tail = ListNode(0, 0)
        self.head.next = self.tail
        self.tail.prev = self.head

    def add(self, node):
        before = self.tail.prev
        after = self.tail
        node.prev = before
        node.next = after
        before.next = node
        after.prev = node

    def remove(self, node):
        before = node.prev
        after = node.next
        before.next = after
        after.prev = before


    def get(self, key: int) -> int:
        if key not in self.dict:
            return -1
        
        node = self.dict[key]
        self.remove(node)
        self.add(node)
        return node.val

    def put(self, key: int, value: int) -> None:
        if key in self.dict:
            self.remove(self.dict[key])

        node = ListNode(key, value)
        self.dict[key] = node
        self.add(node)

        if len(self.dict) > self.capacity:
            node_to_delete = self.head.next
            self.remove(node_to_delete)
            del self.dict[node_to_delete.key]

# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
