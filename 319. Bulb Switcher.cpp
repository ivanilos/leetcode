class Solution {
public:
    int bulbSwitch(int n) {
        if (n <= 0) return 0;
        
        int ans = sqrt(n);
        while(1LL * (ans + 1) * (ans + 1) < n) { ans++; }
        while(1LL * (ans - 1) * (ans - 1) > n) { ans--; }
        
        return ans;
    }
};