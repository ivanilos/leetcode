class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        unordered_set<int> vals;
        for (int num : nums) {
            if (vals.count(num)) return true;
            
            vals.insert(num);
        }
        return false;
    }
};