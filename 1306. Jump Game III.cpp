class Solution {
public:
    bool canReach(vector<int>& arr, int start) {
        int sz = (int)arr.size();
        vector<bool> can(sz, false);
        
        queue<int> next_reachable;
        next_reachable.push(start);
        can[start] = true;
        while(!next_reachable.empty()) {
            int pos = next_reachable.front();
            next_reachable.pop();
            
            if (arr[pos] == 0) return true;
            
            int jump_left_pos = pos - arr[pos];
            if (jump_left_pos >= 0 && !can[jump_left_pos]) {
                can[jump_left_pos] = true;
                next_reachable.push(jump_left_pos);
            }
            
            int jump_right_pos = pos + arr[pos];
            if (jump_right_pos < sz && !can[jump_right_pos]) {
                can[jump_right_pos] = true;
                next_reachable.push(jump_right_pos);
            }
        }
        return false;
        
    }
};