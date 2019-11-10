class Solution {
public:
    int flipgame(vector<int>& fronts, vector<int>& backs) {
        set<int> blocked;
        
        for (int i = 0; i < (int)fronts.size(); i++) {
            if (fronts[i] == backs[i]) {
                blocked.insert(fronts[i]);
            }
        }
        
        int a = solve(fronts, blocked);
        int b = solve(backs, blocked);
        if (a > b) swap(a, b);
        
        int ans = a == 0 ? b : a;
        return ans;
    }
    
private:
    int solve(vector<int>&v, set<int>& blocked) {
        int ans = 0;
        for (int num : v) {
            if (!blocked.count(num)) {
                if (ans == 0 || num < ans) {
                    ans = num;
                }
            }
        }
        return ans;
    }
};