class Solution {
public:
    long long countSubarrays(vector<int>& nums, int minK, int maxK) {
            return countSubarrays(nums, minK, maxK, 0, (int)nums.size() - 1);
    }

    long long countSubarrays(vector<int>& nums, int minK, int maxK, int left, int right) {
        if (left > right) return 0;

        for (int i = left; i <= right; i++) {
            if (nums[i] < minK || nums[i] > maxK) {
                return countSubarrays(nums, minK, maxK, left, i - 1) + countSubarrays(nums, minK, maxK, i + 1, right);
            }
        }

        long long result = 0;
        int rightmostMiniIdx = left - 1;
        int rightmostMaxiIdx = left - 1;
        for (int i = left; i <= right; i++) {
            if (nums[i] == minK) {
                rightmostMiniIdx = i;
            }
            if (nums[i] == maxK) {
                rightmostMaxiIdx = i;
            }
            result += min(rightmostMiniIdx, rightmostMaxiIdx) - left + 1;
        }
        return result;
    }
};
