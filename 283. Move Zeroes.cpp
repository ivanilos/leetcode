class Solution {
public:
    void moveZeroes(vector<int>& nums) {
        int next_pos = 0;
        for (int i = 0; i < (int)nums.size(); i++) {
            if (nums[i] != 0) {
                nums[next_pos++] = nums[i];
            }
        }
        for (int i = next_pos; i < (int)nums.size(); i++) {
            nums[i] = 0;
        }
        return;
    }
};