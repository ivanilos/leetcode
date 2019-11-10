class Solution {
public:
    int movesToMakeZigzag(vector<int>& nums) {
        int even = 0;
        int odd = 0;
        
        for (int i = 0; i < (int)nums.size(); i++) {
            int mini_adj;
            if (i - 1 >= 0 && i + 1 < (int)nums.size()) {
                mini_adj = min(nums[i - 1], nums[i + 1]);
            } else if (i - 1 >= 0) {
                mini_adj = nums[i - 1];
            } else if (i + 1 < (int)nums.size()) {
                mini_adj = nums[i + 1];
            } else {
                continue;
            }
            if (i & 1) {
                even += max(0, nums[i] - mini_adj + 1);
            } else {
                odd += max(0, nums[i] - mini_adj + 1);
            }
            
        }
        return min(even, odd);
    }
};