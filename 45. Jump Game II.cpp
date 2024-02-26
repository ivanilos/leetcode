class Solution {
public:
    int jump(vector<int>& nums) {
        int sz = (int)nums.size();
        set<pair<int, int>> min_steps_to_reach;

        min_steps_to_reach.insert({0, 0});
        for (int pos = 0; pos < sz; pos++) {
            pair<int, int> reach = {-1, -1};
            while(reach.second < pos) {
                min_steps_to_reach.erase(reach);
                reach = *min_steps_to_reach.begin();
            }
            min_steps_to_reach.insert({1 + reach.first, pos + nums[pos]});
        }
        return (*min_steps_to_reach.begin()).first;
    }
};