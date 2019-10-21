class Solution {
public:
    double probabilityOfHeads(vector<double>& prob, int target) {
        vector<double> p(target + 2);
        
        p[0] = 1;
        for (int coin = 0; coin < prob.size(); coin++) {
            for (int i = target; i > 0; i--) {
                p[i] = (1.0 - prob[coin]) * p[i] + prob[coin] * p[i - 1];
            }
            p[0] = (1.0 - prob[coin]) * p[0];
        }
        return p[target];
    }
};