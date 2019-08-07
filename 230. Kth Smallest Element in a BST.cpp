/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    int kthSmallest(TreeNode* root, int k) {
        int ans = -1;
        solve(root, k, ans);
        return ans;
    }
    
private:
    void solve(TreeNode* root, int& k, int& ans) {
        if (k > 0 && root->left) { solve(root->left, k, ans); }
        k--;
        if (k == 0) { ans = root->val; }
        if (k > 0 && root->right) { solve(root->right, k, ans); }
    }
};