class Solution {
public:
    string longestPalindrome(string s) {
        int l = 0;
        int r = -1;
        string ans;
        // odd
        for (int i = 0; i < (int)s.size(); i++) {
            int qt = -1;
            for (int j = i, k = i;; j--, k++) {
                if (j >= 0 && k < (int)s.size() && s[j] == s[k]) { qt += 2; } 
                else {
                    if (qt > r - l + 1) {
                        r = k - 1;
                        l = j + 1;
                        ans = s.substr(j + 1, qt);
                    }
                    break;
                }
            }
        }
    
        // even
        for (int i = 0; i < (int)s.size() - 1; i++) {
            int qt = 0;
            for (int j = i, k = i + 1;; j--, k++) {
                if (j >= 0 && k < (int)s.size() && s[j] == s[k]) { qt += 2; }
                else {
                    if (qt > r - l + 1) {
                        r = k - 1;
                        l = j + 1;
                        ans = s.substr(j + 1, qt);
                    }
                    break;
                }
            }
        }
        return ans;
    }
};