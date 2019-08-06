class Solution {
public:
    bool canJump(vector<int>& nums) {
        vector<int>used(nums.size());
        queue<int> reachedPositions;
        
        if (nums.size() == 0) {
            return true;
        }
        
        reachedPositions.push(0);
        used[0] = true;
        while(!reachedPositions.empty()) {
            int pos = reachedPositions.front();
            reachedPositions.pop();
            
            if (pos == (int)nums.size()) {
                break;
            }
            
            for (int i = -nums[pos]; i <= nums[pos]; i++) {
                if (0 <= i + pos && i + pos < (int)nums.size()) {
                    if (!used[i + pos]) {
                        used[i + pos] = 1;
                        reachedPositions.push(i + pos);
                    }
                }
            }
        }
        return used[nums.size() - 1];
    }
};