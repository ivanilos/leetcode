/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    result := root.Val

    solve(root, &result)

    return result
}

func solve(root *TreeNode, result *int) int {
    if root == nil {
        return 0
    }
    leftRootMaxSum := solve(root.Left, result)
    rightRootMaxSum := solve(root.Right, result)

    leftSum := leftRootMaxSum + root.Val
    rightSum := rightRootMaxSum + root.Val
    NodeSum := leftRootMaxSum + rightRootMaxSum + root.Val

    *result = max(*result, leftSum)
    *result = max(*result, rightSum)
    *result = max(*result, NodeSum)
    *result = max(*result, root.Val)

    return max(leftSum, max(rightSum, root.Val))
}
