class Solution {
public:
    void update(int& low, int& high, int mid, vector<int>& nums) {
        if (mid - 1 >= 0 && mid + 1 < (int)nums.size()) {
            nums[mid - 1] > nums[mid + 1] ? high = mid - 1 : low = mid + 1;
        } else if (mid - 1 < 0) {
            low = mid + 1;
        } else {
            high = mid - 1;
        }
        return;
    }
    
    bool is_peak(int idx, vector<int>& nums) {
        if (idx - 1 >= 0 && nums[idx - 1] > nums[idx]) return false;
        if (idx + 1 < (int)nums.size() && nums[idx + 1] > nums[idx]) return false;
                                               
        return true;
    }
    
    int findPeakElement(vector<int>& nums) {
        int low = 0;
        int high = (int)nums.size() - 1;
        while(low <= high) {
            int mid = (low + high) / 2;
            if (is_peak(mid, nums)) {
                return mid;
            } else {
                update(low, high, mid, nums);
            }
        }
        return -1;
    }
};