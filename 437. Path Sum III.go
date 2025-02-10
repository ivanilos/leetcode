/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    if root == nil {
        return 0
    }
    result := pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)

    nextTargetSum := targetSum - root.Val
    if nextTargetSum == 0 {
        result++
    }

    result += pathSumWithElem(root.Left, targetSum - root.Val)
    result += pathSumWithElem(root.Right, targetSum - root.Val)

    return result
}

func pathSumWithElem(root *TreeNode, targetSum int) int {
    if root == nil {
        return 0
    }

    result := 0

    nextTargetSum := targetSum - root.Val
    if nextTargetSum == 0 {
        result++
    }

    result += pathSumWithElem(root.Left, targetSum - root.Val)
    result += pathSumWithElem(root.Right, targetSum - root.Val)
    
    return result
}
