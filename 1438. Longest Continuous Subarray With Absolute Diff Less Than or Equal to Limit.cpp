class Solution {
public:
    int longestSubarray(vector<int>& nums, int limit) {
        multiset<int> window;

        int result = 1;
        int l = 0;
        for (int r = 0; r < nums.size(); r++) {
            window.insert(nums[r]);
            while(window.size() > 0) {
                int mini = *window.begin();
                int maxi = *window.rbegin();

                if (maxi - mini > limit) {
                    auto leftmost = window.find(nums[l]);
                    l++;
                    window.erase(leftmost);
                } else {
                    break;
                }
            }
            result = max(result, r - l + 1);
        }
        return result;
    }
};
