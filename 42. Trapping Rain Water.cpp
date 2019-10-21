class Solution {
public:
    int trap(vector<int>& height) {
        int n = (int)height.size();
        if (n == 0) return 0;
        
        vector<int> max_left(height.size());
        vector<int> max_right(height.size());
        
        max_left[0] = 0;
        for (int i = 1; i < n; i++) {
            max_left[i] = max(max_left[i - 1], height[i - 1]);
        }
        max_right[n - 1] = 0;
        for (int i = n - 2; i >= 0; i--) {
            max_right[i] = max(max_right[i + 1], height[i + 1]);
        }
        
        int volume = 0;
        for (int i = 0; i < n; i++) {
            int min_side = min(max_left[i], max_right[i]);
            volume += max(0, min_side - height[i]);
        }
        return volume;
    }
};