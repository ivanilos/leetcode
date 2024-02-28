/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    _, val := _findBottomLeftValue(root)
    return val
}

func _findBottomLeftValue(root *TreeNode) (int, int) {
    if root.Left == nil && root.Right == nil {
        return 0, root.Val
    }
    var leftDepth, leftVal, rightDepth, rightVal int

    if root.Left != nil {
        leftDepth, leftVal = _findBottomLeftValue(root.Left)
        leftDepth += 1
    }

    if root.Right != nil {
        rightDepth, rightVal = _findBottomLeftValue(root.Right)
        rightDepth += 1
    }

    if leftDepth >= rightDepth {
        return leftDepth, leftVal
    } else {
        return rightDepth, rightVal
    }
}
