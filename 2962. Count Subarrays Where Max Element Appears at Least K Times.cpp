class Solution {
public:
    long long countSubarrays(vector<int>& nums, int k) {
        int maxi = get_max(nums);
        map<int, int> freq;
        map<int, vector<int>> num_pos;

        int max_acceptable_l = -1;

        long long result = 0;
        for (int r = 0; r < nums.size(); r++) {
            int val = nums[r];
            freq[val]++;
            num_pos[val].push_back(r);

            if (val == maxi && freq[val] >= k) {
                int sz = (int)num_pos[val].size();
                max_acceptable_l = max(max_acceptable_l, num_pos[val][sz - k]);
            }
            result += max_acceptable_l + 1;
        }
        
        return result;
    }

    int get_max(vector<int>& nums) {
        int maxi = nums[0];
        for (auto it : nums) {
            maxi = max(maxi, it);
        }
        return maxi;
    }
};