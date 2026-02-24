/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
    if root == nil {
        return 0
    }

    sum := getSumRootToLeaf(root, 0)

    return sum
}

func getSumRootToLeaf(root *TreeNode, curSum int) int {
    if root.Left == nil && root.Right == nil {
        return curSum + root.Val
    }

    result := 0
    if root.Left != nil {
        result += getSumRootToLeaf(root.Left, 2 * (curSum + root.Val))
    }

    if root.Right != nil {
        result += getSumRootToLeaf(root.Right, 2 * (curSum + root.Val))
    }

    return result
}
