from heapq import heapify, heappop, heappush

class Solution:
    def mostBooked(self, n: int, meetings: list[list[int]]) -> int:
        meetings.sort()

        counts = [0] * n
        available = list(range(n))
        heapify(available)
        occupied = []

        for start, end in meetings:
            while occupied:
                time, room = occupied[0]
                if time > start:
                    break
                heappop(occupied)
                heappush(available, room)

            if available:
                room = heappop(available)
                heappush(occupied, (end, room))
            else:
                time, room = heappop(occupied)
                heappush(occupied, (time + end - start, room))

            counts[room] += 1

        return counts.index(max(counts))
