class Solution {
public:
    long long pickGifts(vector<int>& gifts, int k) {
        priority_queue<int> pq;

        for (auto it : gifts) {
            pq.push(it);
        }

        for (int i = 0; i < k; i++) {
            int next = pq.top();
            pq.pop();

            int left = int(floor(sqrt(next)));
            pq.push(left);
        }

        long long result = 0;
        while (!pq.empty()) {
            result += pq.top();
            pq.pop();
        }
        return result;
    }
};
