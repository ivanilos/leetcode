class Solution {
public:
    bool canConstruct(string ransomNote, string magazine) {
        unordered_map<char, int> freq;
        
        for (char c : ransomNote) {
            freq[c]++;
        }
        for (char c : magazine) {
            freq[c]--;
        }
        for (auto it : freq) {
            if (it.second > 0) return false;
        }
        return true;
    }
};