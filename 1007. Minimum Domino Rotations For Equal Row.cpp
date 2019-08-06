class Solution {
public:
    int minDominoRotations(vector<int>& A, vector<int>& B) {
        map<int, int> top;
        map<int, int> bottom;
        map<int, int> eq;
        
        for (int i = 0; i < (int)A.size(); i++) {
            top[A[i]]++;
            bottom[B[i]]++;
            if (A[i] == B[i]) {
                eq[A[i]]++;
            }
        }
        
        int ans = (int)A.size() + 1;
        ans = min(ans, solve(top, bottom, eq, (int)A.size()));
        ans = min(ans, solve(bottom, top, eq, (int)A.size()));

        if (ans > A.size()) { return -1; }
        return ans;
    }
    
private:
    int solve(map<int,int>& top, map<int,int>& bottom, map<int,int>& eq, int size) {
        int ans = size + 1;
        for (pair<int, int> val : top) {
            int num = val.first;
            if (top[num] + bottom[num] - eq[num] == size) {
                ans = min(ans, min(top[num], bottom[num]) - eq[num]);
            }
        }
        return ans;
    }
};