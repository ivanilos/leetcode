class Solution {
public:
    void sortColors(vector<int>& nums) {
        int left_free = 0;
        int right_free = (int)nums.size() - 1;
        const int RED = 0;
        const int BLUE = 2;
        
        for (int pos = 0; pos < (int)nums.size(); pos++) {
            while(left_free < (int)nums.size() && nums[left_free] == RED) { left_free++; }
            while(right_free >= 0 && nums[right_free] == BLUE) { right_free--; }
            
            if (nums[pos] == RED && pos > left_free) {
                swap(nums[pos], nums[left_free]);
                pos--;
            } else if (nums[pos] == BLUE && pos < right_free) {
                swap(nums[pos], nums[right_free]);
                pos--;
            }
        }
    }
};