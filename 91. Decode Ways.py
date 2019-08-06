class Solution:
    def is_valid(self, s, start, end):
        if end >= len(s) or s[start] == '0':
            return False
        
        num = 0
        for i in range(start, end + 1):
            num = 10 * num + ord(s[i]) - ord('0')
            
        return 1 <= num <= 26
    
    def numDecodings(self, s: str) -> int:
        if (len(s) == 0):
            return 0
        
        dp = [0] * (len(s) + 1)
        dp[len(s)] = 1
        for i in range(len(s) - 1, -1, -1):
            if self.is_valid(s, i, i):
                dp[i] += dp[i + 1]
                
            if self.is_valid(s, i, i + 1):
                dp[i] += dp[i + 2]
                
        return dp[0]
        