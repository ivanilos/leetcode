class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        negative = []
        positive = []

        for num in nums:
            if num < 0:
                negative.append(num * num)
            else:
                positive.append(num * num)

        positive = positive[::-1]

        result = []
        while(len(positive) > 0 and len(negative) > 0):
            if positive[-1] < negative[-1]:
                result.append(positive[-1])
                positive.pop()
            else:
                result.append(negative[-1])
                negative.pop()

        while(len(positive) > 0):
            result.append(positive[-1])
            positive.pop()

        while(len(negative) > 0):
            result.append(negative[-1])
            negative.pop()

        return result
        