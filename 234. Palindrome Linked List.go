/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    midList := getMiddleOfList(head)
    reversedMidList := reverseList(midList)
    return isEqual(head, reversedMidList)
}

func getLinkedListSize(head *ListNode) int {
    result := 0
    for head != nil {
        result += 1
        head = head.Next
    }
    return result
}

func getMiddleOfList(head *ListNode) *ListNode {
    sz := getLinkedListSize(head)

    for i := 0; i < (sz + 1) / 2; i++ {
        head = head.Next
    }
    return head
}

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    var prevNode *ListNode
    curNode := head

    for curNode != nil {
        nextNode := curNode.Next
        curNode.Next = prevNode

        prevNode = curNode
        curNode = nextNode
    }
    return prevNode
}

func isEqual(list1 *ListNode, list2 *ListNode) bool {
    for list2 != nil {
        if list1.Val != list2.Val {
            return false
        }
        list1 = list1.Next
        list2 = list2.Next
    }
    return true
}
