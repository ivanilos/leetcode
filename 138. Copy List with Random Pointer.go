/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    
    curNode := head
    var newHead *Node = nil
    var newNextNode *Node = nil

    randomPointerMap := map[*Node]*Node{}
    randomPointerInNextList := map[*Node]*Node{}


    for curNode != nil {
        if newHead == nil {
            newHead = &Node{curNode.Val, curNode.Next, nil}
            newNextNode = newHead
        } else {
            newNextNode.Next = &Node{curNode.Val, curNode.Next, nil}
            newNextNode = newNextNode.Next
        }

        randomPointerMap[newNextNode] = curNode.Random
        randomPointerInNextList[curNode] = newNextNode

        curNode = curNode.Next
    }

    newNextNode = newHead
    for newNextNode != nil {
        oldListNode := randomPointerMap[newNextNode]
        newNextNode.Random = randomPointerInNextList[oldListNode]

        newNextNode = newNextNode.Next
    }

    return newHead
}
