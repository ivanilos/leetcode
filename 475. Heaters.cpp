class Solution {
public:
    int findRadius(vector<int>& houses, vector<int>& heaters) {
        if (houses.size() == 0) return 0;
        if (heaters.size() == 0) throw BAD_INPUT{};
        
        sort(houses.begin(), houses.end());
        sort(heaters.begin(), heaters.end());
        
        int mini = 0;
        int maxi = max(houses.back(), heaters.back());
        int best = maxi;
        while(mini <= maxi) {
            int rad = (mini + maxi) / 2;
            
            int heat_idx = 0;
            int good = true;
            for (int house : houses) {
                if (heat_idx < heaters.size() && house < heaters[heat_idx] - rad) {
                    good = false;
                    break;
                }
                while(heat_idx < heaters.size() && 
                     !inside(house, heaters[heat_idx] - rad, heaters[heat_idx] + rad)) {
                    heat_idx++;
                }
            }
            if (!good || heat_idx >= heaters.size()) {
                mini = rad + 1;
            } else {
                best = rad;
                maxi = rad - 1;
            }
        }
        return best;
    }
private:
    class BAD_INPUT {};
    
    bool inside(int p, int a, int b) {
        return a <= p && p <= b;
    }
};