// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
        int low = 1;
        int high = n;
        int first_bad = 1;
        while(low <= high) {
            int mid = low + (high - low) / 2;
            if (isBadVersion(mid)) {
                high = mid - 1;
                first_bad = mid;
            } else {
                low = mid + 1;
            }
        }
        return first_bad;
    }
};