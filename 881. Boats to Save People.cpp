class Solution {
public:
    int numRescueBoats(vector<int>& people, int limit) {
        sort(people.begin(), people.end());
        int ans = 0;
        
        for (int l = 0, r = (int)people.size() - 1; l <= r; l++, r--) {
            if (people[l] + people[r] > limit) {
                l--;
            }
            ans++;
        }
        return ans;
    }
};