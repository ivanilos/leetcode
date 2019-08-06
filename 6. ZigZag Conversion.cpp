class Solution {
public:
    string convert(string s, int numRows) {
        if (numRows == 1) return s;
        
        map<pair<int, int>, char> pos; 
        
        int row = 0;
        int col = 0;
        int add_row = 1;
        int add_col = 0;
        for (char c : s) {
            pos[{row, col}] = c;
            printf("row = %d, col = %d, char = %c\n", row, col, c);
        
            row += add_row;
            col += add_col;
            if (row == numRows) {
                col++;
                row -= 2;
                add_row = row == 0 ? 1 : -1;
                add_col = row == 0 ? 0 : 1;
            } else if (row == 0) {
                add_row = 1;
                add_col = 0;
            }
        }
        
        string ans;
        for (auto p : pos) {
            ans += p.second;
        }
        return ans;
    }
};