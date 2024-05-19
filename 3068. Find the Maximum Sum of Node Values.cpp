class Solution {
public:
    vector<int> g[20005];
    vector<int> vals;
    int K;

    long long dp[20005][2];
    
    long long maximumValueSum(vector<int>& nums, int k, vector<vector<int>>& edges) {
        vals = nums;
        K = k;

        for (int i = 0; i <= nums.size(); i++) {
            for (int j = 0; j < 2; j++) {
                dp[i][j] = -1;
            }
        }
        
        for (auto edge : edges) {
            int a = edge[0];
            int b = edge[1];
            
            g[a].push_back(b);
            g[b].push_back(a);
        }
        
        long long ans = DFS(0, 0, -1);
        
        return ans;
        
    }
    
    long long DFS(int node, int xored, int par) {
        if (dp[node][xored] != -1) {
            return dp[node][xored];
        }

        vector<long long> sub;
        vector<long long> sub_xored;
        
        vector<pair<long long, long long>> diff;
        
        for (auto viz : g[node]) {
            if (viz != par) {
                long long a = DFS(viz, 0, node);
                long long b = DFS(viz, 1, node);
                
                sub.push_back(a);
                sub_xored.push_back(b);
            }
        }
        
        for (int i = 0; i < (int)sub.size(); i++) {
            long long difference = sub_xored[i] - sub[i];
            diff.push_back({difference, sub[i]});
        }
        
        sort(diff.begin(), diff.end());
        reverse(diff.begin(), diff.end());
        
        long long cur = vals[node];
        if (xored) {
            cur ^= K;
        }
        
        // leaf node
        if (diff.size() == 0) {
            dp[node][xored] = cur;
            return cur;
        }
        
        
        long long sum_not_xored = 0;
        for (auto it : diff) {
            sum_not_xored += it.second;
        }
        
        long long &best = dp[node][xored];
        best = sum_not_xored + cur;
        
        for (int i = 0; i < (int)diff.size(); i++) {
            long long aux = sum_not_xored + diff[i].first;
            sum_not_xored += diff[i].first;
            if (i % 2 == 0) {
                aux += cur ^ K;
            } else {
                aux += cur;
            }

            if (aux > best) {
                best = aux;
            }
        }
        
        return best;
    }
};
