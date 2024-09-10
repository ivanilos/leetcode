/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    if head != nil {
        solve(head)
    }
    return head
}

func solve(curNode *ListNode) {
    for curNode.Next != nil {
        nextNode := curNode.Next

        d := gcd(curNode.Val, nextNode.Val)
        curNode.Next = &ListNode{d, curNode.Next}
        curNode = nextNode
    }
}

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a % b)
}
