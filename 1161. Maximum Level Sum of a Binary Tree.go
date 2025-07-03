/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    result := BFS(root)

    return result
}

func BFS(root *TreeNode) int {
    result := 1
    resultSum := root.Val

    curLevelNodes := []*TreeNode{root}
    nextLevelNodes := []*TreeNode{}
    level := 0
    for len(curLevelNodes) > 0 {
        level++

        curSum := 0
        for _, node := range curLevelNodes {
            curSum += node.Val
            if node.Left != nil {
                nextLevelNodes = append(nextLevelNodes, node.Left)
            }
            if node.Right != nil {
                nextLevelNodes = append(nextLevelNodes, node.Right)
            }
        }

        if curSum > resultSum {
            resultSum = curSum
            result = level
        }

        curLevelNodes = nextLevelNodes
        nextLevelNodes = []*TreeNode{}
    }

    return result
}
