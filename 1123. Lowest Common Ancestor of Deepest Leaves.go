/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
    _, lca := lcaDeepestLeavesWithDepth(root)

    return lca
}

func lcaDeepestLeavesWithDepth(root *TreeNode) (int, *TreeNode) {
    if root == nil {
        return 0, nil
    }
    leftDepth, leftLCA := lcaDeepestLeavesWithDepth(root.Left)
    rightDepth, rightLCA := lcaDeepestLeavesWithDepth(root.Right)

    if leftDepth == rightDepth {
        return 1 + leftDepth, root
    } else if leftDepth > rightDepth {
        return 1 + leftDepth, leftLCA
    } else {
        return 1 + rightDepth, rightLCA
    }
}
