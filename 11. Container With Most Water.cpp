class Solution {
public:
    int maxArea(vector<int>& height) {
        vector<pair<int, int> > ordered_heights;
        for (int i = 0; i < (int)height.size(); i++) {
            ordered_heights.push_back({height[i], i});
        }
        
        sort(ordered_heights.begin(), ordered_heights.end());
        int leftmost = 0;
        int rightmost = (int)height.size() - 1;
        
        int ans = 0;
        for (int i = 0; i < (int)ordered_heights.size(); i++) {
            height[ordered_heights[i].second] = 0;
            while(leftmost < (int)height.size() && height[leftmost] == 0) { leftmost++; }
            while(rightmost >= 0 && height[rightmost] == 0) { rightmost--; }
            
            if (leftmost < (int)height.size()) { 
                ans = max(ans, ordered_heights[i].first * (ordered_heights[i].second - leftmost));
            }
            if (rightmost >= 0) {
               ans = max(ans, ordered_heights[i].first * (rightmost - ordered_heights[i].second));
            }
        }
        return ans;
    }
};