class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        ALPHA = 26
        freq = [0] * ALPHA
        lastPos = [-n - 1] * ALPHA


        for task in tasks:
            val = ord(task) - ord('A')
            freq[val] += 1

        result = 0
        done = 0
        while done < len(tasks):
            maxi = 0
            idx = -1
            for j in range(ALPHA):
                if freq[j] > maxi and lastPos[j] + n < result:
                    maxi = freq[j]
                    idx = j

            if maxi > 0:
                done += 1
                freq[idx] -= 1
                lastPos[idx] = result
            
            result += 1

        return result
