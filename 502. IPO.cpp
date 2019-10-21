class Solution {
public:
    int findMaximizedCapital(int k, int W, vector<int>& Profits, vector<int>& Capital) {
        int cur = W;
        priority_queue<int> gain;
        priority_queue<pair<int, int>> vals;
        
        for (int i = 0; i < (int)Profits.size(); i++) {
            vals.push({-Capital[i], Profits[i]});
        }
        
        int got = 0;
        while(got < k) {
            while(!vals.empty() && -vals.top().first <= cur) {
                gain.push(vals.top().second);
                vals.pop();
            }
            if (gain.empty()) { break; }
            if (!gain.empty()) {
                int win = gain.top();
                gain.pop();
                if (win <= 0) { break; }
                else {
                    got++;
                    cur += win;
                }
            }
        }
        return cur;
    }
};