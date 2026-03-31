/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    inorderValToIdx := map[int]int{}

    for i, val := range inorder {
        inorderValToIdx[val] = i
    }

    return _buildTree(preorder, inorder, 0, len(inorder) - 1, inorderValToIdx)
}

func _buildTree(preorder, inorder []int, inorderLeft, inorderRight int, inorderValToIdx map[int]int) *TreeNode {
    if inorderLeft > inorderRight {
        return nil
    }

    inorderIdx := inorderValToIdx[preorder[0]]
    leftTreeSize := inorderIdx - inorderLeft
    leftChild := _buildTree(preorder[1:leftTreeSize + 1], inorder, inorderLeft, inorderIdx - 1, inorderValToIdx)
    rightChild := _buildTree(preorder[leftTreeSize + 1:], inorder, inorderIdx + 1, inorderRight, inorderValToIdx)


    return &TreeNode {
        preorder[0],
        leftChild,
        rightChild,
    }
}
