class Solution {
public:
    string intToRoman(int num) {
        string ans;
        
        for (int i = 0; i < symbol.size(); i++) {
            while(num >= values[i]) {
                num -= values[i];
                ans += symbol[i];
            }
        }
        return ans;
        
    }
    
private:
    vector<int> values = {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
    vector<string> symbol = {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};
};