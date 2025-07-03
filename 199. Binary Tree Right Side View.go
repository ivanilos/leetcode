/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    result := BFS(root)

    return result
}

func BFS(root *TreeNode) []int {
    result := []int{}

    curLevel := []*TreeNode{root}
    nextLevel := []*TreeNode{}
    for len(curLevel) > 0 {
        result = append(result, curLevel[len(curLevel) - 1].Val)
        for _, node := range curLevel {
            if node.Left != nil {
                nextLevel = append(nextLevel, node.Left)
            }
            if node.Right != nil {
                nextLevel = append(nextLevel, node.Right)
            }
        }

        curLevel = nextLevel
        nextLevel = []*TreeNode{}
    }

    return result
}
