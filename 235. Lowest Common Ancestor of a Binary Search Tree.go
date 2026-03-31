/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, result := _lowestCommonAncestor(root, p, q)

    return result
}

func _lowestCommonAncestor(root, p, q *TreeNode) (map[*TreeNode]bool, *TreeNode) {
    if root == nil {
        return map[*TreeNode]bool{}, nil
    }

    leftInteresting, leftAnswer := _lowestCommonAncestor(root.Left, p, q)
    rightInteresting, rightAnswer := _lowestCommonAncestor(root.Right, p, q)

    if leftAnswer != nil {
        return map[*TreeNode]bool{}, leftAnswer
    } else if rightAnswer != nil {
        return map[*TreeNode]bool{}, rightAnswer
    }

    interestingQuantity := len(leftInteresting) + len(rightInteresting)
    if interestingQuantity == 2 {
        return map[*TreeNode]bool{}, root 
    } else if (root == p || root == q) && interestingQuantity == 1 {
        return map[*TreeNode]bool{}, root
    } else if (root == p || root == q) {
        return map[*TreeNode]bool{root : true}, nil
    }

    return merge(leftInteresting, rightInteresting), nil
}

func merge(m1, m2 map[*TreeNode]bool) map[*TreeNode]bool {
    for key, val := range m2 {
        m1[key] = val
    }
    return m1
}
