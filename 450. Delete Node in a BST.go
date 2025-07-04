/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root != nil {
        if root.Val > key {
            root.Left = deleteNode(root.Left, key)
        } else if root.Val < key {
            root.Right = deleteNode(root.Right, key)
        } else {
            // 0 children
            if root.Left == nil && root.Right == nil {
                return nil
            } else if root.Left == nil { // 1 child
                return root.Right
            } else if root.Right == nil { // 1 child
                return root.Left
            } else { // 2 children
                replacement := root.Right
                for replacement.Left != nil {
                    replacement = replacement.Left
                }
                root.Val = replacement.Val
                root.Right = deleteNode(root.Right, replacement.Val)
            }

        }
    }

    return root
}
