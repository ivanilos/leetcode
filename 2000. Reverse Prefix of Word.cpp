class Solution {
public:
    string reversePrefix(string word, char ch) {
        int lastPos = 0;
        for (int i = 0; i < word.size(); i++) {
            if (word[i] == ch) {
                lastPos = i + 1;
                break;
            }
        }
        reverse(word.begin(), word.begin() + lastPos);
        return word;
    }
};
