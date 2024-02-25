class Solution {
public:
    bool canTraverseAllPairs(vector<int>& nums) {
        if (nums.size() == 1) {
            return true;
        }

        if (nums.size() <= 200) {
            return solveSmall(nums);
        } else {
            return solveBig(nums);
        }
    }


private:
    const int MAX_NUM = 100000;
    vector<int> subset;
    vector<int> rnk;

    void init() {
        subset.resize(MAX_NUM + 1);
        rnk.resize(MAX_NUM + 1);
        for (int i = 0; i <= MAX_NUM; i++) {
            subset[i] = i;
            rnk[i] = i;
        }
    }

    int find(int i) {
        if (subset[i] == i) return i;
        return subset[i] = find(subset[i]);
    }

    void unite(int x, int y) {
        x = find(x);
        y = find(y);

        if (x == y) return;

        if (rnk[x] < rnk[y]) swap(x, y);

        subset[y] = subset[x];
        if (rnk[x] == rnk[y]) {
            rnk[x]++;
        }
    }

    int gcd(int a, int b) {
        if (b == 0) {
            return a;
        }
        return gcd(b, a % b);
    }

    bool solveSmall(vector<int>& nums) {
        subset.resize(nums.size());
        rnk.resize(nums.size());
        for (int i = 0; i < (int)nums.size(); i++) {
            subset[i] = i;
            rnk[i] = 1;
        }

        for (int i = 0; i < (int)nums.size(); i++) {
            for (int j = i + 1; j < (int)nums.size(); j++) {
                if (gcd(nums[i], nums[j]) > 1) {
                    unite(i, j);
                }
            }
        }

        for (int i = 1; i < (int)nums.size(); i++) {
            if (find(i) != find(0)) {
                return false;
            }
        }
        return true;
    }

    bool solveBig(vector<int>& nums) {
        init();

        vector<int> hasNum(MAX_NUM + 1);
        for (auto num : nums) {
            if (num == 1) {
                return false;
            }
            hasNum[num] = true;
        }

        vector<bool> notVisited(MAX_NUM + 1, true);
        for (int i = 2; i <= MAX_NUM; i++) {
            vector<int> divs;

            if (notVisited[i]) {
                for (int j = i; j <= MAX_NUM; j += i) {
                    notVisited[j] = false;

                    if (hasNum[j]) {
                        divs.push_back(j);
                    }
                }
            }
            for (auto div : divs) {
                unite(div, divs[0]);
            }
        }

        for (auto num : nums) {
            if (find(num) != find(nums[0])) {
                return false;
            }
        }

        return true;
    }
};
