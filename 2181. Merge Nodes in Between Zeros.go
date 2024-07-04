/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    prevZeroNode := head
    curNode := head.Next

    for curNode != nil {
        if curNode.Val != 0 {
            prevZeroNode.Val += curNode.Val
        } else {
            if curNode.Next != nil {
                prevZeroNode.Next = curNode
                prevZeroNode = curNode
            } else {
                prevZeroNode.Next = nil
            }
        }
        curNode = curNode.Next
    }
    return head
}
