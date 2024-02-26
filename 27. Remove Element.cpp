class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        int next_free_pos = 0;
        for (int num : nums) {
            if (num != val) {
                nums[next_free_pos] = num;
                next_free_pos++;
            }
        }
        nums.resize(next_free_pos);
        return next_free_pos;
    }
};