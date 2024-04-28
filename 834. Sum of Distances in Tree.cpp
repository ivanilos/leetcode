class Solution {
public:
    vector<int> sumOfDistancesInTree(int n, vector<vector<int>>& edges) {
        vector<int> result(n);
        vector<int> subSize(n);

        vector<vector<int>> adjList = buildGraph(n, edges);

        DFS(0, adjList, subSize, result);

        for (auto viz : adjList[0]) {
            solve(viz, adjList, subSize, result, 0);
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

    int DFS(int node, vector<vector<int>>& adjList, vector<int> &subSize, vector<int>& result, int p = -1) {
        subSize[node] = 1;

        for (auto viz : adjList[node]) {
            if (viz != p) {
                subSize[node] += DFS(viz, adjList, subSize, result, node);
                result[node] += result[viz] + subSize[viz];
            }
        }
        return subSize[node];
    }

    void solve(int node, vector<vector<int>>& adjList, vector<int> &subSize, vector<int>& result, int p) {
        int otherSubtreeSize = adjList.size() - subSize[node];
        result[node] = result[p] + otherSubtreeSize - subSize[node];

        for (auto viz : adjList[node]) {
            if (viz != p) {
                solve(viz, adjList, subSize, result, node);
            }
        }
    }
};
