/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    sz := getListSize(head)
    tail := getReversedSecondHalf(head, sz)

    for i := 0; i < sz / 2 ; i++ {
        nextTail := tail.Next
        nextHead := head.Next

        head.Next = tail
        tail.Next = nextHead

        tail = nextTail
        head = nextHead
    }
    head.Next = nil
}

func getListSize(head *ListNode) int {
    result := 0
    for head != nil {
        head = head.Next
        result++
    }
    return result
}

func getReversedSecondHalf(head *ListNode, sz int) *ListNode {
    for i := 0; i < sz / 2; i++ {
        head = head.Next
    }

    prevNode := head
    curNode := head.Next
    for curNode != nil {
        nextNode := curNode.Next
        curNode.Next = prevNode

        prevNode = curNode
        curNode = nextNode
    }
    return prevNode
}
