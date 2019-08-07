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
    int rob(TreeNode* root) {
        pair<int, int> ans = solve(root);
        return max(ans.first, ans.second);
    }
    
private:
    pair<int, int> solve(TreeNode* root) {
        if (root == nullptr) return {0, 0};
        
        pair<int, int> left = solve(root->left);
        pair<int, int> right = solve(root->right);
        
        int take = root->val + left.first + right.first;
        int not_take = max(left.first, left.second) + max(right.first, right.second);
        return {not_take, take};
    }
};