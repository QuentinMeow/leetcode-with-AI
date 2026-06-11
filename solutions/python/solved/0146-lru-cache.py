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


def run_assertion_tests():
    cache = LRUCache(2)
    cache.put(1, 1)
    cache.put(2, 2)
    assert cache.get(1) == 1
    cache.put(3, 3)
    assert cache.get(2) == -1
    cache.put(4, 4)
    assert cache.get(1) == -1
    assert cache.get(3) == 3
    assert cache.get(4) == 4

    cache = LRUCache(2)
    cache.put(1, 1)
    cache.put(2, 2)
    cache.put(1, 10)
    cache.put(3, 3)
    assert cache.get(1) == 10
    assert cache.get(2) == -1


if __name__ == "__main__":
    cache = LRUCache(2)
    cache.put(1, 1)
    cache.put(2, 2)
    print(f"Test case 1 (expected: 1): {cache.get(1)}")
    cache.put(3, 3)
    print(f"Test case 2 (expected: -1): {cache.get(2)}")
    cache.put(4, 4)
    print(f"Test case 3 (expected: -1): {cache.get(1)}")
    print(f"Test case 4 (expected: 3): {cache.get(3)}")
    print(f"Test case 5 (expected: 4): {cache.get(4)}")

    run_assertion_tests()
