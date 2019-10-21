class Solution {
public:
    int count(string& s, int left, int right) {
        int ans = 0;
        while(left >= 0 && right < (int)s.size()) {
            if (s[left] == s[right]) {
                ans++;
                left--;
                right++;
            } else {
                break;
            }
        }
        return ans;
    }
    
    int countSubstrings(string s) {
        int counter = 0;
        for (int i = 0; i < (int)s.size(); i++) {
            counter += count(s, i, i);
            counter += count(s, i, i + 1);
        }
        return counter;
    }
};