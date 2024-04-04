class Solution:
    def maxDepth(self, s: str) -> int:
        cur = 0
        best = 0
        for c in s:
            if c == '(':
                cur += 1
            elif c == ')':
                cur -= 1
            
            best = max(best, cur)
        
        return best
