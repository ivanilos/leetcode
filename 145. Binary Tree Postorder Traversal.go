/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    left := postorderTraversal(root.Left)
    right := postorderTraversal(root.Right)

    result := append(left, right...)
    result = append(result, root.Val)
    return result
}
