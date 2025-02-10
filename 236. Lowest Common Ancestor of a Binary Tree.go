/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    parentMap := map[*TreeNode]*TreeNode{}

    DFS(root, nil, parentMap)

    return solve(p, q, parentMap)
}

func DFS(root *TreeNode, parent *TreeNode, parentMap map[*TreeNode]*TreeNode) {
    if root == nil {
        return
    }

    parentMap[root] = parent

    DFS(root.Left, root, parentMap)
    DFS(root.Right, root, parentMap)    
}

func solve(p, q *TreeNode, parentMap map[*TreeNode]*TreeNode) *TreeNode {
    pToRoot := map[*TreeNode]bool{}

    for p != nil {
        pToRoot[p] = true
        p = parentMap[p]
    }

    for q != nil {
        if _, ok := pToRoot[q]; ok {
            return q
        }
        q = parentMap[q]
    }

    panic("Should not reach here")
}
