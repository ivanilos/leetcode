/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isEvenOddTree(root *TreeNode) bool {
    if root == nil {
        return true
    }
    curLevelNodes := []*TreeNode{root}
    return isEvenOddTreeInEvenLevel(curLevelNodes)
}

func isEvenOddTreeInEvenLevel(curLevelNodes []*TreeNode) bool {
    if len(curLevelNodes) == 0 {
        return true
    }
    var nextLevelNodes []*TreeNode

    prevVal := curLevelNodes[0].Val - 1
    for _, node := range(curLevelNodes) {
        if node.Val % 2 == 0 || node.Val <= prevVal {
            fmt.Println(node.Val, prevVal)
            return false
        }
        prevVal = node.Val

        if node.Left != nil {
            nextLevelNodes = append(nextLevelNodes, node.Left)
        }
        if node.Right != nil {
            nextLevelNodes = append(nextLevelNodes, node.Right)
        }
    }
    return isEvenOddTreeInOddLevel(nextLevelNodes)

}

func isEvenOddTreeInOddLevel( curLevelNodes []*TreeNode) bool {
    if len(curLevelNodes) == 0 {
        return true
    }
    var nextLevelNodes []*TreeNode

    prevVal := curLevelNodes[0].Val + 1
    for _, node := range(curLevelNodes) {
        if node.Val % 2 != 0 || node.Val >= prevVal {
            return false
        }
        prevVal = node.Val

        if node.Left != nil {
            nextLevelNodes = append(nextLevelNodes, node.Left)
        }
        if node.Right != nil {
            nextLevelNodes = append(nextLevelNodes, node.Right)
        }
    }
    return isEvenOddTreeInEvenLevel(nextLevelNodes)
}
