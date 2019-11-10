class Solution {
    
public:    
    vector<vector<int>> getSkyline(vector<vector<int>>& buildings) {
        vector<pair<pair<int, int>, int>> events;
        
        for (vector<int>& building : buildings) {
            events.push_back({{building[0], building[2]}, 1});
            events.push_back({{building[1], building[2]}, -1});
        }
        
        vector<vector<int>> skyline;
        multiset<int> heights;
        sort(events.begin(), events.end());
        int prev_height = 0;
        for (int i = 0; i < (int)events.size(); i++) {
            if (events[i].second == 1) {
                heights.insert(events[i].first.second);
            } else {
                auto it = heights.find(events[i].first.second);
                heights.erase(it);
            }
            if (i < (int)events.size() - 1 && events[i].first.first == events[i + 1].first.first) continue;
            
            int max_height = heights.empty() ? 0 : *(heights.rbegin());
            if (max_height != prev_height) {
                skyline.push_back({events[i].first.first, max_height});
                prev_height = max_height;
            }
        }
        return skyline;  
    }
};