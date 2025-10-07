class Solution {
public:
    vector<int> avoidFlood(vector<int>& rains) {
        map<int, int> nextRain;
        map<int, int> lastPos;
        for (int i = rains.size() - 1; i >= 0; i--) {
            if (rains[i] != 0) {
                if (lastPos.find(rains[i]) == lastPos.end()) {
                    lastPos[rains[i]] = rains.size();
                }
                nextRain[i] = lastPos[rains[i]];
                lastPos[rains[i]] = i; 
            }
        }

        vector<int> result = vector<int>(rains.size(), 1);

        map<int, bool> flooded;
        priority_queue<pair<int, int>> nextFlood;  
        for (int i = 0; i < rains.size(); i++) {
            if (rains[i] == 0 && nextFlood.size() > 0) {

                int next = nextFlood.top().second;
                nextFlood.pop();
                flooded[next] = false;
                result[i] = next;
            } else if (rains[i] != 0) {
                result[i] = -1;
                if (flooded[rains[i]] == true) {
                    return vector<int>();
                }

                flooded[rains[i]] = true;
                nextFlood.push({-nextRain[i], rains[i]});
            }
        }

        return result;
    }
};
