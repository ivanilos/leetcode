class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        if (nums.empty()) return 0;
        
        constexpr int MAX_REPEATS = 2;
        int next_pos = 0;
        int repeats = 0;
        int prev_val = nums[0] + 1;
        for (int i = 0; i < (int)nums.size(); i++) {
            if (nums[i] == prev_val) repeats++;
            else {
                repeats = 1;
                prev_val = nums[i];
            }
            
            if (repeats <= MAX_REPEATS) {
                nums[next_pos++] = prev_val;
            }
        }
        return next_pos;
    }
};