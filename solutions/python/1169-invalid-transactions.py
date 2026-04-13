# LeetCode 1169 — Invalid Transactions
# https://leetcode.com/problems/invalid-transactions/

from typing import List
class Solution:
    def invalidTransactions(self, transactions: List[str]) -> List[str]:
        invalid = []
        if not transactions:
            return invalid

        name_to_trans = defaultdict(list)
        for trans_str in transactions:
            trans = Transaction(trans_str)
            name_to_trans[trans.name].append(trans)

        for name, trans in name_to_trans.items():
            city_ct = defaultdict(int)
            trans = sorted(trans)
            # sliding window of + / - 60 min: left <= right
            left = right = 0

            for i in range(len(trans)):
                # print(f"iteration started with i: {i}, trans: {trans[i]}")
                # update sliding window for past transactions, we only include time within 60 min (inclusive)
                while trans[i].time - trans[left].time > 60:
                    city_ct[trans[left].city] -= 1
                    left += 1
                # update sliding window for future transactions, we only include time within 60 min (inclusive)
                while right < len(trans) and trans[right].time - trans[i].time <= 60:
                    city_ct[trans[right].city] += 1
                    right += 1
                # print(f"iteration finished with i: {i}, left: {left}, right: {right}, city: {trans[i].city}, city count: {city_ct[trans[i].city]}")
                if trans[i].amount > 1000:
                    invalid.append(trans[i].value)
                elif city_ct[trans[i].city] != right - left:
                    # if the new city count in sliding window is the same as total count - itself
                    # that means there's no other city in sliding window
                    invalid.append(trans[i].value)
        return invalid

class Transaction:
    def __init__(self, transaction_str: str):
        parts = transaction_str.split(",")
        self.name = parts[0]
        self.time = int(parts[1])
        self.amount = int(parts[2])
        self.city = parts[3]
        self.value = transaction_str

    def __lt__(self, other):
        return self.time < other.time

    def __repr__(self):
        return f"Transaction(name='{self.name}', time={self.time}, amount={self.amount}, city='{self.city}')"

        
