/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    result := []int{}

    solve(root, 0, &result)
    return result
}

func solve(root *TreeNode, curDepth int, result *[]int) {
    if root == nil {
        return
    }

    if len(*result) <= curDepth {
        *result = append(*result, root.Val)
    } else {
        (*result)[curDepth] = max((*result)[curDepth], root.Val)
    }

    solve(root.Left, curDepth + 1, result)
    solve(root.Right, curDepth + 1, result)
}
