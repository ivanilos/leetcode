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
    bool hasPathSum(TreeNode* root, int sum) {
        return solve(root, sum);
    }
private:
    bool solve(TreeNode* root, int sum_needed) {
        if (root == nullptr) return false;
        if (!root->left && !root->right) return sum_needed == root->val;
        
        int next_needed = sum_needed - root->val;
        return solve(root->left, next_needed) || solve(root->right, next_needed);
    }
};