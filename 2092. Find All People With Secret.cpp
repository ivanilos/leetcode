class Solution {
public:
    vector<int> findAllPeople(int n, vector<vector<int>>& meetings, int first_person) {
        vector<vector<pair<int, int>>> person_meetings(n);
        vector<int> know(n);

        for (auto meeting: meetings) {
            int a = meeting[0];
            int b = meeting[1];
            int timer = meeting[2];

            person_meetings[a].push_back({timer, b});
            person_meetings[b].push_back({timer, a});
        }

        for (int i = 0; i < n; i++) {
            sort(person_meetings[i].begin(), person_meetings[i].end());
        }

        priority_queue<pair<int, int>> pq;
        pq.push({0, first_person});
        pq.push({0, 0});

        int cur_timer = 0;
        while(!pq.empty()) {
            int found = -pq.top().first;
            int next = pq.top().second;
            pq.pop();

            if (know[next]) {
                continue;
            }
            know[next] = true;

            for (auto p : person_meetings[next]) {
                int timer = p.first;
                int other = p.second;

                //printf("next = %d, found = %d, timer = %d, other = %d\n", next, found, timer, other);
                if (!know[other] && found <= timer) {
                    pq.push({-timer, other});
                }
            }
        }

        vector<int> ans;
        for (int i = 0; i < n; i++) {
            if (know[i]) {
                ans.push_back(i);
            }
        }
        return ans; 
    }
};
