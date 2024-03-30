class Solution {
public:
    int subarraysWithKDistinct(vector<int>& nums, int k) {
        return solve(nums, k) - solve(nums, k - 1);
    }

    int solve(vector<int>& nums, int k) {
        map<int, int> freq;

        int l = 0;
        int in_segment = 0;

        int result = 0;
        for (int r = 0; r < (int)nums.size(); r++) {
            int val = nums[r];
            freq[val]++;
            if (freq[val] == 1) {
                in_segment++;
            }

            while (l < r && in_segment > k) {
                val = nums[l];
                freq[val]--;
                if (freq[val] == 0) {
                    in_segment--;
                }
                l++;
            }

            if (in_segment <= k) {
                result += r - l + 1;
            }
        }
        return result;
    }
};