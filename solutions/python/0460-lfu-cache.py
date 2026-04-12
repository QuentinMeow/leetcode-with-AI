# LeetCode 460 — LFU Cache
# https://leetcode.com/problems/lfu-cache/

from collections import OrderedDict
class LFUCache:

    def __init__(self, capacity: int):
        self.key_to_val = {}
        self.key_to_freq = {}
        self.freq_to_ordered_keys = collections.defaultdict(collections.OrderedDict)
        self.capacity = capacity
        self.min_freq = 0


    def get(self, key: int) -> int:
        if key not in self.key_to_val:
            return -1
        # update frequency
        old_freq = self.key_to_freq[key]
        self.key_to_freq[key] = old_freq + 1

        # update freq to ordered keys
        self.freq_to_ordered_keys[old_freq].pop(key)
        if not self.freq_to_ordered_keys[old_freq]:
            del self.freq_to_ordered_keys[old_freq]
            if self.min_freq not in self.freq_to_ordered_keys:
                self.min_freq += 1
        self.freq_to_ordered_keys[old_freq + 1][key] = 1
        return self.key_to_val[key]
        

    def put(self, key: int, value: int) -> None:
        if self.capacity < 1:
            return

        # existing item
        if key in self.key_to_val:
            self.get(key) # for update the counter
            self.key_to_val[key] = value
            return

        # new item
        if len(self.key_to_val) >= self.capacity:
            del_key, _ = self.freq_to_ordered_keys[self.min_freq].popitem(last=False)
            del self.key_to_freq[del_key]
            del self.key_to_val[del_key]
        
        self.key_to_val[key] = value
        self.key_to_freq[key] = 1
        self.min_freq = 1
        self.freq_to_ordered_keys[1][key] = 1
        

# Your LFUCache object will be instantiated and called as such:
# obj = LFUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
# obj.put(key,value)
