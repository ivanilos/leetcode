class Solution {
public:
    int numDistinct(string s, string t) {
        vector<vector<int>> dp((int)s.size());
        
        for (int i = 0; i < s.size(); i++) {
            dp[i] = vector<int>((int)t.size(), -1);
        }
        int ans = solve(0, 0, s, t, dp);
        return ans;
    }
private:
    int solve(int pos, int got, string& s, string& t, vector<vector<int>>& dp) {
        if (pos >= s.size() || got >= t.size()) {
            return got >= t.size();
        }
        if (dp[pos][got] != -1) return dp[pos][got];
        
        int& ans = dp[pos][got];
        ans = solve(pos + 1, got, s, t, dp);
        if (got < (int)t.size() && s[pos] == t[got]) {
            ans += solve(pos + 1, got + 1, s, t, dp);
        }
        return ans;
    }
};