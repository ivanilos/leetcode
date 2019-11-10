class Solution {
public:
    int findMaxLength(vector<int>& nums) {
        map<int, int> pos;
        
        int sum = 0;
        int ans = 0;
        pos[0] = -1;
        for (int i = 0; i < (int)nums.size(); i++) {
            sum += nums[i] == 0 ? -1 : 1;
            if (pos.count(sum)) {
                ans = max(ans, i - pos[sum]);
            } else {
                pos[sum] = i;
            }
        }
        return ans;
    }
};