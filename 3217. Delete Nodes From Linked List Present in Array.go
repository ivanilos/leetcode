/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
    forbiddenValues := map[int]bool{}
    for _, num := range nums {
        forbiddenValues[num] = true
    }

    curNode := head
    lastNode := head
    for curNode != nil {
        if _, ok := forbiddenValues[curNode.Val]; ok {
            if curNode == head {
                head = curNode.Next
                lastNode = head
            } else {
                lastNode.Next = curNode.Next
            }
        } else {
            if lastNode != head {
                lastNode.Next = curNode
            }
            lastNode = curNode
        }
        curNode = curNode.Next
    }
    return head
}
