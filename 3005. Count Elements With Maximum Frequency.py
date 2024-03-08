class Solution:
    def maxFrequencyElements(self, nums: List[int]) -> int:
        freq = {}

        for num in nums:
            freq.setdefault(num, 0)
            freq[num] += 1


        max_freq = 0
        result = 0
        for _, val in freq.items():
            if val > max_freq:
                result = val
                max_freq = val
            elif val == max_freq:
                result += val

        return result
