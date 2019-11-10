class Solution {
public:
    void reverseString(vector<char>& s) {
        for (int left = 0, right = (int)s.size() - 1; left < right; left++, right--) {
            swap(s[left], s[right]);
        }
        return;
    }
};