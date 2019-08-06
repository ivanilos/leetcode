class Solution {
public:
    vector<int> grayCode(int n) {
        if (n == 0) return {0};
        vector<int> ans = {0, 1};
        
        for (int bits = 2; bits <= n; bits++) {
            int initial_size = ans.size();
            for (int i = 0; i < initial_size; i++) {
                ans.push_back(ans[i]);
            }
            
            for (int i = 0; i < ans.size() / 2; i++) {
                ans[i] <<= 1;    
            }
            for (int i = ans.size() / 2; i < ans.size(); i++) {
                ans[i] = (ans[i] << 1) + 1;
            }
            reverse(ans.begin() + ans.size() / 2, ans.end());
        }
        return ans;
    }
};