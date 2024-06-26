/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
    nodeVals := []int{}
    getVals(root, &nodeVals)
    return build(nodeVals, 0, len(nodeVals) - 1)
}

func getVals(root *TreeNode, vals *[]int) {
    if root == nil {
        return
    }
    getVals(root.Left, vals)
    *vals = append(*vals, root.Val)
    getVals(root.Right, vals)
}

func build(nodeVals []int, left, right int) *TreeNode {
    if left > right {
        return nil
    }
    mid := (left + right) / 2
    node := TreeNode{nodeVals[mid], nil, nil}
    node.Left = build(nodeVals, left, mid - 1)
    node.Right = build(nodeVals, mid + 1, right)

    return &node
}
