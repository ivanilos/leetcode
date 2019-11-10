class Solution {
public:
    int numRabbits(vector<int>& answers) {
        map<int, int> freq;
        for (int answer : answers) {
            freq[answer]++;
        }
        
        int ans = 0;
        for (auto it : freq) {
            ans += ((it.second + it.first) / (it.first + 1)) * (it.first + 1);
        }
        return ans;
    }
};