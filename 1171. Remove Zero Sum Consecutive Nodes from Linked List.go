/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type Pair struct {
    Sum int
    Node *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
    curNode := head
    curSum := 0
    prefixStack := []Pair{}
    existsPrefix := map[int]bool{}

    for curNode != nil {
        curSum += curNode.Val

        if curSum == 0 {
            prefixStack = []Pair{}
            existsPrefix = map[int]bool{}
        } else if existsPrefix[curSum] {
            for prefixStack[len(prefixStack) - 1].Sum != curSum {
                sz := len(prefixStack) - 1
                existsPrefix[prefixStack[sz].Sum] = false
                prefixStack = prefixStack[:sz]
            }

            if len(prefixStack) > 0 {
                lastOnStack := prefixStack[len(prefixStack)-1]
                lastOnStack.Node.Next = curNode.Next
            }
        } else {
            prefixStack = append(prefixStack, Pair{curSum, curNode})
            existsPrefix[curSum] = true
        }
        curNode = curNode.Next
    }

    if len(prefixStack) == 0 {
        return nil
    } else {
        return prefixStack[0].Node
    }
}
