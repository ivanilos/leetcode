class Solution {
public:
    int openLock(vector<string>& deadends, string target) {
        string start = "0000";
        const int DIGITS = 10;

        set<string> forbidden(deadends.begin(), deadends.end());
        map<string, int> dist;

        if (forbidden.count(start)) {
            return -1;
        }

        queue<string> fila;
        fila.push(start);
        dist[start] = 0;

        while(!fila.empty()) {
            string cur = fila.front();
            fila.pop();

            for (int i = 0; i < 4; i++) {
                for (int delta = -1; delta <= 1; delta += 2) {
                    string next = cur;
                    next[i] = ((next[i] - '0' + delta + DIGITS) % DIGITS) + '0';

                    if (forbidden.count(next)) continue;

                    if (dist.count(next) == 0 || dist[next] > 1 + dist[cur]) {
                        dist[next] = 1 + dist[cur];
                        fila.push(next);
                    }
                }
            }
        }

        if (dist.count(target)) {
            return dist[target];
        } else {
            return -1;
        }
    }
};
