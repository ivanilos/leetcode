/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
    _, _, result := solve(root)

    return result
}

func solve(root *TreeNode) (int, int, int) {
    if root == nil {
        return 0, 0, 0
    }
    leftSz, leftCoins, leftResult := solve(root.Left)
    rightSz, rightCoins, rightResult := solve(root.Right)

    deltaLeft := abs(leftSz - leftCoins)
    deltaRight := abs(rightSz - rightCoins)

    rootSz := 1 + leftSz + rightSz
    rootCoins := root.Val + leftCoins + rightCoins
    rootResult := leftResult + rightResult + deltaLeft + deltaRight

    return rootSz, rootCoins, rootResult
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
