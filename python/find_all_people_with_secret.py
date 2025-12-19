class Solution:
    def findAllPeople(self, n: int, meetings: list[list[int]], firstPerson: int) -> list[int]:
        i = 0
        secret = {0, firstPerson}
        meetings.sort(key=lambda m: m[2])
        while i < len(meetings):
            j = i
            curr_meetings = {}
            while j < len(meetings) and meetings[j][2] == meetings[i][2]:
                a, b, _ = meetings[j]
                curr_meetings.setdefault(a, []).append(b)
                curr_meetings.setdefault(b, []).append(a)
                j += 1
            i = j

            queue = [a for a in curr_meetings if a in secret]
            while queue:
                a = queue.pop()
                for b in curr_meetings[a]:
                    if b not in secret:
                        queue.append(b)
                        secret.add(b)

        return list(secret)
