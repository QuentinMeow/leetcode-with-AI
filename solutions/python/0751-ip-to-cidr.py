# LeetCode 751 — IP to CIDR
# https://leetcode.com/problems/ip-to-cidr/

class Solution:
    def ipToCIDR(self, ip_str: str, n: int) -> List[str]:
        result = []
        ip = self.ipToInt(ip_str)

        while n > 0:
            lowest_set_bit = 1
            while lowest_set_bit & ip == 0 and lowest_set_bit < 2**32:
                lowest_set_bit *= 2

            largest_block_size = 1
            while largest_block_size * 2 <= n:
                largest_block_size *= 2

            # print(f"current ip: {self.intToIP(ip)}, {n}, {lowest_set_bit}, {largest_block_size}")
            actual_block_size = min(lowest_set_bit, largest_block_size)
            prefix = int(32 - math.log2(actual_block_size))
            result.append(f"{self.intToIP(ip)}/{prefix}")
            n -= actual_block_size
            ip += actual_block_size

        return result


    def intToIP(self, n: int) -> str:
        return ".".join(str(n >> shift & 255) for shift in [24, 16, 8, 0])

    def ipToInt(self, ip: str) -> int:
        blocks = ip.split(".")
        if len(blocks) != 4:
            return -1
        return int(blocks[0]) * 2**24 + int(blocks[1]) * 2**16 + int(blocks[2]) * 2**8 + int(blocks[3])

