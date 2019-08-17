class Solution {
public:
    double myPow(double x, long long n) {
        if (abs(x) < EPS) return 0;
        if (n < 0) {
            n *= -1;
            x = 1.0 / x;
        }
        double ans = 1;
        while(n > 0) {
            if (n & 1) {
                ans *= x;
            }
            x *= x;
            n /= 2;
        }
        return ans;
    }
private:
    const double EPS = 1e-8;
};