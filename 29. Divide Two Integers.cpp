class Solution {
public:
    int divide(long long dividend, long long divisor) {
        int s = sign(dividend) != sign(divisor) ? -1 : 1;
        
        dividend = abs(dividend);
        divisor = abs(divisor);
        long long ans = 0;
        for (int i = 31; i >= 0; i--) {
            if (dividend >= ((long long)divisor << i) && (((long long)divisor << i) > 0)) {
                ans += 1 << i;
                dividend -= divisor << i;
            }
        }
        if (s == 1 && ans == (1 << 31)) return (1 << 31) - 1;
        return s == 1 ? (int)ans : (int)-ans;
    }
    
private:
    int sign(int x) {
        if (x < 0) return -1;
        return 1;
    }
};