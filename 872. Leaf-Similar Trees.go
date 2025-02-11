/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    leaf1 := getLeafVal(root1)
    leaf2 := getLeafVal(root2)

    return slices.Equal(leaf1, leaf2)
}

func getLeafVal(root *TreeNode) []int {
    result := []int{}

    DFS(root, &result)

    return result
}

func DFS(root *TreeNode, result *[]int) {
    if root == nil {
        return
    }

    if root.Left == nil && root.Right == nil {
        *result = append(*result, root.Val)
        return
    }

    DFS(root.Left, result)
    DFS(root.Right, result)
}
