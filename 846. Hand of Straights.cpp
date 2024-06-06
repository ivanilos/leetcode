class Solution {
public:
    bool isNStraightHand(vector<int>& hand, int groupSize) {
        multiset<int> freq;

        for (auto h : hand) {
            freq.insert(h);
        }

        bool result = true;
        while(freq.size() > 0) {
            int curFirst = *(freq.begin());
            freq.erase(freq.begin());

            for (int i = 1; i < groupSize; i++) {
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
