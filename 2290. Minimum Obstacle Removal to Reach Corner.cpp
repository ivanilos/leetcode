class Solution {
public:
    vector<int> dx = vector<int>({1, 0, -1, 0});
    vector<int> dy = vector<int>({0, 1, 0, -1});

    int minimumObstacles(vector<vector<int>>& grid) {
        int rows = grid.size();
        int cols = grid[0].size();

        vector<vector<int>> dist = vector<vector<int>>(rows, vector<int>(cols, rows * cols));
        dist[0][0] = 0;
        deque<pair<int, int>> dq;

        dq.push_back({0, 0});
        while (!dq.empty()) {
            pair<int, int> next = dq.front();
            dq.pop_front();

            int x = next.first;
            int y = next.second;

            for (int i = 0; i < 4; i++) {
                int nx = x + dx[i];
                int ny = y + dy[i];

                if (isIn(nx, ny, rows, cols)) {
                    int newDist = dist[x][y] + grid[nx][ny];

                    if (newDist < dist[nx][ny]) {
                        dist[nx][ny] = newDist;
                        if (grid[nx][ny] == 0) {
                            dq.push_front({nx, ny});
                        } else {
                            dq.push_back({nx, ny});
                        }
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
};
