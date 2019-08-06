class Solution {
public:
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        map<int, int> freq;
        
        for (int num : nums) {
            freq[num]++;
        }
        
        vector<int> perm;
        vector<vector<int>> perms;
        gen(0, nums, freq, perm, perms);
        return perms;
    }
    
private:
    void gen(int step, vector<int>& nums, map<int, int>& freq, vector<int>& perm, vector<vector<int>>& perms) {
        if (step >= (int)nums.size()) {
            perms.push_back(perm);
            return;
        }
        
        for (auto val : freq) {
            if (val.second > 0) {
                freq[val.first]--;
                perm.push_back(val.first);
                gen(step + 1, nums, freq, perm, perms);
                perm.pop_back();
                freq[val.first]++;
            }
        }
    }
};