class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        splitted = s.split()
        if (len(splitted) > 0):
            return len(splitted[-1])
        
        return 0