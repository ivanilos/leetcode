/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    n := getLen(head)

    cur := head
    next := head.Next

    cur.Next = nil
    for i := 1; i < n / 2; i++ {
        prev := cur
        cur = next
        next = next.Next

        cur.Next = prev
    }

    result := cur.Val + next.Val
    for i := 1; i < n / 2; i++ {
        cur = cur.Next
        next = next.Next

        result = max(result, cur.Val + next.Val)
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
