/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return DFS(root, root.Val)
}

func DFS(root *TreeNode, largest int) int {
    if root == nil {
        return 0
    }

    result := 0
    if root.Val >= largest {
        result++
    }

    newLargest := max(largest, root.Val)
    result += DFS(root.Left, newLargest)
    result += DFS(root.Right, newLargest)

    return result
}
