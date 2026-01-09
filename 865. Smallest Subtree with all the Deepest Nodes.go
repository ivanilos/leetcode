/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    result, _ := getsubtreeWithAllDeepest(root, 0)

    return result
}

func getsubtreeWithAllDeepest(root *TreeNode, curDepth int) (*TreeNode, int) {
    if root == nil {
        return nil, curDepth
    }

    leftBest, leftDepth := getsubtreeWithAllDeepest(root.Left, curDepth + 1)
    rightBest, rightDepth := getsubtreeWithAllDeepest(root.Right, curDepth + 1)
    
    if leftDepth == rightDepth {
        return root, leftDepth
    } else if leftDepth > rightDepth {
        return leftBest, leftDepth
    } else {
        return rightBest, rightDepth
    }
}
