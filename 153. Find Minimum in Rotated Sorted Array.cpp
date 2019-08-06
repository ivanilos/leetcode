class Solution {
public:
    int findMin(vector<int>& nums) {
        if (nums.empty()) { throw InvalidInputException{}; }
        int mini = nums[0];
        for (int i = 0; i < (int)nums.size() - 1; i++) {
            if (nums[i] > nums[i + 1]) {
                mini = nums[i + 1];
            }
        }
        return mini;
    }
    
private:
    class InvalidInputException {};
};