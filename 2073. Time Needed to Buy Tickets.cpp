class Solution {
public:
    int timeRequiredToBuy(vector<int>& tickets, int k) {
        map<int, int> freq;
        for (auto t : tickets) {
            freq[t]++;
        }

        int result = 0;
        int cur_freq = 0;
        int cur_qt = (int)tickets.size();
        for (auto f : freq) {
            if (f.first <= tickets[k]) {
                result += cur_qt * (f.first - cur_freq);
                cur_freq = f.first;
                cur_qt -= f.second;
            } else {
                break;
            }
        }
        for (int i = k + 1; i < (int)tickets.size(); i++) {
            if (tickets[i] >= tickets[k]) {
                result--;
            }
        }
        return result;
    }
};
