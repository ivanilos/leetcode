class Solution {
public:
    bool reorderedPowerOf2(int N) {
        string s = to_string(N);
        sort(s.begin(), s.end());
        
        for (int i = 0; (1LL << i) <= 10LL * N; i++) {
            string arrange = to_string(1LL << i);
            sort(arrange.begin(), arrange.end());
            if (s == arrange) {
                return true;
            }
        }
        return false;
    }
};