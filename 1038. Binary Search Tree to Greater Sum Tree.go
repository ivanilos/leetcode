/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
    solve(root, 0)
    return root
}

func solve(root *TreeNode, toAdd int) (int, int) {
    if root == nil {
        return 0, 0
    }
    leftChildRightSum, rightChildRightSum := solve(root.Right, toAdd)

    root.Val += toAdd + rightChildRightSum

    leftChildLeftSum, leftChildRightSum := solve(root.Left, root.Val)

    return leftChildLeftSum + leftChildRightSum, leftChildRightSum + root.Val - toAdd
}
