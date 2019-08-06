class Solution {
public:
    int maximalSquare(vector<vector<char>>& matrix) {
        vector < vector <int> > dp;
        for (int i = 0; i < (int)matrix.size(); i++) {
            dp.push_back(vector <int>());
        }
        for (int i = 0; i < (int)dp.size(); i++) {
            dp[i].resize((int)matrix[0].size());
        }
        
        int s = 0;
        for (int i = 0; i < (int)matrix.size(); i++) {
            for (int j = 0; j < (int)matrix[i].size(); j++) {
                dp[i][j] = 0;
                if (matrix[i][j] == '1') {
                    dp[i][j] = 1;
                    if (j - 1 >= 0 && i - 1 >= 0) {
                        dp[i][j] = 1 + min(dp[i][j - 1], min (dp[i - 1][j - 1], dp[i - 1][j]));
                    }
                }
                s = max(s, dp[i][j]);
            }
        }
        return s * s;
    }
};