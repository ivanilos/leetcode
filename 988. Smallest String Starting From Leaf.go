/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
    best := ""
    _smallestFromLeaf(root, []rune{}, &best)
    return best
}

func _smallestFromLeaf(root *TreeNode, cur []rune, best *string) {
    if root.Left == nil && root.Right == nil {
        cur = append(cur,rune(root.Val + 'a'))
        finalString := reverse(cur)
        fmt.Println(finalString)
        if *best == "" || finalString < *best {
            *best = finalString
        }
        cur = cur[:len(cur)-1]
        return
    }
    cur = append(cur, rune(root.Val + 'a'))
    if root.Left != nil {
        _smallestFromLeaf(root.Left, cur, best)
    }
    if root.Right != nil {
        _smallestFromLeaf(root.Right, cur, best)
    }
    cur = cur[:len(cur)-1]
}

func reverse(runes []rune) string {
    r := slices.Clone(runes)
    for i, j := 0, len(r) - 1; i < j; i, j = i + 1, j - 1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
