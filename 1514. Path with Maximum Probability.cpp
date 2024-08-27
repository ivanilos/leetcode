class Solution {
public:
    vector<int> g[10005];
    vector<double> w[10005];

    double maxProbability(int n, vector<vector<int>>& edges, vector<double>& succProb, int start_node, int end_node) {
        for (int i = 0; i < (int)edges.size(); i++) {
            int a = edges[i][0];
            int b = edges[i][1];
            double weight = succProb[i];

            g[a].push_back(b);
            g[b].push_back(a);

            w[a].push_back(weight);
            w[b].push_back(weight);
        }

        double result = dijkstra(n, start_node, end_node);
        return result;
    }

    double dijkstra(int n, int start_node, int end_node) {
        const double EPS = 1e-6;
        vector<double> prob = vector<double>(n);

        for (int i = 0; i < n; i++) prob[i] = 0;

        prob[start_node] = 1;
        priority_queue<pair<double, int>> pq;
        pq.push({1, start_node});

        while(!pq.empty()) {
            int next = pq.top().second;
            double p = pq.top().first;
            pq.pop();

            if (p + EPS < prob[next]) {
                continue;
            }

            for (int i = 0; i < (int)g[next].size(); i++) {
                int viz = g[next][i];
                double newProb = p * w[next][i];

                if (newProb > prob[viz] + EPS) {
                    prob[viz] = newProb;
                    pq.push({newProb, viz});
                }
            }
        }
        return prob[end_node];
    }
};
