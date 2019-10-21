class Solution {
public:
    int findTargetSumWays(vector<int>& nums, int S) {
        int sum = accumulate(nums.begin(), nums.end(), 0);
        if (S > sum) return 0;
        
        int ans = solve(0, nums, S);
        return ans;
    }
private:
    int solve(int pos, vector<int>& nums, int S) {
        if (pos >= nums.size()) return S == 0;
        
        return solve(pos + 1, nums, S - nums[pos]) + 
            solve(pos + 1, nums, S + nums[pos]);        
    }
};