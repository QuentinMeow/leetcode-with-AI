# LeetCode 253 — Meeting Rooms II
# https://leetcode.com/problems/meeting-rooms-ii/

class Solution:
    def minMeetingRooms(self, intervals: List[List[int]]) -> int:
        # 0           10
        #  1               12
        #       5  8
        # --------- time ----->
        if not intervals:
            return 0

        start_times = sorted(interval[0] for interval in intervals)
        end_times = sorted(interval[1] for interval in intervals)
        idx_start = idx_end = 0
        max_rooms = 0

        # This is like a matching parentheses problem, we don't care about which end time matches which start time
        # At the runtime, we are just counting the diff between num of start times & end times (which gives us the number of rooms)
        while idx_start < len(intervals):
            if start_times[idx_start] < end_times[idx_end]:
                # we need a new room
                max_rooms += 1
            else:
                # one end room is freed up and canceled out
                idx_end += 1
            idx_start += 1
        
        return max_rooms

