class Solution {
public:
    int minEatingSpeed(vector<int>& piles, int H) {
        int low = 1;
        int high = *max_element(piles.begin(), piles.end());
        
        int best = 0;
        while(low <= high) {
            int mid = (low + high) / 2;
            int hours = 0;
            for (int i = 0; i < (int)piles.size() && hours <= H; i++) {
                hours += (piles[i] + mid - 1) / mid;
            }
            
            if (hours <= H) {
                best = mid;
                high = mid - 1;
            } else {
                low = mid + 1;
            }
        }
        return best;
    }
};