class Solution {
public:
    bool isPowerOfThree(int n) {
        if (n == 0) return false;
        
        const double EPS = 1e-6;
        double pot = round(log(n) / log(3));
        return abs(pow(3, pot) - n) < EPS;
    }
};