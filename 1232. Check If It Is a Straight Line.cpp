class Solution {
public:
    int cross(vector<int>& a, vector<int>& b, vector<int>& c) {
        return a[0] * (b[1] - c[1]) + b[0] * (c[1] - a[1]) + c[0] * (a[1] - b[1]);
    }
    
    bool checkStraightLine(vector<vector<int>>& coordinates) {
        for (vector<int>& point : coordinates) {
            if (cross(coordinates[0], coordinates[1], point) != 0) {
                return false;
            }
        }
        return true;
    }
};