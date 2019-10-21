class Solution {
public:
    static bool comp(const int& a, const int& b) {
        string s1 = to_string(a) + to_string(b);
        string s2 = to_string(b) + to_string(a);
        return s1 > s2;
    }
    
    string largestNumber(vector<int>& nums) {
        sort(nums.begin(), nums.end(), comp);
        
        string ans;
        for (int num : nums) {
            ans += to_string(num);
        }
        if (ans[0] == '0') return "0";
        return ans;
    }
};