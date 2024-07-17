/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    remainingRoots := []*TreeNode{}
    toDeleteMap := map[int]bool{}
    for _, val := range(to_delete) {
        toDeleteMap[val] = true
    }

    solve(root, toDeleteMap, true, &remainingRoots)

    return remainingRoots
}

func solve(root *TreeNode, toDeleteMap map[int]bool, deletedParent bool, remainingRoots *[]*TreeNode) {
    if root == nil {
        return
    }

    shouldDeleteNode := false
    if _, ok := toDeleteMap[root.Val]; ok {
        shouldDeleteNode = true
    }

    if deletedParent && !shouldDeleteNode {
        *remainingRoots = append(*remainingRoots, root)
    }

    solve(root.Left, toDeleteMap, shouldDeleteNode, remainingRoots)
    solve(root.Right, toDeleteMap, shouldDeleteNode, remainingRoots)

    if (root.Left != nil && toDeleteMap[root.Left.Val]) {
        root.Left = nil
    }
    if (root.Right != nil && toDeleteMap[root.Right.Val]) {
        root.Right = nil
    }
}
