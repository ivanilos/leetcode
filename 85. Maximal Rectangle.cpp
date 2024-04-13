class Solution {
public:
    int maximalRectangle(vector<vector<char>>& matrix) {
        int n = matrix.size();
        int m = matrix[0].size();

        vector<vector<int>> maxOnesToRight(n, vector<int>(m));

        for (int i = 0; i < n; i++) {
            maxOnesToRight[i][m - 1] = matrix[i][m - 1] - '0';
            for (int j = m - 2; j >= 0; j--) {
                if (matrix[i][j] == '1') {
                    maxOnesToRight[i][j] = 1 + maxOnesToRight[i][j + 1];
                } else {
                    maxOnesToRight[i][j] = 0;
                }
            }
        }

        int result = 0;
        for (int j = 0; j < m; j++) {
            vector<int> nums(n);
            for (int i = 0; i < n; i++) {
                nums[i] = maxOnesToRight[i][j];
            }
            result = max(result, maxHistogramArea(nums));
        }

        return result;
    }

    int maxHistogramArea(vector<int>& height) {
        int n = height.size();

        int result = 0;
        for (int i = 0; i < n; i++) {
            int mini = height[i];
            for (int j = i; j < n; j++) {
                mini = min(mini, height[j]);
                result = max(result, mini * (j - i + 1));
            }
        }
        return result;
    }
};
