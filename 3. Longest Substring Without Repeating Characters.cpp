class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int ans = 0;
        
        for (int start = 0; start < (int)s.size(); start++) {
            set<char> got;
            for (int j = start; j < (int)s.size(); j++) {
                if (got.count(s[j])) { break; }
                got.insert(s[j]);
            }
            ans = max(ans, (int)got.size());
        }
        return ans;
    }
};