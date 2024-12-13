class Solution {
public:
    long long findScore(vector<int>& nums) {
        priority_queue<pair<int, int>> pq;
        map<int, bool> used;

        for (int i = 0; i < nums.size(); i++) {
            pq.push({-nums[i], -i});
        }

        long long result = 0;
        while (!pq.empty()) {
            int val = -pq.top().first;
            int idx = -pq.top().second;
            pq.pop();

            if (used[idx]) {
                continue;
            }
            used[idx] = true;
            used[idx - 1] = true;
            used[idx + 1] = true;

            result += val;
        }

        return result;
    }
};
