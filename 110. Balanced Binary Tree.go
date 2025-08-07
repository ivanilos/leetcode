/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    balanced, _ := _isBalanced(root)

    return balanced
}

func _isBalanced(root *TreeNode) (bool, int) {
    if root == nil {
        return true, 0
    }

    leftBalanced, leftDepth := _isBalanced(root.Left)
    rightBalanced, rightDepth := _isBalanced(root.Right)

    return leftBalanced && rightBalanced && diffByAtMostOne(leftDepth, rightDepth), 1 + max(leftDepth, rightDepth)
}

func diffByAtMostOne(a, b int) bool {
    return max(a, b) - min(a, b) <= 1
}
