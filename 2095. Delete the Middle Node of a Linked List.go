/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }

    sz := getSize(head)

    curNode := head
    prev := head

    for i := 0; i < sz / 2; i++ {
        prev = curNode
        curNode = curNode.Next
    }

    prev.Next = curNode.Next

    return head
}

func getSize(root *ListNode) int {
    result := 0
    for root != nil {
        result++
        root = root.Next
    }
    return result
}
