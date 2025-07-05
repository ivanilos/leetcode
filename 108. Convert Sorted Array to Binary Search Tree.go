/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    return _sortedArrayToBST(0, len(nums) - 1, nums)
}

func _sortedArrayToBST(l, r int, nums []int) * TreeNode {
    if l > r {
        return nil
    }

    mid := (l + r) / 2
    val := nums[mid]
    node := &TreeNode{val, nil, nil}
    node.Left = _sortedArrayToBST(l, mid - 1, nums)
    node.Right = _sortedArrayToBST(mid + 1, r, nums)

    return node
}
