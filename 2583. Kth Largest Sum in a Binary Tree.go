/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
    sums := []int64{}
    curNodes := []*TreeNode{root}

    solve(curNodes, &sums)

    sort.Slice(sums, func(a, b int) bool {
        return sums[a] < sums[b]
    })

    if len(sums) - k < 0 {
        return -1
    }

    return sums[len(sums) - k]
}

func solve(curNodes []*TreeNode,  sums *[]int64) {
    if len(curNodes) == 0 {
        return
    }

    curSum := int64(0)
    nextNodes := []*TreeNode{}

    for _, node := range curNodes {
        curSum += int64(node.Val)

        if node.Left != nil {
            nextNodes = append(nextNodes, node.Left)
        }
        if node.Right != nil {
            nextNodes = append(nextNodes, node.Right)
        }
    }

    *sums = append(*sums, curSum)
    solve(nextNodes, sums)
}
