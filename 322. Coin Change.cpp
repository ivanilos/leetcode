class Solution {
public:
    int coinChange(vector<int>& coins, int amount) {
        vector<int> need(amount + 1, amount + 1);
        need[0] = 0;
        
        for (int coin : coins) {
            for (int qt = coin; qt <= amount; qt++) {
                need[qt] = min(need[qt], 1 + need[qt - coin]);
            }
        }
        return need[amount] < amount + 1 ? need[amount] : -1;
    }
};