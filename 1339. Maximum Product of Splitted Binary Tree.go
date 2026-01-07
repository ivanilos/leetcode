/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
const MOD = int64(1e9 + 7)

func maxProduct(root *TreeNode) int {
    subTreeSum := map[*TreeNode]int64{}
    getSum(root, subTreeSum)

    result := getMaxProduct(root, subTreeSum, subTreeSum[root])

    return int(result % MOD)
}

func getSum(root *TreeNode, subTreeSum map[*TreeNode]int64) int64 {
    if root == nil {
        return 0
    }

    subTreeSum[root] = int64(root.Val) + getSum(root.Left, subTreeSum) + getSum(root.Right, subTreeSum)
    return subTreeSum[root]
}

func getMaxProduct(node *TreeNode, subTreeSum map[*TreeNode]int64, rootSum int64) int64 {
    if node == nil {
        return 0
    }

    result := max(getMaxProduct(node.Left, subTreeSum, rootSum), getMaxProduct(node.Right, subTreeSum, rootSum))

    result = max(result, subTreeSum[node] * (rootSum - subTreeSum[node]))

    return result
}
