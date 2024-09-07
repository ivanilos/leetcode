/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type DPNode struct {
    listNode *ListNode
    treeNode *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
    dp := map[DPNode]bool{}
    return _isSubPath(head, head, root, dp)
}

func _isSubPath(head *ListNode, curNode *ListNode, root *TreeNode, dp map[DPNode]bool) bool {
    dpNode := DPNode{curNode, root}
    if _, ok := dp[dpNode]; ok {
        return dp[dpNode]
    }
    if curNode == nil {
        return true
    }
    if root == nil {
        return false
    }

    // start path
    result := _isSubPath(head, head, root.Left, dp) || _isSubPath(head, head, root.Right, dp)

    // continue path
    if curNode.Val == root.Val {
        result = result || _isSubPath(head, curNode.Next, root.Left, dp) || _isSubPath(head, curNode.Next, root.Right, dp)
    }
    dp[dpNode] = result
    return result
}
