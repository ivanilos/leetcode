struct Node {
    int x;
    int y;
    int cost;

    Node() {}
    Node(int x, int y, int cost):x(x), y(y), cost(cost) {}
};

bool operator <(const Node& a, const Node& b) {
    return a.cost > b.cost;
}


class Solution {
public:
    int INF = int(1e9);

    vector<int> dx = {0, 0, 0, 1, -1};
    vector<int> dy = {0, 1, -1, 0, 0};

    int minCost(vector<vector<int>>& grid) {
        int n = grid.size();
        int m = grid[0].size();
        vector<vector<int>> dp(n, vector<int>(m, INF));

        priority_queue<Node> pq;
        pq.push(Node(0, 0, 0));
        dp[0][0] = 0;

        while(!pq.empty()) {
            Node next = pq.top();
            pq.pop();


            int x = next.x;
            int y = next.y;
            int cost = next.cost;

            if (dp[x][y] < cost) {
                continue;
            }

            for (int i = 1; i <= 4; i++) {
                int nx = x + dx[i];
                int ny = y + dy[i];

                int newCost = cost + (grid[x][y] == i ? 0 : 1);

                if (isIn(nx, ny, n, m) && dp[nx][ny] > newCost) {
                    dp[nx][ny] = newCost;
                    pq.push(Node(nx, ny, newCost));
                }
            }
        }

        return dp[n - 1][m - 1];
    }

private:
    bool isIn(int x, int y, int n, int m) {
        return 0 <= x && x < n && 0 <= y && y < m;
    }
};
