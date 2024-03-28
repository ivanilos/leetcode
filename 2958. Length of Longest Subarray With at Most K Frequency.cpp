class Solution {
public:
    int maxSubarrayLength(vector<int>& nums, int k) {
        int l = 0;
        map<int, int> freq;

        int result = 1;
        for (int r = 0; r < (int)nums.size(); r++) {
            int val = nums[r];
            freq[val]++;

            while(l < r && freq[val] > k) {
                int to_delete = nums[l];
                freq[to_delete]--;

                l++;
            }
            result = max(result, r - l + 1);
        }
        return result;
    }
};