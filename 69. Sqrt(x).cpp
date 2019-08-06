class Solution {
public:
    int mySqrt(int x) {
        if (x < 0) {
            throw InvalidInputException{};
        }
        int low = 0;
        int high = x;
        int best = 0;
        while(low <= high) {
            int mid = low + (high - low) / 2;
            if (1LL * mid * mid <= x) {
                best = mid;
                low = mid + 1;
            } else {
                high = mid - 1;
            }
        }
        return best;
    }
private:
    class InvalidInputException {};
};