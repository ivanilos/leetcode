class Solution {
public:
    int beautifulSubsets(vector<int>& nums, int k) {
        vector<int> freq(1005);

        for (int i = 0; i < nums.size(); i++) {
            freq[nums[i]]++;
        }

        vector<vector<int>> v;
        for (int start = 1; start <= 1000; start++) {
            vector<int> cur;
            for (int i = start; i <= 1000; i += k) {
                if (freq[i]) {
                    cur.push_back(freq[i]);
                    freq[i] = 0;
                } else {
                    break;
                }
            }
            if (cur.size()) {
                v.push_back(cur);
            }
        }

        int result = 1;
        for (auto &cur : v) {
            result *= solve(cur);
        }
        return result - 1;
    }

    int solve(vector<int>& cur) {
        vector<int> dp(cur.size(), -1);

        return calc(0, cur, dp);
    }

    int calc(int pos, vector<int>& cur, vector<int>& dp) {
        if (pos >= cur.size()) {
            return 1;
        }
        if (dp[pos] != -1) {
            return dp[pos];
        }

        int& ans = dp[pos];
        ans = calc(pos + 1, cur, dp);
        ans += ((1 << cur[pos]) - 1) * calc(pos + 2, cur, dp);
        return ans;
    }
};
