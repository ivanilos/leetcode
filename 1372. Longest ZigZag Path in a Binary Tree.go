/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
    result := 0

    zigZagLen(root, &result)

    return result - 1
}

func zigZagLen(root *TreeNode, result *int) (int, int) {
    if root == nil {
        return 0, 0
    }

    leftToLeft, leftToRight := zigZagLen(root.Left, result)
    rightToLeft, rightToRight := zigZagLen(root.Right, result)

    *result = max(*result, leftToLeft)
    *result = max(*result, 1 + leftToRight)
    *result = max(*result, 1 + rightToLeft)
    *result = max(*result, rightToRight)

    return 1 + leftToRight, 1 + rightToLeft
}
