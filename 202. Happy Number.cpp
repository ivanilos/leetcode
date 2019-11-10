class Solution {
public:
    bool isHappy(int n) {
        unordered_set<int> found;
        found.insert(n);
        while(n != 1) {
            string num = to_string(n);
            int next_num = 0;
            for (int i = 0; i < num.size(); i++) {
                next_num += (num[i] - '0') * (num[i] - '0');
            }
            n = next_num;
                
            if (found.count(n)) return false;
            found.insert(n);
        }
        return true;
    }
};