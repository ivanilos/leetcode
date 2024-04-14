/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    return solve(root.Left, 1) + solve(root.Right, 0);
}

func solve(root *TreeNode, canAddCur int) int {
    if (root == nil) {
        return 0
    } else if (root.Left == nil && root.Right == nil) {
        return canAddCur * root.Val
    } else {
        return solve(root.Left, 1) + solve(root.Right, 0);
    }
}
