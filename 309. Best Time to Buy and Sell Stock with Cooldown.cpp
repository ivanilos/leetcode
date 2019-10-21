class Solution {
public:
    int maxProfit(vector<int>& prices) {
        if (prices.size() <= 1) return 0;
        
        constexpr int INF = 1e9;
        vector<vector<int>> profit(prices.size(), vector<int>(2, -INF));
        
        profit[0][0] = 0;
        profit[0][1] = -prices[0];
        profit[1][0] = max(0, prices[1] - prices[0]);
        profit[1][1] = -min(prices[0], prices[1]);
        for (int day = 2; day < (int)prices.size(); day++) {
            profit[day][0] = max(profit[day - 1][0], profit[day - 1][1] + prices[day]);
            
            profit[day][1] = max(profit[day - 1][1], profit[day - 2][0] - prices[day]);
        }
        return max(profit[(int)prices.size() - 1][0], profit[(int)prices.size() - 1][1]);
    }
};