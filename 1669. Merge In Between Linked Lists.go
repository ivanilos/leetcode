/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    beforeFirstPosToRemove := list1

    for i := 1; i < a; i++ {
        beforeFirstPosToRemove = beforeFirstPosToRemove.Next 
    }

    lastPosToRemove := beforeFirstPosToRemove
    for i := a; i <= b; i++ {
        lastPosToRemove = lastPosToRemove.Next
    }

    lastPosList2 := list2
    for lastPosList2.Next != nil {
        lastPosList2 = lastPosList2.Next
    }

    beforeFirstPosToRemove.Next = list2
    lastPosList2.Next = lastPosToRemove.Next

    return list1
}
