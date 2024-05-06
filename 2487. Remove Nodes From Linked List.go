/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
    st := []*ListNode{}

    cur := head
    for cur != nil {
        for len(st) != 0 {
            topVal := st[len(st) - 1].Val
            if cur.Val > topVal {
                st = st[:len(st) - 1]
            } else {
                break
            }
        }
        if len(st) > 0 {
            st[len(st) - 1].Next = cur
        }
        st = append(st, cur)
        cur = cur.Next
    }
    st[len(st) - 1].Next = nil
    return st[0]
}
