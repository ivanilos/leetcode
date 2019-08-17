class Solution {
public:
    int canCompleteCircuit(vector<int>& gas, vector<int>& cost) {
        int tank = 0;
        int run = 0;
        int start = 0;
        int stations = (int)gas.size();
        for (int i = 0; i < 2 * (int)gas.size(); i++) {
            int next = i % stations;

            if (tank + gas[next] >= cost[next]) {
                tank += gas[next] - cost[next];
                run++;
            } else {
                run = 0;
                tank = 0;
                start = (i + 1) % stations;
            }
            if (run >= stations) {
                return start;
            }
        }
        return -1;
    }
};