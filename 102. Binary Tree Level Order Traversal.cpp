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
    vector<vector<int>> levelOrder(TreeNode* root) {
        if (!root) return {};
        
        vector<vector<int>> ans = {};
        queue<TreeNode*> cur_level;
        cur_level.push(root);
        while(!cur_level.empty()) {
            queue<TreeNode*> next_level;
            vector<int> level_vals;
            
            while(!cur_level.empty()) {
                TreeNode* node = cur_level.front();
                cur_level.pop();
                level_vals.push_back(node->val);
                
                if (node->left) { next_level.push(node->left); }
                if (node->right) { next_level.push(node->right); }
                
            }
            cur_level = move(next_level);
            ans.push_back(level_vals);
        }
        return ans;
    }
};