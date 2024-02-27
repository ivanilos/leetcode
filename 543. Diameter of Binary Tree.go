/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return longestNodePath(root) - 1;
}

func longestNodePath(root * TreeNode) int {
    if root == nil {
        return 0
    }

    result := max(longestNodePath(root.Left),
            max(longestNodePath(root.Right),
                1 + longestNodePathDown(root.Left) + longestNodePathDown(root.Right)))

    return result 
}

func longestNodePathDown(root *TreeNode) int {
    if root == nil {
        return 0
    }

    result := max(1 + longestNodePathDown(root.Left),
                    1 + longestNodePathDown(root.Right))

    return result 
}
