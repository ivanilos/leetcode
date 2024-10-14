class Solution {
public:
    long long maxKelements(vector<int>& nums, int k) {
        priority_queue<int> pq;

        for (int num : nums) {
            pq.push(num);
        }

        long long result = 0;
        for (int i = 0; i < k; i++) {
            int next = pq.top();
            pq.pop();

            result += next;
            pq.push((next + 2) / 3);
        }

        return result;
    }
};
