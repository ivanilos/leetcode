class Solution {
public:
    int minOperations(vector<int>& nums, int k) {
        int xor_val = 0;
        for (auto num : nums) {
            xor_val ^= num;
        }
        return __builtin_popcount(xor_val ^ k);
    }
};
