struct Node {
    int x, y;
    int val;

    Node() {}
    Node(int x, int y, int val):x(x), y(y), val(val) {}
};

bool operator <(const Node& a, const Node& b) {
    return a.val > b.val;
}

const int MAX_QUERY_VAL = int(1e6);

int dx[] = {0, 1, 0, -1};
int dy[] = {1, 0, -1, 0};

class Solution {
public:
    vector<int> maxPoints(vector<vector<int>>& grid, vector<int>& queries) {
        vector<int> point_result(MAX_QUERY_VAL + 2);
        
        int rows = grid.size();
        int cols = grid[0].size();

        vector<vector<int>> visited = vector<vector<int>>(rows, vector<int>(cols));
        priority_queue<Node> pq;
        pq.push(Node(0, 0, grid[0][0]));
        visited[0][0] = true;

        int maxi = 0;
        while (!pq.empty()) {
            Node cur = pq.top();
            pq.pop();

            maxi = max(maxi, cur.val + 1);

            point_result[maxi]++;

            for (int i = 0; i < 4; i++) {
                int nx = cur.x + dx[i];
                int ny = cur.y + dy[i];
                if (is_in(nx, ny, rows, cols) && !visited[nx][ny]) {
                    visited[nx][ny] = true;
                    pq.push(Node(nx, ny, grid[nx][ny]));
                }
            }
        }

        for (int i = 1; i < MAX_QUERY_VAL + 2; i++) {
            point_result[i] += point_result[i - 1];
        }

        vector<int> result(queries.size());
        for (int i = 0; i < queries.size(); i++) {
            result[i] = point_result[queries[i]];
        }

        return result;
    }

private:
    bool is_in(int x, int y, int rows, int cols) {
        return 0 <= x && x < rows && 0 <= y && y < cols;
    }
};
