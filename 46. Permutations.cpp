class Solution {
public:
    vector<vector<int>> permute(vector<int>& nums) {
        vector<bool> used(nums.size(), false);
        vector<int> perm;
        vector<vector<int>> perms;
        gen(0, perm, used, nums, perms);
        return perms;
    }
    
private:
    void gen(int step, vector<int>& perm, vector<bool>& used, vector<int>& nums, vector<vector<int>>& perms) {
        if (step >= (int)nums.size()) { perms.push_back(perm); }
        
        for (int i = 0; i < (int)nums.size(); i++) {
            if (!used[i]) {
                used[i] = true;
                perm.push_back(nums[i]);
                gen(step + 1, perm, used, nums, perms);
                perm.pop_back();
                used[i] = false;
            }
        }
    }
    
};