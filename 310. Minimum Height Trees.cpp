class Solution {
public:
    vector<int> findMinHeightTrees(int n, vector<vector<int>>& edges) {
        vector<vector<int>> adjList = buildGraph(n, edges);

        vector<int> dist(n);
        dist[0] = 0;
        DFS(0, -1, adjList, dist);

        int farthest = getFarthest(n, dist);

        vector<int> dist2(n);
        dist2[farthest] = 0;
        DFS(farthest, -1, adjList, dist2);

        int farthest2 = getFarthest(n, dist2);

        vector<int> dist3(n);
        dist3[farthest2] = 0;
        DFS(farthest2, -1, adjList, dist3);

        vector<int> result = getMinMaxDistance(n, dist2, dist3);
        return result;
    }

    vector<int> getMinMaxDistance(int n, vector<int> &dist, vector<int> &dist2) {
        vector<int> result = {0};
        int mini = max(dist[0], dist2[0]);

        for (int i = 1; i < n; i++) {
            int height = max(dist[i], dist2[i]);
            if (height < mini) {
                mini = height;
                result.clear();
                result.push_back(i);
            } else if (height == mini) {
                result.push_back(i);
            }
        }
        return result;
    }

    void DFS(int node, int par, vector<vector<int>> &adjList, vector<int> &dist) {
        for (auto viz : adjList[node]) {
            if (viz != par) {
                dist[viz] = 1 + dist[node];
                DFS(viz, node, adjList, dist);
            }
        }
    }

    int getFarthest(int n, vector<int> &dist) {
        int result = 0;
        for (int i = 0; i < n; i++) {
            if (dist[i] > dist[result]) {
                result = i;
            }
        }
        return result;
    }

    vector<vector<int>> buildGraph(int n, vector<vector<int>>& edges) {
        vector<vector<int>> adjList(n);

        for (auto e : edges) {
            int a = e[0];
            int b = e[1];

            adjList[a].push_back(b);
            adjList[b].push_back(a);
        }
        return adjList;
    }
};
