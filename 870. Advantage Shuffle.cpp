class Solution {
public:
    vector<int> advantageCount(vector<int>& A, vector<int>& B) {
        vector<int> ans;
        
        multiset<int> a;
        for (int num : A) {
            a.insert(num);
        }
        
        for (int num : B) {
            auto it = a.lower_bound(num + 1);
            if (it == a.end()) {
                it = a.begin();
            }
            ans.push_back(*it);
            a.erase(it);
        }
        return ans;
    }
};