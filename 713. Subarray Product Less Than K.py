class Solution:
    def numSubarrayProductLessThanK(self, nums: List[int], k: int) -> int:
        result = 0

        l = 0

        cur = 1
        for r in range(len(nums)):
            cur *= nums[r]

            while cur >= k and l < r:
                cur //= nums[l]
                l += 1

            if cur < k:
                result += r - l + 1

        return result
