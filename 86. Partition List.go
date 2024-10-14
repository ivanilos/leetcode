/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    smallerHead := &ListNode{}
    smallerTail := smallerHead
    biggerHead := &ListNode{}
    biggerTail := biggerHead

    curNode := head
    for curNode != nil {
        if curNode.Val < x {
            smallerTail.Next = curNode
            curNode = curNode.Next
            smallerTail.Next.Next = nil
            smallerTail = smallerTail.Next
        } else {
            biggerTail.Next = curNode
            curNode = curNode.Next
            biggerTail.Next.Next = nil
            biggerTail = biggerTail.Next
        }
    }

    if smallerHead.Next == nil {
        return biggerHead.Next
    } else {
        smallerTail.Next = biggerHead.Next
        return smallerHead.Next
    }
}
