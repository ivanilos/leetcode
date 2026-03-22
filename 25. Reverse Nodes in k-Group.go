/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    cur := head

    var prevListLastNode *ListNode = nil
    var result *ListNode = nil
    for cur != nil {
        start := cur
        for i := 0; i < k - 1 && cur != nil; i++ {
            cur = cur.Next
        }

        if cur != nil {
            fmt.Println(start.Val, cur.Val)
            reverseList(prevListLastNode, start, cur)

            if result == nil {
                result = cur
            }

            prevListLastNode = start

            // start became the new end of list
            cur = start.Next
        }
    }

    return result
}

func reverseList(prevListLastNode *ListNode, start, finish *ListNode) *ListNode {
    if start == nil || start == finish {
        return start
    }

    prev := start
    cur := start.Next

    for cur != finish {
        afterCur := cur.Next

        cur.Next = prev
        prev = cur
        cur = afterCur
    }

    start.Next = cur.Next
    cur.Next = prev
    
    if prevListLastNode != nil {
        prevListLastNode.Next = finish
    }

    return finish
}
