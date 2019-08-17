class Solution {
public:
    vector<int> spiralOrder(vector<vector<int>>& matrix) {
        if (matrix.empty()) return {};
        
        vector<int> ans;
        solve(matrix, 0, 0, matrix.size() - 1, matrix.back().size() - 1, ans);
        return ans;
    }
    
    
private:
    void solve(vector<vector<int>>& matrix, int xl, int yl, int xr, int yr, vector<int>& ans) {
        if (xl > xr || yl > yr) { return; }
        if (xl == xr || yl == yr) {
            for (int i = xl; i <= xr; i++) {
                for (int j = yl; j <= yr; j++) {
                    ans.push_back(matrix[i][j]);
                }
            }
            return;
        }
        
        for (int j = yl; j <= yr - 1; j++) { ans.push_back(matrix[xl][j]); }
        for (int i = xl; i <= xr - 1; i++) { ans.push_back(matrix[i][yr]); }
        for (int j = yr; j >= yl + 1; j--) { ans.push_back(matrix[xr][j]); }
        for (int i = xr; i >= xl + 1; i--) { ans.push_back(matrix[i][yl]); }
        
        solve(matrix, xl + 1, yl + 1, xr - 1, yr - 1, ans);
    }
};