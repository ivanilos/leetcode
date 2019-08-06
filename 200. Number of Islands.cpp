class Solution {
public:
    int numIslands(vector<vector<char>>& grid) {
        int ans = 0;
        for (int i = 0; i < grid.size(); i++) {
            for (int j = 0; j < (int)grid[i].size(); j++) {
                if (grid[i][j] == '1') {
                    ans++;
                    BFS(i, j, grid);
                }
            }
        }
        return ans;
    }
    
private:
    vector<int> dx = {0, 1, 0, -1};
    vector<int> dy = {1, 0, -1, 0};
    
    bool is_in(int x, int y, vector<vector<char>>& grid) {
        return 0 <= x && x < grid.size() && 0 <= y && y < grid[0].size();
    }
    
    void BFS(int x, int y, vector<vector<char>>& grid) {
        queue<pair<int, int> > fila;
        grid[x][y] = '0';
        fila.push({x, y});
        
       while(!fila.empty()) {
           x = fila.front().first;
           y = fila.front().second;
           fila.pop();
           
            for (int i = 0; i < (int)dx.size(); i++) {
                int nx = x + dx[i];
                int ny = y + dy[i];
                
                if (is_in(nx, ny, grid) && grid[nx][ny] == '1') {
                    grid[nx][ny] = '0';
                    fila.push({nx, ny});
                }
            }
        }
    }
};