class Solution:
    def climbStairs(self, n: int) -> int:
        if n == 1:
            return 1
        elif n == 2:
            return 2
        else:
            a = 1
            b = 2
            while(n - 2 > 0):
                b = a + b
                a = b - a
                n -= 1
                
            return b;
        
        