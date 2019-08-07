class Solution {
public:
    int rob(vector<int>& nums) {
        if (nums.empty()) return 0;
        if (nums.size() == 1) return nums[0];
        
        int ans = 0;
        vector<int> dp((int)nums.size() + 2);
        
        for (int i = (int)nums.size() - 1; i >= 1; i--) {
            dp[i] = max(dp[i + 1], nums[i] + dp[i + 2]);
        }
        ans = dp[1];
        
        dp[(int)nums.size() - 1] = 0;
        for (int i = (int)nums.size() - 2; i >= 0; i--) {
            dp[i] = max(dp[i + 1], nums[i] + dp[i + 2]);
        }
        ans = max(ans, dp[0]);
        return ans;
    }
};