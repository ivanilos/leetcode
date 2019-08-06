class Solution {
public:
    vector<vector<int>> subsetsWithDup(vector<int>& nums) {
        map<int, int> freq;
        
        for (int num : nums) {
            freq[num]++;
        }
        
        vector<vector<int> > ans = {{}};
        for (pair<int, int> val : freq) {
            int initial_size = ans.size();
            for (int i = 1; i <= val.second; i++) {
                for (int j = 0; j < initial_size; j++) {
                    ans.push_back(ans[j]);
                    for (int k = 0; k < i; k++) {
                        ans.back().push_back(val.first);
                    }
                }
            }
        }
        
        return ans;
    }
};