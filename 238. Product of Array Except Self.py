class Solution(object):
    def productExceptSelf(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        
        ans = [1] * len(nums)
        
        pref = nums[0]
        for i in range(1, len(nums)):
            ans[i] = pref
            pref *= nums[i]
            
        suf = nums[-1]
        for i in range(len(nums) - 2, -1, -1):
            ans[i] *= suf
            suf *= nums[i]
            
        return ans
        