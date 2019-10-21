class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int next_pos = 0;
        for (int i = 0; i < nums.size(); i++) {
            if (i == (int)nums.size() - 1 || nums[i] != nums[i + 1]) {
                nums[next_pos++] = nums[i];
            }
        }
        return next_pos;
    }
};