# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def sumNumbers(self, root: TreeNode) -> int:
        if root == None:
            return 0
        ans = self.solve(root)
        return ans[0]
    
    def solve(self, root: TreeNode) -> int:
        if root.left == None and root.right == None:
            return [root.val, 10]
        
        left = [0, 0]
        right = [0, 0]
        if root.left != None:
            left = self.solve(root.left)
            
        if root.right != None:
            right = self.solve(root.right)
            
        mult = left[1] + right[1]
        ans = [left[0] + right[0] + mult * root.val, mult * 10]
        
        return ans
        