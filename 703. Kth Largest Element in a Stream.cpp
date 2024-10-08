class KthLargest {
public:
    KthLargest(int k, vector<int>& nums) {
        this->k = k;

        for (auto num : nums) {
            pq.push(-num);
            if (pq.size() > k) {
                pq.pop();
            }
        }
    }
    
    int add(int val) {
        pq.push(-val);
        if (pq.size() > k) {
            pq.pop();
        }

        return -pq.top();
    }

private:
    priority_queue<int> pq;
    int k;
};

/**
 * Your KthLargest object will be instantiated and called as such:
 * KthLargest* obj = new KthLargest(k, nums);
 * int param_1 = obj->add(val);
 */
