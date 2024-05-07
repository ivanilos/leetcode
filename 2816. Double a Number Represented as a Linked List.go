/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
    reversed := reverse(head)
    doubleValues(reversed, head)
    doubledList := reverse(reversed)

    return doubledList 
}

func reverse(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    cur := head
    next := head.Next
    cur.Next = nil

    for next != nil {
        third := next.Next
        next.Next = cur
        cur = next
        next = third
    }

    return cur
}

func doubleValues(head *ListNode, tail *ListNode) {
    carry := 0
    for head != nil {
        nextCarry := (head.Val * 2 + carry) / 10
        head.Val = (head.Val * 2 + carry) % 10
        carry = nextCarry

        head = head.Next
    }
    if carry > 0 {
        tail.Next = &ListNode{carry, nil}
    }
}
