/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    if head.Next == nil {
        return head
    }

    oddHead := head
    evenHead := head.Next

    oddTail := head
    evenTail := evenHead

    for oddTail != nil && evenTail != nil {
        oddTail.Next = evenTail.Next

        if oddTail.Next == nil {
            oddTail.Next = evenHead
            break
        } else {
            oddTail = oddTail.Next
            evenTail.Next = oddTail.Next
            evenTail = evenTail.Next
        }
    }

    oddTail.Next = evenHead

    return oddHead
}
