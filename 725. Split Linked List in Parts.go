/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
    len := getLen(head)

    result := []*ListNode{}
    nextHead := head
    for k > 0 {
        nextLen := (len + k - 1) / k
        result = append(result, nextHead)
        nextHead = cropList(nextHead, nextLen)
        len -= nextLen
        k--
    }
    return result
}

func getLen(head *ListNode) int {
    result := 0
    for head != nil {
        result++
        head = head.Next
    }
    return result
}

func cropList(head *ListNode, len int) *ListNode {
    for i := 0; i < len - 1 && head != nil; i++ {
        head = head.Next
    }

    var nextStart *ListNode = nil
    if head != nil {
        nextStart = head.Next
        head.Next = nil
    }
    return nextStart
}
