/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    counter := 0

    cur := head
    for cur != nil {
        counter++
        cur = cur.Next
    }

    cur = head
    for i := 0; i < counter / 2; i++ {
        cur = cur.Next
    }
    return cur
}
