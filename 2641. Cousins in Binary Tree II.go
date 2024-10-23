/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
    curNodes := []*TreeNode{root}
    toSub := []int{root.Val}

    solve(curNodes, toSub)

    return root
}

func solve(curNodes []*TreeNode, toSub []int) {
    if len(curNodes) == 0 {
        return
    }

    curSum := 0
    for i := 0; i < len(curNodes); i++ {
        curSum += curNodes[i].Val
    }

    nextNodes := []*TreeNode{}
    nextToSub := []int{}
    for i := 0; i < len(curNodes); i++ {
        curNodes[i].Val = curSum - toSub[i]

        childrenSum := getChildrenSum(curNodes[i])

        if curNodes[i].Left != nil {
            nextNodes = append(nextNodes, curNodes[i].Left)
            nextToSub = append(nextToSub, childrenSum)
        }
        if curNodes[i].Right != nil {
            nextNodes = append(nextNodes, curNodes[i].Right)
            nextToSub = append(nextToSub, childrenSum)
        }
    }

    solve(nextNodes, nextToSub)
}

func getChildrenSum(node *TreeNode) int {
    result := 0
    if node.Left != nil {
        result += node.Left.Val
    }
    if node.Right != nil {
        result += node.Right.Val
    }
    return result
}
