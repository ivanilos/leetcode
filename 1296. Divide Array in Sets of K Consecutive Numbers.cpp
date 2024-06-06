class Solution {
public:
    bool isPossibleDivide(vector<int>& nums, int k) {
        multiset<int> freq;

        for (auto h : nums) {
            freq.insert(h);
        }

        bool result = true;
        while(freq.size() > 0) {
            int curFirst = *(freq.begin());
            freq.erase(freq.begin());

            for (int i = 1; i < k; i++) {
                int nextVal = curFirst + 1;
                auto next = freq.find(nextVal);
                if (next != freq.end()) {
                    freq.erase(next);
                    curFirst = nextVal;
                } else {
                    result = false;
                    break;
                }
            }
        }
        return result;
    }
};
