/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
    sz := len(preorder)
    return constructFromPrePostWithIndexes(preorder, postorder, 0, sz - 1, 0, sz - 1)
}

func constructFromPrePostWithIndexes(preorder, postorder []int, preLeft, preRight, postLeft, postRight int) *TreeNode {
    if preLeft > preRight || postLeft > postRight {
        return nil
    }
    root := &TreeNode{preorder[preLeft], nil, nil}

    if preLeft == preRight {
        return root
    }

    for i := postLeft; i <= postRight; i++ {
        if postorder[i] == preorder[preLeft + 1] {
            elems := i - postLeft + 1
            root.Left = constructFromPrePostWithIndexes(preorder, postorder, preLeft + 1, preLeft + elems, postLeft, i)
            root.Right = constructFromPrePostWithIndexes(preorder, postorder, preLeft + elems + 1, preRight, i + 1, postRight - 1)

            break
        }
    }

    return root
}
