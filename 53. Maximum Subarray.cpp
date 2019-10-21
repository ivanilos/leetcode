class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        if (nums.empty()) return 0;
        
        int maxi = nums[0];
        int suf = nums[0];
        for (int i = 1; i < nums.size(); i++) {
            suf = max(nums[i], nums[i] + suf);
            maxi = max(maxi, suf);
        }
        return maxi;
    }
};