class Solution {
public:
    int minSubArrayLen(int s, vector<int>& nums) {
        int ans = (int)nums.size() + 1;
        int sum = 0;
        int l = 0;
        int r = 0;
        while(r < (int)nums.size()) {
            while (r <= l) {
                sum += nums[r];
                r++;
            }
            while(r < (int)nums.size() && sum < s) {
                sum += nums[r];
                r++;
            }
            while(l < r && sum >= s) {
                ans = min(ans, r - l);
                sum -= nums[l];
                l++;
            }
        }
        return ans == (int)nums.size() + 1 ? 0 : ans;
    }
};