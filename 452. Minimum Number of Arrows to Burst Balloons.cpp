class Solution {
public:
    int findMinArrowShots(vector<vector<int>>& points) {
        sort(points.begin(), points.end(),
            [](const vector<int>& a, const vector<int>& b) {
                return a.back() < b.back();});
        
        int ans = 0;
        if (points.size() > 0) {
            int shoot = points[0].back();
            ans++;
            for (auto balloon : points) {
                if (shoot < balloon[0] || balloon[1] < shoot) {
                    ans++;
                    shoot = balloon[1];
                }
            }
        }
        return ans;
        
    }
};