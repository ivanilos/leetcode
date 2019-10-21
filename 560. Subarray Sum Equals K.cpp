class Solution {
public:
    int subarraySum(vector<int>& nums, int k) {
        map<int, int> pref_sum;
        
        pref_sum[0] = 1;
        int pref = 0;
        int ans = 0;
        for (int num : nums) {
            pref += num;
            if (pref_sum.count(pref - k)) {
                ans += pref_sum[pref - k];
            }
            pref_sum[pref]++;
        }
        return ans;
    }
};