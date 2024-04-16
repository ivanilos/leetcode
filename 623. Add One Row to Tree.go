/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if (depth == 1) {
        newRoot := &TreeNode{val, root, nil}
        return newRoot
    }
    _addOneRow(root, val, depth - 1)
    return root
}

func _addOneRow(root *TreeNode, val int, depth int) {
    if root == nil {
        return
    }
    if depth == 1 {
        leftNode := &TreeNode{val, root.Left, nil}
        root.Left = leftNode
        rightNode := &TreeNode{val, nil, root.Right}
        root.Right = rightNode

        return
    } else {
        _addOneRow(root.Left, val, depth - 1)
        _addOneRow(root.Right, val, depth - 1)
    }
}
