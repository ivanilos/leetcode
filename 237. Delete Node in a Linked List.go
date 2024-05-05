/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    cur := node
    next := node.Next

    for next.Next != nil {
        cur.Val = next.Val
        cur = next
        next = next.Next
    }
    cur.Val = next.Val
    cur.Next = nil
}
