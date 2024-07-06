/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
    if lessThan3Nodes(head) {
        return []int{-1, -1}
    }
    first := head
    second := first.Next
    third := second.Next

    criticalPos := []int{}
    curPos := 1
    for third != nil {
        if isMidCritical(first, second, third) {
            criticalPos = append(criticalPos, curPos)
        }
        third = third.Next
        second = second.Next
        first = first.Next
        curPos++
    }

    fmt.Println(criticalPos)

    if len(criticalPos) < 2 {
        return []int{-1, -1}
    }

    result := []int{curPos, 0}
    for i := 1; i < len(criticalPos); i++ {
        result[0] = min(result[0], criticalPos[i] - criticalPos[i - 1])
    }
    result[1] = criticalPos[len(criticalPos) - 1] - criticalPos[0]
    
    return result
}

func lessThan3Nodes(head *ListNode) bool {
    counter := 0
    for head != nil && counter < 3 {
        counter++
        head = head.Next
    }
    return counter < 3
}

func isMidCritical(first, second, third *ListNode) bool {
    return first.Val < second.Val && second.Val > third.Val || 
            first.Val > second.Val && second.Val < third.Val
}
