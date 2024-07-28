class Solution {
public:

    const static int MAXN = 10005;
    const static int INF = int(1e9);

    vector<vector<int>> adjList[MAXN];

    int secondMinimum(int n, vector<vector<int>>& edges, int time, int change) {
        vector<vector<int>> adjList = buildGraph(n, edges);

        int result = solve(n, adjList, time, change);
        return result;
    }

    vector<vector<int>> buildGraph(int n, vector<vector<int>> &edges) {
        vector<vector<int>> adjList = vector<vector<int>>(n + 1);

        for (auto edge : edges) {
            int a = edge[0];
            int b = edge[1];
            
            adjList[a].push_back(b);
            adjList[b].push_back(a);
        }
        return adjList;
    }

    int solve(int n, vector<vector<int>> adjList, int time, int change) {
        int initial = INF;
        vector<vector<int>> dist = vector<vector<int>>(n + 1, vector<int>(2, initial));

        dist[1] = {0, INF};

        queue<int> fila;
        fila.push(1);

        while (!fila.empty()) {
            int next = fila.front();

            fila.pop();

            for (auto viz : adjList[next]) {
                for (int i = 0; i < 2; i++ ) {
                    int nextGreen = getNextGreen(dist[next][i], change);
                    int nextDist = nextGreen + time;
                    if (dist[viz][0] > nextDist) {
                        dist[viz][1] = dist[viz][0];
                        dist[viz][0] = nextDist;
                        fila.push(viz);
                    } else if (dist[viz][1] > nextDist && nextDist > dist[viz][0]) {
                        dist[viz][1] = nextDist;
                        fila.push(viz);
                    }
                }
            }
        }

        return dist[n][1];
    }

    int getNextGreen(int curTime, int change) {
        int timesChanged = curTime / change;

        if (timesChanged % 2 == 1) {
            return ((curTime + change) / change) * change;
        } else {
            return curTime;
        }
    }
};
