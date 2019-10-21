class Solution {
public:
    bool intersect(const vector<int>& interval1, const vector<int>& interval2) {
        return min(interval1[1], interval2[1]) >= max(interval1[0], interval2[0]);
    }
    
    vector<vector<int>> insert(vector<vector<int>>& intervals, vector<int>& newInterval) {
        set<vector<int>> inters;
        
        inters.insert(newInterval);
        for (auto& it : intervals) {
            inters.insert(it);
        }
        
        auto it = inters.find(newInterval);
        while (it != inters.begin()) {
            auto prev = it;
            prev--;
            if (intersect(*prev, *it)) {
                vector<int> replace = {min((*it)[0], (*prev)[0]), max((*it)[1], (*prev)[1])};
                inters.erase(it);
                inters.erase(prev);
                inters.insert(replace);
                it = inters.find(replace);
            } else {
                break;
            }
        }
        auto next = it;
        next++;
        while(next != inters.end()) {
            if (intersect(*next, *it)) {
                vector<int> replace = {min((*it)[0], (*next)[0]), max((*it)[1], (*next)[1])};
                inters.erase(it);
                inters.erase(next);
                inters.insert(replace);
                it = inters.find(replace);
                next = it;
                next++;
            } else {
                break;
            }
        }
        return vector<vector<int>>(inters.begin(), inters.end());
    }
};