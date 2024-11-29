struct Node {
    int x;
    int y;
    int val;

    Node() {}
    Node(int x, int y, int val):x(x), y(y), val(val) {}
};

bool operator <(const Node& a, const Node& b) {
    return tie(a.val, a.x, a.y) < tie(b.val, b.x, b.y);
}

class Solution {
public:

    vector<int> dx = {0, 1, 0, -1};
    vector<int> dy = {1, 0, -1, 0};

    int minimumTime(vector<vector<int>>& grid) {
        if (grid[0][1] > 1 && grid[1][0] > 1) {
            return -1;
        }


        int rows = grid.size();
        int cols = grid[0].size();
        const int INF = (int)1e9;

        vector<vector<int>> dist = vector<vector<int>>(rows, vector<int>(cols, INF));

        dist[0][0] = 0;
        
        priority_queue<Node> pq;
        pq.push(Node(0, 0, 0));

        while (!pq.empty()) {
            Node next = pq.top();
            pq.pop();

            if (-next.val > dist[next.x][next.y]) {
                continue;
            }

            for (int i = 0; i < 4; i++) {
                int nx = next.x + dx[i];
                int ny = next.y + dy[i];

                if (isIn(nx, ny, rows, cols)) {
                    int newDist = getNewDist(nx, ny, dist[next.x][next.y], grid[nx][ny]);

                    if (dist[nx][ny] > newDist) {
                        dist[nx][ny] = newDist;
                        pq.push(Node(nx, ny, -newDist));
                    }
                }
            }
        }

        return dist[rows - 1][cols - 1];
    }

private:
    bool isIn(int x, int y, int rows, int cols) {
        return 0 <= x && x < rows && 0 <= y && y < cols;
    }

    int getNewDist(int x, int y, int oldDist, int minDist) {
        int minNewDist = minDist + (minDist % 2 != (x + y) % 2);
        int fromOldDist = oldDist + 1 + ((oldDist + 1) % 2 != (x + y) % 2);
        return max(minNewDist, fromOldDist);
    }
};
