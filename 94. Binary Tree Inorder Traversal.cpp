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
    vector<int> inorderTraversal(TreeNode* root) {
        if (root == nullptr) return {};
        stack<pair<TreeNode*, int> > next_node;
        next_node.push({root, 0});
        
        vector<int> ans;
        while(!next_node.empty()) {
            pair<TreeNode*, int> next = next_node.top();
            next_node.pop();
            
            if (next.second == 0) {
                next_node.push({next.first, 1});
                if (next.first->left != nullptr) { next_node.push({next.first->left, 0}); }
            } else {
                ans.push_back(next.first->val);
                if (next.first->right != nullptr) { next_node.push({next.first->right, 0}); }
            }
        }
        return ans;
    }
};