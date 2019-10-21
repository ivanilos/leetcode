class Solution {
public:
    long long titleToNumber(string s) {
        constexpr int BASE = 26;
        long long column = 0;
        for (char c : s) {
            column = BASE * column + c - 'A' + 1;
        }
        return column;
    }
};