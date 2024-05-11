class Solution {
public:
    double mincostToHireWorkers(vector<int>& quality, vector<int>& wage, int k) {
        int n = quality.size();
        vector<pair<double, int>> workers(n);

        for (int i = 0; i < n; i++) {
            workers[i] = {1.0 * wage[i] / quality[i], quality[i]};
        }

        sort(workers.begin(), workers.end());

        priority_queue<int> quality_chosen;
        double best = 1e9;

        double cur_slope = 0;
        double old_slope = 0;
        double cur_result = 0;
        int quality_sum = 0;
        for (int i = 0; i < n; i++) {
            cur_slope = workers[i].first;
            cur_result = cur_result + quality_sum * (cur_slope - old_slope);
            cur_result += cur_slope * workers[i].second;

            old_slope = cur_slope;

            quality_chosen.push(workers[i].second);
            quality_sum += workers[i].second;
            if (quality_chosen.size() > k) {
                int top = quality_chosen.top();
                quality_chosen.pop();
                cur_result -= cur_slope * top;
                quality_sum -= top;
            }
            if (quality_chosen.size() == k && cur_result < best) {
                best = cur_result;
            }
        }
        return best;
    }
};
