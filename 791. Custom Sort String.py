class Solution:
    def customSortString(self, order: str, s: str) -> str:
        posMap = {}

        next = 0
        for c in order:
            posMap[c] = next
            next += 1
        
        for c in s:
            if c not in posMap:
                posMap[c] = next
                next += 1


        result = sorted(s, key = lambda x: posMap[x])

        return ''.join(result)
        