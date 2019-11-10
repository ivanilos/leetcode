class Solution {
public:
    vector<int> build_fail(string s) {
        vector<int> lps(s.size());

        for (int i = 1; i < (int)s.size(); i++) {
            int sz = lps[i - 1];
            while(s[i] != s[sz] && sz > 0) {
                sz = lps[sz - 1];
            }
            if (s[i] == s[sz]) {
                sz++;
            }
            lps[i] = sz;
        }
        return lps;
    }
    
    int strStr(string haystack, string needle) {
        if (needle == "") return 0;
        vector<int> fail = build_fail(needle);
        
        int match = 0;
        for (int i = 0; i < haystack.size(); i++) {
            while(match > 0 && haystack[i] != needle[match]) {
                match = fail[match - 1];
            }
            if (haystack[i] == needle[match]) {
                match++;
                if (match >= needle.size()) return i - match + 1;
            }
        }
        return -1;
    }
};