class Solution {
public:
    vector<int> majorityElement(vector<int>& nums) {
        map<int, int> freq;
        
        for (int num : nums) {
            freq[num]++;
            
            if (freq.size() == 3) {
                for (auto it : freq) {
                    if (it.second == 0) {
                        freq.erase(it.first);
                    }
                }
            }
        }
        vector<int> cand;
        for (auto it : freq) {
            if (it.second > 0) {
                cand.push_back(it.first);
            }
        }
        
        vector<int> ans;
        vector<int> qt(cand.size());
        for (int num : nums) {
            for (int i = 0; i < cand.size(); i++) {
                if (num == cand[i]) { qt[i]++; }
            }
        }
        
        for (int i = 0; i < (int)qt.size(); i++) {
            if (qt[i] > nums.size() / 3) { ans.push_back(cand[i]); }
        }
        return ans;
    }
};