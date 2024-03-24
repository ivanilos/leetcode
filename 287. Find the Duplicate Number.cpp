class Solution {
public:
    int findDuplicate(vector<int>& nums) {
        int sz = (int)nums.size();

        bitset<100005> got;

        for (int i = 0; i < sz; i++) {
            if (got.test(nums[i])) {
                return nums[i];
            } else {
                got.set(nums[i]);
            }
        }

        return -1;
    }
};
