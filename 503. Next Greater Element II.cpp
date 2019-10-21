class Solution {
public:
    vector<int> nextGreaterElements(vector<int>& nums) {
        if (nums.size() == 0) return {};
        
        int maxi = *max_element(nums.begin(), nums.end());
        int start = 0;
        for (int i = 0; i < (int)nums.size(); i++) {
            if (nums[i] == maxi) {
                start = i;
                break;
            }
        }
        
        int pos = start;
        vector<int> ans(nums.size());
        do {
            if (nums[pos] == maxi) ans[pos] = -1;
            else {           
                ans[pos] = (pos + 1) % nums.size();
                while(nums[ans[pos]] <= nums[pos]) {
                    ans[pos] = ans[ans[pos]];
                }
            }
            pos = (pos - 1 + (int)nums.size()) % nums.size();
        } while(pos != start);
        
        for (int& val : ans) {
            if (val != -1) val = nums[val];
        }
        
        return ans;
    }
};