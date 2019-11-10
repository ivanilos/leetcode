class Solution {
public:
    int longestSubstring(string s, int k) {
        return longestSubstring(0, (int)s.size() - 1, s, k);
    }
private:
    int longestSubstring(int left, int right, const string& s, int k) {
        if (left > right) return 0;
        
        unordered_map<char, int> freq;
        unordered_set<char> bad;
        
        for (int i = left; i <= right; i++) {
            freq[s[i]]++;
        }
        for (auto& it : freq) {
            if (it.second < k) bad.insert(it.first);
        }
        
        int max_len = 0;
        if (bad.empty()) return right - left + 1;
        else {
            int next_left = 0;
            for (int i = left; i <= right; i++) {
                if (bad.count(s[i])) {
                    max_len = max(max_len, longestSubstring(next_left, i - 1, s, k));
                    next_left = i + 1;
                }
            }
            max_len = max(max_len, longestSubstring(next_left, right, s, k));
        }    
        return max_len;   
    }
};