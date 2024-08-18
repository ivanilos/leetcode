class Solution {
public:
    int nthUglyNumber(int n) {
        set<long long> ugly;

        ugly.insert(1);

        int got = 0;
        long long result = 1;
        while(got < n) {
            got++;
            long long small = *ugly.begin();
            result = small;
            ugly.erase(small);
            ugly.insert(small * 2);
            ugly.insert(small * 3);
            ugly.insert(small * 5);
        }
        return result;
    }
};
